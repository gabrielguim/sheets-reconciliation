package infra

import (
	"bufio"
	"log"
	"os"
)

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Printf("failed to close file: %v\n", err)
	}
}

func ReadFileStream(fileName string, rowProcessor func(row string)) bool {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer closeFile(file)

	scanner := bufio.NewScanner(file)

	// jump first line
	scanner.Scan()
	for scanner.Scan() {
		rowProcessor(scanner.Text())
	}

	return true
}
