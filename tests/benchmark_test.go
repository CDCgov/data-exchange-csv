package test

import (
	"testing"
)

func BenchmarkCliFlags(b *testing.B)                  {}
func BenchmarkSetupEnvironment(b *testing.B)          {}
func BenchmarkFileValidate(b *testing.B)              {}
func BenchmarkStoreFileValidationResult(b *testing.B) {}
func BenchmarkRowValidate(b *testing.B)               {}
func BenchmarkRowTransform(b *testing.B)              {}
func BenchmarkStoreResult(b *testing.B)               {}
