package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "slim",
		Usage: "Slim down vtt files",
		Action: func(cCtx *cli.Context) error {
			return doIt(cCtx)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func doIt(cCtx *cli.Context) error {
	// open the file at the path given
	path := cCtx.Args().First()
	in, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not open file: %s", path)
	}
	defer func() {
		if err := in.Close(); err != nil {
			panic(err)
		}
	}()

	// Create a new scanner for the file
	scanner := bufio.NewScanner(in)
	// currentSpeaker := ""

	// Read each line of the file
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return nil
}
