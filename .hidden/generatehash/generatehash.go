// Copyright (C) 2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://www.bluice.com.br

package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"os"
	"fmt"
)

func getHash(sTipoHash string, file io.Reader) string {
	var hashsum hash.Hash

	switch sTipoHash {
	case "md5":
		hashsum = md5.New()
	case "sha1":
		hashsum = sha1.New()
	case "sha256":
		hashsum = sha256.New()
	case "sha512":
		hashsum = sha512.New()
	default:
		hashsum = md5.New()
	}

	if _, err := io.Copy(hashsum, file); err != nil {
		return ""
	}

	hashInBytes := hashsum.Sum(nil)
	return hex.EncodeToString(hashInBytes)
}

func main() {
	args := os.Args
	if len(args) > 2  {
		file, _ := os.Open(args[1])
		defer file.Close()

		fmt.Println(getHash(args[2], file))
		return
	}

	fmt.Println("File not found!")
}
