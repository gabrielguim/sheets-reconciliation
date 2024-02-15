package use_case

import (
	"sheets-reconciliation/commons"
	"sheets-reconciliation/consolidation/model"
	"sheets-reconciliation/consolidation/repository"
	"sync"
)

type ConsolidatedArea struct {
	Payments int64
	Receipts int64
	Balance  int64
}

func dispatchConsolidation(area commons.Area) ConsolidatedArea {
	areaWaitGroup := new(sync.WaitGroup)
	areaWaitGroup.Add(2)
	payments, receipts := int64(0), int64(0)

	go func() {
		defer areaWaitGroup.Done()
		payments = model.Consolidate(area.Payments)
	}()

	go func() {
		defer areaWaitGroup.Done()
		receipts = model.Consolidate(area.Receipts)
	}()

	areaWaitGroup.Wait()

	return ConsolidatedArea{
		Payments: payments,
		Receipts: receipts,
		Balance:  receipts - payments,
	}
}

func ConsolidateArea(areaName string, repository repository.DataRepository) ConsolidatedArea {
	area := repository.Read(areaName)
	return dispatchConsolidation(area)
}
