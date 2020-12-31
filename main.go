package main

import (
	//"crypto/x509"
	"./salt256"
	"crypto/rand"
	"encoding/base64"
	"encoding/pem"
	"fmt"
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
		if strings.HasPrefix(block.Type, "SALT256 ") {
			block.Type = block.Type[8:]
			salt, err := base64.StdEncoding.DecodeString(block.Headers["Salt"])
			if err != nil {
				log.Fatal("Unable to parse salt value", err)
			}
			delete(block.Headers, "Salt")
			block.Bytes, _ = salt256.Apply([]byte(salt), block.Bytes, passwd)
			pem.Encode(os.Stdout, block)
		} else {
			block.Type = "SALT256 " + block.Type
			salt := make([]byte, 15)
			rand.Read(salt)
			block.Headers["Salt"] = base64.StdEncoding.EncodeToString(salt)
			block.Bytes, _ = salt256.Apply(salt, block.Bytes, passwd)
			pem.Encode(os.Stdout, block)
		}
		block, rest = pem.Decode(rest)
		if block == nil {
			break
		}
	}

	//if block == nil || block.Type != "PUBLIC KEY" {
	//	log.Fatal("failed to decode PEM block containing public key")
	//}

	//pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
