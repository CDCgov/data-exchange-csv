package main

// TODO: Move run() into a separate app package? That way we can organize how we define different execution envs - Thomas
import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"strings"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/file"
	"golang.org/x/text/encoding/charmap"
)

func run(
	ctx context.Context,
	args []string,
	stdout, stderr io.Writer,
) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	logger := slog.New(slog.NewTextHandler(stdout, nil))
	slog.SetDefault(logger)

	//source := "data/file-with-headers-100-rows.csv"
	//source := "data/file-with-headers-100-rows_with_BOM.csv"
	//source := "data/file-with-headers-100-rows_US_ASCII.csv"
	//source := "data/file-with-headers-rows_iso8859-1.csv"
	source := "data/file-with-headers-windows1252.csv"

	fileValidationResult := file.Validate(source)

	// TODO: FileValidationResult will contain unescaped quotation marks when converted to a string, this will
	// confuse any upstream consumers of logs b/c they are not readable in JSON
	slog.Info("file validation result: ", fileValidationResult)

	//detect encoding with random sample data
	//enc := utils.DetectEncoding(randomSampleData)
	decoder := charmap.Windows1252.NewDecoder()

	//tempcode-> check if windows1252 decoder can correctly parse the csv file
	file, err := os.Open(source)
	if err != nil {
		slog.Error(err.Error())
	}
	reader := csv.NewReader(decoder.Reader(file))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		slog.Info(strings.Join(record, ""))
	}

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
