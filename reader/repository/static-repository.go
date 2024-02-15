package repository

import (
	"fmt"
	"sheets-reconciliation/commons"
	"sheets-reconciliation/reader/infra"
	"sync"
)

type StaticAreaRepository struct{}

const StaticFilePath = "/home/gabriel/pessoal/sheets-reconciliation/static/"

func getFileName(area, kind string) string {
	return fmt.Sprintf("%s%s_%s.csv", StaticFilePath, area, kind)
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
