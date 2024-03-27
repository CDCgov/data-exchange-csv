using Microsoft.VisualBasic.FileIO;
using System.Diagnostics;

internal class Program
{
    private static void Main(string[] args)
    {
        string filePath = "data/file-with-headers-100-rows.csv";

        Stopwatch totalTimeStopwatch = Stopwatch.StartNew();
        //change the number of iteration as per needs
        for (int i = 0; i < 100; i++)
        {
            Stopwatch iterationStopwatch = Stopwatch.StartNew();
            _ = ReadCsv(filePath);
            iterationStopwatch.Stop();
            Console.WriteLine($"Iteration {i + 1}: Time taken: {iterationStopwatch.ElapsedMilliseconds} ms");
        }
        totalTimeStopwatch.Stop();
        Console.WriteLine($"Total time taken for 100 iterations: {totalTimeStopwatch.ElapsedMilliseconds} ms");

        static List<string[]> ReadCsv(string filePath)
        {
            List<string[]> data = new List<string[]>();
            List<string> errors = new List<string>();
            using TextFieldParser parser = new TextFieldParser(filePath);
            parser.TextFieldType = FieldType.Delimited;
            parser.SetDelimiters(",");
            while (!parser.EndOfData)
            {
                string[] fields = parser.ReadFields();
                if (fields !=null){
                    data.Add(fields);
                }else{
                    errors.Add($"Error was found in row {parser.LineNumber}");
                }

                //you can add custom error handling here testing against RFC4180 compliance
            }
            return data;
        }
    }
}