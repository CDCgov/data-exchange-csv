package gov.cdc.ocio.functions

import com.microsoft.azure.functions.ExecutionContext
import com.microsoft.azure.functions.HttpRequestMessage
import com.microsoft.azure.functions.HttpResponseMessage
import com.microsoft.azure.functions.HttpStatus
import com.opencsv.CSVParserBuilder
import com.opencsv.CSVReaderBuilder
import com.opencsv.enums.CSVReaderNullFieldIndicator
import com.opencsv.validators.RowMustHaveSameNumberOfColumnsAsFirstRowValidator
import io.github.oshai.kotlinlogging.KotlinLogging
import java.io.StringReader


class ValidateRawCSV {
    fun run(request: HttpRequestMessage<String>, context: ExecutionContext): HttpResponseMessage {

        val logger = KotlinLogging.logger {}
        return  if (request.body.isNullOrEmpty()) {
            request.createResponseBuilder(HttpStatus.BAD_REQUEST).body("CSV data not provided.").build()
        } else {
            val csvContent = request.body
            return if(isValidCSVStructure(csvContent)) {
                request.createResponseBuilder(HttpStatus.OK).body("CSV Structure is valid.").build()
            } else {
                request.createResponseBuilder(HttpStatus.BAD_REQUEST).body("CSV structure is invalid.").build()
            }
        }
    }

    fun isValidCSVStructure(csvContent: String): Boolean {
        try {
            StringReader(csvContent).use {
                    stringReader ->
                val csvParser = CSVParserBuilder().withSeparator(',').build()
                val rowValidator = RowMustHaveSameNumberOfColumnsAsFirstRowValidator();
                val csvReader = CSVReaderBuilder(stringReader)
                    .withCSVParser(csvParser)
                    .withFieldAsNull(CSVReaderNullFieldIndicator.BOTH)
                    .withKeepCarriageReturn(true)
                    .withRowValidator(rowValidator)
                    .build()

                var line: Array<String?>?
                while (csvReader.readNext().also { line = it } != null) {
                    continue
                }
                return true
            }
        } catch(e: Exception) {
            return false;
        }
    }
}
