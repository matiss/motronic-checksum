package motronic

import (
	"io/ioutil"
	"os"
	"testing"
)

// Note: All tests should be performed against original binary files

func TestBMW402C098Checksum(t *testing.T) {
	file, err := os.Open("./firmwares/402_ori.bin")
	defer file.Close()
	if err != nil {
		t.Error(err)
		return
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error(err)
		return
	}

	if err = BMW402C098Validate(buf); err != nil {
		t.Errorf("Invalid test binary file")
		return
	}

	sum, _ := BMW402C098Checksum(buf, false)
	if sum != 0xF7E7 {
		t.Errorf("Checksum failed, expected: %X got: %X", 0xF7E7, sum)
	}
}

func TestBMW402C599Checksum(t *testing.T) {
	file, err := os.Open("./firmwares/402SW599_ori.bin")
	defer file.Close()
	if err != nil {
		t.Error(err)
		return
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error(err)
		return
	}

	if err = BMW402C599Validate(buf); err != nil {
		t.Errorf("Invalid test binary file")
		return
	}

	sum, _ := BMW402C599Checksum(buf, false)
	if sum != 0x0985 {
		t.Errorf("Checksum failed, expected: %X got: %X", 0x0985, sum)
	}
}

func TestBMW403Checksum(t *testing.T) {
	file, err := os.Open("./firmwares/403_ori.bin")
	defer file.Close()
	if err != nil {
		t.Error(err)
		return
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error(err)
		return
	}

	if err = BMW403Validate(buf); err != nil {
		t.Errorf("Invalid test binary file")
		return
	}

	sum, _ := BMW403Checksum(buf, false)
	if sum != 0xBD68 {
		t.Errorf("Checksum failed, expected: %X got: %X", 0xBD68, sum)
	}
}

func TestBMW403C950Checksum(t *testing.T) {
	file, err := os.Open("./firmwares/403C950_ori.bin")
	defer file.Close()
	if err != nil {
		t.Error(err)
		return
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error(err)
		return
	}

	if err = BMW403C950Validate(buf); err != nil {
		t.Errorf("Invalid test binary file")
		return
	}

	sum, _ := BMW403C950Checksum(buf, false)
	if sum != 0xC7A9 {
		t.Errorf("Checksum failed, expected: %X got: %X", 0xC7A9, sum)
	}
}

func TestBMW405C951Checksum(t *testing.T) {
	file, err := os.Open("./firmwares/405C951_ori.bin")
	defer file.Close()
	if err != nil {
		t.Error(err)
		return
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error(err)
		return
	}

	if err = BMW405C951Validate(buf); err != nil {
		t.Errorf("Invalid test binary file")
		return
	}

	sum, _ := BMW405C951Checksum(buf, false)
	if sum != 0xCFAF {
		t.Errorf("Checksum failed, expected: %X got: %X", 0xCFAF, sum)
	}
}
