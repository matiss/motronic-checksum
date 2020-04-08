package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

const infoText = `
Motronic Checksum Correction Tool
---------------------------------

All bugs and issue should be reported at https://github.com/matiss/motronic_checksum/issues
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

	if err = BMW402C098Validate(buf); err == nil {
		color.Blue("BMW DME\nHW: 402\nChip: 098")
		found = true
		BMW402C098Checksum(buf)
	}

	if err = BMW402C599Validate(buf); !found && err == nil {
		color.Blue("BMW DME\nHW: 402\nChip: 599")
		found = true
		BMW402C599Checksum(buf)
	}

	if err = BMW403Validate(buf); !found && err == nil {
		color.Blue("BMW DME\nHW: 403\nChip: TBC")
		found = true
		BMW403Checksum(buf)
	}

	if err = BMW403C950Validate(buf); !found && err == nil {
		color.Blue("BMW DME\nHW:403\nChip: 950")
		found = true
	}

	if err = BMW405C951Validate(buf); !found && err == nil {
		color.Blue("BMW DME\nHW: 405\nChip: 951")
		found = true
		BMW405C951Checksum(buf)
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
	time.Sleep(2 * time.Second)
}
