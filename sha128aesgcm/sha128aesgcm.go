package sha128aesgcm

//pbeWithSHAAnd128BitAES-GCM

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	//"errors"
)

/* Basic example of Salt256 operation:
input := []byte("this is my message this is my messagethis is my messagethis is my message")
fmt.Println("input:", input)
// input: [116 104 105 115 32 105 115 32 109 121 32 109 101 115 115 97 103 101 32 116 104 105 115 32 105 115 32 109 121 32 109 101 115 115 97 103 101 116 104 105 115 32 105 115 32 109 121 32 109 101 115 115 97 103 101 116 104 105 115 32 105 115 32 109 121 32 109 101 115 115 97 103 101]

// In this example the salt is "abc" and the password is "eeeee"
output,_ := salt256.Apply([]byte("abc"), input, []byte("eeeee"))
fmt.Println("output:", output)
// output: [158 162 148 107 26 31 171 175 251 186 101 117 252 63 166 139 247 41 238 158 126 146 173 45 212 125 3 81 204 79 107 244 165 108 255 200 169 111 129 168 59 28 85 90 63 17 84 72 247 52 121 63 77 125 210 210 206 137 216 4 197 36 223 126 215 224 74 253 111 242 254 36 105]

// Note: Calling the same function again on the output, reproduces the original input
output2,_ := salt256.Apply([]byte("abc"), output, []byte("eeeee"))
fmt.Println("output2:", output2)
// output2: [116 104 105 115 32 105 115 32 109 121 32 109 101 115 115 97 103 101 32 116 104 105 115 32 105 115 32 109 121 32 109 101 115 115 97 103 101 116 104 105 115 32 105 115 32 109 121 32 109 101 115 115 97 103 101 116 104 105 115 32 105 115 32 109 121 32 109 101 115 115 97 103 101]
*/

func zeroing20(buf [20]byte) {
	for i := 0; i < len(buf); i++ {
		buf[i] = 0
	}
}

func Decrypt(nonce []byte, blob []byte, passwd []byte) (output []byte, err error) {
	h := sha1.Sum(passwd)
	defer zeroing20(h)

	block, err := aes.NewCipher(h[0:16])
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	output, err = aesgcm.Open(nil, nonce, blob, nil)
	if err != nil {
		return nil, err
	}
	return
}
func Encrypt(nonce []byte, blob []byte, passwd []byte) (output []byte, err error) {
	h := sha1.Sum(passwd)
	defer zeroing20(h)

	block, err := aes.NewCipher(h[0:16])
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	output = aesgcm.Seal(nil, nonce, blob, nil)
	return
}
