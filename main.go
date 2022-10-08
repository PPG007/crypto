package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"hash"
	"io"
	"os"
)

var rootCommand *cobra.Command

func init() {
	rootCommand = &cobra.Command{
		Use: "crypto",
	}
	rootCommand.AddCommand(&cobra.Command{
		Use: "md5",
		RunE: func(cmd *cobra.Command, args []string) error {
			fileHash, err := getFileHash(args[0], "md5")
			if err != nil {
				return err
			}
			fmt.Println(fileHash)
			return nil
		},
	})
	rootCommand.AddCommand(&cobra.Command{
		Use: "sha1",
		RunE: func(cmd *cobra.Command, args []string) error {
			fileHash, err := getFileHash(args[0], "sha1")
			if err != nil {
				return err
			}
			fmt.Println(fileHash)
			return nil
		},
	})
	rootCommand.AddCommand(&cobra.Command{
		Use: "sha256",
		RunE: func(cmd *cobra.Command, args []string) error {
			fileHash, err := getFileHash(args[0], "sha256")
			if err != nil {
				return err
			}
			fmt.Println(fileHash)
			return nil
		},
	})
	rootCommand.AddCommand(&cobra.Command{
		Use: "sha512",
		RunE: func(cmd *cobra.Command, args []string) error {
			fileHash, err := getFileHash(args[0], "sha512")
			if err != nil {
				return err
			}
			fmt.Println(fileHash)
			return nil
		},
	})
}

func getFileHash(filePath, method string) (string, error) {
	var h hash.Hash
	switch method {
	case "md5":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	case "sha256":
		h = sha256.New()
	case "sha512":
		h = sha512.New()
	default:
		return "", errors.New("Invalid method ")
	}
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	buffer := make([]byte, h.Size())
	needBreak := false
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			needBreak = true
		} else if err != nil {
			return "", err
		}
		_, err = h.Write(buffer[:n])
		if err != nil {
			return "", err
		}
		if needBreak {
			break
		}
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func main() {
	err := rootCommand.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
