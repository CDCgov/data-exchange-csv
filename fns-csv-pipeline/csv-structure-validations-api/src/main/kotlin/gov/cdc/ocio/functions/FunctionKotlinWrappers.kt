package gov.cdc.ocio.functions
import com.azure.storage.blob.BlobClientBuilder
import com.microsoft.azure.functions.*
import com.microsoft.azure.functions.HttpMethod
import com.microsoft.azure.functions.annotation.AuthorizationLevel
import com.microsoft.azure.functions.annotation.FunctionName
import com.microsoft.azure.functions.annotation.HttpTrigger
import java.util.Optional

class FunctionKotlinWrappers {

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
//        val blobClient = BlobClientBuilder()
//            .endpoint(System.getenv("DexStorageEndpoint"))
//            .connectionString(System.getenv("DexStorageConnectionString"))
//            .containerName(System.getenv("TusHooksContainerName"))
//            .blobName(System.getenv("DestinationsFileName"))
//            .buildClient()
        return HealthCheckFunction().run(request, context);
    }

//    @FunctionName("ValidateRawCSV")
//    fun validateRawCSV(
//        @HttpTrigger(
//            name = "req",
//            methods = [HttpMethod.GET],
//            route = "validateRawCsv",
//            authLevel = AuthorizationLevel.FUNCTION
//        ) request: HttpRequestMessage<Optional<String>>,
//        context: ExecutionContext
//    ) : HttpResponseMessage {
//        val blobClient = BlobClientBuilder()
//            .endpoint(System.getenv("DexStorageEndpoint"))
//            .connectionString(System.getenv("DexStorageConnectionString"))
//            .containerName(System.getenv("TusHooksContainerName"))
//            .blobName(System.getenv("DestinationsFileName"))
//            .buildClient()
//
//        return ValidateMultiPartFileUpload().run(request, context, blobClient)
//    }
//
//    @FunctionName("ValidateMultiPartFileUpload")
//    fun validateMultiPartFileUpload(
//            @HttpTrigger(
//                    name = "req",
//                    methods = [HttpMethod.GET],
//                    route = "destination",
//                    authLevel = AuthorizationLevel.FUNCTION
//            ) request: HttpRequestMessage<Optional<String>>,
//            context: ExecutionContext
//    ) : HttpResponseMessage {
//        val blobClient = BlobClientBuilder()
//                .endpoint(System.getenv("DexStorageEndpoint"))
//                .connectionString(System.getenv("DexStorageConnectionString"))
//                .containerName(System.getenv("TusHooksContainerName"))
//                .blobName(System.getenv("DestinationsFileName"))
//                .buildClient()
//
//        return ValidateMultiPartFileUpload().run(request, context, blobClient)
//    }
}
