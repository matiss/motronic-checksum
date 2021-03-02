package motronic

// Checksum16bit
func Checksum16bit(initial uint16, start int, end int, buf []byte) uint16 {
	var sum uint16 = initial

	for i := start; i <= end; i++ {
		sum += uint16(buf[i])
	}

	return sum
}

// PatchBuffer
func PatchBuffer(start int, patch []byte, buf []byte) {
	for i := 0; i < len(patch); i++ {
		buf[start+i] = patch[i]
	}
}
