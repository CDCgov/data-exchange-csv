# CSV File Validation results using Rust csv library

| **Number of executions** | **Total Time(MS)** | **Total CPU Time(MS)** |**Memory Allocation(MB)**|**Filename**|
| --------------- | --------------- | --------------- |--------------- | --------------- | 
| 100 |1880|||file-with-headers-100-rows.csv|
| 5000|130865|||file-with-headers-100-rows.csv|
| 100 |283355|||file-with-headers-10000-rows.csv|
| 5000||||file-with-headers-10000-rows.csv|

## How to Install Rust:
- Download and install `Rust` from https://www.rust-lang.org/ (Note:`Cargo` Rust package manager will be installed during this step) 
- 'Microsoft C++ Build Tools' (Some Rust crates (packages) rely on native code written in C or C++ and to be able to use it, this tool needs to be installed)

## How to run the code:
#### Run `git clone` to clone the repository.
#### cd `cd csv_validator/src`
#### run `cargo build` to compile the Rust code and its dependencies
#### run `cargo run` to run the application


## To generate and view Profiling metrics with perf:
- TO DO