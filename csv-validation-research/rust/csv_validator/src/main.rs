use std::error::Error;
use std::fs::File;
use std::path::Path;
use std::time::Instant;


fn main() -> Result<(), Box<dyn Error>> {
  
        read_csv("data/file-with-headers-100-rows.csv")
    
}

fn read_csv<P: AsRef<Path> + Copy>(csv_filename: P) -> Result<(), Box<dyn Error>>{
    let start_time = Instant::now();
    for _ in 0..100 {
        let file = File::open(csv_filename)?;
        let mut rdr: csv::Reader<File> = csv::Reader::from_reader(file);

        for result in rdr.records() {
            match result {
                Ok(record) => println!("{:?}", record),
                Err(err) => {eprintln!("Error reading CSV record: {}", err);
            }
        }
    }
    }
    let end_time = Instant::now(); 
    let elapsed_time = end_time.duration_since(start_time);

    println!("Elapsed time: {} milliseconds", elapsed_time.as_millis());
    Ok(())

   

}
