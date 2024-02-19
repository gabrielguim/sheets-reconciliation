package main

import (
	"fmt"
	"sheets-reconciliation/internal/consolidation"
	"sheets-reconciliation/internal/reader"
)

func main() {
	staticAreaRepository := new(reader.StaticAreaRepository)
	fmt.Println(consolidation.ConsolidateArea("area1", staticAreaRepository))
	fmt.Println(consolidation.ConsolidateArea("area2", staticAreaRepository))
	fmt.Println(consolidation.ConsolidateArea("area3", staticAreaRepository))
}
