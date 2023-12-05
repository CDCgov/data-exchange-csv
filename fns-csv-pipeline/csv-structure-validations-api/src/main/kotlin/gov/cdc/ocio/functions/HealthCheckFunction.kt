package gov.cdc.ocio.functions

import com.microsoft.azure.functions.ExecutionContext
import com.microsoft.azure.functions.HttpRequestMessage
import com.microsoft.azure.functions.HttpStatus
import java.util.*
import com.microsoft.azure.functions.HttpMethod
import com.microsoft.azure.functions.HttpResponseMessage
import com.microsoft.azure.functions.annotation.AuthorizationLevel
import com.microsoft.azure.functions.annotation.HttpTrigger
class HealthCheckFunction {

    fun run(
        @HttpTrigger(name="req",
                methods = [HttpMethod.GET],
                authLevel = AuthorizationLevel.ANONYMOUS)
        request: HttpRequestMessage<Optional<String>>, context: ExecutionContext):
            HttpResponseMessage {
        context.logger.info("Health check request received.");

        return if (performHealthCheck()) {
            request.createResponseBuilder(HttpStatus.OK).body("Health Check Passed").build()
        } else {
            request.createResponseBuilder(HttpStatus.INTERNAL_SERVER_ERROR).body("Health Check Failed").build();
        }
    }

    private fun performHealthCheck(): Boolean {
        return true;
    }
}