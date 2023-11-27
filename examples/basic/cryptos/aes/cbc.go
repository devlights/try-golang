package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"

	"github.com/devlights/gomy/output"
)

// Cbc -- crypto/aes のサンプルです (CBC) .
//
// # REFERENCES
//   - https://blog.takuchalle.dev/post/2019/06/06/how_to_aes_encrypt_golang/
//   - https://www.developer.com/languages/cryptography-in-go/
//   - https://ja.wikipedia.org/wiki/Advanced_Encryption_Standard
func Cbc() error {
	const (
		txt = "123456789012345_123456789012345_123456789012345_123456789012345_" // 64 bytes
		key = "this_must_be_of_32_byte_length!!"
	)

	var (
		iv  = make([]byte, aes.BlockSize)
		blk cipher.Block
		err error
	)

	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return err
	}

	blk, err = aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	var (
		plain     = []byte(txt)
		encBytes  = make([]byte, len(plain))
		decBytes  = make([]byte, len(plain))
		encrypter = cipher.NewCBCEncrypter(blk, iv)
		decrypter = cipher.NewCBCDecrypter(blk, iv)
	)

	encrypter.CryptBlocks(encBytes, plain)
	decrypter.CryptBlocks(decBytes, encBytes)

	output.Stdoutl("[AES][Encoding]", hex.EncodeToString(encBytes))
	output.Stdoutl("[AES][Decoding]", string(decBytes))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: crypto_aes_cbc

	   [Name] "crypto_aes_cbc"
	   [AES][Encoding]      41b9e7408e1df7bbf637cc8c933c18f589ea342f87141ed070acfbf3b71b1922b34dc0c2e036c832af3155535170dae8462859a789cf730737cce6e59bb08103
	   [AES][Decoding]      123456789012345_123456789012345_123456789012345_123456789012345_


	   [Elapsed] 72.28µs
	*/

}
