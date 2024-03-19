import com.github.doyaaaaaken.kotlincsv.dsl.csvReader
import com.opencsv.CSVParserBuilder
import com.opencsv.CSVReaderBuilder
import com.sun.jna.StringArray

import com.univocity.parsers.csv.CsvParser
import com.univocity.parsers.csv.CsvParserSettings

import io.deephaven.csv.CsvSpecs
import io.deephaven.csv.reading.CsvReader
import io.deephaven.csv.sinks.SinkFactory

import org.apache.commons.csv.CSVFormat
import org.apache.commons.csv.CSVParser

import java.io.BufferedReader
import java.io.File
import java.io.InputStream
import java.io.InputStreamReader

import java.time.LocalDateTime
import java.time.temporal.ChronoUnit
import java.util.*


class CSVValidator {

    fun validateWithKotlinCSV(filename: String) {
        val startTime = LocalDateTime.now()
         val endTime: LocalDateTime?

        try {
            val resourceURL = object {}.javaClass.getResource(filename)
            val file: File? = resourceURL?.toURI()?.let { File(it) }

            if (file != null) {
                csvReader().open(file) {
                    readAllAsSequence().forEach { row: List<String> ->
                        for (field in row){
                            continue
                        }
                    }
                }
            }
        }catch (e: Exception){
            println (e.printStackTrace())
        }finally {
            endTime = LocalDateTime.now()
        }
        val totalTime= ChronoUnit.MILLIS.between(startTime,endTime)
        println ("Total processing time $filename took in milliseconds: $totalTime")
    }
    fun validateWithOpenCSV(filename:String): List<String> {

        val startTime = LocalDateTime.now()
        val endTime: LocalDateTime?
        val parsedData = mutableListOf<String>()
        try {
            val inputStream: InputStream? =
                CSVValidator::class.java.getResourceAsStream("/$filename")

            val csvParserBuilder = CSVParserBuilder()
                .withSeparator(',')
                .withQuoteChar('"')
                .withIgnoreLeadingWhiteSpace(false)
                .withIgnoreQuotations(false)
                .withEscapeChar('\r')

                .build()
            val builder = CSVReaderBuilder(inputStream?.let { InputStreamReader(it) })
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
        }catch (e: Exception) {
            e.printStackTrace()
        }finally {
            endTime = LocalDateTime.now()
        }

        val totalTime= ChronoUnit.MILLIS.between(startTime,endTime)
        println ("Total processing time $filename took in milliseconds: $totalTime")
        return parsedData
    }
    fun validateWithApacheCommons(filename:String) {
        val startTime = LocalDateTime.now()
        val endTime: LocalDateTime?
        try {
            val inputStream: InputStream? =
                CSVValidator::class.java.getResourceAsStream("/$filename")
            val inputStreamReader = InputStreamReader(inputStream!!)
            val bufferedReader =  BufferedReader(inputStreamReader)

            val csvParser = CSVParser(bufferedReader,CSVFormat.RFC4180)
            println (csvParser.records)


            for (csvRecord in csvParser){
                println (csvRecord)
            }

            inputStream.close()
            inputStreamReader.close()
            bufferedReader.close()
            csvParser.close()
        }catch (e: Exception) {
            println ("Error parsing CSV: ${e.message}")
            e.printStackTrace()
        }finally {
            endTime = LocalDateTime.now()
        }

        val totalTime= ChronoUnit.MILLIS.between(startTime,endTime)
        println ("Total processing time $filename took in milliseconds: $totalTime")

    }

    fun univocityCSVParser(filename: String){
        val startTime = LocalDateTime.now()
        val endTime: LocalDateTime?
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
            while (parser.parseNext() !=null){
                record = arrayOf(parser.context.currentParsedContent())

                for (row in record){
                    println (row)
                }
            }

            // Close resources
            parser.stopParsing()
            bufferedReader.close()
            inputStreamReader.close()
            inputStream.close()
        }catch (e: Exception){
            e.printStackTrace()
        }finally {
            endTime = LocalDateTime.now()
        }
        val totalTime= ChronoUnit.MILLIS.between(startTime,endTime)
        println ("Total processing time $filename took  in milliseconds: $totalTime")
    }

    fun  deepHavenCSVParser(filename: String) {
        val startTime = LocalDateTime.now()
        val endTime: LocalDateTime?
        try {
            val inputStream = object {}.javaClass.getResourceAsStream(filename)
            val specs = CsvSpecs.builder()
                .ignoreEmptyLines(false)
                .hasHeaderRow(true)
                .ignoreSurroundingSpaces(false)
                .delimiter(',')
                .build()

            val result = CsvReader.read(specs, inputStream, SinkFactory.arrays())


            //val numOfRows = result.numRows() //returns number of rows

            for (col in result) {
                println("The data type for ${col.name()} is ${col.dataType()}")
                if (col.data() is Array<*>) {
                    println("Column: ${col.name()} fields:")
                    for (element in col.data() as Array<*>) {
                        println(element)
                    }
                } else if (col.data() is IntArray) {
                    println("Column: ${col.name()} fields:")
                    for (element in col.data() as IntArray) {
                        println(element)
                    }
                } else if (col.data() is StringArray) {
                    for (element in col.data() as String) {
                        println(element)
                    }
                }
            }
        }catch(e: Exception){
            e.printStackTrace()
        }   finally {
            endTime = LocalDateTime.now()
        }
        val totalTime= ChronoUnit.MILLIS.between(startTime,endTime)
        println ("Total processing time $filename took  in milliseconds: $totalTime")
    }
}



fun main(){
    /*
    to run with the profiler, you can surround each function logic with outer
    for loop with number of iterations, then
    val startTime = LocalDateTime.now()
    val validate = CSVValidator()
    validate.validateWithOpenCSV("file-with-headers-100-rows.csv")
    val endTime: LocalDateTime?
     */
    val validate = CSVValidator()

    validate.validateWithOpenCSV("file-with-headers-100-rows.csv")
    validate.univocityCSVParser("file-with-headers-100-rows.csv")
    validate.deepHavenCSVParser("file-with-headers-100-rows.csv")
    validate.validateWithKotlinCSV("file-with-headers-100-rows.csv")
    validate.validateWithApacheCommons("file-with-headers-100-rows.csv")
    validate.validateWithOpenCSV("file-with-headers-10000-rows.csv")
    validate.univocityCSVParser("file-with-headers-10000-rows.csv")
    validate.deepHavenCSVParser("file-with-headers-10000-rows.csv")
    validate.validateWithKotlinCSV("file-with-headers-10000-rows.csv")
    validate.validateWithApacheCommons("file-with-headers-10000-rows.csv")


}