# CSV File Validation results using the [GoLang CSV library](https://pkg.go.dev/encoding/csv)

| **Number of executions** | **Total Time(MS)** | **Total CPU Time(MS)** |**Memory Allocation(MB)**|**Filename**|
| --------------- | --------------- | --------------- |--------------- | --------------- | 
| 100 |26|NA|1.16|file-with-headers-100-rows.csv|
| 5000|706|710|3.66|file-with-headers-100-rows.csv|
| 100 |628|410|1.15|file-with-headers-10000-rows.csv|
| 5000|31992|29040|6.6|file-with-headers-10000-rows.csv|



## To generate and view Profiling metrics with pprof:
- run the Go code that will create cpu.pprof and memory.pprof files
- using terminal run `go tool pprof cpu.pprof or memory.pprof` that will start interactive shell with pprof.
- type `help` to see a list of available commands.

## Unit Tests created for Compliance with RFC 4180
We have conducted an analysis to ensure that the Go CSV package complies fully with the [RFC 4180](https://www.rfc-editor.org/rfc/rfc4180) standards for CSV files. Through a suite of unit tests, we have verified that the Go CSV package adheres to the guidelines outlined in RFC 4180.
 

### Running Unit Tests:
1. Clone the repository to your local machine.
2. Navigate to the root project directory.
3. run `go test .\tests`

### Running Benchmark:
1. Clone the repository to your local machine
2. Navigate to the root project directory
3. Run `cd tests/benchmark`
4. Run `go test -bench . -benchmem` to see results
