package consolidation

import (
	"sheets-reconciliation/internal/commons"
	"sync"
)

type ConsolidatedArea struct {
	Payments int64
	Receipts int64
	Balance  int64
}

func ConsolidateEntries(entries []commons.Entry) int64 {
	sum := int64(0)

	for _, entry := range entries {
		sum += int64(entry.Value)
	}

	return sum
}

func dispatchConsolidation(area commons.Area) ConsolidatedArea {
	areaWaitGroup := new(sync.WaitGroup)
	areaWaitGroup.Add(2)
	payments, receipts := int64(0), int64(0)

	go func() {
		defer areaWaitGroup.Done()
		payments = ConsolidateEntries(area.Payments)
	}()

	go func() {
		defer areaWaitGroup.Done()
		receipts = ConsolidateEntries(area.Receipts)
	}()

	areaWaitGroup.Wait()

	return ConsolidatedArea{
		Payments: payments,
		Receipts: receipts,
		Balance:  receipts - payments,
	}
}

func ConsolidateArea(areaName string, repository AreaRepository) ConsolidatedArea {
	area := repository.Read(areaName)
	return dispatchConsolidation(area)
}
