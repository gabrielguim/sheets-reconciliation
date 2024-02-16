package generator

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Printf("failed to close file: %v\n", err)
	}
}

func GenerateAreaFile(areaName string, lines uint32) {
	staticFilePath, _ := os.Getwd()

	paymentsFile, err := os.OpenFile(fmt.Sprintf("%s/static/%s_payments.csv", staticFilePath, areaName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer closeFile(paymentsFile)

	if err != nil {
		log.Fatal(err)
	}

	receiptsFile, err := os.OpenFile(fmt.Sprintf("%s/static/%s_receipts.csv", staticFilePath, areaName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer closeFile(receiptsFile)

	if err != nil {
		log.Fatal(err)
	}

	paymentsWriter := bufio.NewWriter(paymentsFile)
	receiptsWriter := bufio.NewWriter(receiptsFile)

	paymentsWriter.WriteString("name,value")
	receiptsWriter.WriteString("name,value")

	writeWaitGroup := new(sync.WaitGroup)
	writeWaitGroup.Add(2)

	go func() {
		startTime := time.Now()
		for i := uint32(0); i < lines; i++ {
			paymentsWriter.WriteString(fmt.Sprintf("\npayment %d,%d", i+1, i+500))
		}

		writeWaitGroup.Done()
		fmt.Printf("Time elapsed for payment: %dms\n", time.Since(startTime).Milliseconds())
	}()

	go func() {
		startTime := time.Now()
		for i := uint32(0); i < lines; i++ {
			receiptsWriter.WriteString(fmt.Sprintf("\nreceipt %d,%d", i+1, i+1_000))
		}

		writeWaitGroup.Done()
		fmt.Printf("Time elapsed for receipt: %dms\n", time.Since(startTime).Milliseconds())

	}()

	writeWaitGroup.Wait()

	paymentsWriter.Flush()
	receiptsWriter.Flush()
}
