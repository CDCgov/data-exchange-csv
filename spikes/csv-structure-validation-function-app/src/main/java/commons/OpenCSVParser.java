package commons;


import com.opencsv.CSVReader;
import parse.Iparser;

import java.io.FileReader;
import java.nio.file.Path;
import java.time.LocalDateTime;
import java.time.temporal.ChronoUnit;

public class OpenCSVParser implements Iparser {
    @Override
    public void parseCSVFileWithHeader(Path filePath, String[] header) {
        LocalDateTime startTime = LocalDateTime.now();
        LocalDateTime endTime = LocalDateTime.MIN;
        try {
            // Create an object of filereader
            // class with CSV file as a parameter.
            FileReader filereader = new FileReader(filePath.toString());

            // create csvReader object passing
            // file reader as a parameter
            CSVReader csvReader = new CSVReader(filereader);
            String[] nextRecord;

            // we are going to read data line by line
            while ((nextRecord = csvReader.readNext()) != null) {
                for (String cell : nextRecord) {
                    String val = cell;
                }
            }
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            endTime = LocalDateTime.now();
        }

        long elapsedSeconds = ChronoUnit.MILLIS.between(startTime, endTime);
        System.out.println("Total Processing Time (in millis): "+ elapsedSeconds + "ms");
    }
}