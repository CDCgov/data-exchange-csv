package parse;

import java.nio.file.Path;

public interface Iparser {
    void parseCSVFileWithHeader(Path filePath, String[] header);
}
