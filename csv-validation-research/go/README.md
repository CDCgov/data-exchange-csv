# CSV File Validation results using GoLang csv library

| **Number of executions** | **Total Time(MS)** | **Total CPU Time(MS)** |**Memory Allocation(MB)**|**Filename**|
| --------------- | --------------- | --------------- |--------------- | --------------- | 
| 100 |26|NA|1.16|file-with-headers-100-rows.csv|
| 5000|706|710|3.66|file-with-headers-100-rows.csv|
| 100 |628|410|1.15|file-with-headers-100000-rows.csv|
| 5000|31992|29040|6.6|file-with-headers-100000-rows.csv|


## To generate and view Profiling metrics with pprof:
- run the Go code that will create cpu.pprof and memory.pprof files
- using terminal run 'go tool pprof cpu.pprof or memory.pprof' that will start interactive shell with pprof.
- type 'help' to see a list of available commands.
