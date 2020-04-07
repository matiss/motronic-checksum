package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

const (
	regionStart  = 0x69FE
	regionEnd    = 0x797F
	storeAddress = 0x7980
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		color.Red("Please input path to file")
		return
	}

	filePath := args[0]

	file, err := os.Open(filePath)
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
		color.Blue("BMW DMW\nHW: 403\nChip: TBC")
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
		color.Red("Unsupported binary file")
		return
	}

	// Write output file
	err = ioutil.WriteFile("corrected.bin", buf, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	color.Green("Done!\n")
}
