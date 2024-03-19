import com.opencsv.CSVParserBuilder
import com.opencsv.CSVReader
import com.opencsv.CSVReaderBuilder
import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import java.io.StringReader
import java.util.*

class OpenCsvRfc4180ComplianceTest {

    @Test
    fun lineBreakCRLF(){
        // CSV data with CRLF line breaks
        val csvData = "Name,Email\r\nJane Doe,johndoe@example.com\r\nJane Smith,janesmith@example.com\r\nChris Mallok,cmallok@example.com"
        // Create string  reader
        val reader = StringReader(csvData)
        // Create csv  reader
        val csvReader = CSVReader(reader)
        val expectedNumberOfRecords = 3
        var actualNumberOfRecords = -1 // Subtract header row
        while (csvReader.readNext() !=null){
            actualNumberOfRecords +=1
        }
        assertEquals(expectedNumberOfRecords, actualNumberOfRecords, "Expected $expectedNumberOfRecords records, but got $actualNumberOfRecords")
    }
    @Test
    fun lineBreakCR(){
        // CSV data with CR line breaks
        val csvData = "Name,Email\rJane Doe,johndoe@example.com\rJane Smith,janesmith@example.com\rChris Mallok,cmallok@example.com"
        // Create string  reader
        val reader = StringReader(csvData)
        // Create csv  reader
        val csvReader = CSVReader(reader)
        val expectedNumberOfRecords = 3
        var actualNumberOfRecords = -1 // Subtract header row
        while (csvReader.readNext() !=null){
            actualNumberOfRecords +=1
        }
        assertEquals(expectedNumberOfRecords, actualNumberOfRecords, "Expected $expectedNumberOfRecords records, but got $actualNumberOfRecords")
    }
    @Test
    fun lineBreakLF(){
        // CSV data with LF line breaks
        val csvData = "Name,Email\nJane Doe,johndoe@example.com\nJane Smith,janesmith@example.com\nChris Mallok,cmallok@example.com"
        // Create string  reader
        val reader = StringReader(csvData)
        // Create csv  reader
        val csvReader = CSVReader(reader)
        val expectedNumberOfRecords = 3
        var actualNumberOfRecords = -1 // Subtract header row
        while (csvReader.readNext() !=null){
            actualNumberOfRecords +=1
        }
        assertEquals(expectedNumberOfRecords, actualNumberOfRecords, "Expected $expectedNumberOfRecords records, but got $actualNumberOfRecords")
    }
    @Test
    fun lineBreakAtTheEndOfFile(){
        // CSV data with Line breaks on the last record
        val csvData = "Name,Email\nJane Doe,johndoe@example.com\nJane Smith,janesmith@example.com\nChris Mallok, cmallok@example.com\n"
        // Create string  reader
        val reader = StringReader(csvData)
        // Create csv  reader
        val csvReader = CSVReader(reader)
        val expectedNumberOfRecords = 3
        var actualNumberOfRecords = -1 // Subtract header row
        while (csvReader.readNext() !=null){
            actualNumberOfRecords +=1
        }
        assertEquals(expectedNumberOfRecords, actualNumberOfRecords, "Expected $expectedNumberOfRecords records, but got $actualNumberOfRecords")
    }
    @Test
    fun noHeader(){
        // CSV data no header
        val csvData = "Jane Doe,johndoe@example.com\nJane Smith,janesmith@example.com\nChris Mallok, cmallok@example.com\n"
        // Create string  reader
        val reader = StringReader(csvData)
        // Create csv  reader
        val csvReader = CSVReader(reader)
        val expectedNumberOfRecords = 3
        var actualNumberOfRecords = 0
        while (csvReader.readNext() !=null){
            actualNumberOfRecords +=1
        }
        assertEquals(expectedNumberOfRecords, actualNumberOfRecords, "Expected $expectedNumberOfRecords records, but got $actualNumberOfRecords")
    }
    @Test
    fun fieldOneOrMoreFields(){
        //Within the header and within each record, there may be one or more fields, separated by commas.
        val csvData = "Name,Email\nJohn,john@example.com\nJane,jane@example.com\nAlice,alice@example.com\nBob,bob@example.com\nCharlie,charlie@example.com\nDiana,diana@example.com\nEva,eva@example.com\nFrank,frank@example.com\nGrace,grace@example.com\nHenry,henry@example.com"
        // Create string  reader
        val reader = StringReader(csvData)
        // Create csv  reader
        val csvReader = CSVReader(reader)
        val expectedNumberOfRecords = 10
        var actualNumberOfRecords = -1
        while (csvReader.readNext() !=null){
            actualNumberOfRecords +=1
        }
        assertEquals(expectedNumberOfRecords, actualNumberOfRecords, "Expected $expectedNumberOfRecords records, but got $actualNumberOfRecords")
    }
    @Test
    fun recordWithDifferentNumberOfFields(){
        //Each record should contain the same number of fields throughout the file
        val csvData = "Name,Email\nJohn,john@example.com\nJane,jane@example.com\nAlice\nBob,bob@example.com\nCharlie\nDiana,diana@example.com\nEva,eva@example.com\nFrank,frank@example.com\nGrace,grace@example.com\nHenry,henry@example.com"

        // Create string  reader
        val reader = StringReader(csvData)
        // Create csv  reader
        val csvReader = CSVReader(reader)

        val expectedNumberOfRecords = 10
        var actualNumberOfRecords = -1
        while (csvReader.readNext() !=null){
            actualNumberOfRecords +=1
        }
        assertEquals(expectedNumberOfRecords, actualNumberOfRecords, "Expected $expectedNumberOfRecords records, but got $actualNumberOfRecords")
    }
    @Test
    fun recordFieldsWithSpaces(){
        //Spaces are considered part of a field and should not be ignored.
        val csvData = "Name,Email\nJohn,john@example.com    \nJane    ,jane@example.com\nAlice,alice@example.com\nBob,bob@example.com\nCharlie,charlie@example.com\nDiana,diana@example.com       \nEva,eva@example.com\nFrank,frank@example.com\nGrace,grace@example.com\nHenry    ,henry@example.com"

        // Create StringReader
        val reader = StringReader(csvData)

        val parsedData = mutableListOf<String>()

        val csvParserBuilder = CSVParserBuilder()
            .withSeparator(',')
            .withQuoteChar('"')
            .withIgnoreLeadingWhiteSpace(false)
            .withIgnoreQuotations(false)
            .withEscapeChar('\r')

            .build()
        val builder = CSVReaderBuilder(reader)
            .withCSVParser(csvParserBuilder)
            .withErrorLocale(Locale.US)
            .withKeepCarriageReturn(true)
            .build()

        var record: Array<String>?
        while (builder.readNext().also {record=it} !=null){
            val row = record?.joinToString(",")
            println (row)
            if (row != null) {
                parsedData.add(row)
            }
        }
        builder.close()

        val expectedOutput = mutableListOf<String>()
        expectedOutput.add("Name,Email")
        expectedOutput.add("John,john@example.com    ")
        expectedOutput.add("Jane    ,jane@example.com")
        expectedOutput.add("Alice,alice@example.com")
        expectedOutput.add("Bob,bob@example.com")
        expectedOutput.add("Charlie,charlie@example.com")
        expectedOutput.add("Diana,diana@example.com       ")
        expectedOutput.add("Eva,eva@example.com")
        expectedOutput.add("Frank,frank@example.com")
        expectedOutput.add("Grace,grace@example.com")
        expectedOutput.add("Henry    ,henry@example.com")

        assertEquals(expectedOutput, parsedData, "Unexpected $parsedData")
    }
    @Test
    fun lastFieldInRecordFollowedByComma(){
        //The last field in the record must not be followed by a comma.
        // DOES NOT THROW ERROR WHEN file ends with ','
        val csvData = "Name,Email\nJohn,john@example.com\nJane,jane@example.com\nAlice\nBob,bob@example.com\nCharlie\nDiana,diana@example.com\nEva,eva@example.com\nFrank,frank@example.com\nGrace,grace@example.com\nHenry,henry@example.com,"

        // Create string  reader
        val reader = StringReader(csvData)
        // Create csv  reader
        val csvReader = CSVReader(reader)

        println (csvReader)
    }
    @Test
    fun quotesInFieldNotEnclosedWithDoubleQuotes(){
        //
    }
    @Test
    fun doubleQuotesInsideQuotedField(){
        //
    }
    @Test
    fun fieldsWithLineBreaksInsideQuotes(){
        //
    }
    @Test
    fun fieldsWithCommasInsideQuotes() {
        //
    }
}