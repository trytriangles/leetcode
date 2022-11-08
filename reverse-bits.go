func reverseBits(num uint32) uint32 {
	var result uint32 = 0
	var shifter uint32 = 31
	for num != 0 {
		result += (num & 1) << shifter
		num = num >> 1
		shifter -= 1
	}
	return ret
}
