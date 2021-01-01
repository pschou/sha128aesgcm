package sha128aesgcm

//pbeWithSHAAnd128BitAES-GCM

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
)

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
