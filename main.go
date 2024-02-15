package main

import (
	"fmt"
	"sheets-reconciliation/consolidation"
	"sheets-reconciliation/reader"
)

func main() {
	staticAreaRepository := new(reader.StaticAreaRepository)
	fmt.Println(consolidation.ConsolidateArea("area1", staticAreaRepository))
}
