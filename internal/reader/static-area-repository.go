package reader

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sheets-reconciliation/internal/commons"
	"sheets-reconciliation/internal/reader/infra"
	"strconv"
	"strings"
	"sync"
)

type StaticAreaRepository struct{}

func getFileName(area, kind string) string {
	staticFilePath, _ := os.Getwd()
	return fmt.Sprintf("%s/static/%s_%s.csv", staticFilePath, area, kind)
}

func parseRow(row string) (commons.Entry, error) {
	values := strings.Split(row, ",")

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

func getEntries(waitGroup *sync.WaitGroup, entries *[]commons.Entry, area, kind string) {
	infra.ReadFileStream(getFileName(area, kind), func(row string) {
		entry, err := parseRow(row)

		if err != nil {
			log.Fatal(err)
		}

		*entries = append(*entries, entry)
	})

	waitGroup.Done()
}

func (staticDataRepository StaticAreaRepository) Read(areaName string) commons.Area {
	payments := new([]commons.Entry)
	receipts := new([]commons.Entry)

	fileWaitGroup := new(sync.WaitGroup)
	fileWaitGroup.Add(2)

	go getEntries(fileWaitGroup, payments, areaName, "payments")
	go getEntries(fileWaitGroup, receipts, areaName, "receipts")

	fileWaitGroup.Wait()

	return commons.Area{
		Name:     areaName,
		Receipts: *receipts,
		Payments: *payments,
	}
}
