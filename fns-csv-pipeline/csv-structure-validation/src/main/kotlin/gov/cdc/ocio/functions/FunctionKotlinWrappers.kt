package gov.cdc.ocio.functions
import com.azure.storage.blob.BlobClientBuilder
import com.microsoft.azure.functions.*
import com.microsoft.azure.functions.HttpMethod
import com.microsoft.azure.functions.annotation.AuthorizationLevel
import com.microsoft.azure.functions.annotation.FunctionName
import com.microsoft.azure.functions.annotation.HttpTrigger
import java.util.Optional

public class FunctionKotlinWrappers {

    @FunctionName("HealthCheck")
    fun healthCheck(
            @HttpTrigger(
                    name = "req",
                    methods = [HttpMethod.GET],
                    route = "status/health",
                    authLevel = AuthorizationLevel.ANONYMOUS
            ) request: HttpRequestMessage<Optional<String>>,
            context: ExecutionContext
    ): HttpResponseMessage {
        return HealthCheckFunction().run(request, context);
    }

    @FunctionName("ValidateRawCSV")
    fun validateRawCSV(
        @HttpTrigger(
            name = "req",
            methods = [HttpMethod.GET, HttpMethod.POST],
            route = "validateRawCsv",
            authLevel = AuthorizationLevel.ANONYMOUS
        ) request: HttpRequestMessage<String>,
        context: ExecutionContext
    ) : HttpResponseMessage {
        return ValidateRawCSV().run(request, context);
    }

    @FunctionName("ValidateMultiPartCSVUpload")
    fun ValidateMultiPartCSVUpload(
            @HttpTrigger(
                    name = "req",
                    methods = [HttpMethod.POST],
                    route = "validateMultiPartCSVUpload",
                    authLevel = AuthorizationLevel.ANONYMOUS
            ) request: HttpRequestMessage<Optional<String>>,
            context: ExecutionContext
    ) : HttpResponseMessage {
        return ValidateMultiPartCSVUpload().run(request, context)
    }
}
