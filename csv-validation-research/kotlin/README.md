# CSV Parsing Performance Metrics

| **Library**       | **Number of executions** | **Total Time(MS)** | **Total CPU Time(MS)** | **Memory Allocation** | **Filename** | **Source** |
|-------------------|--------------------------|---------------------|------------------------|---------------------------|--------------|-----------------|
| **OpenCSV**      | 100|98|144|34.75MB|file-with-headers-100-rows.csv|[OpenCSV](https://opencsv.sourceforge.net/)|
| **Apache Commons**| 100|104|44|35.76MB|file-with-headers-100-rows.csv|[Apache Commons](https://commons.apache.org/proper/commons-csv/)|
| **Deephaven**     | 100|NA|NA|NA|NA|[DeepHaven](https://github.com/deephaven/deephaven-csv)|
| **Univocity Parser**|100|292|276|226.4MB|file-with-headers-100-rows.csv|[Univocity Parsers](https://www.baeldung.com/java-univocity-parsers)|
| **Kotlin-csv**    |100|158|99|31.54MB|file-with-headers-100-rows.csv||
| **OpenCSV**      |5000|1464|1400|776.16MB|file-with-headers-100-rows.csv||
| **Apache Commons**|5000|2651|2393|836.39MB|file-with-headers-100-rows.csv||
| **Deephaven**     |5000|NA|NA|NA|NA||
| **Univocity Parser**|5000|4817|3960|11.44GB|file-with-headers-100-rows.csv||
| **Kotlin-csv**    |5000|3496|2667|1.53GB|file-with-headers-100-rows.csv||
| **OpenCSV**      |100|2293|2427|1.13GB|file-with-headers-10000-rows.csv||
| **Apache Commons**|100|3666|4502|1.12GB|file-with-headers-10000-rows.csv||
| **Deephaven**     |100|NA|NA|NA|file-with-headers-10000-rows.csv||
| **Univocity Parser**|100|1448|1236|1.21GB|file-with-headers-10000-rows.csv||
| **Kotlin-csv**    |100|5605|6346|2.78GB|file-with-headers-10000-rows.csv||
| **OpenCSV**      |5000|70190|81432|57.23GB|file-with-headers-10000-rows.csv||
| **Apache Commons**|5000|202819|191314|55.29GB|file-with-headers-10000-rows.csv||
| **Deephaven**     |5000|NA|NA|NA|file-with-headers-10000-rows.csv||
| **Univocity Parser**|5000|53316|45934|59.45GB|file-with-headers-10000-rows.csv||
| **Kotlin-csv**    |5000|279743|356336|138.72GB|file-with-headers-10000-rows.csv||


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