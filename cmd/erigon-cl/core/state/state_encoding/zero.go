package state_encoding

// ZeroHashes is a representation of all zerohashes of
// varying depths till h=100.
var ZeroHashes = [100][32]byte{
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{245, 165, 253, 66, 209, 106, 32, 48, 39, 152, 239, 110, 211, 9, 151, 155, 67, 0, 61, 35, 32, 217, 240, 232, 234, 152, 49, 169, 39, 89, 251, 75},
	{219, 86, 17, 78, 0, 253, 212, 193, 248, 92, 137, 43, 243, 90, 201, 168, 146, 137, 170, 236, 177, 235, 208, 169, 108, 222, 96, 106, 116, 139, 93, 113},
	{199, 128, 9, 253, 240, 127, 197, 106, 17, 241, 34, 55, 6, 88, 163, 83, 170, 165, 66, 237, 99, 228, 76, 75, 193, 95, 244, 205, 16, 90, 179, 60},
	{83, 109, 152, 131, 127, 45, 209, 101, 165, 93, 94, 234, 233, 20, 133, 149, 68, 114, 213, 111, 36, 109, 242, 86, 191, 60, 174, 25, 53, 42, 18, 60},
	{158, 253, 224, 82, 170, 21, 66, 159, 174, 5, 186, 212, 208, 177, 215, 198, 77, 166, 77, 3, 215, 161, 133, 74, 88, 140, 44, 184, 67, 12, 13, 48},
	{216, 141, 223, 238, 212, 0, 168, 117, 85, 150, 178, 25, 66, 193, 73, 126, 17, 76, 48, 46, 97, 24, 41, 15, 145, 230, 119, 41, 118, 4, 31, 161},
	{135, 235, 13, 219, 165, 126, 53, 246, 210, 134, 103, 56, 2, 164, 175, 89, 117, 226, 37, 6, 199, 207, 76, 100, 187, 107, 229, 238, 17, 82, 127, 44},
	{38, 132, 100, 118, 253, 95, 197, 74, 93, 67, 56, 81, 103, 201, 81, 68, 242, 100, 63, 83, 60, 200, 91, 185, 209, 107, 120, 47, 141, 125, 177, 147},
	{80, 109, 134, 88, 45, 37, 36, 5, 184, 64, 1, 135, 146, 202, 210, 191, 18, 89, 241, 239, 90, 165, 248, 135, 225, 60, 178, 240, 9, 79, 81, 225},
	{255, 255, 10, 215, 230, 89, 119, 47, 149, 52, 193, 149, 200, 21, 239, 196, 1, 78, 241, 225, 218, 237, 68, 4, 192, 99, 133, 209, 17, 146, 233, 43},
	{108, 240, 65, 39, 219, 5, 68, 28, 216, 51, 16, 122, 82, 190, 133, 40, 104, 137, 14, 67, 23, 230, 160, 42, 180, 118, 131, 170, 117, 150, 66, 32},
	{183, 208, 95, 135, 95, 20, 0, 39, 239, 81, 24, 162, 36, 123, 187, 132, 206, 143, 47, 15, 17, 35, 98, 48, 133, 218, 247, 150, 12, 50, 159, 95},
	{223, 106, 245, 245, 187, 219, 107, 233, 239, 138, 166, 24, 228, 191, 128, 115, 150, 8, 103, 23, 30, 41, 103, 111, 139, 40, 77, 234, 106, 8, 168, 94},
	{181, 141, 144, 15, 94, 24, 46, 60, 80, 239, 116, 150, 158, 161, 108, 119, 38, 197, 73, 117, 124, 194, 53, 35, 195, 105, 88, 125, 167, 41, 55, 132},
	{212, 154, 117, 2, 255, 207, 176, 52, 11, 29, 120, 133, 104, 133, 0, 202, 48, 129, 97, 167, 249, 107, 98, 223, 157, 8, 59, 113, 252, 200, 242, 187},
	{143, 230, 177, 104, 146, 86, 192, 211, 133, 244, 47, 91, 190, 32, 39, 162, 44, 25, 150, 225, 16, 186, 151, 193, 113, 211, 229, 148, 141, 233, 43, 235},
	{141, 13, 99, 195, 158, 186, 222, 133, 9, 224, 174, 60, 156, 56, 118, 251, 95, 161, 18, 190, 24, 249, 5, 236, 172, 254, 203, 146, 5, 118, 3, 171},
	{149, 238, 200, 178, 229, 65, 202, 212, 233, 29, 227, 131, 133, 242, 224, 70, 97, 159, 84, 73, 108, 35, 130, 203, 108, 172, 213, 185, 140, 38, 245, 164},
	{248, 147, 233, 8, 145, 119, 117, 182, 43, 255, 35, 41, 77, 187, 227, 161, 205, 142, 108, 193, 195, 91, 72, 1, 136, 123, 100, 106, 111, 129, 241, 127},
	{205, 219, 167, 181, 146, 227, 19, 51, 147, 193, 97, 148, 250, 199, 67, 26, 191, 47, 84, 133, 237, 113, 29, 178, 130, 24, 60, 129, 158, 8, 235, 170},
	{138, 141, 127, 227, 175, 140, 170, 8, 90, 118, 57, 168, 50, 0, 20, 87, 223, 185, 18, 138, 128, 97, 20, 42, 208, 51, 86, 41, 255, 35, 255, 156},
	{254, 179, 195, 55, 215, 165, 26, 111, 191, 0, 185, 227, 76, 82, 225, 201, 25, 92, 150, 155, 212, 231, 160, 191, 213, 29, 92, 91, 237, 156, 17, 103},
	{231, 31, 10, 168, 60, 195, 46, 223, 190, 250, 159, 77, 62, 1, 116, 202, 133, 24, 46, 236, 159, 58, 9, 246, 166, 192, 223, 99, 119, 165, 16, 215},
	{49, 32, 111, 168, 10, 80, 187, 106, 190, 41, 8, 80, 88, 241, 98, 18, 33, 42, 96, 238, 200, 240, 73, 254, 203, 146, 216, 200, 224, 168, 75, 192},
	{33, 53, 43, 254, 203, 237, 221, 233, 147, 131, 159, 97, 76, 61, 172, 10, 62, 227, 117, 67, 249, 180, 18, 177, 97, 153, 220, 21, 142, 35, 181, 68},
	{97, 158, 49, 39, 36, 187, 109, 124, 49, 83, 237, 157, 231, 145, 215, 100, 163, 102, 179, 137, 175, 19, 197, 139, 248, 168, 217, 4, 129, 164, 103, 101},
	{124, 221, 41, 134, 38, 130, 80, 98, 141, 12, 16, 227, 133, 197, 140, 97, 145, 230, 251, 224, 81, 145, 188, 192, 79, 19, 63, 44, 234, 114, 193, 196},
	{132, 137, 48, 189, 123, 168, 202, 197, 70, 97, 7, 33, 19, 251, 39, 136, 105, 224, 123, 184, 88, 127, 145, 57, 41, 51, 55, 77, 1, 123, 203, 225},
	{136, 105, 255, 44, 34, 178, 140, 193, 5, 16, 217, 133, 50, 146, 128, 51, 40, 190, 79, 176, 232, 4, 149, 232, 187, 141, 39, 31, 91, 136, 150, 54},
	{181, 254, 40, 231, 159, 27, 133, 15, 134, 88, 36, 108, 233, 182, 161, 231, 180, 159, 192, 109, 183, 20, 62, 143, 224, 180, 242, 176, 197, 82, 58, 92},
	{152, 94, 146, 159, 112, 175, 40, 208, 189, 209, 169, 10, 128, 143, 151, 127, 89, 124, 124, 119, 140, 72, 158, 152, 211, 189, 137, 16, 211, 26, 192, 247},
	{198, 246, 126, 2, 230, 228, 225, 189, 239, 185, 148, 198, 9, 137, 83, 243, 70, 54, 186, 43, 108, 162, 10, 71, 33, 210, 178, 106, 136, 103, 34, 255},
	{28, 154, 126, 95, 241, 207, 72, 180, 173, 21, 130, 211, 244, 228, 161, 0, 79, 59, 32, 216, 197, 162, 183, 19, 135, 164, 37, 74, 217, 51, 235, 197},
	{47, 7, 90, 226, 41, 100, 107, 111, 106, 237, 25, 165, 227, 114, 207, 41, 80, 129, 64, 30, 184, 147, 255, 89, 155, 63, 154, 204, 12, 13, 62, 125},
	{50, 137, 33, 222, 181, 150, 18, 7, 104, 1, 232, 205, 97, 89, 33, 7, 181, 198, 124, 121, 184, 70, 89, 92, 198, 50, 12, 57, 91, 70, 54, 44},
	{191, 185, 9, 253, 178, 54, 173, 36, 17, 180, 228, 136, 56, 16, 160, 116, 184, 64, 70, 70, 137, 152, 108, 63, 138, 128, 145, 130, 126, 23, 195, 39},
	{85, 216, 251, 54, 135, 186, 59, 164, 159, 52, 44, 119, 245, 161, 248, 155, 236, 131, 216, 17, 68, 110, 26, 70, 113, 57, 33, 61, 100, 11, 106, 116},
	{247, 33, 13, 79, 142, 126, 16, 57, 121, 14, 123, 244, 239, 162, 7, 85, 90, 16, 166, 219, 29, 212, 185, 93, 163, 19, 170, 168, 139, 136, 254, 118},
	{173, 33, 181, 22, 203, 198, 69, 255, 227, 74, 181, 222, 28, 138, 239, 140, 212, 231, 248, 210, 181, 30, 142, 20, 86, 173, 199, 86, 60, 218, 32, 111},
	{107, 254, 141, 43, 204, 66, 55, 183, 74, 80, 71, 5, 142, 244, 85, 51, 158, 205, 115, 96, 203, 99, 191, 187, 142, 229, 68, 142, 100, 48, 186, 4},
	{167, 242, 60, 233, 24, 23, 64, 220, 34, 12, 129, 71, 130, 101, 79, 238, 106, 206, 185, 241, 236, 146, 34, 196, 226, 70, 125, 10, 177, 104, 8, 55},
	{174, 249, 71, 108, 137, 89, 10, 44, 140, 201, 179, 183, 79, 73, 103, 199, 87, 196, 157, 152, 102, 164, 75, 172, 242, 31, 162, 237, 103, 93, 223, 162},
	{154, 66, 188, 173, 130, 246, 169, 228, 18, 132, 216, 8, 234, 211, 25, 242, 159, 59, 8, 32, 157, 104, 15, 14, 44, 231, 21, 16, 208, 113, 226, 5},
	{209, 166, 109, 53, 74, 103, 185, 207, 23, 149, 113, 216, 229, 249, 119, 146, 113, 110, 141, 212, 236, 68, 25, 104, 57, 163, 247, 198, 183, 79, 139, 172},
	{250, 250, 48, 37, 242, 248, 149, 9, 194, 199, 28, 116, 251, 160, 205, 146, 133, 142, 244, 155, 7, 128, 251, 84, 121, 116, 108, 138, 155, 252, 179, 70},
	{51, 52, 167, 193, 231, 246, 112, 90, 166, 1, 26, 106, 148, 150, 69, 1, 109, 180, 172, 222, 12, 169, 171, 214, 109, 199, 157, 130, 102, 66, 48, 86},
	{7, 150, 253, 117, 102, 79, 174, 247, 68, 238, 78, 82, 215, 39, 30, 43, 187, 118, 159, 145, 237, 111, 155, 116, 216, 182, 148, 245, 102, 6, 133, 44},
	{123, 163, 174, 74, 65, 127, 232, 84, 91, 20, 43, 200, 159, 74, 220, 215, 174, 19, 148, 28, 186, 183, 117, 11, 131, 233, 240, 166, 109, 22, 190, 100},
	{120, 143, 175, 204, 74, 165, 32, 57, 154, 219, 174, 209, 149, 248, 177, 44, 78, 179, 30, 193, 1, 104, 229, 10, 171, 198, 89, 166, 174, 165, 22, 220},
	{232, 51, 215, 166, 113, 96, 230, 139, 244, 201, 4, 74, 83, 7, 125, 242, 114, 122, 208, 12, 243, 111, 73, 73, 199, 182, 129, 169, 18, 20, 12, 187},
	{48, 158, 171, 240, 149, 220, 103, 20, 249, 244, 216, 100, 187, 165, 175, 250, 224, 179, 90, 226, 245, 227, 86, 91, 204, 58, 71, 178, 18, 118, 119, 1},
	{34, 106, 142, 190, 250, 40, 134, 101, 166, 68, 165, 2, 115, 51, 94, 251, 182, 16, 81, 15, 36, 27, 91, 114, 12, 138, 54, 141, 89, 166, 154, 93},
	{65, 171, 253, 153, 84, 37, 130, 118, 37, 147, 129, 49, 175, 12, 79, 51, 254, 11, 212, 104, 140, 34, 44, 33, 250, 157, 168, 232, 156, 170, 3, 248},
	{68, 44, 100, 46, 245, 15, 161, 166, 103, 166, 230, 209, 5, 199, 124, 92, 195, 254, 200, 215, 170, 37, 112, 207, 26, 48, 119, 181, 3, 195, 128, 105},
	{160, 160, 141, 252, 155, 66, 217, 108, 45, 225, 155, 109, 18, 123, 138, 225, 54, 221, 207, 62, 90, 208, 220, 228, 34, 196, 90, 86, 246, 31, 106, 116},
	{125, 52, 131, 130, 175, 9, 109, 190, 11, 240, 134, 199, 187, 57, 178, 162, 192, 188, 54, 182, 33, 171, 12, 115, 142, 152, 133, 215, 49, 216, 23, 64},
	{58, 177, 52, 117, 29, 25, 18, 105, 2, 108, 134, 153, 78, 170, 139, 67, 168, 59, 74, 209, 246, 208, 231, 115, 129, 196, 226, 151, 74, 251, 200, 246},
	{154, 116, 82, 97, 29, 178, 210, 62, 174, 38, 249, 189, 187, 136, 149, 142, 244, 76, 100, 208, 254, 152, 123, 233, 247, 38, 173, 249, 56, 245, 15, 108},
	{114, 92, 127, 129, 96, 55, 191, 228, 82, 205, 30, 123, 163, 90, 196, 126, 220, 180, 154, 154, 43, 39, 174, 202, 112, 220, 228, 131, 203, 125, 237, 31},
	{44, 234, 26, 245, 31, 178, 139, 98, 136, 124, 57, 153, 138, 201, 254, 244, 223, 222, 218, 31, 7, 224, 113, 186, 85, 138, 23, 58, 253, 6, 203, 195},
	{255, 29, 89, 249, 139, 108, 85, 29, 149, 8, 147, 87, 5, 125, 92, 139, 226, 100, 2, 39, 158, 157, 240, 177, 223, 26, 16, 183, 43, 243, 146, 127},
	{47, 138, 24, 31, 124, 153, 221, 33, 90, 117, 41, 191, 226, 150, 169, 96, 58, 20, 70, 115, 113, 134, 210, 26, 235, 139, 199, 174, 89, 225, 253, 33},
	{236, 197, 2, 201, 177, 20, 95, 57, 80, 203, 125, 62, 56, 66, 68, 111, 129, 164, 240, 223, 29, 245, 55, 206, 225, 57, 239, 100, 234, 152, 75, 217},
	{200, 133, 194, 54, 20, 2, 73, 201, 225, 100, 14, 94, 153, 251, 151, 45, 129, 251, 179, 30, 165, 226, 159, 189, 222, 6, 54, 39, 240, 214, 189, 200},
	{48, 60, 227, 136, 9, 186, 122, 119, 182, 96, 173, 11, 7, 74, 249, 198, 188, 213, 192, 43, 191, 242, 243, 176, 36, 134, 51, 176, 184, 118, 228, 73},
	{47, 212, 195, 43, 10, 101, 97, 109, 75, 206, 185, 226, 242, 189, 77, 207, 117, 53, 84, 111, 67, 58, 62, 29, 69, 206, 84, 171, 192, 89, 200, 103},
	{47, 141, 48, 4, 136, 171, 79, 116, 100, 217, 238, 158, 89, 216, 10, 170, 138, 32, 57, 175, 85, 19, 243, 32, 229, 163, 8, 60, 99, 234, 104, 239},
	{72, 86, 47, 42, 177, 135, 58, 97, 32, 245, 117, 38, 122, 55, 219, 71, 13, 74, 107, 200, 62, 209, 173, 144, 62, 100, 247, 179, 117, 87, 102, 173},
	{40, 32, 249, 7, 55, 7, 206, 255, 106, 14, 94, 43, 207, 202, 141, 115, 210, 53, 173, 231, 13, 10, 253, 83, 92, 145, 119, 251, 146, 102, 201, 247},
	{243, 246, 168, 21, 31, 100, 246, 189, 221, 196, 184, 192, 150, 60, 87, 18, 238, 244, 125, 110, 180, 50, 241, 38, 153, 197, 41, 89, 20, 240, 142, 162},
	{46, 169, 65, 177, 1, 217, 158, 123, 107, 24, 166, 166, 42, 15, 87, 60, 75, 128, 208, 198, 140, 161, 209, 95, 136, 93, 233, 206, 11, 79, 196, 136},
	{130, 18, 196, 155, 9, 73, 105, 48, 145, 198, 103, 42, 6, 36, 31, 61, 248, 101, 166, 118, 204, 205, 203, 209, 111, 6, 21, 238, 166, 6, 131, 131},
	{231, 38, 228, 13, 189, 47, 152, 65, 41, 59, 91, 60, 21, 233, 24, 168, 114, 174, 210, 186, 73, 31, 78, 17, 30, 160, 145, 58, 4, 255, 225, 101},
	{119, 41, 71, 94, 26, 206, 150, 141, 9, 106, 124, 191, 11, 136, 52, 129, 88, 163, 126, 239, 100, 185, 9, 148, 203, 211, 125, 220, 58, 223, 83, 112},
	{252, 72, 126, 70, 108, 43, 244, 141, 135, 181, 214, 111, 90, 162, 79, 156, 91, 63, 153, 14, 33, 15, 80, 101, 5, 13, 82, 8, 213, 155, 107, 135},
	{119, 252, 136, 88, 44, 17, 79, 183, 126, 155, 198, 102, 110, 63, 63, 198, 232, 158, 76, 189, 61, 209, 89, 10, 111, 103, 173, 181, 71, 243, 72, 213},
	{220, 36, 54, 20, 255, 218, 121, 227, 215, 85, 110, 243, 253, 234, 11, 68, 223, 23, 87, 186, 218, 224, 23, 176, 95, 81, 51, 253, 21, 194, 122, 234},
	{231, 5, 9, 95, 227, 237, 108, 178, 218, 69, 134, 100, 34, 156, 81, 88, 216, 138, 160, 247, 117, 82, 141, 251, 39, 255, 63, 254, 39, 15, 192, 201},
	{119, 159, 30, 203, 250, 227, 208, 225, 231, 50, 136, 23, 68, 109, 191, 75, 203, 106, 103, 140, 108, 164, 226, 114, 110, 159, 62, 110, 101, 252, 218, 202},
	{94, 92, 87, 126, 64, 31, 176, 237, 60, 5, 25, 121, 32, 28, 252, 94, 193, 0, 218, 233, 148, 34, 120, 233, 158, 52, 52, 207, 229, 96, 194, 118},
	{24, 148, 206, 97, 36, 101, 156, 176, 214, 249, 10, 132, 66, 78, 55, 43, 89, 221, 126, 114, 32, 189, 177, 104, 4, 83, 155, 81, 186, 124, 110, 241},
	{100, 30, 223, 25, 101, 232, 122, 112, 66, 135, 18, 114, 126, 209, 219, 19, 178, 222, 145, 237, 87, 76, 131, 178, 131, 57, 188, 117, 2, 174, 195, 242},
	{10, 54, 220, 77, 27, 23, 10, 236, 101, 76, 214, 227, 136, 111, 106, 227, 162, 68, 141, 48, 169, 94, 58, 229, 197, 107, 122, 91, 0, 203, 143, 61},
	{87, 25, 104, 166, 165, 136, 87, 120, 75, 19, 225, 191, 148, 128, 40, 90, 122, 231, 13, 223, 46, 50, 29, 6, 177, 124, 240, 169, 51, 17, 156, 221},
	{229, 104, 19, 98, 142, 197, 126, 141, 224, 89, 188, 167, 135, 209, 239, 124, 177, 203, 159, 96, 140, 107, 160, 191, 120, 2, 119, 190, 221, 189, 63, 39},
	{94, 241, 229, 58, 5, 244, 61, 1, 71, 36, 71, 217, 20, 64, 148, 89, 206, 95, 43, 13, 36, 86, 97, 135, 239, 219, 239, 194, 88, 100, 218, 43},
	{173, 25, 183, 152, 36, 178, 168, 10, 139, 137, 20, 80, 210, 199, 252, 207, 158, 81, 194, 89, 121, 179, 125, 186, 142, 213, 151, 251, 147, 186, 157, 157},
	{24, 15, 237, 202, 6, 101, 108, 185, 16, 7, 112, 19, 173, 38, 121, 105, 80, 144, 38, 159, 173, 21, 137, 226, 144, 22, 47, 233, 14, 151, 212, 170},
	{115, 244, 114, 101, 14, 141, 160, 155, 207, 44, 140, 185, 152, 219, 82, 187, 39, 187, 165, 198, 51, 93, 72, 68, 50, 61, 104, 9, 76, 16, 40, 248},
	{255, 116, 102, 150, 149, 92, 250, 145, 174, 213, 231, 97, 2, 113, 13, 254, 46, 39, 130, 121, 191, 84, 61, 233, 195, 223, 138, 113, 52, 163, 59, 100},
	{46, 152, 174, 75, 110, 105, 185, 91, 113, 4, 142, 193, 25, 113, 152, 81, 192, 80, 82, 154, 32, 198, 122, 44, 64, 23, 35, 179, 77, 233, 116, 138},
	{59, 92, 57, 215, 39, 76, 42, 116, 212, 237, 45, 158, 132, 16, 234, 210, 243, 81, 223, 214, 21, 129, 72, 241, 177, 13, 151, 214, 13, 115, 146, 7},
	{66, 63, 103, 35, 17, 193, 164, 76, 229, 164, 71, 101, 246, 33, 126, 171, 247, 89, 222, 43, 216, 130, 182, 191, 15, 123, 9, 105, 78, 75, 128, 67},
	{163, 186, 13, 133, 4, 147, 149, 51, 87, 24, 136, 60, 116, 81, 205, 61, 16, 163, 38, 177, 215, 62, 34, 58, 165, 84, 158, 230, 140, 59, 234, 191},
	{211, 212, 125, 255, 170, 153, 160, 51, 200, 195, 187, 98, 220, 10, 67, 139, 113, 249, 27, 24, 163, 140, 135, 164, 159, 168, 132, 116, 221, 232, 121, 36},
	{65, 137, 210, 1, 238, 211, 244, 235, 203, 63, 213, 164, 199, 63, 40, 41, 35, 240, 16, 231, 83, 195, 197, 231, 227, 208, 25, 144, 55, 163, 176, 135},
	{20, 42, 3, 176, 157, 102, 203, 229, 180, 95, 234, 143, 162, 93, 106, 41, 112, 132, 174, 217, 224, 59, 238, 114, 134, 47, 157, 107, 61, 186, 187, 160},
	{76, 216, 234, 82, 172, 165, 36, 8, 251, 135, 176, 230, 151, 205, 36, 99, 45, 148, 166, 223, 117, 127, 163, 242, 156, 192, 165, 213, 73, 78, 220, 197},
	{175, 162, 199, 93, 194, 197, 168, 70, 44, 23, 9, 46, 225, 239, 226, 29, 233, 177, 15, 13, 192, 222, 160, 121, 79, 48, 69, 19, 225, 170, 51, 210},
}
