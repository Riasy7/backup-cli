// Package backup provides the backup functionality
package backup

import (
    "fmt"
    "io"
    "os"
    "path/filepath"
)

// ValidateSource checks if the source exists and whether it's a file or dir
//
// Parameters: source - the source path to validate
// Returns:
// - nil if the source is valid
func ValidateSource(source string) error {

    // check if source exists
    info, err := os.Stat(source)
    if os.IsNotExist(err) {
        return fmt.Errorf("source path '%s' does not exist", source)
    }

    // check dir and mode if legit
    if !info.IsDir() && !info.Mode().IsRegular() {
        return fmt.Errorf("source path '%s' is not a valid file or directory", source)
    }

    return nil
}

// HandleFiles handles the logic for copying files from the source to the dest
//
// Parameters:
// - source: the source path
// - dest: the destination path
// Returns:
// - nil if successful
func HandleFiles(source, dest string) error {

    // Validate source
    if err := ValidateSource(source); err != nil {
        return err
    }

    info, err := os.Stat(source)
    if err != nil {
        return err
    }

    // if directory recursively
    if info.IsDir() {
        return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }

            // create dest path
            relPath, _ := filepath.Rel(source, path)
            destPath := filepath.Join(dest, relPath)

            if info.IsDir() {
                // create dir in dest (file logic)
                return os.MkdirAll(destPath, os.ModePerm)
            }

            // copy file
            return copyFile(path, destPath)
        })
    }

    // if file
    destPath := filepath.Join(dest, info.Name())
    return copyFile(source, destPath)
}

// copyFile copies a file from src to dst
//
// Parameters:
// - src: the source file to copy
// - dst: the destination file to copy to
// Returns:
// - nil if the file is copied successfully
func copyFile(src, dst string) error {

    input, err := os.Open(src)
    if err != nil {
        return err
    }
    defer input.Close()

    output, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer output.Close()

    _, err = io.Copy(output, input)
    return err
}
