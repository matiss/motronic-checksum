package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/matiss/motronic-checksum/v1"
)

const infoText = `
Motronic Checksum Correction Tool
---------------------------------

All bugs and issues should be reported at https://github.com/matiss/motronic-checksum/issues

Created by Matiss Kiris

---------------------------------

`

func main() {
	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("Filename must have more than 3 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "File",
		Validate: validate,
		Default:  "",
	}

	fmt.Printf(infoText)

	filePath, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)

		// Delay exiting
		time.Sleep(2 * time.Second)
		return
	}

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	found := false
	corrNeeded := false
	romSize := len(buf)

	if romSize == 0x8000 {
		// ECU with ROM size 0x8000

		if err = motronic.BMW173Validate(buf); err == nil {
			found = true
			fmt.Printf("BMW DME\nHW: 0 261 200 173\nSW: 794\n")

			csNew, csOld := motronic.BMW173Checksum(buf, true)
			corrNeeded = (csNew != csOld)

			fmt.Printf("Checksum: %X -> %X\n", csOld, csNew)
		} else if err = motronic.BMW402C098Validate(buf); err == nil {
			found = true
			fmt.Printf("BMW DME\nHW: 0 261 200 402\nSW: 098\n")

			csNew, csOld := motronic.BMW402C098Checksum(buf, true)
			corrNeeded = (csNew != csOld)

			fmt.Printf("Checksum: %X -> %X\n", csOld, csNew)
		} else if err = motronic.BMW402C599Validate(buf); err == nil {
			found = true
			fmt.Printf("BMW DME\nHW: 0 261 200 402\nSW: 599\n")

			csNew, csOld := motronic.BMW402C599Checksum(buf, true)
			corrNeeded = (csNew != csOld)

			fmt.Printf("Checksum: %X -> %X\n", csOld, csNew)
		} else if err = motronic.BMW403C547Validate(buf); err == nil {
			found = true
			fmt.Printf("BMW DME\nHW: 0 261 200 403\nSW: 547\n")

			csNew, csOld := motronic.BMW403C547Checksum(buf, true)
			corrNeeded = (csNew != csOld)

			fmt.Printf("Checksum: %X -> %X\n", csOld, csNew)
		} else if err = motronic.BMW403C950Validate(buf); err == nil {
			found = true
			fmt.Printf("BMW DME\nHW: 0 261 200 403\nSW: 950\n")

			csNew, csOld := motronic.BMW403C950Checksum(buf, true)
			corrNeeded = (csNew != csOld)

			fmt.Printf("Checksum: %X -> %X\n", csOld, csNew)
		} else if err = motronic.BMW590C597Validate(buf); err == nil {
			found = true
			fmt.Printf("BMW DME\nHW: 0 261 203 590\nSW: 597\n")

			csNew, csOld := motronic.BMW173Checksum(buf, true)
			corrNeeded = (csNew != csOld)

			fmt.Printf("Checksum: %X -> %X\n", csOld, csNew)
		}
	} else if romSize == 0x10000 {
		// ECU with ROM size 0x10000

		if err = motronic.BMW404C689Validate(buf); err == nil {
			found = true
			fmt.Printf("BMW DME\nHW: 0 261 200 404\nSW: 689\n")

			csNew, csOld := motronic.BMW404C689Checksum(buf, true)
			corrNeeded = (csNew != csOld)

			fmt.Printf("Checksum: %X -> %X\n", csOld, csNew)
		} else if err = motronic.BMW405C951Validate(buf); err == nil {
			found = true
			fmt.Printf("BMW DME\nHW: 0 261 200 405\nSW: 951\n")

			csNew, csOld := motronic.BMW405C951Checksum(buf, true)
			corrNeeded = (csNew != csOld)

			fmt.Printf("Checksum: %X -> %X\n", csOld, csNew)
		} else if err = motronic.BMW413C609Validate(buf); err == nil {
			found = true
			fmt.Printf("BMW DME\nHW: 0 261 200 413\nSW: 609\n")

			csNew, csOld := motronic.BMW413C609Checksum(buf, true)
			corrNeeded = (csNew != csOld)

			fmt.Printf("Checksum: %X -> %X\n", csOld, csNew)
		} else if err = motronic.BMW413C623Validate(buf); err == nil {
			found = true
			fmt.Printf("BMW DME\nHW: 0 261 200 413\nSW: 623\n")

			csNew, csOld := motronic.BMW413C623Checksum(buf, true)
			corrNeeded = (csNew != csOld)

			fmt.Printf("Checksum: %X -> %X\n", csOld, csNew)
		} else if err = motronic.BMW413C715Validate(buf); err == nil {
			found = true
			fmt.Printf("BMW DME\nHW: 0 261 200 413\nSW: 715\n")

			csNew, csOld := motronic.BMW413C715Checksum(buf, true)
			corrNeeded = (csNew != csOld)

			fmt.Printf("Checksum: %X -> %X\n", csOld, csNew)
		} else if err = motronic.BMW484C582Validate(buf); err == nil {
			found = true
			fmt.Printf("BMW DME\nHW: 0 261 203 484\nSW: 582\n")

			csNew, csOld := motronic.BMW484C582Checksum(buf, true)
			corrNeeded = (csNew != csOld)

			fmt.Printf("Checksum: %X -> %X\n", csOld, csNew)
		}

	} else {
		fmt.Printf("Invalid file size\n")
	}

	// Check if firmware is detected correctly
	if !found {
		fmt.Printf("Unsupported firmware file\n")

		// Delay exiting
		time.Sleep(2 * time.Second)
		return
	}

	// Check if correction is needed
	if !corrNeeded {
		fmt.Printf("No checksum correction needed!\n")

		// Delay exiting
		time.Sleep(2 * time.Second)
		return
	}

	prompt = promptui.Prompt{
		Label:     "Save to file",
		IsConfirm: true,
	}

	_, err = prompt.Run()
	if err != nil {
		return
	}

	// Overwrite file
	_, err = file.WriteAt(buf, 0) // Write at 0 beginning
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Done!\n")

	// Delay exiting
	time.Sleep(1 * time.Second)
}
