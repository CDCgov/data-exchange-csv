package benchmark

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func BenchmarkCSVValidator(b *testing.B) {
	b.ReportAllocs() //report allocations during benchmark
	f, err := os.Open("../../data/file-with-headers-10000-rows.csv")
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
