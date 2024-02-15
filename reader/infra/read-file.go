package infra

import (
	"bufio"
	"errors"
	"log"
	"os"
	"sheets-reconciliation/commons"
	"strconv"
	"strings"
)

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Println("failed to close file: %v", err)
	}
}

func getEntry(row, delimiter string) (commons.Entry, error) {
	values := strings.Split(row, delimiter)

	if len(values) < 2 {
		return commons.Entry{}, errors.New("invalid row")
	}

	rowValue, err := strconv.ParseUint(values[1], 0, 32)
	if err != nil {
		return commons.Entry{}, err
	}

	return commons.Entry{
		Name:  values[0],
		Value: uint32(rowValue),
	}, nil
}

func ReadFileStream(fileName string, delimiter string, rowProcessor func(entry commons.Entry)) bool {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer closeFile(file)

	scanner := bufio.NewScanner(file)

	// jump first line
	scanner.Scan()
	for scanner.Scan() {
		row := scanner.Text()
		entry, err := getEntry(row, delimiter)

		if err != nil {
			log.Fatal(err)
		}

		rowProcessor(entry)
	}

	return true
}
