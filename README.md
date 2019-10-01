# GoParanoia
Code from the workshop "Adventures in Paranoia with Go and SQLite"


## AES Encryption and Decryption

$ AES_KEY="0123456789012345" go run 01_encrypt_aes.go "Hello World"

[189 66 213 32 192 240 163 198 89 18 112 190 150 234 133 248 10 214 159 253 11 245 129 197 68 225 13 23 141 226 137 82]

$ AES_KEY="0123456789012345" go run 02_decrypt_aes.go 189 66 213 32 192 240 163 198 89 18 112 190 150 234 133 248 10 214 159 253 11 245 129 197 68 225 13 23 141 226 137 82

Hello World
  		
