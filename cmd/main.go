// Package main is the entry point for the application
// it defines the command-line flags and parses them
package main

import (

    "flag"
    "fmt"
    "os"

    "github.com/riasy7/backup-cli/pkg/backup"
    "github.com/riasy7/backup-cli/pkg/log"
)

func main() {

    // Init logging
    if err := log.Init("backup.log"); err != nil {
        fmt.Println("Error initializing log:", err)
        os.Exit(1)
    }
    defer log.Close()

    // Define command-line flags
    source := flag.String("source", "", "Source file or directory to back up")
    dest :=  flag.String("dest", "", "Destination directory for back up")
    help := flag.Bool("help", false, "Display help")

    // Parse flags
    flag.Parse()

    // Display help
    if *help {
        fmt.Println("Usage: backup-cli --source=<source> --dest=<destination>")
        fmt.Println("Flags:")
        flag.PrintDefaults()
        os.Exit(0)
    }

    // Check if source is provided
    if *source == "" {
        exitWithError("Error: Source path is requried")
    }

    // Check if destination is provided
    if *dest == "" {
        exitWithError("Error: Destination path is requried")
    }


    // handle file copying to dest
    if err := backup.HandleFiles(*source, *dest); err != nil {
        exitWithError(err.Error())
    }

    // log successful back up + user also needs to see
    log.Log(fmt.Sprintf("Backup completed successfully from '%s' to '%s'", *source, *dest))
    fmt.Println("Backup completed successfully!")

    // print args for testing
    fmt.Printf("Source: %s\n", *source)
    fmt.Printf("Destination: %s\n", *dest)
}

// exitWithError handles error printing and exiting
func exitWithError(message string) {
    fmt.Println(message)
    log.Log(message)
    os.Exit(1)
}
