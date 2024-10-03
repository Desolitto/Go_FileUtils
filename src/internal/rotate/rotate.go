package rotate

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

func createArchiveFile(filename, archiveDir string) (string, error) {
	info, err := os.Stat(filename)
	if err != nil {
		return "", fmt.Errorf("error retrieving information for file %s: %v", filename, err)
	}

	if info.IsDir() {
		return "", fmt.Errorf("file %s is a directory and cannot be archived", filename)
	}

	timestamp := info.ModTime().Unix()
	baseName := filepath.Base(filename)
	nameWithoutExt := baseName[:len(baseName)-len(filepath.Ext(baseName))]
	archiveFilePath := filepath.Join(archiveDir, fmt.Sprintf("%s_%d.tag.gz", nameWithoutExt, timestamp))

	outFile, err := os.Create(archiveFilePath)
	if err != nil {
		return "", fmt.Errorf("error creating archive %s: %v", archiveFilePath, err)
	}
	defer outFile.Close()

	gz := gzip.NewWriter(outFile)
	defer gz.Close()

	tw := tar.NewWriter(gz)
	defer tw.Close()

	header, err := tar.FileInfoHeader(info, filename)
	if err != nil {
		return "", fmt.Errorf("error creating header for file %s: %v", filename, err)
	}
	header.Name = baseName

	if err := tw.WriteHeader(header); err != nil {
		return "", fmt.Errorf("error writing header to archive %s: %v", archiveFilePath, err)
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("error opening file %s: %v", filename, err)
	}
	defer file.Close()

	if _, err := io.Copy(tw, file); err != nil {
		return "", fmt.Errorf("error copying file %s to archive %s: %v", filename, archiveFilePath, err)
	}

	return archiveFilePath, nil
}

func ArchiveFile(filename, archiveDir string, wg *sync.WaitGroup) {
	defer wg.Done()

	archivePath, err := createArchiveFile(filename, archiveDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s - Successfully archived\n", archivePath)
}
