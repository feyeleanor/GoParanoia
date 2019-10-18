# GoParanoia
Code from the workshop "Adventures in Paranoia with Go and SQLite"


## AES Symmetric Ciphers

### Encryption

$ AES_KEY=0123456789012345 go run 01_encrypt_aes.go 'Hello World''

[189 66 213 32 192 240 163 198 89 18 112 190 150 234 133 248 10 214 159 253 11 245 129 197 68 225 13 23 141 226 137 82]

### Decryption

$ AES_KEY=0123456789012345 go run 02_decrypt_aes.go 189 66 213 32 192 240 163 198 89 18 112 190 150 234 133 248 10 214 159 253 11 245 129 197 68 225 13 23 141 226 137 82

Hello World

$ AES_KEY=0123456789012345 go run 03_encrypt_aes_to_base64.go 'Hello World!'

3Q9B81dbetzrdptKdv0TBgbanOZX9wgSViRGoF6YxUs=

$ AES_KEY=0123456789012345 go run 03_encrypt_aes_to_base64.go 'Hello World'
4dlr5z3SkGP4jykKekFYs6J3IabHjecnz3leLbeM0DI=

$ AES_KEY=0123456789012345 go run 04_decrypt_aes_from_base64.go 3Q9B81dbetzrdptKdv0TBgbanOZX9wgSViRGoF6YxUs=
Hello World

$ AES_KEY=0123456789012345 go run 04_decrypt_aes_from_base64.go 4dlr5z3SkGP4jykKekFYs6J3IabHjecnz3leLbeM0DI=
Hello World

## RSA Public Key Cryptography

### Key Generation

$ PEM_KEY=0123456789012345 go run 05_generate_rsa_keypair.go 1024

LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpQcm9jLVR5cGU6IDQsRU5DUllQVEVECkRFSy1JbmZvOiBBRVMtMjU2LUNCQywyZGYwNmEyMTUwY2VmODM2MDc0YTk2ZjNhMDA4N2NlYgoKeG9JSDJZVFhIVjQ4ckJCSWxZdmdtQ21Bc1M4bno2OHEwbFNLZHFJeHhralZNbjZJb3pQVndZN1A5SCtmOVIzRApYUmkwTVI5NnpHdkZYaXpjNEpLRk1odHFablhlRmJONnk3RXBnMWZObllhdzJ0MEtidlFjbXhiOEsyOEhRMXNJCjBlWjM4bEZzblFtWURWaHNpSlB0ZGpXeGxodjdrMFJlNEF0ZGk1NVhlZFhMaHVVSTkyeWhlZ2ttemV1NzhhbGIKWklLZUd0aUxCNWU4T2M5T3dEZnYyc0IwSEltaGFnbzRLeTJ3Yit2UG1SbElRR0lRb3FzcExIbWFxVUFHblpEZgpqcXY4SzJxRW9ITGluQVVLb3pVQm1CZnFFaUptRVdiOG45YUpvVWJDYlV1REM0dDZRRWRhS0lQbi9QeWZ3N1pOCnZuMERUZi9xditsd2JORElENUVhK2tENDhKcHdNODFpcmZwQlBncjlJd0E2cWtIZTExU2tCNDhhZmZYd1kvZzQKblBrRGhUc0Q4bExGeEJnVVJIaVJXQmhrMXNvdU5VRmVGQ2dhTTJwZkttMU1kNXdFVDBCdHVsNUZGbzV1QW1mVAp1RGp5TjlPTzJBaW5YRVpFeGZYOEV0MHJLODlEdWlaSmRrTzBKSks1MnhtL01NVkdIYkc3aXcreXYvYzlHTEtiCmF1REU4MDVHald6aTdUQXNxdlcxTzQ2ZkhWMEt6UUFuS1JoblJqOTdkRE83MkdGWUJKSWh0R3VPdkxwbEVKMUUKSWtkOWxpa0MrTmI5Q0YzMWMyTXNZQnRJTnh4L0UyZmtJRXhlRm50TkZxaFdNWW44Q1JWSVZGTTJXMlV2OFV5cQpQOFE3ZlNBeEhYYkhLUDlUTFNpcTZaeDdtM3JGeVBqV1hyNXJ6Zk16TXlXUWpBNk1hVmhNZ3ZPdHVFcFd0TWUrClV0c3FCL0ZBTStJbkI1MkhCUkdNT1dXeVM3bEF3SGVkL0pldWRxbGkzZ3llL29xTTFPQ0dzd296M05UUGsvaFcKVGJ1b2JSclJJS21namFnM3YzallSOC9aRnI5QS82YU84bTAxN0dGWVJadzJaMjR1L3lZR2FkK0Q0WUFPWEJBMwotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=

$ PEM_KEY=0123456789012345 go run 06_generate_rsa_keypair_to_pem.go test.pem 1024

$ cat test.pem
```
-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: AES-256-CBC,6a5f5b4d6607c56e3463821e53b977c8

N7hJVkMZ66DMNftmVawomEEQpb6vQ4s2WTj/OcDoVulq2fWrKFwneBlBLZhY1cVB
VpN5zPr1ici1G48H91xOXivYCi2ofwmKZ5RhTIABrU5gEA+gVhBRbARGW+pA5wee
2MogHCFd6kV+taTWxVB6weQ9lO4jcU8UqBxoJaHwjfqZs2gKLb68Mxdx4D0ZjN6k
uILzv36PafqVyip3+9b5jak5GAASuzAxwyLLPDKVZvqfXkHrBEl65lVYZ58nwmmC
ZN/RBFSp5kX2KVZ6twR8sEDYgSvAKfMOv/8lXM/9L4eZ9xRYt9rWATlY52dN5gor
40MjFvYfAR3u3jYKtdz5XPQlgZCil8RjlkcGn9M2cEPk7HQEgSZn7eLJ7YJyD+F6
wubcf+TJoIo3hwT7eREZtZW54XHRYzjmx+Utc4T4tsa2BsvF5QitFCDGBHp+F7Hs
v2JL5ncdaqMxj61+i997DgH8w5le+8chiak1Iv/t86hOE6GQKURgIqIIVmqkcYMc
6XDozo30p5pV8wp7skDBuRkZR0NNsYdozFglIcXAOoi1bGivehc+29990SY2mkAs
uKj/B08u0h21KeULTuRi6Zim/GV9dINzRXQQZP0QbiJ8U4/DPThv3vAbBtrrApzO
bqhh7Cx7gPUBSIshYx9eFuf8ev2iuHZ6buXwGJNEkTbSwI18WaAwqQMT0fSH9oEH
bnR85B0iDbXw1vIMrmfNCx9BTrkrVIWcDU7T/cfC2Bk5TRxu0SFPwjiKcQgycTBb
JkK77NmZ7rtv1r5AMLBZxwyKBWCSytMYOH/iBa2jeWIJJisP2elI8TsHtj250MoX
-----END RSA PRIVATE KEY-----
```

$ go run 06_generate_rsa_keypair_to_pem.go test2.pem 1024

$ cat test2.pem
```
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDJaCPleRozqN9dvYfkHlVNbqm7DAMaitpRXuItJeGL2r33VG8N
F6xPIgiS2jDFVRVeKXKZyry2uu0kw8pG5dg84O/tK23outls13+mmMozARA9gW4r
WVvPzlyob1mmfPWoc3YIlvRKP/TOYJ2r+s49nirJe0T/20TnDsJdB1WQ2QIDAQAB
AoGAMXZ48l9Gdw9vI9drKzPj4StfVceCb59QNJGn3EykUrN62eYLi9yXfauvDVm2
ho/5unFVar2mkP3hRZkr8TN7YjAUCzDk1EMUqWbGPaujBNb62GBvBzMmFL0AksoW
x8vE+T5rZhzJWJOnKDs+9w5kzHb0UxbyWSLTuLQcA8NsNjECQQDOR7BsjrceZzwc
PXM+H5VoC8uvvK7ztJQgbN+HK4q38RjlqYJ33qYeTszrtvXy7GA8lraDnIka4gOO
W0s5xHfNAkEA+fPAqtSdNrMLqiJNiWLThIETLTUcOrAYJSTVSW79NM5GXbztkXhU
UT38M47cHnBNIzdP6cgNDmfDZBRMjRoZPQJBAMXsX0kp/mX7o99szsCSyWZVuxBZ
uHw5jc8255rommc3vC3QYVDmnvmBnnmkyYRWbVh1O5Y4ggz7Q3I6AonhCxkCQQD0
U57b3iUnWPFoB0v/pcjY10slf6ruez++3zhWJdYBDjmeW5VmA1d4arzisRzd7Cya
5fCJt3F7yVYnt+f0bbSJAkBrJYpE9P9ikj/FWxp8IXuBjEQCjNwvyXnRO/CNBqjJ
32MGH2+wOQFbMVG5N9cdjISydkwu/IpWXfqFLFi1uo2y
-----END RSA PRIVATE KEY-----
```

$ PEM_KEY=0123456789012345 go run 07_extract_rsa_pubkey.go test.pem

$ cat test.pem.pub
```
-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAOJeXdQw2dRkYxPwAgGfLh8cDKlMW7I1yksd4sqYOGb9TF4UYg3IOm9Q
+gLhn8WWt72ua7pcEy2MPDUTXUB6WohhwiP+b384W2M63KlBBWh23S5z4mwOfRU4
IHO/qjMQT6j2o1zbGd7Fie0wiujQREegta5jnb67zt2OdDxAQpeNAgMBAAE=
-----END RSA PUBLIC KEY-----
```

$ go run 07_extract_rsa_pubkey.go test2.pem

$ cat test2.pem.pub
```
-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAMloI+V5GjOo3129h+QeVU1uqbsMAxqK2lFe4i0l4YvavfdUbw0XrE8i
CJLaMMVVFV4pcpnKvLa67STDykbl2Dzg7+0rbei62WzXf6aYyjMBED2BbitZW8/O
XKhvWaZ89ahzdgiW9Eo/9M5gnav6zj2eKsl7RP/bROcOwl0HVZDZAgMBAAE=
-----END RSA PUBLIC KEY-----
```

### PKCS#1 Encryption

$ go run 08_rsa_encrypt_data_pkcs1_to_base64.go test2.pem.pub 'Hello World'

F4nG2/g0k87S4CiEjUCPKCl4uYTM3WwyhcYs/RhKDx0y54bqMV0EtLW/CWomAE7VqYxI2w1Hl7XPbEB77HT/Cgmv/Z6XR8VBcm3v5n8RuvZyFr98/sSJrZynEnM0hIXRGoqYRWgo+i7kcufKAccHtAGvF1qCiPCi+Zf+mkdxEVY=

$ go run 08_rsa_encrypt_data_pkcs1_to_base64.go test.pem.pub 'Hello World'

Jmkib6RpL0AxLfEQh1zkGjVvvSEPaRxGyhx+xeLCmiprXS9PL1XX5U9H9bbE8uqfqjAKCzRGD4cDq8rrJaUx51MNK/1Jg4ZgbUg0NwJurL7R0cAG9AeNlwYwtMk2+zCcD8TsHdnTpqsFzrDcs9Qax2lsxfe4kXUPwhcZcj+zVJ0=

### PKCS#1 Decryption

$ go run 09_rsa_decrypt_data_pkcs1_from_base64.go test2.pem F4nG2/g0k87S4CiEjUCPKCl4uYTM3WwyhcYs/RhKDx0y54bqMV0EtLW/CWomAE7VqYxI2w1Hl7XPbEB77HT/Cgmv/Z6XR8VBcm3v5n8RuvZyFr98/sSJrZynEnM0hIXRGoqYRWgo+i7kcufKAccHtAGvF1qCiPCi+Zf+mkdxEVY=

Hello World

$ PEM_KEY=0123456789012345 go run 09_rsa_decrypt_data_pkcs1_from_base64.go test.pem Jmkib6RpL0AxLfEQh1zkGjVvvSEPaRxGyhx+xeLCmiprXS9PL1XX5U9H9bbE8uqfqjAKCzRGD4cDq8rrJaUx51MNK/1Jg4ZgbUg0NwJurL7R0cAG9AeNlwYwtMk2+zCcD8TsHdnTpqsFzrDcs9Qax2lsxfe4kXUPwhcZcj+zVJ0=

Hello World

### PKCS#1 Signing

$ go run 10_rsa_sign_data_pkcs1_to_base64.go test2.pem 'Hello World'

RRoTsrAQWMkKFLe/PfIPgCb+FkakkL62Zr+W8L500yz2BFQ6XKV7cqfcmKEJWlA0dmFO/HbBfe5DVE5FE6/OpQb+zYv3HXs4FJL+bMfE1T7wjrC2zQGb5N9evq6AWU1JFQbmOFPM6tnHsFsOr5d8n2h/xh66CANsV/W2XHBJbL0=

$ go run 11_rsa_verify_data_pkcs1_from_base64.go test2.pem.pub 'Hello World' RRoTsrAQWMkKFLe/PfIPgCb+FkakkL62Zr+W8L500yz2BFQ6XKV7cqfcmKEJWlA0dmFO/HbBfe5DVE5FE6/OpQb+zYv3HXs4FJL+bMfE1T7wjrC2zQGb5N9evq6AWU1JFQbmOFPM6tnHsFsOr5d8n2h/xh66CANsV/W2XHBJbL0=

Signature Verification Succeeded

$ go run 11_rsa_verify_data_pkcs1_from_base64.go test2.pem.pub 'Hello World' AAoTsrAQWMkKFLe/PfIPgCb+FkakkL62Zr+W8L500yz2BFQ6XKV7cqfcmKEJWlA0dmFO/HbBfe5DVE5FE6/OpQb+zYv3HXs4FJL+bMfE1T7wjrC2zQGb5N9evq6AWU1JFQbmOFPM6tnHsFsOr5d8n2h/xh66CANsV/W2XHBJbL0=

Signature Verification Failed

exit status 8

### OAEP Encryption

$ go run 12_rsa_encrypt_data_oaep_to_base64.go test3.pem.pub first 'Hello World'

WcEisRyeQaiGYa8ArThh9n0FJkTOW7gO0lluQSyMReW0GFw44jUwrlTD1eVjyQDEo4+p6rxkKsfDlKR7n8zxJDMN5nLtpFTxGxjTc4ugqrRjfPEY0aIP+9jkcuSFTkdR525vVww5vdC1J03crM01sqVh1Jgu67DOkUcY32toafNBuZ/lutPR+5yWBBvSQsQKAzIPlTplWpa3YuXJYjURXlfmE0ea1GMnqzfesf/tF7Qt/ejTFHkYUdWwSazOqTVj70vSG15cr4h+0W+x9K7vwSM8PrpajHnFq8/jIXFd0jFc9hWN+XzeojyHVeLnzEBszIorFgU8SC+4aekU0KC2NQ==

### OAEP Decryption

$ go run 13_rsa_decrypt_data_oaep_from_base64.go test3.pem first WcEisRyeQaiGYa8ArThh9n0FJkTOW7gO0lluQSyMReW0GFw44jUwrlTD1eVjyQDEo4+p6rxkKsfDlKR7n8zxJDMN5nLtpFTxGxjTc4ugqrRjfPEY0aIP+9jkcuSFTkdR525vVww5vdC1J03crM01sqVh1Jgu67DOkUcY32toafNBuZ/lutPR+5yWBBvSQsQKAzIPlTplWpa3YuXJYjURXlfmE0ea1GMnqzfesf/tF7Qt/ejTFHkYUdWwSazOqTVj70vSG15cr4h+0W+x9K7vwSM8PrpajHnFq8/jIXFd0jFc9hWN+XzeojyHVeLnzEBszIorFgU8SC+4aekU0KC2NQ==

Hello World

### PSS Signing

$ go run 14_rsa_sign_data_pss_to_base64.go test3.pem 'Hello World'

D18jN+2N9+cXXGKc8/4YvxYh0zgcM+p4U0wgczrwh3fa8hLM1+hoW4YC3ANMOzmTuDRsmh8iusDCI3Kx3huuHKLLiHRRK5qUFbSZkITZWjxE8zCNA2Ebe5ZyOeV8T5iqoNGT7xjt2efAi+GZOk4IpXB1pxdVsv6zF6zJCeumCHVgJx6cRJ7cyeqKr15Ry9sJhzmA0qmTkO2YiriKt+IFoulxYiJXwDgXWmADym00B4A3Di4qINEdLJwqbIuXfKQPm93RBXpqw7Pc621KEF/FJR725T6/95vrOqP59qntMmSV+oLVIqLPca/PwCEqhi7FJKeFNEEpwhqmYxUIWpBcfQ==

$ go run 15_rsa_verify_data_pss_from_base64.go test3.pem.pub 'Hello World' D18jN+2N9+cXXGKc8/4YvxYh0zgcM+p4U0wgczrwh3fa8hLM1+hoW4YC3ANMOzmTuDRsmh8iusDCI3Kx3huuHKLLiHRRK5qUFbSZkITZWjxE8zCNA2Ebe5ZyOeV8T5iqoNGT7xjt2efAi+GZOk4IpXB1pxdVsv6zF6zJCeumCHVgJx6cRJ7cyeqKr15Ry9sJhzmA0qmTkO2YiriKt+IFoulxYiJXwDgXWmADym00B4A3Di4qINEdLJwqbIuXfKQPm93RBXpqw7Pc621KEF/FJR725T6/95vrOqP59qntMmSV+oLVIqLPca/PwCEqhi7FJKeFNEEpwhqmYxUIWpBcfQ==

Signature Verification Succeeded

$ go run 15_rsa_verify_data_pss_from_base64.go test3.pem.pub 'Hello World' D18jN+2N9+cXXGKc8/4YvxYh0zgcM+p4U0wgczrwh3fa8hLM1+hoW4YC3ANMOzmTuDRsmh8iusDCI3Kx3huuHKLLiHRRK5qUFbSZkITZWjxE8zCNA2Ebe5ZyOeV8T5iqoNGT7xjt2efAi+GZOk4IpXB1pxdVsv6zF6zJCeumCHVgJx6cRJ7cyeqKr15Ry9sJhzmA0qmTkO2YiriKt+IFoulxYiJXwDgXWmADym00B4A3Di4qINEdLJwqbIuXfKQPm93RBXpqw7Pc621KEF/FJR725T6/95vrOqP59qntMmSV+oLVIqLPca/PwCEqhi7FJKeFNEEpwhqmYxUIWpBcfq==

Signature Verification Failed

exit status 8

## Hashing

### SHA512

$ go run 16_sha512_hash_to_base64.go 'Hello World'

LHT9F+2v2A6ER7DUZ0HuJDt+t03SFJoKsbkkb7MDgvJ+hT2FhXGeDmfL2g2qj1FnEGRhXWRa4nrLFb+xRH9Fmw==

### HMAC

$ HMAC_KEY=0123456789012345 go run 17_hmac_to_base64.go 'Hello World'

rxeUOuUKx3uVVO6qKQt+jawwRMvglw62D02h3ZTFfb3fz2gL0k29nmmYpe6n09X+LtZFdG4tTp8EOlM1jY75+A==

$ HMAC_KEY='0123456789012345' go run 18_verify_hmac_from_base64.go 'Hello World' rxeUOuUKx3uVVO6qKQt+jawwRMvglw62D02h3ZTFfb3fz2gL0k29nmmYpe6n09X+LtZFdG4tTp8EOlM1jY75+A==

Signature Verification Succeeded

$ HMAC_KEY=0123456789012345 go run 18_verify_hmac_from_base64.go 'Hello World' rxeUOuUKx3uVVO6qKQt+jawwRMvglw62D02h3ZTFfb3fz2gL0k29nmmYpe6n09X+LtZFdG4tTp8EOlM1jY75+Z==

Signature Verification Failed

exit status 1

### HMAC Content Signatures

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 19_aes_encrypt_hmac_sign_content_to_base64.go 'Hello World'

jz7vcg+B0kjBhTp5TfuV2XgrpA0JOKyAZcB9o+Qp7Jj4cE01W8hFfTLaqs3WFNM89MTfwAn9stUI9KfTRw3UOQ==f+tGhML1HsD0qqfQ+Gi+wNLRCskGtlXi+nhPEne73qM=

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 20_aes_decrypt_hmac_verify_content_from_base64.go 'Hello World' jz7vcg+B0kjBhTp5TfuV2XgrpA0JOKyAZcB9o+Qp7Jj4cE01W8hFfTLaqs3WFNM89MTfwAn9stUI9KfTRw3UOQ==f+tGhML1HsD0qqfQ+Gi+wNLRCskGtlXi+nhPEne73qM=

Signature Verification Succeeded

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 20_aes_decrypt_hmac_verify_content_from_base64.go 'Hello World' jz7vcg+B0kjBhTp5TfuV2XgrpA0JOKyAZcB9o+Qp7Jj4cE01W8hFfTLaqs3WFNM89MTfwAn9stUI9KfTRw3UOQ==f+tGhML1HsD0qqfQ+Gi+wNLRCskGtlXi+nhPEne73qm=

error: content doesn't match

exit status 2

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 20_aes_decrypt_hmac_verify_content_from_base64.go 'Hello World' jz7vcg+B0kjBhTp5TfuV2XgrpA0JOKyAZcB9o+Qp7Jj4cE01W8hFfTLaqs3WFNM89MTfwAn9stUI9KfTRw3UOq==f+tGhML1HsD0qqfQ+Gi+wNLRCskGtlXi+nhPEne73qM=

Signature Verification Failed

exit status 3

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 21_aes_decrypt_hmac_verify_content_from_base64.go jz7vcg+B0kjBhTp5TfuV2XgrpA0JOKyAZcB9o+Qp7Jj4cE01W8hFfTLaqs3WFNM89MTfwAn9stUI9KfTRw3UOQ==f+tGhML1HsD0qqfQ+Gi+wNLRCskGtlXi+nhPEne73qM=

Signature Verification Succeeded

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 21_aes_decrypt_hmac_verify_content_from_base64.go jz7vcg+B0kjBhTp5TfuV2XgrpA0JOKyAZcB9o+Qp7Jj4cE01W8hFfTLaqs3WFNM89MTfwAn9stUI9KfTRw3UOQ==f+tGhML1HsD0qqfQ+Gi+wNLRCskGtlXi+nhPEne73qm=

Signature Verification Failed

exit status 2

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 21_aes_decrypt_hmac_verify_content_from_base64.go jz7vcg+B0kjBhTp5TfuV2XgrpA0JOKyAZcB9o+Qp7Jj4cE01W8hFfTLaqs3WFNM89MTfwAn9stUI9KfTRw3UOq==f+tGhML1HsD0qqfQ+Gi+wNLRCskGtlXi+nhPEne73qM=

Signature Verification Failed

exit status 2

### HMAC Wrapper Signatures

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 22_aes_encrypt_hmac_sign_wrapper_to_base64.go 'Hello World'

dj82aiViZEVUZr++Q3Zz5Ix5Q7D0G7NX3DjyMm+i9O72/nXGcc6WKuqukg2n+c8TmlkC2t/rjXpLFmq5kNRsyw==tBpwJs9FZ0A+QPWoks7udfz6P0YA8h+vbwlduRcitS4=

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 23_aes_decrypt_hmac_verify_wrapper_from_base64.go dj82aiViZEVUZr++Q3Zz5Ix5Q7D0G7NX3DjyMm+i9O72/nXGcc6WKuqukg2n+c8TmlkC2t/rjXpLFmq5kNRsyw==tBpwJs9FZ0A+QPWoks7udfz6P0YA8h+vbwlduRcitS4=

Signature Verification Succeeded

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 23_aes_decrypt_hmac_verify_wrapper_from_base64.go dj82aiViZEVUZr++Q3Zz5Ix5Q7D0G7NX3DjyMm+i9O72/nXGcc6WKuqukg2n+c8TmlkC2t/rjXpLFmq5kNRsyw==tBpwJs9FZ0A+QPWoks7udfz6P0YA8h+vbwlduRcitS0=

Signature Verification Failed

exit status 2

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 23_aes_decrypt_hmac_verify_wrapper_from_base64.go dj82aiViZEVUZr++Q3Zz5Ix5Q7D0G7NX3DjyMm+i9O72/nXGcc6WKuqukg2n+c8TmlkC2t/rjXpLFmq5kNRsyW==tBpwJs9FZ0A+QPWoks7udfz6P0YA8h+vbwlduRcitS4=

Signature Verification Failed

exit status 2

### HMAC signed value chains

$ HMAC_KEY=0123456789012345 go run 24_chain_list_hmac_sign_to_base64.go Hello World Goodbye World

nDo7/2Jel9aP4RtT/v2vzkqvAuZC5mGjvsNKGGXofqZEMumfFZoIZDqFHLdMqsCQr2n1ujeKyXlvURO/iDAa5w== Hello

RX7rJOTD89VN7Kn/mWo6RTRaQPMkRexAjYH8w27NTG3Vfn9ST2oeguGuraOmHM8KXQLCu08tgaUYnCcy0wH/fQ== World

vt1b+boo+Qnpnx2E/knshgPAHmvXph3dHRacUXtn+bIOS61v2UyPpBFuK/ck5Uz40e7q+mC3kWcT3zWtMzKOXg== Goodbye

wl0TqHTzm5s3+FZYh5R7+rfkFfy4ffjY1LVPUEHO3DlHXOnRbRusEcfiTpPz41QEjvQ6Ywqb3RI3ugHL89DLyg== World

$ HMAC_KEY=0123456789012345 go run 24_chain_list_hmac_sign_to_base64.go Goodbye World Hello World

36Mee0SiQQr4yeJkgBJZfP4eMlRVG9O+FFcwNG6GlfkQOgvBXabzW0P+3ZqJ9qvppJzww2pMXZg4XsfOpaitpA== Goodbye

vuAFbTAZgmuFdo3xpZqY6n7TAq/HNA3PqXrCmgA+SsCMDzlc4GgfPsURx5IFiDY3RIdFoXnivRih/KIHQr+vwQ== World

p18Nb4M1msr7eIv27eIrMDaOUTX4qopGm5nL9T6QS6/PrXMI3at3Wq0vqbrkb9kRNFuWSzgFERyRax7GvBsajw== Hello

wl0TqHTzm5s3+FZYh5R7+rfkFfy4ffjY1LVPUEHO3DlHXOnRbRusEcfiTpPz41QEjvQ6Ywqb3RI3ugHL89DLyg== World

$ HMAC_KEY=0123456789012345 go run 25_chain_list_hmac_verify_from_base64.go Hello nDo7/2Jel9aP4RtT/v2vzkqvAuZC5mGjvsNKGGXofqZEMumfFZoIZDqFHLdMqsCQr2n1ujeKyXlvURO/iDAa5w== World RX7rJOTD89VN7Kn/mWo6RTRaQPMkRexAjYH8w27NTG3Vfn9ST2oeguGuraOmHM8KXQLCu08tgaUYnCcy0wH/fQ== Goodbye vt1b+boo+Qnpnx2E/knshgPAHmvXph3dHRacUXtn+bIOS61v2UyPpBFuK/ck5Uz40e7q+mC3kWcT3zWtMzKOXg== World wl0TqHTzm5s3+FZYh5R7+rfkFfy4ffjY1LVPUEHO3DlHXOnRbRusEcfiTpPz41QEjvQ6Ywqb3RI3ugHL89DLyg==

Signature Verification Succeeded

$ HMAC_KEY=0123456789012345 go run 25_chain_list_hmac_verify_from_base64.go Hello nDo7/2Jel9aP4RtT/v2vzkqvAuZC5mGjvsNKGGXofqZEMumfFZoIZDqFHLdMqsCQr2n1ujeKyXlvURO/iDAa5w== World RX7rJOTD89VN7Kn/mWo6RTRaQPMkRexAjYH8w27NTG3Vfn9ST2oeguGuraOmHM8KXQLCu08tgaUYnCcy0wH/fQ== Goodbye vt1b+boo+Qnpnx2E/knshgPAHmvXph3dHRacUXtn+bIOS61v2UyPpBFuK/ck5Uz40e7q+mC3kWcT3zWtMzKOXg== World wl0TqHTzm5s3+FZYh5R7+rfkFfy4ffjY1LVPUEHO3DlHXOnRbRusEcfiTpPz41QEjvQ6Ywqb3RI3ugHL89DLyG==

Signature Verification Failed

wl0TqHTzm5s3+FZYh5R7+rfkFfy4ffjY1LVPUEHO3DlHXOnRbRusEcfiTpPz41QEjvQ6Ywqb3RI3ugHL89DLyg== != wl0TqHTzm5s3+FZYh5R7+rfkFfy4ffjY1LVPUEHO3DlHXOnRbRusEcfiTpPz41QEjvQ6Ywqb3RI3ugHL89DLyG==

exit status 1

$ HMAC_KEY=0123456789012345 go run 25_chain_list_hmac_verify_from_base64.go Hello nDo7/2Jel9aP4RtT/v2vzkqvAuZC5mGjvsNKGGXofqZEMumfFZoIZDqFHLdMqsCQr2n1ujeKyXlvURO/iDAa5w== World RX7rJOTD89VN7Kn/mWo6RTRaQPMkRexAjYH8w27NTG3Vfn9ST2oeguGuraOmHM8KXQLCu08tgaUYnCcy0wH/fQ== Goodbye vt1b+boo+Qnpnx2E/knshgPAHmvXph3dHRacUXtn+bIOS61v2UyPpBFuK/ck5Uz40e7q+mC3kWcT3zWtMzKOXG== World wl0TqHTzm5s3+FZYh5R7+rfkFfy4ffjY1LVPUEHO3DlHXOnRbRusEcfiTpPz41QEjvQ6Ywqb3RI3ugHL89DLyg==

Signature Verification Failed

vt1b+boo+Qnpnx2E/knshgPAHmvXph3dHRacUXtn+bIOS61v2UyPpBFuK/ck5Uz40e7q+mC3kWcT3zWtMzKOXg== != vt1b+boo+Qnpnx2E/knshgPAHmvXph3dHRacUXtn+bIOS61v2UyPpBFuK/ck5Uz40e7q+mC3kWcT3zWtMzKOXG==

exit status 1

$ HMAC_KEY=0123456789012345 go run 25_chain_list_hmac_verify_from_base64.go Hello nDo7/2Jel9aP4RtT/v2vzkqvAuZC5mGjvsNKGGXofqZEMumfFZoIZDqFHLdMqsCQr2n1ujeKyXlvURO/iDAa5w== World RX7rJOTD89VN7Kn/mWo6RTRaQPMkRexAjYH8w27NTG3Vfn9ST2oeguGuraOmHM8KXQLCu08tgaUYnCcy0wH/fq== Goodbye vt1b+boo+Qnpnx2E/knshgPAHmvXph3dHRacUXtn+bIOS61v2UyPpBFuK/ck5Uz40e7q+mC3kWcT3zWtMzKOXg== World wl0TqHTzm5s3+FZYh5R7+rfkFfy4ffjY1LVPUEHO3DlHXOnRbRusEcfiTpPz41QEjvQ6Ywqb3RI3ugHL89DLyg==

Signature Verification Failed

RX7rJOTD89VN7Kn/mWo6RTRaQPMkRexAjYH8w27NTG3Vfn9ST2oeguGuraOmHM8KXQLCu08tgaUYnCcy0wH/fQ== != RX7rJOTD89VN7Kn/mWo6RTRaQPMkRexAjYH8w27NTG3Vfn9ST2oeguGuraOmHM8KXQLCu08tgaUYnCcy0wH/fq==

exit status 1

$ HMAC_KEY=0123456789012345 go run 25_chain_list_hmac_verify_from_base64.go Hello nDo7/2Jel9aP4RtT/v2vzkqvAuZC5mGjvsNKGGXofqZEMumfFZoIZDqFHLdMqsCQr2n1ujeKyXlvURO/iDAa5W== World RX7rJOTD89VN7Kn/mWo6RTRaQPMkRexAjYH8w27NTG3Vfn9ST2oeguGuraOmHM8KXQLCu08tgaUYnCcy0wH/fQ== Goodbye vt1b+boo+Qnpnx2E/knshgPAHmvXph3dHRacUXtn+bIOS61v2UyPpBFuK/ck5Uz40e7q+mC3kWcT3zWtMzKOXg== World wl0TqHTzm5s3+FZYh5R7+rfkFfy4ffjY1LVPUEHO3DlHXOnRbRusEcfiTpPz41QEjvQ6Ywqb3RI3ugHL89DLyg==

Signature Verification Failed

nDo7/2Jel9aP4RtT/v2vzkqvAuZC5mGjvsNKGGXofqZEMumfFZoIZDqFHLdMqsCQr2n1ujeKyXlvURO/iDAa5w== != nDo7/2Jel9aP4RtT/v2vzkqvAuZC5mGjvsNKGGXofqZEMumfFZoIZDqFHLdMqsCQr2n1ujeKyXlvURO/iDAa5W==

exit status 1

### Merkle Trees

$ HMAC_KEY=0123456789012345 go run 26_merkle_tree_hmac_sign_to_base64.go

XD3NWQD4H35xBS0E1cfrsFpECV8tTGN2JsCxbiEeO8LI1Sudg+5DrM2vBJSupUE++cBDTcNovQoLiX2u21OkSw== +

LFYhJE5EgA5MgEVIhbu3V3YGCknUmw1bcaOtqmcUwIT6In/wnM4kir2ONj+cXixGPLYoTwHBfkfoB0K1aOskUQ== *

KKpcA3TbQEjqg8Xj5wxOyVmjqA25hyhxl9mpR/dIFA+LFaO+788Y+7S3n4ddz+F/RKL2YsvNcTchNHDyWAOnUA== 3

IqEsHAfsvx4w/06gqbvQoU0iD+sIUXTtcunF4nyEWBPLUEsBeg7aUyID3Kmy1uutz5To/Mrjb3pu5BqhdLX/rA== 2

lCo6zpPpmoEaF9xW7a+ch0dKC+HzlOUGxpVvIorjXIen0H2iphehmZSEhK9LxMbWxE+D64N8/Xvi2cl91DDr2Q== 1

$ HMAC_KEY=0123456789012345 go run 27_merkle_tree_hmac_verify_from_base64.go XD3NWQD4H35xBS0E1cfrsFpECV8tTGN2JsCxbiEeO8LI1Sudg+5DrM2vBJSupUE++cBDTcNovQoLiX2u21OkSw== LFYhJE5EgA5MgEVIhbu3V3YGCknUmw1bcaOtqmcUwIT6In/wnM4kir2ONj+cXixGPLYoTwHBfkfoB0K1aOskUQ== KKpcA3TbQEjqg8Xj5wxOyVmjqA25hyhxl9mpR/dIFA+LFaO+788Y+7S3n4ddz+F/RKL2YsvNcTchNHDyWAOnUA== IqEsHAfsvx4w/06gqbvQoU0iD+sIUXTtcunF4nyEWBPLUEsBeg7aUyID3Kmy1uutz5To/Mrjb3pu5BqhdLX/rA== lCo6zpPpmoEaF9xW7a+ch0dKC+HzlOUGxpVvIorjXIen0H2iphehmZSEhK9LxMbWxE+D64N8/Xvi2cl91DDr2Q==

Signature Verification Succeeded

$ HMAC_KEY=0123456789012345 go run 27_merkle_tree_hmac_verify_from_base64.go XD3NWQD4H35xBS0E1cfrsFpECV8tTGN2JsCxbiEeO8LI1Sudg+5DrM2vBJSupUE++cBDTcNovQoLiX2u21OkSw== lCo6zpPpmoEaF9xW7a+ch0dKC+HzlOUGxpVvIorjXIen0H2iphehmZSEhK9LxMbWxE+D64N8/Xvi2cl91DDr2Q== LFYhJE5EgA5MgEVIhbu3V3YGCknUmw1bcaOtqmcUwIT6In/wnM4kir2ONj+cXixGPLYoTwHBfkfoB0K1aOskUQ== KKpcA3TbQEjqg8Xj5wxOyVmjqA25hyhxl9mpR/dIFA+LFaO+788Y+7S3n4ddz+F/RKL2YsvNcTchNHDyWAOnUA== IqEsHAfsvx4w/06gqbvQoU0iD+sIUXTtcunF4nyEWBPLUEsBeg7aUyID3Kmy1uutz5To/Mrjb3pu5BqhdLX/rA==

Signature Verification Failed

LFYhJE5EgA5MgEVIhbu3V3YGCknUmw1bcaOtqmcUwIT6In/wnM4kir2ONj+cXixGPLYoTwHBfkfoB0K1aOskUQ== != lCo6zpPpmoEaF9xW7a+ch0dKC+HzlOUGxpVvIorjXIen0H2iphehmZSEhK9LxMbWxE+D64N8/Xvi2cl91DDr2Q==

exit status 1

## Hybrid Cryptography

$ go run 28_hybrid_crypto_between_goroutines.go test3.pem session_label 0123456789012345 A B C

Bob heard: session_label

Bob heard: MDEyMzQ1Njc4OTAxMjM0NQ==

Bob wants to say: A

Alice heard: A

Alice wants to say: A received

Bob heard: A received

Bob wants to say: B

Alice heard: B

Alice wants to say: B received

Bob heard: B received

Bob wants to say: C

Alice heard: C

Alice wants to say: C received

Bob heard: C received

## SQLite3 Databases

### Create Database

$ ls -al test.db

ls: test.db: No such file or directory

$ go run 40_sqlite3_create_database.go test.db

$ ls -al test.db

-rw-r--r--  1 eleanor  admin  0 16 Oct 02:11 test.db

$ go run 41_sqlite3_create_database.go test2.db

$ ls -al *.db

-rw-r--r--  1 eleanor  admin  0 16 Oct 02:11 test.db

-rw-r--r--  1 eleanor  admin  0 16 Oct 13:03 test2.db

$ go run 42_sqlite3_create_database.go test3

$ ls -al *.db

-rw-r--r--  1 eleanor  admin  0 16 Oct 19:26 test.db

-rw-r--r--  1 eleanor  admin  0 16 Oct 19:26 test2.db

-rw-r--r--  1 eleanor  admin  0 16 Oct 20:37 test3.db

$ go run 43_sqlite3_create_database.go test4

$ ls -al *.db

-rw-r--r--  1 eleanor  admin  0 16 Oct 19:26 test.db

-rw-r--r--  1 eleanor  admin  0 16 Oct 19:26 test2.db

-rw-r--r--  1 eleanor  admin  0 16 Oct 20:37 test3.db

-rw-r--r--  1 eleanor  admin  0 16 Oct 21:06 test4.db

### Create Table

$ go run 44_sqlite3_create_table.go test.db

$ go run 44_sqlite3_create_table.go test.db

$ sqlite3 test.db

SQLite version 3.29.0 2019-07-10 17:32:03

Enter ".help" for usage hints.

sqlite> .schema

	CREATE TABLE Account (

		id    INTEGER PRIMARY KEY,

	    Name  TEXT NOT NULL,

	    Email TEXT UNIQUE NOT NULL

	);

sqlite> .quit

### Drop Table, Insert, Select

$ go run 45_sqlite3_create_table.go test.db

rows in Account table = 3

$ sqlite3 test.db

SQLite version 3.29.0 2019-07-10 17:32:03

Enter ".help" for usage hints.

sqlite> select * from Account;

a|Ellie|a@someserver.com

b|Ellie|b@someserver.com

c|Ellie|c@someserver.com

sqlite> select count(*) from Account;

3

### Create Table Index

$ go run 46_sqlite3_create_index.go test.db

UNIQUE constraint failed: Account.Name

exit status 4

$ sqlite3 test.db

SQLite version 3.29.0 2019-07-10 17:32:03

Enter ".help" for usage hints.

sqlite> select count(*) from Account;

1

sqlite> select * from Account;

a|Ellie|a@someserver.com

sqlite> .quit

$ go run 47_sqlite3_create_index.go test.db

rows in Account table = 3

$ sqlite3 test.db

SQLite version 3.29.0 2019-07-10 17:32:03

Enter ".help" for usage hints.

sqlite> select count(*) from Account;

3

sqlite> select * from Account;

a|Alpha|a@someserver.com

b|Beta|b@someserver.com

g|Gamma|g@someserver.com

sqlite> .quit


### Update in Table


### Delete from Table


### Join Tables
