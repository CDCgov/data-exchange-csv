package commons;


import commons.OpenCSVParser;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import java.net.URISyntaxException;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class OpenCSVParserTest {
    private OpenCSVParser parser = new OpenCSVParser();

    @BeforeAll
    public static void setup() {
        System.out.println("------------------");
        System.out.println("Open CSV");
        System.out.println("------------------");
    }
    // Parse CSV file with 10 headers and 100 alphanumeric rows
    @Test
    @DisplayName("Test to parse CSV file with 10 columns and 100 rows")
    public void parseCSVFile1Test() throws URISyntaxException {
        System.out.println("Test to parse CSV file with 10 columns and 100 rows");
        Path filePath = Path.of(this.getClass().getResource("file-with-headers-100-rows.csv").toURI());
        String[] headerRow = new String[] {"Index","Account Id","Lead Owner","First Name","Last Name","Company","Phone 1","Phone 2","Email","Website","Notes"};
        parser.parseCSVFileWithHeader(filePath, headerRow);
    }

    @Test
    @DisplayName("Test to parse CSV file with 10 columns and 100000 rows")
    public void parseCSVFileT2est() throws URISyntaxException {
        System.out.println("Test to parse CSV file with 10 columns and 100000 rows");
        Path filePath = Path.of(this.getClass().getResource("file-with-headers-100000-rows.csv").toURI());
        String[] headerRow = new String[] {"Index","Account Id","Lead Owner","First Name","Last Name","Company","Phone 1","Phone 2","Email","Website","Notes"};
        parser.parseCSVFileWithHeader(filePath, headerRow);
    }

    @Test
    @DisplayName("Test to parse CSV file with 100 columns and 100000 rows")
    public void parseCSVFileT3est() throws URISyntaxException {
        System.out.println("Test to parse CSV file with 100 columns and 100000 rows");
        Path filePath = Path.of(this.getClass().getResource("new.csv").toURI());
        List<String> records = new ArrayList<>();
        for (int i = 0; i < 100; i++) {
            records.add("column"+(i));
        }
        String[] headerRow = Arrays.stream(records.toArray()).map(Object::toString)
                .toArray(String[]::new);
        parser.parseCSVFileWithHeader(filePath, headerRow);
    }


    @Test
    @DisplayName("Test to parse CSV file with 10 columns and 100 rows with missing column values")
    public void parseCSVFile4Test() throws URISyntaxException {
        System.out.println("Test to parse CSV file with 10 columns and 100000 rows with missing column values");
        Path filePath = Path.of(this.getClass().getResource("file-with-headers-less-columns.csv").toURI());
        String[] headerRow = new String[] {"Index","Account Id","Lead Owner","First Name","Last Name","Company","Phone 1","Phone 2","Email","Website","Notes"};
        parser.parseCSVFileWithHeader(filePath, headerRow);
    }

    @Test
    @DisplayName("Test to parse CSV file with 10 columns and 100 rows with blank row")
    public void parseCSVFile5Test() throws URISyntaxException {
        System.out.println("Test to parse CSV file with 10 columns and 100000 rows with blank row");
        Path filePath = Path.of(this.getClass().getResource("file-with-blank-row.csv").toURI());
        String[] headerRow = new String[] {"Index","Account Id","Lead Owner","First Name","Last Name","Company","Phone 1","Phone 2","Email","Website","Notes"};
        parser.parseCSVFileWithHeader(filePath, headerRow);
    }

}