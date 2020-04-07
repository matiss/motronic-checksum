package main

// SimpleSum16bit
func SimpleSum16bit(initial uint32, start int, end int, buf []byte) uint16 {
	var sum uint32 = initial

	for i := start; i <= end; i += 1 {
		sum += uint32(buf[i])
	}

	sum = (sum & 0xFFFF)

	return uint16(sum)
}

// PatchBuffer
func PatchBuffer(start int, patch []byte, buf []byte) {
	for i := 0; i < len(patch); i++ {
		buf[start+i] = patch[i]
	}
}
