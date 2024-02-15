package model

import "sheets-reconciliation/commons"

func Consolidate(entries []commons.Entry) int64 {
	sum := int64(0)

	for _, entry := range entries {
		sum += int64(entry.Value)
	}

	return sum
}
