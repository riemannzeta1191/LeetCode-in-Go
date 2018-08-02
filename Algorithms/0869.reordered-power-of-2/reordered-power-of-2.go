package problem0869

// 对 [1, 1e9] 中所有的 2^n 进行编码
// 把 key 写成 16 进制， 对应关系更明显
var isValid = map[int]bool{
	//9876543210
	0x0000000010: true, // 2^0 = 1
	//9876543210
	0x0000000100: true, // 2^1 = 2
	//9876543210
	0x0000010000: true, // 2^2 = 4
	//9876543210
	0x0100000000: true, // 2^3 = 8
	//9876543210
	0x0001000010: true, // 2^4 = 16
	//9876543210
	0x0000001100: true, // 2^5 = 32
	//9876543210
	0x0001010000: true, // 2^6 = 64
	//9876543210
	0x0100000110: true, // 2^7 = 128
	//9876543210
	0x0001100100: true, // 2^8 = 256
	//9876543210
	0x0000100110: true, // 2^9 = 512
	//9876543210
	0x0000010111: true, // 2^10 = 1024
	//9876543210
	0x0100010101: true, // 2^11 = 2048
	//9876543210
	0x1001010001: true, // 2^12 = 4096
	//9876543210
	0x1100000110: true, // 2^13 = 8192
	//9876543210
	0x0101011010: true, // 2^14 = 16384
	//9876543210
	0x0111001100: true, // 2^15 = 32768
	//9876543210
	0x0002201000: true, // 2^16 = 65536
	//9876543210
	0x0010001121: true, // 2^17 = 131072
	//9876543210
	0x0001020210: true, // 2^18 = 262144
	//9876543210
	0x0200110200: true, // 2^19 = 524288
	//9876543210
	0x0111110011: true, // 2^20 = 1048576
	//9876543210
	0x1010100211: true, // 2^21 = 2097152
	//9876543210
	0x1000031011: true, // 2^22 = 4194304
	//9876543210
	0x0401001001: true, // 2^23 = 8388608
	//9876543210
	0x0032000120: true, // 2^24 = 16777216
	//9876543210
	0x0000223100: true, // 2^25 = 33554432
	//9876543210
	0x0212010011: true, // 2^26 = 67108864
	//9876543210
	0x0120011220: true, // 2^27 = 134217728
	//9876543210
	0x0102221100: true, // 2^28 = 268435456
	//9876543210
	0x1111101111: true, // 2^29 = 536870912
}

func reorderedPowerOf2(N int) bool {
	return isValid[encode(N)]
}

// 由于 n 的范围是 [1,1e9]
// 每个数字出现的次数不会超过 10 次，可以用 4 bit 的二进制数记录
// go 的 int 类型是 64 bit，足够记录 10 个数字的出现次数
// 所以， res 的 [4*digit, 4*digit+4) 位代表的二进制数，表示 digit 在 n 中出现的次数
func encode(N int) int {
	res := 0
	n := uint(N)
	for n > 0 {
		digit := n % 10
		res += 1 << (4 * digit)
		n /= 10
	}
	return res
}
