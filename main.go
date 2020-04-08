package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Prompt{
		Label:     "Save to file",
		IsConfirm: true,
	}

	args := os.Args[1:]

	if len(args) == 0 {
		color.Red("Please input path to file")
		return
	}

	filePath := args[0]

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

	if err = BMW405Validate(buf); !found && err == nil {
		color.Blue("BMW DME\nHW: 405")
		found = true
		// WIP
	}

	if !found {
		color.Red("Unsupported firmware file")
		return
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
}
