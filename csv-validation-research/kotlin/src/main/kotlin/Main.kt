import com.fasterxml.jackson.dataformat.csv.CsvMapper
import com.github.doyaaaaaken.kotlincsv.dsl.csvReader
import com.opencsv.CSVParserBuilder
import com.opencsv.CSVReaderBuilder
import com.sun.jna.StringArray

import com.univocity.parsers.csv.CsvParser
import com.univocity.parsers.csv.CsvParserSettings
import de.siegmar.fastcsv.reader.CsvRecord

import io.deephaven.csv.CsvSpecs
import io.deephaven.csv.reading.CsvReader as DeephavenCSVReader
import io.deephaven.csv.sinks.SinkFactory

import de.siegmar.fastcsv.reader.CsvReader as FastCSVReader
import org.apache.commons.csv.CSVFormat
import org.apache.commons.csv.CSVParser
import org.supercsv.io.CsvListReader
import org.supercsv.prefs.CsvPreference


import java.io.BufferedReader
import java.io.File
import java.io.InputStream
import java.io.InputStreamReader
import java.nio.file.Path
import java.nio.file.Paths

import java.time.LocalDateTime



class CSVValidator {

    fun validateWithKotlinCSV(filename: String) {
        try {
            val resourceURL = object {}.javaClass.getResource(filename)
            val file: File? = resourceURL?.toURI()?.let { File(it) }

            if (file != null) {
                csvReader().open(file) {
                    readAllAsSequence().forEach { row: List<String> ->
                        for (field in row){
                            val f = field
                        }
                    }
                }
            }
        }catch (e: Exception){
            e.printStackTrace()
        }
    }
    fun validateWithOpenCSV(filename:String) {

        try {
            val inputStream: InputStream? =
                CSVValidator::class.java.getResourceAsStream("/$filename")

            val csvParserBuilder = CSVParserBuilder()
                .withIgnoreLeadingWhiteSpace(false)
                .withIgnoreQuotations(false)

                .build()
            val builder = CSVReaderBuilder(inputStream?.let { InputStreamReader(it) })
                .withCSVParser(csvParserBuilder)
                .build()

            var record: Array<String>?
            while (builder.readNext().also {record=it} !=null){
                for (row in record!!){
                    val field = row
                }
            }
            builder.close()
        }catch (e: Exception) {
            e.printStackTrace()
        }

    }
    fun validateWithApacheCommons(filename:String) {
        try {
            val inputStream: InputStream? =
                CSVValidator::class.java.getResourceAsStream("/$filename")
            val inputStreamReader = InputStreamReader(inputStream!!)
            val bufferedReader =  BufferedReader(inputStreamReader)

            val csvParser = CSVParser(bufferedReader,CSVFormat.RFC4180)
            for (csvRecord in csvParser){
                for (field in csvRecord){
                    val f = field
                }
            }
            inputStream.close()
            inputStreamReader.close()
            bufferedReader.close()
            csvParser.close()
        }catch (e: Exception) {
            e.printStackTrace()
        }
    }

    fun univocityCSVParser(filename: String){
        try {
            val settings = CsvParserSettings()
            settings.detectFormatAutomatically()
            settings.skipEmptyLines = false

            val parser = CsvParser(settings)
            val inputStream: InputStream? =
                CSVValidator::class.java.getResourceAsStream("/$filename")

            val inputStreamReader = InputStreamReader(inputStream!!)
            val bufferedReader = BufferedReader(inputStreamReader)

            parser.beginParsing(bufferedReader)

            var record: Array<String>?
            while (parser.parseNext()!=null){

                record = arrayOf(parser.context.currentParsedContent())
                //convert record to csv row
                val csvRecord = record.joinToString(",")
                // get fields
                for (field in csvRecord.split(",")){
                    val f = field
                }
            }

            // Close resources
            parser.stopParsing()
            bufferedReader.close()
            inputStreamReader.close()
            inputStream.close()
        }catch (e: Exception){
            e.printStackTrace()
        }
    }

    fun  deepHavenCSVParser(filename: String) {
        try {
            val inputStream = object {}.javaClass.getResourceAsStream(filename)
            val specs = CsvSpecs.builder()
                .ignoreEmptyLines(false)
                .hasHeaderRow(true)
                .ignoreSurroundingSpaces(false)
                .delimiter(',')
                .build()

            val result = DeephavenCSVReader.read(specs, inputStream, SinkFactory.arrays())
            for (col in result) {
                if (col.data() is Array<*>) {
                    for (element in col.data() as Array<*>) {
                        val field = element
                    }
                } else if (col.data() is IntArray) {
                    for (element in col.data() as IntArray) {
                        val field = element
                    }
                } else if (col.data() is StringArray) {
                    for (element in col.data() as String) {
                        val field = element
                    }
                }
            }
        }catch(e: Exception){
            e.printStackTrace()
        }
    }

    fun fastCSVParser(filename: String){

        val pathToCSV: Path? = CSVValidator::class.java.getResource("/$filename")?.toURI()?.let { Paths.get(it) }
        try {
            val reader = FastCSVReader.builder().ofCsvRecord(pathToCSV).iterator()
            while (reader.hasNext()) {
                val csvRecord: CsvRecord = reader.next()
                for (field in csvRecord.fields){
                    val f = field
                }
            }

        } catch (e: Exception){
            e.printStackTrace()
        }

    }
    fun superCSVParser(filename: String) {
        val csvInputSteam: InputStream? = CSVValidator::class.java.getResourceAsStream("/$filename")
        val reader = csvInputSteam?.let { InputStreamReader(it) }
        val csvReader = CsvListReader(reader, CsvPreference.STANDARD_PREFERENCE) // you can also configure with Tab

        try {

            var rowRecord: List<String>?
            while (csvReader.read().also { rowRecord = it } != null) {
                rowRecord?.let {
                    for (field in rowRecord!!){
                        val f = field
                    }
                }
            }
        } finally {
            csvReader.close()
            reader?.close()
            csvInputSteam?.close()
        }
    }
    fun jacksonCSVParser(filename: String){
        val mapper = CsvMapper() // create mapper to provide reading or writing functionality

        val csvFile = File("src/main/resources/$filename")
        val iter = mapper
            .readerForListOf(String::class.java)
            .with(com.fasterxml.jackson.dataformat.csv.CsvParser.Feature.WRAP_AS_ARRAY)
            .readValues<List<String>>(csvFile)// converts each row list of strings

        while (iter.hasNext()) {
            val row = iter.next()
            for (field in row){
                val f = field
            }
        }
    }
}


fun main(){
    val startTime = LocalDateTime.now()
    val validate = CSVValidator()
    repeat(1){
        validate.superCSVParser("file-with-headers-100-rows.csv")
    }

    val endTime = LocalDateTime.now()
    val duration = java.time.Duration.between(startTime, endTime).toMillis()

    println("The Total processing time: $duration milliseconds")

}

