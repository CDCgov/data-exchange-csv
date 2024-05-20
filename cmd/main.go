package main

import (
	"fmt"
	"os"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/file"
)

func main() {
	data := "data/file-with-headers-rows_iso8859-1.csv"
	//data := "data/file-with-headers-100-rows.csv"
	//data := "data/file-with-headers-100-rows.csv"
	//data := "data/file-with-headers-windows1252.csv"
	//data := "data/file-with-headers-100-rows_with_BOM.csv"
	//data := "data/tabDelimited.tsv"
	//data := "data/badFile.csv"
	//fmt.Println(data)
	//fileTest := "data/file-with-headers-100-rows.csv"
	//fileTest := "data/file-with-headers-rows_iso8859-1.csv"
	//fileTest := "data/file-with-headers-100-rows_US_ASCII.csv"
	fileTest := "data/file-with-headers-windows1252.csv"
	file.IsValid(data)
	file1, err := os.Open(fileTest)
	if err != nil {
		fmt.Println("Error occurred while opening a file")
	}
	defer file1.Close()

}
