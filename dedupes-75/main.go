package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type fileInfo struct {
	info         os.FileInfo
	relativePath string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please provide a directory path")
	}
	dir := os.Args[1]
	if dir == "" {
		log.Fatal("please provide a directory path")
	}

	// list all files recursively inside dir
	files := listFiles(dir)
	sizeMap := identifyDuplicates(files, dir)

	for _, files := range sizeMap {
		if len(files) > 1 {
			msg := fmt.Sprintf("Potential duplicates:")
			for _, file := range files {
				msg = fmt.Sprintf("%s  %s", msg, file.relativePath)
			}
			log.Println(msg)
		}
	}

}

func identifyDuplicates(files []fileInfo, dir string) map[int64][]fileInfo {
	sizeMap := make(map[int64][]fileInfo)
	hashMap := make(map[string][]fileInfo)
	for _, file := range files {
		sizeMap[file.info.Size()] = append(sizeMap[file.info.Size()], file)
		fp := filepath.Join(dir, file.relativePath)
		hash := calculateHash(fp)
		hashMap[hash] = append(hashMap[hash], file)
	}
	return sizeMap
}

func calculateHash(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("error opening file %q: %v", path, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("error closing file %q: %v", path, err)
		}
	}(file)
	hash := md5.New()
	for {
		_, err := io.CopyN(hash, file, 1024)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error hashing file %q: %v", path, err)
		}
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func listFiles(dir string) []fileInfo {
	var files []fileInfo
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			relPath, _ := strings.CutPrefix(path, dir)
			relPath = strings.TrimPrefix(relPath, "/")
			file := fileInfo{
				info:         info,
				relativePath: relPath,
			}
			files = append(files, file)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("error walking the path %q: %v", dir, err)
	}
	return files
}
