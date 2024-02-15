package main

import (
	"fmt"
	use_case "sheets-reconciliation/consolidation/use-case"
	"sheets-reconciliation/reader/repository"
)

func main() {
	staticAreaRepository := new(repository.StaticAreaRepository)
	fmt.Println(use_case.ConsolidateArea("area1", staticAreaRepository))
}
