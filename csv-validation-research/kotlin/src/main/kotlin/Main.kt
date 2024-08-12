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
            val rows = mutableListOf<Map<Int, String>>()
            if (file != null) {
                csvReader().open(file) {
                    readAllAsSequence().forEach { row: List<String> ->
                        val rowObject = mutableMapOf<Int,String>()
                        var i =0
                        for (field in row){
                            rowObject[i] = field
                            i +=1
                        }
                        rows.add(rowObject)
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
            val rows = mutableListOf<Map<Int, String>>()
            var record: Array<String>?
            while (builder.readNext().also {record=it} !=null){
                val row = mutableMapOf<Int,String>()
                var i = 0
                for (field in record!!){
                    row[i] =field
                    i +=1
                }
                rows.add(row)
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
            val rows = mutableListOf<Map<Int, String>>()
            val csvParser = CSVParser(bufferedReader,CSVFormat.RFC4180)
            for (csvRecord in csvParser){
                val row = mutableMapOf<Int,String>()
                var i = 0
                for (field in csvRecord){
                    row[i] = field
                    i+=1
                }
                rows.add(row)
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
            val rows = mutableListOf<Map<Int, String>>()
            parser.beginParsing(bufferedReader)

            var record: Array<String>?
            while (parser.parseNext()!=null){

                record = arrayOf(parser.context.currentParsedContent())
                val row = mutableMapOf<Int, String>()
                //convert record to csv row
                val csvRecord = record.joinToString(",")

                // set index to 0
                var i = 0
                // get fields
                for (field in csvRecord.split(",")){
                    row[i] = field
                    i +=1
                }
                rows.add(row)
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
        val rows = mutableListOf<Map<Int, String>>()
        try {
            val reader = FastCSVReader.builder()
                .fieldSeparator(',')
                .quoteCharacter('"')
                .ignoreDifferentFieldCount(false)
                .ofCsvRecord(pathToCSV)
                .iterator()

            while (reader.hasNext()) {
                val row = mutableMapOf<Int,String>()
                val csvRecord: CsvRecord = reader.next()
                var i =0
                for (field in csvRecord.fields){
                    row[i] = field
                    i+=1
                }
                rows.add(row)
            }
        } catch (e: Exception){
            e.printStackTrace()
        }

    }
    fun superCSVParser(filename: String) {
        val csvInputSteam: InputStream? = CSVValidator::class.java.getResourceAsStream("/$filename")
        val reader = csvInputSteam?.let { InputStreamReader(it) }
        val csvReader = CsvListReader(reader, CsvPreference.STANDARD_PREFERENCE) // you can also configure with Tab, standard is CSV
        val rows = mutableListOf<Map<Int, String>>()
        try {

            var rowRecord: List<String>?
            while (csvReader.read().also { rowRecord = it } != null) {
                var i = 0
                val row = mutableMapOf<Int,String>()
                rowRecord?.let {
                    for (field in rowRecord!!){
                        row[i] = field
                        i+=1
                    }
                }
                rows.add(row)
            }
        } finally {
            csvReader.close()
            reader?.close()
            csvInputSteam?.close()
        }
    }
    fun jacksonCSVParser(filename: String){
        val mapper = CsvMapper()

        val csvFile = File("src/main/resources/$filename")
        val rows = mutableListOf<Map<Int, String>>()
        val iter = mapper
            .readerForListOf(String::class.java)
            .with(com.fasterxml.jackson.dataformat.csv.CsvParser.Feature.WRAP_AS_ARRAY)
            .readValues<List<String>>(csvFile)

        while (iter.hasNext()) {
            val row = iter.next()
            val rowObject = mutableMapOf<Int,String>()
            var i = 0
            for (field in row){
                rowObject[i] = field
                i+=1
            }
            rows.add(rowObject)
        }
    }
}


fun main(){
    val startTime = LocalDateTime.now()
    val validate = CSVValidator()
    var totalDuration: Long =0
    val numOfRuns = 60
        repeat(numOfRuns){
            val innerStartTime = LocalDateTime.now()
            validate.validateWithKotlinCSV("file-with-headers-10000-rows.csv")
            val innerEndTime = LocalDateTime.now()
            totalDuration += java.time.Duration.between(innerStartTime, innerEndTime).toMillis()
    }

    val endTime = LocalDateTime.now()

    val totalProcessingTime = java.time.Duration.between(startTime, endTime).toMillis()
    val averageDuration = totalDuration / numOfRuns


    println("Total processing time for $numOfRuns runs: $totalProcessingTime  milliseconds")
    println("Average processing time per run: $averageDuration milliseconds")

}

