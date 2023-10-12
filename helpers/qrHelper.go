package helpers

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ZipFolder(folderName string) error {
	zipFile, err := os.Create(folderName + ".zip")
	if err != nil {
		fmt.Println("Error creating zip file:", err)
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(folderName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create a zip header for the current file or directory.
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Set the name of the file or directory within the zip archive.
		// The path should be relative to the source folder.
		// For example, if the source folder is "myfolder" and we're adding "myfolder/file.txt",
		// we set the name as "file.txt".
		header.Name, _ = filepath.Rel(folderName, path)

		// If the file is a directory, mark it as such.
		if info.IsDir() {
			header.Name += "/"
		}

		// Create a zip writer for the current file or directory.
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// If the item is a file, copy its contents to the zip writer.
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
