package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s fastq\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}
	defer f.Close()

	//####################################

	fileScanner := bufio.NewScanner(f)

	output := ""
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line[0:1] != "@" {
			output = output + line + "\n"
		}
	}
	fmt.Print(output)
	//####################################
}
