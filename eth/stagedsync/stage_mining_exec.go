package stagedsync

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sync/atomic"
	"time"

	"github.com/holiman/uint256"
	libcommon "github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon-lib/kv"
	"github.com/ledgerwatch/erigon-lib/kv/memdb"
	"github.com/ledgerwatch/erigon-lib/txpool"
	types2 "github.com/ledgerwatch/erigon-lib/types"
	"github.com/ledgerwatch/log/v3"
	"golang.org/x/net/context"

	"github.com/ledgerwatch/erigon/rlp"
	"github.com/ledgerwatch/erigon/turbo/services"

	"github.com/ledgerwatch/erigon/common"
	"github.com/ledgerwatch/erigon/consensus"
	"github.com/ledgerwatch/erigon/consensus/misc"
	"github.com/ledgerwatch/erigon/core"
	"github.com/ledgerwatch/erigon/core/rawdb"
	"github.com/ledgerwatch/erigon/core/state"
	"github.com/ledgerwatch/erigon/core/systemcontracts"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/core/vm"
	"github.com/ledgerwatch/erigon/eth/stagedsync/stages"
	"github.com/ledgerwatch/erigon/params"
	"github.com/ledgerwatch/erigon/turbo/snapshotsync"
)

type MiningExecCfg struct {
	db          kv.RwDB
	miningState MiningState
	notifier    ChainEventNotifier
	chainConfig params.ChainConfig
	engine      consensus.Engine
	blockReader services.FullBlockReader
	vmConfig    *vm.Config
	tmpdir      string
	interrupt   *int32
	payloadId   uint64
	txPool2     *txpool.TxPool
	txPool2DB   kv.RoDB
}

func StageMiningExecCfg(
	db kv.RwDB,
	miningState MiningState,
	notifier ChainEventNotifier,
	chainConfig params.ChainConfig,
	engine consensus.Engine,
	vmConfig *vm.Config,
	tmpdir string,
	interrupt *int32,
	payloadId uint64,
	txPool2 *txpool.TxPool,
	txPool2DB kv.RoDB,
) MiningExecCfg {
	return MiningExecCfg{
		db:          db,
		miningState: miningState,
		notifier:    notifier,
		chainConfig: chainConfig,
		engine:      engine,
		blockReader: snapshotsync.NewBlockReader(),
		vmConfig:    vmConfig,
		tmpdir:      tmpdir,
		interrupt:   interrupt,
		payloadId:   payloadId,
		txPool2:     txPool2,
		txPool2DB:   txPool2DB,
	}
}

// SpawnMiningExecStage
// TODO:
// - resubmitAdjustCh - variable is not implemented
func SpawnMiningExecStage(s *StageState, tx kv.RwTx, cfg MiningExecCfg, quit <-chan struct{}) error {
	cfg.vmConfig.NoReceipts = false
	chainID, _ := uint256.FromBig(cfg.chainConfig.ChainID)
	logPrefix := s.LogPrefix()
	current := cfg.miningState.MiningBlock
	noempty := true

	stateReader := state.NewPlainStateReader(tx)
	ibs := state.New(stateReader)
	stateWriter := state.NewPlainStateWriter(tx, tx, current.Header.Number.Uint64())
	if cfg.chainConfig.DAOForkSupport && cfg.chainConfig.DAOForkBlock != nil && cfg.chainConfig.DAOForkBlock.Cmp(current.Header.Number) == 0 {
		misc.ApplyDAOHardFork(ibs)
	}
	systemcontracts.UpgradeBuildInSystemContract(&cfg.chainConfig, current.Header.Number, ibs)

	// Create an empty block based on temporary copied state for
	// sealing in advance without waiting block execution finished.
	if !noempty {
		log.Info("Commit an empty block", "number", current.Header.Number)
		return nil
	}

	getHeader := func(hash common.Hash, number uint64) *types.Header { return rawdb.ReadHeader(tx, hash, number) }
	executionAt, err := s.ExecutionAt(tx)
	if err != nil {
		return err
	}

	// Short circuit if there is no available pending transactions.
	// But if we disable empty precommit already, ignore it. Since
	// empty block is necessary to keep the liveness of the network.
	if noempty {
		killSwitch := 0
		var cumLogs types.Logs

		// some code paths pre-load transactions such as from the integration tool
		// we need to make sure these transactions are processed instead of
		// looking in the pool
		if current.PreparedTxs != nil && !current.PreparedTxs.Empty() {
			logs, err := addTransactionsToMiningBlock(logPrefix, current, cfg.chainConfig, cfg.vmConfig, getHeader, cfg.engine, current.PreparedTxs, cfg.miningState.MiningConfig.Etherbase, ibs, quit, cfg.interrupt, cfg.payloadId)
			if err != nil {
				return err
			}
			cumLogs = append(cumLogs, logs...)
		} else {
			// keep looping until we're out of gas for the block or the txpool has nothing left
			simulationTx := memdb.NewMemoryBatch(tx, cfg.tmpdir)
			defer simulationTx.Rollback()
			for {
				nextBatch, err := getNextTransactions(cfg, chainID, current.Header, 50, executionAt, simulationTx)
				if err != nil {
					return err
				}

				if !nextBatch.Empty() {
					logs, err := addTransactionsToMiningBlock(logPrefix, current, cfg.chainConfig, cfg.vmConfig, getHeader, cfg.engine, nextBatch, cfg.miningState.MiningConfig.Etherbase, ibs, quit, cfg.interrupt, cfg.payloadId)
					if err != nil {
						return err
					}
					cumLogs = append(cumLogs, logs...)
					log.Debug(fmt.Sprintf("[%s] Added transactions to mining block", logPrefix), "count", len(logs), "gas_remaining", current.Header.GasLimit-current.Header.GasUsed)

				} else {
					break
				}
				killSwitch++
				if killSwitch >= 1000 {
					log.Debug(fmt.Sprintf("[%s] Mining exec killswitch limit hit", logPrefix))
					break
				}
			}
		}

		// We don't push the pendingLogsEvent while we are mining. The reason is that
		// when we are mining, the worker will regenerate a mining block every 3 seconds.
		// In order to avoid pushing the repeated pendingLog, we disable the pending log pushing.
		//if !w.isRunning() {
		//}
		log.Info(fmt.Sprintf("[%s] Mined transactions", logPrefix), "count", len(current.Txs), "gas_used", current.Header.GasUsed)
		NotifyPendingLogs(logPrefix, cfg.notifier, cumLogs)
	}

	log.Debug("SpawnMiningExecStage", "block txn", current.Txs.Len(), "payload", cfg.payloadId)
	if current.Uncles == nil {
		current.Uncles = []*types.Header{}
	}
	if current.Txs == nil {
		current.Txs = []types.Transaction{}
	}
	if current.Receipts == nil {
		current.Receipts = types.Receipts{}
	}

	_, current.Txs, current.Receipts, err = core.FinalizeBlockExecution(cfg.engine, stateReader, current.Header, current.Txs, current.Uncles, stateWriter,
		&cfg.chainConfig, ibs, current.Receipts, current.Withdrawals, epochReader{tx: tx}, chainReader{config: &cfg.chainConfig, tx: tx, blockReader: cfg.blockReader}, true)
	if err != nil {
		return err
	}
	log.Debug("FinalizeBlockExecution", "current txn", current.Txs.Len(), "current receipt", current.Receipts.Len(), "payload", cfg.payloadId)

	/*
		if w.isRunning() {
			if interval != nil {
				interval()
			}

			select {
			case w.taskCh <- &task{receipts: receipts, state: s, tds: w.env.tds, block: block, createdAt: time.Now(), ctx: ctx}:
				log.Debug("mining: worker task event",
					"number", block.NumberU64(),
					"hash", block.Hash().String(),
					"parentHash", block.ParentHash().String(),
				)

				log.Info("Commit new mining work", "number", block.Number(), "sealhash", w.engine.SealHash(block.Header()),
					"uncles", len(uncles), "txs", w.env.tcount,
					"gas", block.GasUsed(), "fees", totalFees(block, receipts),
					"elapsed", common.PrettyDuration(time.Since(start)))

			case <-w.exitCh:
				log.Info("Worker has exited")
			}
		}
		if update {
			w.updateSnapshot()
		}
	*/

	// hack: pretend that we are real execution stage - next stages will rely on this progress
	if err := stages.SaveStageProgress(tx, stages.Execution, current.Header.Number.Uint64()); err != nil {
		return err
	}
	return nil
}

func getNextTransactions(
	cfg MiningExecCfg,
	chainID *uint256.Int,
	header *types.Header,
	amount uint16,
	executionAt uint64,
	simulationTx *memdb.MemoryMutation,
) (types.TransactionsStream, error) {
	txSlots := types2.TxsRlp{}
	var onTime bool
	if err := cfg.txPool2DB.View(context.Background(), func(poolTx kv.Tx) error {
		var err error
		counter := 0
		for !onTime && counter < 1000 {
			remainingGas := header.GasLimit - header.GasUsed
			if onTime, err = cfg.txPool2.YieldBest(amount, &txSlots, poolTx, executionAt, remainingGas); err != nil {
				return err
			}
			time.Sleep(1 * time.Millisecond)
			counter++
		}
		return nil
	}); err != nil {
		return nil, err
	}

	var txs []types.Transaction //nolint:prealloc
	reader := bytes.NewReader([]byte{})
	stream := new(rlp.Stream)
	for i := range txSlots.Txs {
		reader.Reset(txSlots.Txs[i])
		stream.Reset(reader, uint64(len(txSlots.Txs[i])))

		transaction, err := types.DecodeTransaction(stream)
		if err == io.EOF {
			continue
		}
		if err != nil {
			return nil, err
		}
		if !transaction.GetChainID().IsZero() && transaction.GetChainID().Cmp(chainID) != 0 {
			continue
		}

		var sender common.Address
		copy(sender[:], txSlots.Senders.At(i))

		// Check if tx nonce is too low
		txs = append(txs, transaction)
		txs[len(txs)-1].SetSender(sender)
	}

	blockNum := executionAt + 1
	txs, err := filterBadTransactions(txs, cfg.chainConfig, blockNum, header.BaseFee, simulationTx)
	if err != nil {
		return nil, err
	}

	return types.NewTransactionsFixedOrder(txs), nil
}

func addTransactionsToMiningBlock(logPrefix string, current *MiningBlock, chainConfig params.ChainConfig, vmConfig *vm.Config, getHeader func(hash common.Hash, number uint64) *types.Header, engine consensus.Engine, txs types.TransactionsStream, coinbase common.Address, ibs *state.IntraBlockState, quit <-chan struct{}, interrupt *int32, payloadId uint64) (types.Logs, error) {
	header := current.Header
	tcount := 0
	gasPool := new(core.GasPool).AddGas(header.GasLimit - header.GasUsed)
	signer := types.MakeSigner(&chainConfig, header.Number.Uint64())

	var coalescedLogs types.Logs
	noop := state.NewNoopWriter()

	var miningCommitTx = func(txn types.Transaction, coinbase common.Address, vmConfig *vm.Config, chainConfig params.ChainConfig, ibs *state.IntraBlockState, current *MiningBlock) ([]*types.Log, error) {
		ibs.Prepare(txn.Hash(), common.Hash{}, tcount)
		gasSnap := gasPool.Gas()
		snap := ibs.Snapshot()
		log.Debug("addTransactionsToMiningBlock", "txn hash", txn.Hash())
		receipt, _, err := core.ApplyTransaction(&chainConfig, core.GetHashFn(header, getHeader), engine, &coinbase, gasPool, ibs, noop, header, txn, &header.GasUsed, *vmConfig)
		if err != nil {
			ibs.RevertToSnapshot(snap)
			gasPool = new(core.GasPool).AddGas(gasSnap) // restore gasPool as well as ibs
			return nil, err
		}

		current.Txs = append(current.Txs, txn)
		current.Receipts = append(current.Receipts, receipt)
		return receipt.Logs, nil
	}

	var stopped *time.Ticker
	defer func() {
		if stopped != nil {
			stopped.Stop()
		}
	}()
LOOP:
	for {
		// see if we need to stop now
		if stopped != nil {
			select {
			case <-stopped.C:
				break LOOP
			default:
			}
		}

		if err := libcommon.Stopped(quit); err != nil {
			return nil, err
		}

		if interrupt != nil && atomic.LoadInt32(interrupt) != 0 && stopped == nil {
			log.Debug("Transaction adding was requested to stop", "payload", payloadId)
			// ensure we run for at least 500ms after the request to stop comes in from GetPayload
			stopped = time.NewTicker(500 * time.Millisecond)
		}
		// If we don't have enough gas for any further transactions then we're done
		if gasPool.Gas() < params.TxGas {
			log.Debug(fmt.Sprintf("[%s] Not enough gas for further transactions", logPrefix), "have", gasPool, "want", params.TxGas)
			break
		}
		// Retrieve the next transaction and abort if all done
		txn := txs.Peek()
		if txn == nil {
			break
		}

		// We use the eip155 signer regardless of the env hf.
		from, err := txn.Sender(*signer)
		if err != nil {
			log.Warn(fmt.Sprintf("[%s] Could not recover transaction sender", logPrefix), "hash", txn.Hash(), "err", err)
			txs.Pop()
			continue
		}

		// Check whether the txn is replay protected. If we're not in the EIP155 (Spurious Dragon) hf
		// phase, start ignoring the sender until we do.
		if txn.Protected() && !chainConfig.IsSpuriousDragon(header.Number.Uint64()) {
			log.Debug(fmt.Sprintf("[%s] Ignoring replay protected transaction", logPrefix), "hash", txn.Hash(), "eip155", chainConfig.SpuriousDragonBlock)

			txs.Pop()
			continue
		}

		// Start executing the transaction
		logs, err := miningCommitTx(txn, coinbase, vmConfig, chainConfig, ibs, current)

		if errors.Is(err, core.ErrGasLimitReached) {
			// Pop the env out-of-gas transaction without shifting in the next from the account
			log.Debug(fmt.Sprintf("[%s] Gas limit exceeded for env block", logPrefix), "hash", txn.Hash(), "sender", from)
			txs.Pop()
		} else if errors.Is(err, core.ErrNonceTooLow) {
			// New head notification data race between the transaction pool and miner, shift
			log.Debug(fmt.Sprintf("[%s] Skipping transaction with low nonce", logPrefix), "hash", txn.Hash(), "sender", from, "nonce", txn.GetNonce())
			txs.Shift()
		} else if errors.Is(err, core.ErrNonceTooHigh) {
			// Reorg notification data race between the transaction pool and miner, skip account =
			log.Debug(fmt.Sprintf("[%s] Skipping transaction with high nonce", logPrefix), "hash", txn.Hash(), "sender", from, "nonce", txn.GetNonce())
			txs.Pop()
		} else if err == nil {
			// Everything ok, collect the logs and shift in the next transaction from the same account
			log.Debug(fmt.Sprintf("[%s] addTransactionsToMiningBlock Successful", logPrefix), "sender", from, "nonce", txn.GetNonce(), "payload", payloadId)
			coalescedLogs = append(coalescedLogs, logs...)
			tcount++
			txs.Shift()
		} else {
			// Strange error, discard the transaction and get the next in line (note, the
			// nonce-too-high clause will prevent us from executing in vain).
			log.Debug(fmt.Sprintf("[%s] Skipping transaction", logPrefix), "hash", txn.Hash(), "sender", from, "err", err)
			txs.Shift()
		}
	}

	/*
		// Notify resubmit loop to decrease resubmitting interval if env interval is larger
		// than the user-specified one.
		if interrupt != nil {
			w.resubmitAdjustCh <- &intervalAdjust{inc: false}
		}
	*/
	return coalescedLogs, nil

}

func NotifyPendingLogs(logPrefix string, notifier ChainEventNotifier, logs types.Logs) {
	if len(logs) == 0 {
		return
	}

	if notifier == nil {
		log.Debug(fmt.Sprintf("[%s] rpc notifier is not set, rpc daemon won't be updated about pending logs", logPrefix))
		return
	}
	notifier.OnNewPendingLogs(logs)
}
