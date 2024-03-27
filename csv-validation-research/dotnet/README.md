# CSV File Validation results using .NET csv library

| **Number of executions** | **Total Time(MS)** | **Total CPU Time(MS)** |**Memory Allocation(MB)**|**Filename**|
| --------------- | --------------- | --------------- |--------------- | --------------- | 
| 100 |158|||file-with-headers-100-rows.csv|
| 5000|7980|||file-with-headers-100-rows.csv|
| 100 |15100|||file-with-headers-10000-rows.csv|
| 5000|847707|||file-with-headers-10000-rows.csv|

##How to run the code
#### Run `git clone` to clone the repository.
#### Run `cd csv_validation-research/dotnet`
#### Run `dotnet build` to compile the Rust code and its dependencies
#### Run `dotnet run` to run the application
> **Note:** Make sure .NET SDK is installed on your machine, as well as C# extension if using `Visual Studio Code`

## To generate and view Profiling metrics with perf:
- TO DO