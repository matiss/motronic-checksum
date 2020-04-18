package motronic

import (
	"io/ioutil"
	"os"
	"testing"
)

// Note: All tests should be performed against original binary files

func TestBMW173Checksum(t *testing.T) {
	file, err := os.Open("./firmwares/BMW/173_ori.bin")
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

	if err = BMW173Validate(buf); err != nil {
		t.Errorf("Invalid test binary file")
		return
	}

	sum, _ := BMW173Checksum(buf, false)
	if sum != 0x40CE {
		t.Errorf("Checksum failed, expected: %X got: %X", 0x40CE, sum)
	}
}

func TestBMW402C098Checksum(t *testing.T) {
	file, err := os.Open("./firmwares/BMW/402_ori.bin")
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
	file, err := os.Open("./firmwares/BMW/402SW599_ori.bin")
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

func TestBMW403C547Checksum(t *testing.T) {
	file, err := os.Open("./firmwares/BMW/403C547_ori.bin")
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

	if err = BMW403C547Validate(buf); err != nil {
		t.Errorf("Invalid test binary file")
		return
	}

	sum, _ := BMW403C547Checksum(buf, false)
	if sum != 0xBD68 {
		t.Errorf("Checksum failed, expected: %X got: %X", 0xBD68, sum)
	}
}

func TestBMW403C950Checksum(t *testing.T) {
	file, err := os.Open("./firmwares/BMW/403C950_ori.bin")
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

func TestBMW404C689Checksum(t *testing.T) {
	file, err := os.Open("./firmwares/BMW/404C689_ori.bin")
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

	if err = BMW404C689Validate(buf); err != nil {
		t.Errorf("Invalid test binary file")
		return
	}

	sum, _ := BMW404C689Checksum(buf, false)
	if sum != 0x86B9 {
		t.Errorf("Checksum failed, expected: %X got: %X", 0x86B9, sum)
	}
}

func TestBMW405C951Checksum(t *testing.T) {
	file, err := os.Open("./firmwares/BMW/405C951_ori.bin")
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

func TestBMW413C609Checksum(t *testing.T) {
	file, err := os.Open("./firmwares/BMW/413C609_ori.bin")
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

	if err = BMW413C609Validate(buf); err != nil {
		t.Errorf("Invalid test binary file")
		return
	}

	sum, _ := BMW413C609Checksum(buf, false)
	if sum != 0xE53A {
		t.Errorf("Checksum failed, expected: %X got: %X", 0xE53A, sum)
	}
}

func TestBMW413C623Checksum(t *testing.T) {
	file, err := os.Open("./firmwares/BMW/413C623_ori.bin")
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

	if err = BMW413C623Validate(buf); err != nil {
		t.Errorf("Invalid test binary file")
		return
	}

	sum, _ := BMW413C623Checksum(buf, false)
	if sum != 0xDFC3 {
		t.Errorf("Checksum failed, expected: %X got: %X", 0xDFC3, sum)
	}
}

func TestBMW413C715Checksum(t *testing.T) {
	file, err := os.Open("./firmwares/BMW/413C715_ori.bin")
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

	if err = BMW413C715Validate(buf); err != nil {
		t.Errorf("Invalid test binary file")
		return
	}

	sum, _ := BMW413C715Checksum(buf, false)
	if sum != 0xE76E {
		t.Errorf("Checksum failed, expected: %X got: %X", 0xE76E, sum)
	}
}
