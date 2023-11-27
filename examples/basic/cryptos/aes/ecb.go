package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"github.com/devlights/gomy/output"
)

// Ecb -- crypto/aes のサンプルです (ECB) .
//
// # REFERENCES
//   - https://stackoverflow.com/questions/24072026/golang-aes-ecb-encryption
//   - https://blog.takuchalle.dev/post/2019/06/06/how_to_aes_encrypt_golang/
//   - https://www.developer.com/languages/cryptography-in-go/
//   - https://ja.wikipedia.org/wiki/Advanced_Encryption_Standard
func Ecb() error {
	const (
		size = aes.BlockSize
		txt  = "123456789012345_123456789012345_123456789012345_123456789012345_" // 64 bytes
		key  = "this_must_be_of_32_byte_length!!"
	)

	var (
		blk cipher.Block
		err error
	)

	blk, err = aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	var (
		plain    = []byte(txt)
		encBytes = make([]byte, len(plain))
		decBytes = make([]byte, len(plain))
	)

	// https://stackoverflow.com/questions/24072026/golang-aes-ecb-encryption
	for bs, be := 0, size; bs < len(plain); bs, be = bs+size, be+size {
		blk.Encrypt(encBytes[bs:be], plain[bs:be])
		blk.Decrypt(decBytes[bs:be], encBytes[bs:be])
	}

	output.Stdoutl("[AES][Encoding]", hex.EncodeToString(encBytes))
	output.Stdoutl("[AES][Decoding]", string(decBytes))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: crypto_aes_ecb

	   [Name] "crypto_aes_ecb"
	   [AES][Encoding]      b418fa0c8115c71ee2ba4680a1339033b418fa0c8115c71ee2ba4680a1339033b418fa0c8115c71ee2ba4680a1339033b418fa0c8115c71ee2ba4680a1339033
	   [AES][Decoding]      123456789012345_123456789012345_123456789012345_123456789012345_


	   [Elapsed] 42.24µs
	*/

}
