const fileSystem = require('fs');
const papaParse = require('papaparse');
const Benchmark = require('benchmark');


const csvFilePath = './data/file-with-headers-10000-rows.csv';
const suite = new Benchmark.Suite;

// Add a benchmark for reading and parsing the CSV file
suite.add('Papaparse benchmark results: ', {

    defer: true, // defer execution of benchmark until its manually stopped
    
    fn: deferred => {
        // read the csv file async
        fileSystem.readFile(csvFilePath, 'utf8', (err, data) => {
            if (err) {
                console.error('Error reading the file:', err);
                return;
            }
            //create array to hold row objects
            const rows = [];
            papaParse.parse(data, {
                header: true,
                step: function(row){ 
                    const rowObject = {};
                    let i =0;
                    for (const field in row.data) {
                        const value = row.data[field]
                        rowObject[i] = value;
                        i++;
                    }
                    if (Object.keys(row.errors).length !=0) {
                        console.log("errors ", row.errors)
                    }
                    rows.push(row);
                },
                complete: function() {
                    deferred.resolve(); // end of the benchmark execution
                },
            });
           
        });
    }
   
});

// Run the benchmark
suite.on('complete', function() {
    this.forEach(result => {
        console.log(`${result.name}: ${result.hz} ops/sec ±${result.stats.rme}%`); 
    });
    const memoryAllocated = process.memoryUsage().heapUsed;
    const memoryAsMegabytes = memoryAllocated  / (1024 * 1024);
    console.log("Memory usage in MB :", memoryAsMegabytes);
}).run({ async: true });