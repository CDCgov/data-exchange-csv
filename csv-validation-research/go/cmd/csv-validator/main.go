package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {

	defer stopCPUProfile()
	startCPUProfile("cpu.pprof") // call function to start cpu profile

	totalStartTime := time.Now()

	for i := 0; i < 100; i++ {
		f, err := os.Open("data/file-with-headers-100000-rows.csv")
		if err != nil {
			fmt.Println("Error:", err)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(f)

		reader := csv.NewReader(f)

		for {
			record, err := reader.Read()

			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
			}

			for _, field := range record {
				var _ = field
			}
		}
	}
	writeMemoryProfile("memory.pprof")

	totalEndTime := time.Now()
	totalDuration := totalEndTime.Sub(totalStartTime).Milliseconds()
	fmt.Println("Total process took", totalDuration, "ms")
}

func startCPUProfile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
}

func stopCPUProfile() {
	pprof.StopCPUProfile()
}

func writeMemoryProfile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
	f.Close()
}
