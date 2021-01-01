package main

import (
	//"crypto/x509"
	"./sha128aesgcm"
	"crypto/rand"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:\n " + os.Args[0] + " file.pem passphrase")
		return
	}
	file := os.Args[1]
	passwd := []byte(os.Args[2])

	rest, err := ioutil.ReadFile(file) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	block, rest := pem.Decode(rest)
	if block == nil {
		log.Fatal("Unable to find PEM section")
	}
	for {
		if strings.HasPrefix(block.Type, "128-AES-GCM ENCRYPTED ") {
			block.Type = block.Type[22:]
			nonce, err := base64.StdEncoding.DecodeString(block.Headers["nonce"])
			if err != nil {
				log.Fatal("Unable to parse nonce value", err)
			}
			delete(block.Headers, "nonce")
			block.Bytes, _ = sha128aesgcm.Decrypt([]byte(nonce), block.Bytes, passwd)
			pem.Encode(os.Stdout, block)
		} else {
			block.Type = "128-AES-GCM ENCRYPTED " + block.Type

			// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
			nonce := make([]byte, 12)
			if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
				panic(err.Error())
			}

			block.Headers["nonce"] = base64.StdEncoding.EncodeToString(nonce)
			block.Bytes, _ = sha128aesgcm.Encrypt(nonce, block.Bytes, passwd)
			pem.Encode(os.Stdout, block)
		}
		block, rest = pem.Decode(rest)
		if block == nil {
			break
		}
	}
}
