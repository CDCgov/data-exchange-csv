package test

import com.azure.cosmos.*
import com.microsoft.azure.functions.ExecutionContext
import com.microsoft.azure.functions.HttpRequestMessage
import com.microsoft.azure.functions.HttpStatus


import org.mockito.Mockito.mock

import org.testng.Assert.*

import org.testng.annotations.BeforeMethod
import org.testng.annotations.Test

import java.util.*

import gov.cdc.functions.HealthCheckFunction


class HealthCheckFunctionTest {

    private lateinit var request: HttpRequestMessage<Optional<String>>
    private lateinit var context: ExecutionContext

    @BeforeMethod
    fun setUp() {
        // Initialize any mock objects or dependencies needed for testing
        request = mock(HttpRequestMessage::class.java) as HttpRequestMessage<Optional<String>>
        context = mock(ExecutionContext::class.java)
    }

     @Test
    fun testFailureStatusBack() {
        // Create a HealthCheckFunction instance
        val healthCheckFunction = HealthCheckFunction()

        // call HealthCheckFunction
//        val response = healthCheckFunction.run(request, context, mockCosmosClient)

//        assert(response == HttpStatus.INTERNAL_SERVER_ERROR)
    }
}