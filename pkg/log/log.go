// Package log provides logging functionality for the application
package log

import (
    "fmt"
    "os"
    "time"
)

var logFile *os.File

// Init initializes the log file
func Init(filePath string) error {

    var err error
    logFile, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        return fmt.Errorf("could not open log file %v", err)
    }
    return nil
}

// Logs logs a message with a timestamp
func Log(message string) {

    if logFile != nil {
        timestamp := time.Now().Format("2006-01-02 15:04:05")
        logFile.WriteString(fmt.Sprintf("%s: %s\n", timestamp, message))
    }
}

// Close closes the log file
func Close() {
    if logFile != nil {
        logFile.Close()
    }
}
