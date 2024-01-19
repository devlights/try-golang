# サンプルリスト

このディレクトリには以下のサンプルがあります。

| file               | example name           | note                                                                                  |
| ------------------ | ---------------------- | ------------------------------------------------------------------------------------- |
| checksum/md5.go    | crypto_md5_checksum    | crypto/md5 のサンプルです.                                                            |
| aes/ecb.go         | crypto_aes_ecb         | crypto/aes のサンプルです (ECB) .                                                     |
| aes/cbc.go         | crypto_aes_cbc         | crypto/aes のサンプルです (CBC) .                                                     |
| rand/reader.go     | crypto_rand_reader     | crypto/rand.Reader を用いてセキュリティ的に安全な乱数を生成するサンプルです.          |
| rand/read.go       | crypto_rand_read       | crypto/rand.Read を用いてセキュリティ的に安全な乱数を生成するサンプルです.            |
| bcrypt/generate.go | crypto_bcrypt_generate | golang.org/x/crypto/bcrypt を使って bcrypt パスワードハッシュ を生成するサンプルです. |
| bcrypt/compare.go  | crypto_bcrypt_compare  | golang.org/x/crypto/bcrypt を使って生成したパスワードハッシュと比較するサンプルです.  |
