# CSV File Validation results using .NET csv library

| **Number of executions** | **Total Time(MS)** | **Total CPU Time(MS)** |**Memory Allocation(MB)**|**Filename**|
| --------------- | --------------- | --------------- |--------------- | --------------- | 
| 100 |158|||file-with-headers-100-rows.csv|
| 5000|7980|||file-with-headers-100-rows.csv|
| 100 |15100|||file-with-headers-10000-rows.csv|
| 5000|847707|||file-with-headers-10000-rows.csv|

## How to run the code
1. `git clone` to clone the repository.
2. `cd csv_validation-research/dotnet`
3. `dotnet build` to compile the .NET code and its dependencies
4. `dotnet run` to run the application
 
 > Ensure the [.NET SDK 8.0](https://dotnet.microsoft.com/en-us/download/dotnet/8.0) is installed on your machine. The [C# Dev Kit](https://marketplace.visualstudio.com/items?itemName=ms-dotnettools.csdevkit) extension for Visual Studio Code is also recommended.

