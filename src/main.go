package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"hash/crc32"
	"io"
	"os"
)

// 通用函数：接收一个 hash.Hash 接口（如 md5.New(), sha256.New()）
func computeFileHash(filePath string, h hash.Hash) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func crc32File(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	table := crc32.MakeTable(crc32.IEEE)
	hasher := crc32.New(table)
	if _, err := io.Copy(hasher, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%08x", hasher.Sum32()), nil // 通常 CRC32 用 8 位十六进制表示
}

func main() {
	fmt.Println("Welcome to use ChecksumAssistant!")

	if len(os.Args) < 2 {
		fmt.Println("Usage: program <file>")
		fmt.Println("No file path provided. Please provide a file path as an argument.")
		os.Exit(1)
	}

	filePath := os.Args[1]
	fmt.Printf("Processing file: %s\n", filePath)

	success := true

	// MD5
	if md5sum, err := computeFileHash(filePath, md5.New()); err != nil {
		fmt.Printf("Error computing MD5: %v\n", err)
		success = false
	} else {
		fmt.Printf("MD5:      %s\n", md5sum)
	}

	// SHA1
	if sha1sum, err := computeFileHash(filePath, sha1.New()); err != nil {
		fmt.Printf("Error computing SHA1: %v\n", err)
		success = false
	} else {
		fmt.Printf("SHA1:     %s\n", sha1sum)
	}

	// SHA256
	if sha256sum, err := computeFileHash(filePath, sha256.New()); err != nil {
		fmt.Printf("Error computing SHA256: %v\n", err)
		success = false
	} else {
		fmt.Printf("SHA256:   %s\n", sha256sum)
	}

	// CRC32
	if crc32sum, err := crc32File(filePath); err != nil {
		fmt.Printf("Error computing CRC32: %v\n", err)
		success = false
	} else {
		fmt.Printf("CRC32:    %s\n", crc32sum)
	}

	if !success {
		os.Exit(1)
	}
}