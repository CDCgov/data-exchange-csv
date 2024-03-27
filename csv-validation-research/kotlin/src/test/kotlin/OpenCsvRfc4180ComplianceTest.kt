import com.opencsv.CSVParserBuilder
import com.opencsv.CSVReader
import com.opencsv.CSVReaderBuilder
import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Disabled
import org.junit.jupiter.api.DisplayName
import org.junit.jupiter.api.Test
import java.io.StringReader
import java.util.*

class OpenCsvRfc4180ComplianceTest {

    @Test
    @DisplayName("Test for CSV data with CRLF line breaks")
    fun lineBreakCRLF(){
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
    @DisplayName("Test for CSV data with CR line breaks")
    fun lineBreakCR(){
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
    @DisplayName("Test for CSV data with LF line breaks")
    fun lineBreakLF(){
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
    @DisplayName("Test for CSV data with Line breaks on the last record")
    fun lineBreakAtTheEndOfFile(){
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
    @DisplayName("Test for CSV data with no header")
    fun noHeader(){
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
    @DisplayName("Test for within the header and within each record, there may be one or more fields, separated by commas ")
    fun fieldOneOrMoreFields(){
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
    @DisplayName("Test for each record should contain the same number of fields throughout the file")
    fun recordWithDifferentNumberOfFields(){
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
    @DisplayName("Test for spaces are considered part of a field and should not be ignored")
    fun recordFieldsWithSpaces(){
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
    @DisplayName("Test for The last field in the record must not be followed by a comma")
    @Disabled
    fun lastFieldInRecordFollowedByComma(){
        // DOES NOT THROW ERROR WHEN file ends with ',' therefore does not adhere to RFC4180 standard
        val csvData = "Name,Email\nJohn,john@example.com\nJane,jane@example.com\nAlice\nBob,bob@example.com\nCharlie\nDiana,diana@example.com\nEva,eva@example.com\nFrank,frank@example.com\nGrace,grace@example.com\nHenry,henry@example.com,"

        // Create string  reader
        val reader = StringReader(csvData)
        // Create csv  reader
        val csvReader = CSVReader(reader)
        println (csvReader)
    }
    @Test
    @DisplayName("Test for If fields are not enclosed with double quotes, then double quotes may not appear inside the fields")
    fun quotesInFieldNotEnclosedWithDoubleQuotes(){
        val csvData = "Name,Email\nJohn,john@example.com\nJane,jane@example.com\nAlice,alice@example.com\nBob,bob@example.com\nCharlie,charlie@example.com\nDiana,diana@example.com\nEva,eva@example.com\nFrank,frank@example.com\nGrace,grace@example.com\nHenry,henry@example.com"

        // Create StringReader
        val reader = StringReader(csvData)

        val parsedData = mutableListOf<String>()

        val csvParserBuilder = CSVParserBuilder()
            .build()
        val builder = CSVReaderBuilder(reader)
            .withCSVParser(csvParserBuilder)
            .withErrorLocale(Locale.US)
            .withKeepCarriageReturn(true)
            .build()

        var record: Array<String>?
        while (builder.readNext().also {record=it} !=null){
            val row = record?.joinToString(",")
            if (row != null) {
                parsedData.add(row)
            }
        }
        builder.close()

        val expectedOutput = mutableListOf<String>()
        expectedOutput.add("Name,Email")
        expectedOutput.add("John,john@example.com")
        expectedOutput.add("Jane,jane@example.com")
        expectedOutput.add("Alice,alice@example.com")
        expectedOutput.add("Bob,bob@example.com")
        expectedOutput.add("Charlie,charlie@example.com")
        expectedOutput.add("Diana,diana@example.com")
        expectedOutput.add("Eva,eva@example.com")
        expectedOutput.add("Frank,frank@example.com")
        expectedOutput.add("Grace,grace@example.com")
        expectedOutput.add("Henry,henry@example.com")

        assertEquals(expectedOutput, parsedData, "Unexpected $parsedData")
    }
    @Test
    @DisplayName("Test for CSV data If double-quotes are used to enclose fields, then a double-quote appearing inside a field must be escaped by preceding it with another double quote")
    fun doubleQuotesInsideQuotedField(){
        val csvData = "Name,Email\nJohn,john@example.com\nJane,\"jane\"\"@example.com\"\nAlice,alice@example.com\nBob,bob@example.com\nCharlie,charlie@example.com\nDiana,diana@example.com\nEva,eva@example.com\nFrank,frank@example.com\nGrace,grace@example.com\nHenry,henry@example.com"

        val reader = CSVReaderBuilder(StringReader(csvData))
            .withCSVParser(CSVParserBuilder().withQuoteChar('"').build())
            .build()

        val records = reader.readAll()

        val expectedOutput = listOf(
            listOf("Name", "Email"),
            listOf("John", "john@example.com"),
            listOf("Jane", "jane\"@example.com"),
            listOf("Alice", "alice@example.com"),
            listOf("Bob", "bob@example.com"),
            listOf("Charlie", "charlie@example.com"),
            listOf("Diana", "diana@example.com"),
            listOf("Eva", "eva@example.com"),
            listOf("Frank", "frank@example.com"),
            listOf("Grace", "grace@example.com"),
            listOf("Henry", "henry@example.com")
        )

        for ((rowIndex, record) in records.withIndex()) {
            for ((fieldIndex, field) in record.withIndex()) {
                assert(field == expectedOutput[rowIndex][fieldIndex]) {
                    "Expected field: ${expectedOutput[rowIndex][fieldIndex]}, and actual field: $field"
                }

            }
        }

    }
    @Test
    @DisplayName("Test for CSV data  with line breaks inside quotes")
    fun fieldsWithLineBreaksInsideQuotes(){
        val csvData = "Name,Email\nJohn,john@example.com\nJane,\"jane   \n@example.com\"\nAlice,alice@example.com\nBob,bob@example.com\nCharlie,charlie@example.com\nDiana,diana@example.com\nEva,eva@example.com\nFrank,frank@example.com\nGrace,grace@example.com\nHenry,\"henry@example.com       \n\""

        val reader = CSVReaderBuilder(StringReader(csvData))
            .withCSVParser(CSVParserBuilder().withQuoteChar('"').build())
            .build()

        val records = reader.readAll()

        val expectedOutput = listOf(
            listOf("Name", "Email"),
            listOf("John", "john@example.com"),
            listOf("Jane", "jane   \n@example.com"),
            listOf("Alice", "alice@example.com"),
            listOf("Bob", "bob@example.com"),
            listOf("Charlie", "charlie@example.com"),
            listOf("Diana", "diana@example.com"),
            listOf("Eva", "eva@example.com"),
            listOf("Frank", "frank@example.com"),
            listOf("Grace", "grace@example.com"),
            listOf("Henry", "henry@example.com       \n")
        )

        for ((rowIndex, record) in records.withIndex()) {
            for ((fieldIndex, field) in record.withIndex()) {
                assert(field == expectedOutput[rowIndex][fieldIndex]) {
                    "Expected field: ${expectedOutput[rowIndex][fieldIndex]}, and actual field: $field"
                }
            }
        }
    }
    @Test
    @DisplayName("Test for CSV data with fields and within fields commas inside quotes")
    fun fieldsWithCommasInsideQuotes() {
        val csvData = "Name,Email\nJohn,john@example.com\nJane,\"jane,@example.com\"\nAlice,alice@example.com\nBob,bob@example.com\nCharlie,charlie@example.com\nDiana,diana@example.com\nEva,eva@example.com\nFrank,frank@example.com\nGrace,grace@example.com\nHenry,\"henry,@example.com\""

        val reader = CSVReaderBuilder(StringReader(csvData))
            .withCSVParser(CSVParserBuilder().withQuoteChar('"').build())
            .build()

        val records = reader.readAll()

        val expectedOutput = listOf(
            listOf("Name", "Email"),
            listOf("John", "john@example.com"),
            listOf("Jane", "jane,@example.com"),
            listOf("Alice", "alice@example.com"),
            listOf("Bob", "bob@example.com"),
            listOf("Charlie", "charlie@example.com"),
            listOf("Diana", "diana@example.com"),
            listOf("Eva", "eva@example.com"),
            listOf("Frank", "frank@example.com"),
            listOf("Grace", "grace@example.com"),
            listOf("Henry", "henry,@example.com")
        )

        for ((rowIndex, record) in records.withIndex()) {
            for ((fieldIndex, field) in record.withIndex()) {
                assert(field == expectedOutput[rowIndex][fieldIndex]) {
                    "Expected field: ${expectedOutput[rowIndex][fieldIndex]}, and actual field: $field"
                }
            }
        }
    }
}