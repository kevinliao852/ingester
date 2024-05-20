package store

import (
	"fmt"
	"log"
	"os"
	"time"
)

type AppendLogger struct {
	logger   *log.Logger
	fileName string
}

func NewAppendLogger(fileName string) *AppendLogger {
	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening log file for writing:", err)
	}

	return &AppendLogger{
		logger:   log.New(logFile, "", 0),
		fileName: fileName,
	}
}

func (al *AppendLogger) LogMessage(message string) {
	al.logger.Println(time.Now().Format(time.RFC3339), message)
}

func (al *AppendLogger) RetrieveLogs(cursor *int) {

	logFile, err := os.Open(al.fileName)
	if err != nil {
		log.Fatal("Error opening log file for reading:", err)
	}
	defer logFile.Close()

	// Seek to the cursor position in the file
	_, err = logFile.Seek(int64(*cursor), 0)
	if err != nil {
		log.Fatal("Error seeking to cursor position:", err)
	}

	// Read and print logs from the cursor position
	buffer := make([]byte, 1024)
	for {
		n, err := logFile.Read(buffer)
		if err != nil {
			break
		}

		fmt.Print(string(buffer[:n]))

		// Update the cursor position
		*cursor += n
	}
}
