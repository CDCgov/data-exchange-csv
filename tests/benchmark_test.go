package test

import (
	"testing"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func BenchmarkConfig(b *testing.B) {

}
func RunBenchmark(b *testing.B, benchmarkType constants.BenchmarkType) {
	switch benchmarkType {
	case constants.CLI_PARSE_FLAGS_BENCHMARK:
		benchmarkCliFlags(b)
	case constants.SETUP_ENVIRONMENT_BENCHMARK:
		benchmarkSetupEnvironment(b)
	case constants.FILE_VALIDATION_BENCHMARK:
		benchmarkFileValidate(b)
	case constants.STORE_FILE_VALIDATION_BENCHMARK:
		benchmarkStoreFileValidationResult(b)
	case constants.ROW_VALIDATION_BENCHMARK:
		benchmarkRowValidate(b)
	case constants.ROW_TRANSFORMATION_BENCHMARK:
		benchmarkRowTransform(b)
	case constants.STORE_RESULTS_BENCHMARK:
		benchmarkStoreResult(b)
	}
}

func benchmarkCliFlags(b *testing.B)                  {}
func benchmarkSetupEnvironment(b *testing.B)          {}
func benchmarkFileValidate(b *testing.B)              {}
func benchmarkStoreFileValidationResult(b *testing.B) {}
func benchmarkRowValidate(b *testing.B)               {}
func benchmarkRowTransform(b *testing.B)              {}
func benchmarkStoreResult(b *testing.B)               {}
