package main

const (
	_ = iota
	FILE_UNREADABLE
	FILE_WRITE_FAILED
	MISSING_FILENAME
	MISSING_KEYSIZE
	INVALID_KEYSIZE
	CREATE_KEY_FAILED
	NOT_A_PEM_FILE
	NOT_A_PRIVATE_KEY
	NOT_A_PUBLIC_KEY
	PEM_PASSWORD_REQUIRED
	PEM_ENCRYPTION_FAILED
	PEM_DECRYPTION_FAILED
	INVALID_PRIVATE_KEY
	INVALID_PUBLIC_KEY
	RSA_ENCRYPTION_FAILED
	RSA_DECRYPTION_FAILED
	AES_ENCRYPTION_FAILED
	AES_DECRYPTION_FAILED
	SIGNING_FAILED
	VERIFICATION_FAILED
	CONTENT_MISMATCH
	UNEVEN_PARAMETERS
  MISSING_HASHES
	NOT_AN_RSA_KEY
  NOT_ENOUGH_RANDOMNESS
  WEB_REQUEST_FAILED
  WEB_NO_BODY
)
