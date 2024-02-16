package reader

import (
	"fmt"
	"os"
	"sheets-reconciliation/internal/commons"
	"sheets-reconciliation/internal/reader/infra"
	"sync"
)

type StaticAreaRepository struct{}

func getFileName(area, kind string) string {
	staticFilePath, _ := os.Getwd()
	return fmt.Sprintf("%s/static/%s_%s.csv", staticFilePath, area, kind)
}

func getEntries(waitGroup *sync.WaitGroup, entries *[]commons.Entry, area, kind string) {
	infra.ReadFileStream(getFileName(area, kind), ",", func(entry commons.Entry) {
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
