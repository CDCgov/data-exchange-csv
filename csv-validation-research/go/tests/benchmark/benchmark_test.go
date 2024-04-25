package benchmark

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

type Row struct {
	Data map[int]string
}

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

	var rows []Row
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}

		//create row object
		row := Row{Data: make(map[int]string)}
		for i, field := range record {
			row.Data[i] = field

		}
		rows = append(rows, row)

	}

}
