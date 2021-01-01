package main

import (
	"./sha128aesgcm"
	"crypto/rand"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func zeroing(buf []byte) {
	for i := 0; i < len(buf); i++ {
		buf[i] = 0
	}
}

func main() {
	var passwd []byte
	var err error
	defer zeroing(passwd)
	if len(os.Args) < 2 {
		fmt.Println("Usage:\n " + os.Args[0] + " file.pem passphrase")
		return
	} else if len(os.Args) < 3 {
		fmt.Print("Enter Password: ")
		passwd, err = terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return
		}
		fmt.Println()
	} else {
		passwd = []byte(os.Args[2])
	}
	file := os.Args[1]

	rest, err := ioutil.ReadFile(file) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	block, rest := pem.Decode(rest)
	if block == nil {
		log.Fatal("Unable to find PEM section")
	}
	for {
		if _, ok := block.Headers["Proc-Type"]; ok {
			if block.Headers["Proc-Type"] == "4,ENCRYPTED" {
				dekinfo, ok := block.Headers["DEK-Info"]
				if !ok {
					log.Fatal("Missing the DEK-Info field for decrypting")
				}
				parts := strings.SplitN(dekinfo, ",", 2)
				if parts[0] != "AES-128-GCM" {
					log.Fatal("This tool only demonstrates the AES-128-GCM model")
				}

				nonce, err := hex.DecodeString(parts[1])
				if err != nil {
					log.Fatal("Unable to parse DEK-Info value", err)
				}
				delete(block.Headers, "Proc-Type")
				delete(block.Headers, "DEK-Info")
				block.Bytes, _ = sha128aesgcm.Decrypt([]byte(nonce), block.Bytes, passwd)
				pem.Encode(os.Stdout, block)
			} else {
				log.Fatal("Unknown attribute Proc-Type:", block.Headers["Proc-Type"])
			}
		} else {
			// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
			nonce := make([]byte, 12)
			if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
				panic(err.Error())
			}

			block.Headers["Proc-Type"] = "4,ENCRYPTED"
			block.Headers["DEK-Info"] = "AES-128-GCM," + strings.ToUpper(hex.EncodeToString(nonce))
			block.Bytes, _ = sha128aesgcm.Encrypt(nonce, block.Bytes, passwd)
			pem.Encode(os.Stdout, block)
		}
		block, rest = pem.Decode(rest)
		if block == nil {
			break
		}
	}
}
