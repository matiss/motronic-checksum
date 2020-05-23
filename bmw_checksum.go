package motronic


// BMW173Checksum calculates checksum for ECU with code ending 173
func BMW173Checksum(buf []byte, patch bool) (uint16, uint16) {
	const (
		reg1Start = 0x0000
		reg1End   = 0x1EFF
		reg2Start = 0x2000
		reg2End   = 0x7FFF
		store     = 0x1F00
	)

	// Stored checksum
	checksumStored := uint16((uint16(buf[store]) << 8) | uint16(buf[store+1]))

	// Calculate new checksum for region 1
	sum := Checksum16bit(0xB51F, reg1Start, reg1End, buf)

	// Calculate new checksum for region 2
	sum = Checksum16bit(sum, reg2Start, reg2End, buf)

	// Patch buffer
	if patch {
		PatchBuffer(store, []byte{byte(sum >> 8), byte(sum)}, buf)
	}

	return sum, checksumStored
}

// BMW402C098Checksum calculates checksum for ECU with code ending 402 and chip ending with 098
func BMW402C098Checksum(buf []byte, patch bool) (uint16, uint16) {
	const (
		reg1Start = 0x6AFE
		reg1End   = 0x7B15
		reg1Store = 0x7B16
	)

	// Stored checksum
	checksumStored := uint16((uint16(buf[reg1Store]) << 8) | uint16(buf[reg1Store+1]))

	// Calculate new checksum
	sum := Checksum16bit(0, reg1Start, reg1End, buf)

	// Patch buffer
	if patch {
		PatchBuffer(reg1Store, []byte{byte(sum >> 8), byte(sum)}, buf)
	}

	return sum, checksumStored
}

// BMW402C599Checksum calculates checksum for ECU with code ending 402 and chip ending with 599
func BMW402C599Checksum(buf []byte, patch bool) (uint16, uint16) {
	const (
		reg1Start = 0x6AFE
		reg1End   = 0x7B31
		reg1Store = 0x7B32
	)

	// Stored checksum
	checksumStored := uint16((uint16(buf[reg1Store]) << 8) | uint16(buf[reg1Store+1]))

	// Calculate new checksum
	sum := Checksum16bit(0, reg1Start, reg1End, buf)

	// Patch buffer
	if patch {
		PatchBuffer(reg1Store, []byte{byte(sum >> 8), byte(sum)}, buf)
	}

	return sum, checksumStored
}

// BMW403C547Checksum calculates checksum for ECU with code ending 403 and chip ending with 547
func BMW403C547Checksum(buf []byte, patch bool) (uint16, uint16) {
	const (
		reg1Start = 0x69FE
		reg1End   = 0x797F
		reg1Store = 0x7980
	)

	// Stored checksum
	checksumStored := uint16((uint16(buf[reg1Store]) << 8) | uint16(buf[reg1Store+1]))

	// Calculate new checksum
	sum := Checksum16bit(0, reg1Start, reg1End, buf)

	// Patch buffer
	if patch {
		PatchBuffer(reg1Store, []byte{byte(sum >> 8), byte(sum)}, buf)
	}

	return sum, checksumStored
}

// BMW403C950Checksum calculates checksum for ECU with code ending 403 and chip ending with 950
func BMW403C950Checksum(buf []byte, patch bool) (uint16, uint16) {
	const (
		reg1Start = 0x69FE
		reg1End   = 0x799B
		reg1Store = 0x799C
	)

	// Stored checksum
	checksumStored := uint16((uint16(buf[reg1Store]) << 8) | uint16(buf[reg1Store+1]))

	// Calculate new checksum
	sum := Checksum16bit(0, reg1Start, reg1End, buf)

	// Patch buffer
	if patch {
		PatchBuffer(reg1Store, []byte{byte(sum >> 8), byte(sum)}, buf)
	}

	return sum, checksumStored
}

// BMW404C689Checksum calculates checksum for ECU with code ending 404 and chip ending 689
func BMW404C689Checksum(buf []byte, patch bool) (uint16, uint16) {
	const (
		reg1Start = 0xBFFE
		reg1End   = 0xD147
		reg1Store = 0xD148
	)

	// Stored checksum
	checksumStored := uint16((uint16(buf[reg1Store]) << 8) | uint16(buf[reg1Store+1]))

	// Calculate new checksum
	sum := Checksum16bit(0, reg1Start, reg1End, buf)

	// Patch buffer
	if patch {
		PatchBuffer(reg1Store, []byte{byte(sum >> 8), byte(sum)}, buf)
	}

	return sum, checksumStored
}

// BMW405C951Checksum calculates checksum for ECU with code ending 405 and chip ending 951
func BMW405C951Checksum(buf []byte, patch bool) (uint16, uint16) {
	const (
		reg1Start = 0x69FE
		reg1End   = 0x799B
		reg1Store = 0x799C
	)

	// Stored checksum
	checksumStored := uint16((uint16(buf[reg1Store]) << 8) | uint16(buf[reg1Store+1]))

	// Calculate new checksum
	sum := Checksum16bit(0, reg1Start, reg1End, buf)

	// Patch buffer
	if patch {
		PatchBuffer(reg1Store, []byte{byte(sum >> 8), byte(sum)}, buf)
	}

	return sum, checksumStored
}

// BMW413C609Checksum calculates checksum for ECU with code ending 413 and chip ending 609
func BMW413C609Checksum(buf []byte, patch bool) (uint16, uint16) {
	const (
		reg1Start = 0xE000
		reg1End   = 0xF7A5
		reg1Store = 0xF7A6
	)

	// Stored checksum
	checksumStored := uint16((uint16(buf[reg1Store]) << 8) | uint16(buf[reg1Store+1]))

	// Old checksum
	sum := Checksum16bit(0, reg1Start, reg1End, buf)

	// Patch buffer
	if patch {
		PatchBuffer(reg1Store, []byte{byte(sum >> 8), byte(sum)}, buf)
	}

	return sum, checksumStored
}

// BMW413C623Checksum calculates checksum for ECU with code ending 413 and chip ending 623
func BMW413C623Checksum(buf []byte, patch bool) (uint16, uint16) {
	const (
		reg1Start = 0xD000
		reg1End   = 0xE7AF
		reg1Store = 0xE7B0
	)

	// Stored checksum
	checksumStored := uint16((uint16(buf[reg1Store]) << 8) | uint16(buf[reg1Store+1]))

	// Old checksum
	sum := Checksum16bit(0, reg1Start, reg1End, buf)

	// Patch buffer
	if patch {
		PatchBuffer(reg1Store, []byte{byte(sum >> 8), byte(sum)}, buf)
	}

	return sum, checksumStored
}

// BMW413C715Checksum calculates checksum for ECU with code ending 413 and chip ending 715
func BMW413C715Checksum(buf []byte, patch bool) (uint16, uint16) {
	const (
		reg1Start = 0xD000
		reg1End   = 0xE7C1
		reg1Store = 0xE7C2
	)

	// Stored checksum
	checksumStored := uint16((uint16(buf[reg1Store]) << 8) | uint16(buf[reg1Store+1]))

	// Calculate new checksum
	sum := Checksum16bit(0, reg1Start, reg1End, buf)

	// Patch buffer
	if patch {
		PatchBuffer(reg1Store, []byte{byte(sum >> 8), byte(sum)}, buf)
	}

	return sum, checksumStored
}

// BMW484C582Checksum calculates checksum for ECU with code ending 484 and chip ending 582
func BMW484C582Checksum(buf []byte, patch bool) (uint16, uint16) {
	const (
		reg1Start = 0xDBFE
		reg1End   = 0xED5F
		reg1Store = 0xED60
	)

	// Stored checksum
	checksumStored := uint16((uint16(buf[reg1Store]) << 8) | uint16(buf[reg1Store+1]))

	// Calculate new checksum
	sum := Checksum16bit(0, reg1Start, reg1End, buf)

	// Patch buffer
	if patch {
		PatchBuffer(reg1Store, []byte{byte(sum >> 8), byte(sum)}, buf)
	}

	return sum, checksumStored
}
