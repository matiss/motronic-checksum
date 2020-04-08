package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/fatih/color"
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

	if err = motronic.BMW402C098Validate(buf); err == nil {
		found = true
		color.Blue("BMW DME\nHW: 402\nChip: 098")

		csNew, csOld := motronic.BMW402C098Checksum(buf, true)

		fmt.Printf("Checksum old: %X new: %X\n", csOld, csNew)
	}

	if err = motronic.BMW402C599Validate(buf); !found && err == nil {
		found = true
		color.Blue("BMW DME\nHW: 402\nChip: 599")

		csNew, csOld := motronic.BMW402C599Checksum(buf, true)

		fmt.Printf("Checksum old: %X new: %X\n", csOld, csNew)
	}

	if err = motronic.BMW403Validate(buf); !found && err == nil {
		found = true
		color.Blue("BMW DME\nHW: 403\nChip: TBC")

		csNew, csOld := motronic.BMW403Checksum(buf, true)

		fmt.Printf("Checksum old: %X new: %X\n", csOld, csNew)
	}

	if err = motronic.BMW403C950Validate(buf); !found && err == nil {
		found = true
		color.Blue("BMW DME\nHW:403\nChip: 950")

		csNew, csOld := motronic.BMW403C950Checksum(buf, true)

		fmt.Printf("Checksum old: %X new: %X\n", csOld, csNew)
	}

	if err = motronic.BMW405C951Validate(buf); !found && err == nil {
		found = true
		color.Blue("BMW DME\nHW: 405\nChip: 951")

		csNew, csOld := motronic.BMW405C951Checksum(buf, true)

		fmt.Printf("Checksum old: %X new: %X\n", csOld, csNew)
	}

	if !found {
		color.Red("Unsupported firmware file")

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

	color.Green("Done!\n")

	// Delay exiting
	time.Sleep(1 * time.Second)
}
