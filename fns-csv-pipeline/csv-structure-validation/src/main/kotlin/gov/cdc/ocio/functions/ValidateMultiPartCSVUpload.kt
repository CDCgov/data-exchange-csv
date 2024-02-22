package gov.cdc.ocio.functions

import com.microsoft.azure.functions.ExecutionContext
import com.microsoft.azure.functions.HttpRequestMessage
import com.microsoft.azure.functions.HttpResponseMessage
import com.microsoft.azure.functions.HttpStatus
import io.ktor.http.content.*
import io.ktor.server.application.*
import io.ktor.server.engine.*
import io.ktor.server.netty.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import java.io.InputStreamReader
import java.util.*
import java.util.regex.Pattern


class ValidateMultiPartCSVUpload {
    fun run(request: HttpRequestMessage<Optional<String>>,
            context: ExecutionContext): HttpResponseMessage {
        return try {
           val server = embeddedServer(Netty, port = 0) {
               routing {
                   post("/validatePartCsv") {
                       val multiPart = call.receiveMultipart()
                       val csvPart = multiPart.readPart() ?: throw Exception("CSV file not found in the request")
                       val isValidCSV = validateCSV(csvPart)

                       if (isValidCSV) {
                           call.respond(HttpStatus.OK)
                       } else {
                            call.respond(HttpStatus.BAD_REQUEST)
                       }
                   }
               }
           }
            server.start(wait = false)
            request.createResponseBuilder(HttpStatus.OK).body("Function is running..").build()
        } catch (e: Exception) {
            request.createResponseBuilder(HttpStatus.INTERNAL_SERVER_ERROR)
                .body("Error processing the request")
                .build()
        }
    }

    private fun validateCSV(csvPart: PartData) : Boolean {
        if (csvPart is PartData.FileItem) {
            // validate csv content for the uploaded part
            val csvContent = InputStreamReader(csvPart.streamProvider()).readText()
            // example: Ensure each row has exactly three columns
            val csvPattern = Pattern.compile(("^[^,]+,[^,]+,[^,]+\$"))
            return csvContent.lines().all {
                csvPattern.matcher(it).matches()
            }
        }
        csvPart.dispose()
        return false
    }
}
