# CSV Parsing Performance Metrics

| **Library**       | **Number of executions** | **Total Time(MS)** | **Total CPU Time(MS)** | **Memory Allocation** | **Filename** | **Source** |
|-------------------|--------------------------|---------------------|------------------------|---------------------------|--------------|-----------------|
| **OpenCSV**      | 100|151|133|44.99MB|file-with-headers-100-rows.csv|[OpenCSV](https://opencsv.sourceforge.net/)|
| **Apache Commons**| 100|170|156|46.91MB|file-with-headers-100-rows.csv|[Apache Commons](https://commons.apache.org/proper/commons-csv/)|
| **Univocity Parser**|100|433|408|255.84MB|file-with-headers-100-rows.csv|[Univocity Parsers](https://www.baeldung.com/java-univocity-parsers)|
| **Kotlin-csv**    |100|211|122|62.46MB|file-with-headers-100-rows.csv|[Kotlin CSV](https://github.com/doyaaaaaken/kotlin-csv)|
| **FastCSV**    |100|180|120|45.08MB|file-with-headers-100-rows.csv|[Fast CSV](https://github.com/osiegmar/FastCSV)|
| **SuperCSV**    |100|121|110|44.5MB|file-with-headers-100-rows.csv|[Super CSV](https://github.com/super-csv/super-csvj)|
| **JacksonCSV**    |100|492|348|55.78MB|file-with-headers-100-rows.csv|[Jackson CSV](https://github.com/FasterXML/jackson-dataformats-text/tree/2.18/csv)|
| **OpenCSV**      |5000|2392|951|1.08GB|file-with-headers-100-rows.csv|[OpenCSV](https://opencsv.sourceforge.net/)|
| **Apache Commons**|5000|4024|2542|1.14GB|file-with-headers-100-rows.csv|[Apache Commons](https://commons.apache.org/proper/commons-csv/)|
| **Univocity Parser**|5000|5407|2477|12.48GB|file-with-headers-100-rows.csv|[Univocity Parsers](https://www.baeldung.com/java-univocity-parsers)|
| **Kotlin-csv**    |5000|5169|3593|1.84GB|file-with-headers-100-rows.csv|[Kotlin CSV](https://github.com/doyaaaaaken/kotlin-csv)|
| **FastCSV**    |5000|1898|637|1.05GB|file-with-headers-100-rows.csv|[Fast CSV](https://github.com/osiegmar/FastCSV)|
| **SuperCSV**    |5000|2671|1335|1.01GB|file-with-headers-100-rows.csv|[Super CSV](https://github.com/super-csv/super-csvj)|
| **JacksonCSV**    |5000|2700|1232|899.72MB|file-with-headers-100-rows.csv|[Jackson CSV](https://github.com/FasterXML/jackson-dataformats-text/tree/2.18/csv)|
| **OpenCSV**      |100|1804|1665|1.74GB|file-with-headers-10000-rows.csv|[OpenCSV](https://opencsv.sourceforge.net/)|
| **Apache Commons**|100|5204|5084|1.72GB|file-with-headers-10000-rows.csv|[Apache Commons](https://commons.apache.org/proper/commons-csv/)|
| **Univocity Parser**|100|1920|1732|3.13GB|file-with-headers-10000-rows.csv|[Univocity Parsers](https://www.baeldung.com/java-univocity-parsers)|
| **Kotlin-csv**    |100|6830|6673|3.39GB|file-with-headers-10000-rows.csv|[Kotlin CSV](https://github.com/doyaaaaaken/kotlin-csv)|
| **FastCSV**    |100|864|763|1.45GB|file-with-headers-100-rows.csv|[Fast CSV](https://github.com/osiegmar/FastCSV)|
| **SuperCSV**    |100|2131|1984|1.62GB|file-with-headers-100-rows.csv|[Super CSV](https://github.com/super-csv/super-csvj)|
| **JacksonCSV**    |100|1665|1496|1.45GB|file-with-headers-100-rows.csv|[Jackson CSV](https://github.com/FasterXML/jackson-dataformats-text/tree/2.18/csv)|
| **OpenCSV**      |5000|103020|94175|86.88GB|file-with-headers-10000-rows.csv|[OpenCSV](https://opencsv.sourceforge.net/)|
| **Apache Commons**|5000|260236|252665|84.94GB|file-with-headers-10000-rows.csv|[Apache Commons](https://commons.apache.org/proper/commons-csv/)|
| **Univocity Parser**|5000|91813|83598|155.71GB|file-with-headers-10000-rows.csv|[Univocity Parsers](https://www.baeldung.com/java-univocity-parsers)|
| **Kotlin-csv**    |5000|367716|360520|168.38GB|file-with-headers-10000-rows.csv|[Kotlin CSV](https://github.com/doyaaaaaken/kotlin-csv)|
| **FastCSV**    |5000|29210|34311|71.02GB|file-with-headers-100-rows.csv|[Fast CSV](https://github.com/osiegmar/FastCSV)|
| **SuperCSV**    |5000|122951|116569|80.09GB|file-with-headers-100-rows.csv|[Super CSV](https://github.com/super-csv/super-csvj)|
| **JacksonCSV**    |5000|75316|70233|70.3GB|file-with-headers-100-rows.csv|[Jackson CSV](https://github.com/FasterXML/jackson-dataformats-text/tree/2.18/csv)|

# OpenCSV RFC-4180 Compliance Results
| Rule | Example | Compliant (Yes/No) |
|------|---------|--------------------|
|Each record is located on a separate line, delimited by a line break (CR)|"Name,Email\rJane Doe,johndoe@example.com"|Yes|
|Each record is located on a separate line, delimited by a line break (LF)|"Name,Email\nJane Doe,johndoe@example.com"|Yes|
|Each record is located on a separate line, delimited by a line break (CRLF)|"Name,Email\r\nJane Doe,johndoe@example.com"|Yes|
|Within the header and within each record, there may be one or more fields, separated by commas|"Name,Email\rJane Doe,johndoe@example.com"|Yes|
|Each record should contain the same number of fields throughout the file|"Name,Email\nJane Doe,johndoe@example.com"|No|
|Spaces are considered part of a field and should not be ignored|"Name,Email\r\nJane       Doe,johndoe@example.com"|Yes|
|The last field in the record must not be followed by a comma|"Name,Email\r\nJaneDoe,johndoe@example.com,"|No|
|If fields are not enclosed with double quotes, then double quotes may not appear inside the fields|"Name,Email\nJohn,john@example.com\nJane,jane@example.com\nAlice,alice@example.com"|Yes
|If double-quotes are used to enclose fields, then a double-quote appearing inside a field must be escaped by preceding it with another double quote|"Name,Email\nJohn,john@example.com\nJane,\"jane\"\"@example.com\"\nAlice,alice@example.com"|Yes
|Fields with line breaks inside quotes|"Name,Email\nJohn,john@example.com\nJane,\"jane   \n@example.com\"\nAlice,alice@example.com"|Yes
|Fields with commas inside quotes|"Name,Email\nJohn,john@example.com\nJane,\"jane,@example.com\"\nAlice,alice@example.com|Yes