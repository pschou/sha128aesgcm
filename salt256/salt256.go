package salt256

import (
	"crypto/sha256"
	"errors"
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

func zeroing(buf []byte) {
	for i := 0; i < len(buf); i++ {
		buf[i] = 0
	}
}
func zeroing32(buf [32]byte) {
	for i := 0; i < len(buf); i++ {
		buf[i] = 0
	}
}

func Apply(salt []byte, block []byte, passwd []byte) (output []byte, err error) {
	block_s := len(block)
	if block_s > 32*255*256 {
		return nil, errors.New("Algorithm not built for large data sets")
	}
	salt_s := len(salt)
	passwd_s := len(passwd)
	saltpasswd_s := salt_s + passwd_s
	counter_s := 2

	seed := make([]byte, saltpasswd_s+4+counter_s)
	defer zeroing(seed)

	copy(seed[0:salt_s], salt)
	copy(seed[salt_s:saltpasswd_s], passwd)
	seed[saltpasswd_s] = byte((block_s >> 24) & 255)
	seed[saltpasswd_s+1] = byte((block_s >> 16) & 255)
	seed[saltpasswd_s+2] = byte((block_s >> 8) & 255)
	seed[saltpasswd_s+3] = byte(block_s & 255)

	output = make([]byte, block_s)
	sum := [32]byte{0}
	defer zeroing32(sum)

	for i := 0; i <= block_s/32; i++ {
		seed[saltpasswd_s+4] = byte(i >> 8)
		seed[saltpasswd_s+5] = byte(i & 255)

		sum = sha256.Sum256(seed)

		offset := i * 32
		stop := 32
		if offset+stop > block_s {
			stop = block_s - offset
		}
		for j := 0; j < stop; j++ {
			output[offset+j] = block[offset+j] ^ sum[j]
		}
	}
	return
}
