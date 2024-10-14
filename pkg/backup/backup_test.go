package backup

import (
    "os"
    "testing"
)

func TestHandleFiles_Success(t *testing.T) {

    // create temporary directories for testing
    srcDir := t.TempDir()
    destDir := t.TempDir()

    // creaqte a sample file in src dir
    srcFile := srcDir + "/testfile.txt"
    err := os.WriteFile(srcFile, []byte("test"), 0644)
    if err != nil {
        t.Fatalf("failed to create source file: %v", err)
    }

    // Test: Attempt to back up file to dest directory
    err = HandleFiles(srcDir, destDir)
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }

    // Verify: Check if the file was copied successfully
    destFile := destDir + "/testfile.txt"
    if _, err := os.Stat(destFile); os.IsNotExist(err) {
        t.Errorf("expected file to be copied to %s but it does not exist", destFile)
    }
}

func TestHandleFiles_SourceNotFound(t *testing.T) {

    // Test: Attemp to back up a non-existent file
    err := HandleFiles("non-existent-dir", t.TempDir())
    if err == nil {
        t.Errorf("expected error but got nil")
    }
}
