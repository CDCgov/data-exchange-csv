const fileSystem = require('fs');
const papaParse = require('papaparse');
const Benchmark = require('benchmark');


const csvFilePath = './data/file-with-headers-10000-rows.csv';
const suite = new Benchmark.Suite;

// Add a benchmark for reading and parsing the CSV file
suite.add('Papaparse benchmark results: ', {

    defer: true, // This will defer the execution of the benchmark
    
    fn: deferred => {
        // Measure memory usage before
        const memoryBefore = process.memoryUsage().heapUsed;

        fileSystem.readFile(csvFilePath, 'utf8', (err, data) => {
            if (err) {
                console.error('Error reading the file:', err);
                return;
            }

            papaParse.parse(data, {
                header: true,
                step: function(row){ 
                    for (const field in row.data) {
                        const value = row.data[field]
                    }
                    if (Object.keys(row.errors).length !=0) {
                        console.log("errors ", row.errors)
                    }
                }
            });

            // Measure memory usage after parsing
            const memoryAfter = process.memoryUsage().heapUsed;

            // Calculate the memory usage difference
            const memoryDifference = memoryAfter - memoryBefore;

            console.log(`Memory usage difference: ${memoryDifference} bytes`);

            deferred.resolve(); // signal the end of the benchmark
        });
    }
   
});

// Run the benchmark
suite.on('complete', function() {
    this.forEach(result => {
        console.log(`${result}`); //result.hz-> number of times per second the benchmarked operation was executed
    });
}).run({ async: true });