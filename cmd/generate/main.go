package main

import (
	"fmt"
	"sheets-reconciliation/internal/generator"
	"time"
)

func main() {
	startTime := time.Now()
	fmt.Println("Start generator for area1")
	generator.GenerateAreaFile("area1", 2_000_000)
	fmt.Printf("Time elapsed for generate area1: %dms\n", time.Since(startTime).Milliseconds())

	startTime = time.Now()
	fmt.Println("Start generator for area2")
	generator.GenerateAreaFile("area2", 5_000_000)
	fmt.Printf("Time elapsed for generate area2: %dms\n", time.Since(startTime).Milliseconds())

	startTime = time.Now()
	fmt.Println("Start generator for area3")
	generator.GenerateAreaFile("area3", 10_000_000)
	fmt.Printf("Time elapsed for generate area3: %dms\n", time.Since(startTime).Milliseconds())
}
