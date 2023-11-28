package commons;

import com.opencsv.CSVReader;
import org.apache.commons.csv.CSVFormat;
import org.apache.commons.csv.CSVParser;
import org.apache.commons.csv.CSVPrinter;
import org.apache.commons.csv.CSVRecord;
import parse.Iparser;

import javax.xml.validation.Validator;
import java.io.BufferedWriter;
import java.io.FileReader;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.time.LocalDateTime;
import java.time.temporal.ChronoUnit;
import java.util.Arrays;

public class ApacheCommonsCSVParser implements Iparser {
    @Override
    public void parseCSVFileWithHeader(Path filePath, String[] headerRow) {
        LocalDateTime startTime = LocalDateTime.now();
        LocalDateTime endTime = LocalDateTime.MIN;
        try (CSVParser csvParser = new CSVParser(Files.newBufferedReader(filePath),
                CSVFormat.RFC4180.withFirstRecordAsHeader())) {
            FileReader filereader = new FileReader(filePath.toString());
            final CSVFormat csvFormat = CSVFormat.Builder.create()
                    .setHeader(headerRow)
                    .setAllowMissingColumnNames(true)
                    .build();
            final Iterable<CSVRecord> records = csvFormat.parse(filereader);

            csvParser.stream().forEach((record) -> {
                for (String cell : record) {
                    String val = cell;
                }
            });


        } catch (IOException e) {
            throw new RuntimeException(e);
        } finally {
            endTime = LocalDateTime.now();
        }

        long elapsedMillisSeconds = ChronoUnit.MILLIS.between(startTime, endTime);
        System.out.println("Total Processing Time (in millis): "+ elapsedMillisSeconds + "ms");
    }
}