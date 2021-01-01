package main

import (
	"./sha128aesgcm"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {

	input := []byte("This is my message.  This is my message again.")
	fmt.Println("input:", input)
	// input: [84 104 105 115 32 105 115 32 109 121 32 109 101 115 115 97 103 101 46 32 32 84 104 105 115 32 105 115 32 109 121 32 109 101 115 115 97 103 101 32 97 103 97 105 110 46]

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	fmt.Println("a good nonce:", nonce)
	fmt.Println("this nonce:", []byte("123456789012"))

	// In this example the nonce is "123456789012" and the password is "eeeee"
	output, _ := sha128aesgcm.Encrypt([]byte("123456789012"), input, []byte("eeeee"))
	fmt.Println("output:", output)
	// output: [203 153 155 192 238 213 114 164 208 196 183 216 111 244 57 10 206 181 49 156 5 124 209 24 57 186 178 125 238 75 137 250 192 171 158 198 223 103 213 91 241 59 183 172 136 35 115 61 111 59 223 43 100 76 223 73 182 89 9 185 247 244]

	// Note: Calling the Decrypt function on the output, reproduces the original input
	output2, _ := sha128aesgcm.Decrypt([]byte("123456789012"), output, []byte("eeeee"))
	fmt.Println("output2:", output2)
	// output2: [84 104 105 115 32 105 115 32 109 121 32 109 101 115 115 97 103 101 46 32 32 84 104 105 115 32 105 115 32 109 121 32 109 101 115 115 97 103 101 32 97 103 97 105 110 46]

}
