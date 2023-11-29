import commons.ApacheCommonsCSVParser;
import commons.OpenCSVParser;
import parse.Iparser;

public class Main {
    public static void main(String[] args) {
        Iparser apacheParser = new ApacheCommonsCSVParser();
        Iparser openCsvParser = new OpenCSVParser();
    }
}
