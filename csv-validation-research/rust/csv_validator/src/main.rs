use std::error::Error;
use std::time::Instant;

use csv;

fn main() -> Result<(), Box<dyn Error>> {
    
    read_csv("data/file-with-headers-10000-rows.csv")
}

fn read_csv(csv_file_path: &str) -> Result<(), Box<dyn Error>> {
    let start_time = Instant::now();
    /*
    //example of configuring csv reader
    let mut rdr = csv::ReaderBuilder::new()
        .has_headers(true)
        .delimiter(b',')
        .double_quote(true)
        .escape(Some(b'\\'))
        .comment(Some(b'#'))
        .from_path(csv_file_path);
     */
    
    let mut reader = csv::Reader::from_path(csv_file_path)?;
    for _ in 0..5000 {
        
        for result in reader.records(){
       
            match result {
                Ok(record) => {
                println!("{:?}", record);
                //process rows here  
                }
                Err(err) => {
                eprintln!("Error reading CSV record: {}", err);
                // Handle the error here
                }
            }
        }
    }

    
    let end_time = Instant::now();
    let elapsed_time = end_time.duration_since(start_time);
    println!("Elapsed time: {} milliseconds", elapsed_time.as_millis());
    Ok(())
}