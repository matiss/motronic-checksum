package motronic

import (
	"fmt"
)

// BMW173Validate validates ECU with code ending 173
func BMW173Validate(buf []byte) error {
	// Validate length
	if len(buf) != 0x8000 {
		return fmt.Errorf("Invalid file length")
	}

	const (
		reg1Start = 0x0000
		reg1End   = 0x1EFF
		reg2Start = 0x2000
		reg2End   = 0x7FFF
	)

	// Validate Data Region 1
	if buf[reg1End] != 0x01 || buf[reg1End-1] != 0x23 || buf[reg1End-2] != 0x9F || buf[reg1End-3] != 0x03 {
		return fmt.Errorf("Invalid end of Data Region 1")
	}

	// Validate Data Region 2
	if buf[reg2Start] != 0x02 || buf[reg2Start+1] != 0x25 || buf[reg2Start+2] != 0x52 || buf[reg2Start+3] != 0xFF {
		return fmt.Errorf("Invalid start of Data Region 2")
	}

	return nil
}

// BMW402Validate validates ECU with code ending 402
func BMW402C098Validate(buf []byte) error {
	// Validate length
	if len(buf) != 0x8000 {
		return fmt.Errorf("Invalid file length")
	}

	const (
		reg1Start = 0x6AFE
		reg1End   = 0x7B15
	)

	// Validate Data Region 1
	if buf[reg1Start] != 0x00 || buf[reg1Start+1] != 0x00 || buf[reg1Start+2] != 0x00 {
		return fmt.Errorf("Invalid start of Data Region 1")
	}

	if buf[reg1End] != 0x4F || buf[reg1End-1] != 0x13 || buf[reg1End-2] != 0x22 || buf[reg1End-3] != 0x40 {
		return fmt.Errorf("Invalid end of Data Region 1")
	}

	return nil
}

// BMW402SW599Validate validates ECU with code ending 402 and chip number ending with 599
func BMW402C599Validate(buf []byte) error {
	// Validate length
	if len(buf) != 0x8000 {
		return fmt.Errorf("Invalid file length")
	}

	const (
		reg1Start = 0x6AFE
		reg1End   = 0x7B31
		// reg1Store = 0x7B32
	)

	// Validate Data Region 1
	if buf[reg1Start] != 0x00 || buf[reg1Start+1] != 0x00 || buf[reg1Start+2] != 0x00 {
		return fmt.Errorf("Invalid start of Data Region 1")
	}

	if buf[reg1End] != 0x44 || buf[reg1End-1] != 0xFF || buf[reg1End-2] != 0xFF || buf[reg1End-3] != 0xFF {
		return fmt.Errorf("Invalid end of Data Region 1")
	}

	return nil
}

// BMW403C547Validate validates ECU with code ending 403 and chip number ending with C547
func BMW403C547Validate(buf []byte) error {
	// Validate length
	if len(buf) != 0x8000 {
		return fmt.Errorf("Invalid file length")
	}

	const (
		reg1Start = 0x69FE
		reg1End   = 0x797F
	)

	// Validate Data Region 1
	if buf[reg1Start] != 0x00 || buf[reg1Start+1] != 0x00 || buf[reg1Start+2] != 0x00 {
		return fmt.Errorf("Invalid start of Data Region 1")
	}

	if buf[reg1End] != 0x4F || buf[reg1End-1] != 0x13 || buf[reg1End-2] != 0x22 || buf[reg1End-3] != 0x40 {
		return fmt.Errorf("Invalid end of Data Region 1")
	}

	return nil
}

// BMW403C950Validate validates ECU with code ending 403 and chip code ending 950
func BMW403C950Validate(buf []byte) error {
	// Validate length
	if len(buf) != 0x8000 {
		return fmt.Errorf("Invalid file length")
	}

	const (
		reg1Start = 0x69FE
		reg1End   = 0x799B
	)

	// Validate Data Region 1
	if buf[reg1Start] != 0x00 || buf[reg1Start+1] != 0x00 || buf[reg1Start+2] != 0x00 {
		return fmt.Errorf("Invalid start of Data Region 1")
	}

	if buf[reg1End] != 0x4F || buf[reg1End-1] != 0x13 || buf[reg1End-2] != 0x22 || buf[reg1End-3] != 0x40 {
		return fmt.Errorf("Invalid end of Data Region 1")
	}

	return nil
}

// BMW404C689Validate validates ECU with code ending 404 and chip ending 689
func BMW404C689Validate(buf []byte) error {
	// Validate length
	if len(buf) != 0x10000 {
		return fmt.Errorf("Invalid file length")
	}

	const (
		reg1Start = 0xBFFE
		reg1End   = 0xD147
	)

	// Validate Data Region 1
	if buf[reg1Start] != 0x00 || buf[reg1Start+1] != 0x00 || buf[reg1Start+2] != 0x00 {
		return fmt.Errorf("Invalid start of Data Region 1")
	}

	if buf[reg1End] != 0x44 || buf[reg1End-1] != 0xFF || buf[reg1End-2] != 0xFF || buf[reg1End-3] != 0xFF {
		return fmt.Errorf("Invalid end of Data Region 1")
	}

	return nil
}

// BMW405C951Validate validates ECU with code ending 405 and chip ending 951
func BMW405C951Validate(buf []byte) error {
	// Validate length
	if len(buf) != 0x10000 {
		return fmt.Errorf("Invalid file length")
	}

	const (
		reg1Start = 0x69FE
		reg1End   = 0x799B
	)

	// Validate Data Region 1
	if buf[reg1Start] != 0x00 || buf[reg1Start+1] != 0x00 || buf[reg1Start+2] != 0x00 {
		return fmt.Errorf("Invalid start of Data Region 1")
	}

	if buf[reg1End] != 0x48 || buf[reg1End-1] != 0x53 || buf[reg1End-2] != 0xA1 || buf[reg1End-3] != 0x45 {
		return fmt.Errorf("Invalid end of Data Region 1")
	}

	return nil
}

// BMW413C609Validate validates ECU with code ending 413 and chip ending 609
func BMW413C609Validate(buf []byte) error {
	// Validate length
	if len(buf) != 0x10000 {
		return fmt.Errorf("Invalid file length")
	}

	const (
		reg1Start = 0xE000
		reg1End   = 0xF7A5
	)

	// Validate Data Region 1
	if buf[reg1Start] != 0x00 || buf[reg1Start+1] != 0x00 || buf[reg1Start+2] != 0x02 {
		return fmt.Errorf("Invalid start of Data Region 1")
	}

	if buf[reg1End] != 0x44 || buf[reg1End-1] != 0xFF || buf[reg1End-2] != 0xFF || buf[reg1End-3] != 0xFF {
		return fmt.Errorf("Invalid end of Data Region 1")
	}

	return nil
}

// BMW413C623Validate validates ECU with code ending 413 and chip ending 623
func BMW413C623Validate(buf []byte) error {
	// Validate length
	if len(buf) != 0x10000 {
		return fmt.Errorf("Invalid file length")
	}

	const (
		reg1Start = 0xD000
		reg1End   = 0xE7AF
	)

	// Validate Data Region 1
	if buf[reg1Start] != 0x00 || buf[reg1Start+1] != 0x00 || buf[reg1Start+2] != 0x02 {
		return fmt.Errorf("Invalid start of Data Region 1")
	}

	if buf[reg1End] != 0x44 || buf[reg1End-1] != 0xFF || buf[reg1End-2] != 0xFF || buf[reg1End-3] != 0xFF {
		return fmt.Errorf("Invalid end of Data Region 1")
	}

	return nil
}

// BMW413C715Validate validates ECU with code ending 413 and chip ending 715
func BMW413C715Validate(buf []byte) error {
	// Validate length
	if len(buf) != 0x10000 {
		return fmt.Errorf("Invalid file length")
	}

	const (
		reg1Start = 0xD000
		reg1End   = 0xE7C1
	)

	// Validate Data Region 1
	if buf[reg1Start] != 0x00 || buf[reg1Start+1] != 0x00 || buf[reg1Start+2] != 0x02 {
		return fmt.Errorf("Invalid start of Data Region 1")
	}

	if buf[reg1End] != 0x44 || buf[reg1End-1] != 0xFF || buf[reg1End-2] != 0xFF || buf[reg1End-3] != 0xFF {
		return fmt.Errorf("Invalid end of Data Region 1")
	}

	return nil
}

// BMW484C582Validate validates ECU with code ending 484 and chip ending 582
func BMW484C582Validate(buf []byte) error {
	// Validate length
	if len(buf) != 0x10000 {
		return fmt.Errorf("Invalid file length")
	}

	const (
		reg1Start = 0xDBFE
		reg1End   = 0xED5F
	)

	// Validate Data Region 1
	if buf[reg1Start] != 0x00 || buf[reg1Start+1] != 0x00 || buf[reg1Start+2] != 0x00 {
		return fmt.Errorf("Invalid start of Data Region 1")
	}

	if buf[reg1End] != 0xEC || buf[reg1End-1] != 0x49 || buf[reg1End-2] != 0xEC || buf[reg1End-3] != 0x37 {
		return fmt.Errorf("Invalid end of Data Region 1")
	}

	return nil
}

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
	sum1 := Checksum16bit(0, reg1Start, reg1End, buf)

	// Calculate new checksum for region 2
	sum2 := Checksum16bit(0, reg2Start, reg2End, buf)

	// Calculate final checksum
	sum := sum1 + sum2 + 0xB51F

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
