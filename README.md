# GoParanoia
Code from the workshop "Adventures in Paranoia with Go and SQLite"


## AES Symmetric Ciphers

### Encryption

```bash
$ AES_KEY=0123456789012345 go run 01_encrypt_aes.go 'Hello World'
[189 66 213 32 192 240 163 198 89 18 112 190 150 234 133 248 10 214 159 253 11 245 129 197 68 225 13 23 141 226 137 82]
```

### Decryption

```bash
$ AES_KEY=0123456789012345 go run 02_decrypt_aes.go 189 66 213 32 192 240 163 198 89 18 112 190 150 234 133 248 10 214 159 253 11 245 129 197 68 225 13 23 141 226 137 82
Hello World

$ AES_KEY=0123456789012345 go run 03_encrypt_aes_to_base64.go aes.go 'Hello World!'
3Q9B81dbetzrdptKdv0TBgbanOZX9wgSViRGoF6YxUs=

$ AES_KEY=0123456789012345 go run 03_encrypt_aes_to_base64.go aes.go 'Hello World'
4dlr5z3SkGP4jykKekFYs6J3IabHjecnz3leLbeM0DI=

$ AES_KEY=0123456789012345 go run 04_decrypt_aes_from_base64.go aes.go 3Q9B81dbetzrdptKdv0TBgbanOZX9wgSViRGoF6YxUs=
Hello World

$ AES_KEY=0123456789012345 go run 04_decrypt_aes_from_base64.go aes.go 4dlr5z3SkGP4jykKekFYs6J3IabHjecnz3leLbeM0DI=
Hello World
```

## RSA Public Key Cryptography

### Key Generation

```bash
$ PEM_KEY=0123456789012345 go run 05_generate_rsa_keypair.go helpers.go errors.go 1024
LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpQcm9jLVR5cGU6IDQsRU5DUllQVEVECkRFSy1JbmZvOiBBRVMtMjU2LUNCQywyZGYwNmEyMTUwY2VmODM2MDc0YTk2ZjNhMDA4N2NlYgoKeG9JSDJZVFhIVjQ4ckJCSWxZdmdtQ21Bc1M4bno2OHEwbFNLZHFJeHhralZNbjZJb3pQVndZN1A5SCtmOVIzRApYUmkwTVI5NnpHdkZYaXpjNEpLRk1odHFablhlRmJONnk3RXBnMWZObllhdzJ0MEtidlFjbXhiOEsyOEhRMXNJCjBlWjM4bEZzblFtWURWaHNpSlB0ZGpXeGxodjdrMFJlNEF0ZGk1NVhlZFhMaHVVSTkyeWhlZ2ttemV1NzhhbGIKWklLZUd0aUxCNWU4T2M5T3dEZnYyc0IwSEltaGFnbzRLeTJ3Yit2UG1SbElRR0lRb3FzcExIbWFxVUFHblpEZgpqcXY4SzJxRW9ITGluQVVLb3pVQm1CZnFFaUptRVdiOG45YUpvVWJDYlV1REM0dDZRRWRhS0lQbi9QeWZ3N1pOCnZuMERUZi9xditsd2JORElENUVhK2tENDhKcHdNODFpcmZwQlBncjlJd0E2cWtIZTExU2tCNDhhZmZYd1kvZzQKblBrRGhUc0Q4bExGeEJnVVJIaVJXQmhrMXNvdU5VRmVGQ2dhTTJwZkttMU1kNXdFVDBCdHVsNUZGbzV1QW1mVAp1RGp5TjlPTzJBaW5YRVpFeGZYOEV0MHJLODlEdWlaSmRrTzBKSks1MnhtL01NVkdIYkc3aXcreXYvYzlHTEtiCmF1REU4MDVHald6aTdUQXNxdlcxTzQ2ZkhWMEt6UUFuS1JoblJqOTdkRE83MkdGWUJKSWh0R3VPdkxwbEVKMUUKSWtkOWxpa0MrTmI5Q0YzMWMyTXNZQnRJTnh4L0UyZmtJRXhlRm50TkZxaFdNWW44Q1JWSVZGTTJXMlV2OFV5cQpQOFE3ZlNBeEhYYkhLUDlUTFNpcTZaeDdtM3JGeVBqV1hyNXJ6Zk16TXlXUWpBNk1hVmhNZ3ZPdHVFcFd0TWUrClV0c3FCL0ZBTStJbkI1MkhCUkdNT1dXeVM3bEF3SGVkL0pldWRxbGkzZ3llL29xTTFPQ0dzd296M05UUGsvaFcKVGJ1b2JSclJJS21namFnM3YzallSOC9aRnI5QS82YU84bTAxN0dGWVJadzJaMjR1L3lZR2FkK0Q0WUFPWEJBMwotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=

$ PEM_KEY=0123456789012345 go run 06_generate_rsa_keypair_to_pem.go helpers.go errors.go test.pem 1024
$ cat test.pem
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

$ go run 06_generate_rsa_keypair_to_pem.go helpers.go errors.go test2.pem 1024
$ cat test2.pem
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

$ go run 06_generate_rsa_keypair_to_pem.go helpers.go errors.go test3.pem 2048
$ cat test3.pem
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA+g9QHQztv2/761iqF+cf8GRcp8iReFv0uU38YuAaXE5rE0Hs
St6UQMkDbK4yq9VTq4NvXdo+FLopz/+JOi1ao29YRim15MU6ZnaPILN8l+6u4Pjb
+kkLr2AhVIIgy5tw7VCHuf8bv4zrLNe2wH3SSUKINDXdGE0X4FthsS2hy8/nGcZq
xPPBGaEvNXwRPV+gsa2ebSc1tA9Bn9ot5EKnLfZTDR2uIBRyr0h9Shc1p7bh7Ib+
qa2kRmmoZ0PM1twc+Eh1qUk/JE74bQbnGJY3Ib3PHz+D3s+nyXj4UARF+Zet9PnK
DQXUhF0PnBHT834HrA3z26VnIoO1fyJ6j8MmZwIDAQABAoIBAQDHYq5ee5CWtJQK
dAgsRv4Qj/XRUvurHKdlmelEbayXLW/zSv0+NXvArIcugTemC9NLTuHd08bsgAQW
9YpV+RPsfTvFtjxyzVBjouU9DeCqxIZsnAHK77OBgwUQfTHpJ2hD0ZJnnCJj+dkL
SHVPxwOFpqcz+j7wwYfoPPUIcXoxJ2F+gOZivRVEgjMwHf/TM3weHOEH492EqY7T
PBUGgrQj5SIdwn2E3tbrLWe6rQ5xpEuEbcabpx+Y2rs62/dHfKNkPsciodHklmzN
XigveiEPVM7SGInoP09bddSf25xTIIcrccKVyF4RR2Ky8tqBIdP/637DUG4X15Xv
7aqajbkBAoGBAPsHZNlCUB8RcRsXEPkmJFrfIXLxKY7+J1G6xfJ16vDX/MF7bxDa
Tc9sCT5b108m9Fi8q1C73SEfe2HobX+NygTCUbL3PgsXiASvblHIQOTwi4ZIi6af
oWT7VyWdbqlf9OlrOJOmMHchCuayqc/f8f2gldbkcBjB+puI6FeVCQ2LAoGBAP8D
AZptnY7mDfGyeokxKQNEZPt9iNVioNOJKQqp9FkvITN+sNULoiG5bY5fLTNTNK0L
NSOdU2+ObLAM6S6K0KtpDykOpiUgLOwKW6b2LlymMaqglQAzsQumIZZkl4L5l6pM
pn5SFitA9lp6T1ZwDq3bgEV+5MywxbPescbCpF4VAoGBAPOnzZ+D2jkNuPdOgE8V
UpxRkCn9IxWEY63U2GXrRsvXGaLDqJ53Bqeaea5pfqd5bBrlpnOKpayM3jm+XOo+
OiQ8aQiJK1Og14mUrkP7V+HgCc2hhMuKjiFyRLhiAxOr7BDU3emSmeBH1kjuih0X
hc9km4wl1xhw0Es/rCW2pz7ZAoGBANWnMySrSiJ7ZJtnqXpgEbiHCQQbvRTquUwB
0rz4f5xo/CwkrQsR4BjZPozV+QzTXYLNrQvVuB3y+eTWLFeuPf4e23DOsSzRtFFe
rENKZrWBGIOYXr34kRDtc56JUNePPh6UcbnMFH7QwpYze100LVhEW0fedNt6D2oP
3mpkkUmxAoGAOaECRKq5zibXrDQzRceNHZzfxFfGiZgXz4n3nFwW35mHmud0Du2a
8/bbdbWCeN4h11pD9C7lSHjmfY3LDW/4ENnqIgR2HWQ887kIRN2O/DazrJBkT58A
yMuDjojWXbMeeM+yAvocE2mUm67q5VoEr8BO3CIEdjzPW0gg0UVmRB4=

$ PEM_KEY=0123456789012345 go run 07_extract_rsa_pubkey.go pem.go helpers.go errors.go test.pem
$ cat test.pem.pub
-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAOJeXdQw2dRkYxPwAgGfLh8cDKlMW7I1yksd4sqYOGb9TF4UYg3IOm9Q
+gLhn8WWt72ua7pcEy2MPDUTXUB6WohhwiP+b384W2M63KlBBWh23S5z4mwOfRU4
IHO/qjMQT6j2o1zbGd7Fie0wiujQREegta5jnb67zt2OdDxAQpeNAgMBAAE=
-----END RSA PUBLIC KEY-----

$ go run 07_extract_rsa_pubkey.go pem.go helpers.go errors.go test2.pem
$ cat test2.pem.pub
-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAMloI+V5GjOo3129h+QeVU1uqbsMAxqK2lFe4i0l4YvavfdUbw0XrE8i
CJLaMMVVFV4pcpnKvLa67STDykbl2Dzg7+0rbei62WzXf6aYyjMBED2BbitZW8/O
XKhvWaZ89ahzdgiW9Eo/9M5gnav6zj2eKsl7RP/bROcOwl0HVZDZAgMBAAE=
-----END RSA PUBLIC KEY-----

$ go run 07_extract_rsa_pubkey.go pem.go helpers.go errors.go test3.pem
-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEAsMy4eOgDV0bScY9JW2ePbc1YysCqk4vqkye+4DM2ZCyxXA9Ngjco
FiqQTqgc7HX9C//6SKjEO8sJ8Ja1LYrc8Inf9wS3LG2TKU4EjRRjt7jPyHG6dJHu
gwvk+3aDeluV5H3gJJbAo5YEVHDKF6Mw2YxBMAijKdzETLleSEdqXQkF6hjy+3Bm
fX9pJDHyiC9aWMZ3oP2wqNh0EA++pFqQaWMLUtIZA+B4BY5MN8mTm3Ll0EWOOTHA
Tfl0dOQTzgUGM0N371JJ2Ar4UrHUusRhhxvFeJnCJ9Yb/l9yG3uJ7qljzzDOLvc/
BBpjM63N6wnMgvs2u22QehsEIyNqkRybIwIDAQAB
-----END RSA PUBLIC KEY-----
```

### PKCS#1 Encryption

```bash
$ go run 08_rsa_encrypt_data_pkcs1_to_base64.go pem.go helpers.go errors.go test2.pem.pub 'Hello World'
F4nG2/g0k87S4CiEjUCPKCl4uYTM3WwyhcYs/RhKDx0y54bqMV0EtLW/CWomAE7VqYxI2w1Hl7XPbEB77HT/Cgmv/Z6XR8VBcm3v5n8RuvZyFr98/sSJrZynEnM0hIXRGoqYRWgo+i7kcufKAccHtAGvF1qCiPCi+Zf+mkdxEVY=

$ go run 08_rsa_encrypt_data_pkcs1_to_base64.go pem.go helpers.go errors.go test.pem.pub 'Hello World'
Jmkib6RpL0AxLfEQh1zkGjVvvSEPaRxGyhx+xeLCmiprXS9PL1XX5U9H9bbE8uqfqjAKCzRGD4cDq8rrJaUx51MNK/1Jg4ZgbUg0NwJurL7R0cAG9AeNlwYwtMk2+zCcD8TsHdnTpqsFzrDcs9Qax2lsxfe4kXUPwhcZcj+zVJ0=
```

### PKCS#1 Decryption

```bash
$ go run 09_rsa_decrypt_data_pkcs1_from_base64.go pem.go helpers.go errors.go test2.pem F4nG2/g0k87S4CiEjUCPKCl4uYTM3WwyhcYs/RhKDx0y54bqMV0EtLW/CWomAE7VqYxI2w1Hl7XPbEB77HT/Cgmv/Z6XR8VBcm3v5n8RuvZyFr98/sSJrZynEnM0hIXRGoqYRWgo+i7kcufKAccHtAGvF1qCiPCi+Zf+mkdxEVY=
Hello World

$ PEM_KEY=0123456789012345 go run 09_rsa_decrypt_data_pkcs1_from_base64.go pem.go helpers.go errors.go test.pem Jmkib6RpL0AxLfEQh1zkGjVvvSEPaRxGyhx+xeLCmiprXS9PL1XX5U9H9bbE8uqfqjAKCzRGD4cDq8rrJaUx51MNK/1Jg4ZgbUg0NwJurL7R0cAG9AeNlwYwtMk2+zCcD8TsHdnTpqsFzrDcs9Qax2lsxfe4kXUPwhcZcj+zVJ0=
Hello World
```

### PKCS#1 Signing

```bash
$ go run 10_rsa_sign_data_pkcs1_to_base64.go pem.go helpers.go errors.go test2.pem 'Hello World'
pYTAAufA3nBKgdd9OyF+zQ0Vgp962VvdjD2y9M3lfgL1K5dOFp1eEGZ+ORpZ44mAXhN8f2OJOwiy2lil2adle/B6GVb0/Za1JDoxUSDs7g66gsu7FBMrquhfPOrpQdu8+q02n+dmhIpzvn7+c4inbp+qJ6GKuvTGzU2X/y8ad+E=

$ go run 11_rsa_verify_data_pkcs1_from_base64.go pem.go helpers.go errors.go test2.pem.pub 'Hello World' pYTAAufA3nBKgdd9OyF+zQ0Vgp962VvdjD2y9M3lfgL1K5dOFp1eEGZ+ORpZ44mAXhN8f2OJOwiy2lil2adle/B6GVb0/Za1JDoxUSDs7g66gsu7FBMrquhfPOrpQdu8+q02n+dmhIpzvn7+c4inbp+qJ6GKuvTGzU2X/y8ad+E=
Signature Verification Succeeded

$ go run 11_rsa_verify_data_pkcs1_from_base64.go pem.go helpers.go errors.go test2.pem.pub 'Hello World' pYTAAufA3nBKgdd9OyF+zQ0Vgp962VvdjD2y9M3lfgL1K5dOFp1eEGZ+ORpZ44mAXhN8f2OJOwiy2lil2adle/B6GVb0/Za1JDoxUSDs7g66gsu7FBMrquhfPOrpQdu8+q02n+dmhIpzvn7+c4inbp+qJ6GKuvTGzU2X/y8ad+e=
crypto/rsa: verification error
exit status 20
```

### OAEP Encryption

```bash
$ go run 12_rsa_encrypt_data_oaep_to_base64.go pem.go helpers.go errors.go test3.pem.pub first 'Hello World'
WcEisRyeQaiGYa8ArThh9n0FJkTOW7gO0lluQSyMReW0GFw44jUwrlTD1eVjyQDEo4+p6rxkKsfDlKR7n8zxJDMN5nLtpFTxGxjTc4ugqrRjfPEY0aIP+9jkcuSFTkdR525vVww5vdC1J03crM01sqVh1Jgu67DOkUcY32toafNBuZ/lutPR+5yWBBvSQsQKAzIPlTplWpa3YuXJYjURXlfmE0ea1GMnqzfesf/tF7Qt/ejTFHkYUdWwSazOqTVj70vSG15cr4h+0W+x9K7vwSM8PrpajHnFq8/jIXFd0jFc9hWN+XzeojyHVeLnzEBszIorFgU8SC+4aekU0KC2NQ==
```

### OAEP Decryption

```bash
$ go run 13_rsa_decrypt_data_oaep_from_base64.go pem.go helpers.go errors.go test3.pem first WcEisRyeQaiGYa8ArThh9n0FJkTOW7gO0lluQSyMReW0GFw44jUwrlTD1eVjyQDEo4+p6rxkKsfDlKR7n8zxJDMN5nLtpFTxGxjTc4ugqrRjfPEY0aIP+9jkcuSFTkdR525vVww5vdC1J03crM01sqVh1Jgu67DOkUcY32toafNBuZ/lutPR+5yWBBvSQsQKAzIPlTplWpa3YuXJYjURXlfmE0ea1GMnqzfesf/tF7Qt/ejTFHkYUdWwSazOqTVj70vSG15cr4h+0W+x9K7vwSM8PrpajHnFq8/jIXFd0jFc9hWN+XzeojyHVeLnzEBszIorFgU8SC+4aekU0KC2NQ==
Hello World
```

### PSS Signing

```
$ go run 14_rsa_sign_data_pss_to_base64.go pem.go helpers.go errors.go test3.pem 'Hello World'
U0J3zvhiOlwn7qlQLjCyUvGZhu9j7w2jICTYq7jpuxqRAM/1Rk9aCFQRZdW10MblX+fzc3lNDSfiR1QkFu7NHdPW40Ai3bAq7SPyzn38GX38pVqmHIui3DwbtTHrcYhYZhEw5GHQOlTWhS2giS7Oyvd4iqCPWrcP0e9NP4L1z1ZDC2KLOpNJkgon5hfiryGR0YCa9TtrxJaWmFySV4z2dpJUpcvy+Ca8fwcgr/JLhomP7JQdv7iSvp0Tlubw3v48tiKxcYak2cj27gbAa4d7E+798QWE09OIdcFbmio3IuJuC0e6E9OMDqEuK/dzLHHqw2A5BfHhve8eXnDbH5D8JA==

$ go run 15_rsa_verify_data_pss_from_base64.go pem.go helpers.go errors.go test3.pem.pub 'Hello World' U0J3zvhiOlwn7qlQLjCyUvGZhu9j7w2jICTYq7jpuxqRAM/1Rk9aCFQRZdW10MblX+fzc3lNDSfiR1QkFu7NHdPW40Ai3bAq7SPyzn38GX38pVqmHIui3DwbtTHrcYhYZhEw5GHQOlTWhS2giS7Oyvd4iqCPWrcP0e9NP4L1z1ZDC2KLOpNJkgon5hfiryGR0YCa9TtrxJaWmFySV4z2dpJUpcvy+Ca8fwcgr/JLhomP7JQdv7iSvp0Tlubw3v48tiKxcYak2cj27gbAa4d7E+798QWE09OIdcFbmio3IuJuC0e6E9OMDqEuK/dzLHHqw2A5BfHhve8eXnDbH5D8JA==
Signature Verification Succeeded

$ go run 15_rsa_verify_data_pss_from_base64.go pem.go helpers.go errors.go test3.pem.pub 'Hello World' U0J3zvhiOlwn7qlQLjCyUvGZhu9j7w2jICTYq7jpuxqRAM/1Rk9aCFQRZdW10MblX+fzc3lNDSfiR1QkFu7NHdPW40Ai3bAq7SPyzn38GX38pVqmHIui3DwbtTHrcYhYZhEw5GHQOlTWhS2giS7Oyvd4iqCPWrcP0e9NP4L1z1ZDC2KLOpNJkgon5hfiryGR0YCa9TtrxJaWmFySV4z2dpJUpcvy+Ca8fwcgr/JLhomP7JQdv7iSvp0Tlubw3v48tiKxcYak2cj27gbAa4d7E+798QWE09OIdcFbmio3IuJuC0e6E9OMDqEuK/dzLHHqw2A5BfHhve8eXnDbH5D8Ja==
crypto/rsa: verification error
exit status 20
```

## Hashing

### SHA512

```bash
$ go run 16_sha512_hash_to_base64.go helpers.go 'Hello World'
LHT9F+2v2A6ER7DUZ0HuJDt+t03SFJoKsbkkb7MDgvJ+hT2FhXGeDmfL2g2qj1FnEGRhXWRa4nrLFb+xRH9Fmw==
```

### HMAC

```bash
$ HMAC_KEY=0123456789012345 go run 17_hmac_to_base64.go helpers.go 'Hello World'
rxeUOuUKx3uVVO6qKQt+jawwRMvglw62D02h3ZTFfb3fz2gL0k29nmmYpe6n09X+LtZFdG4tTp8EOlM1jY75+A==

$ HMAC_KEY='0123456789012345' go run 18_verify_hmac_from_base64.go helpers.go errors.go 'Hello World' rxeUOuUKx3uVVO6qKQt+jawwRMvglw62D02h3ZTFfb3fz2gL0k29nmmYpe6n09X+LtZFdG4tTp8EOlM1jY75+A==
Signature Verification Succeeded

$ HMAC_KEY=0123456789012345 go run 18_verify_hmac_from_base64.go helpers.go errors.go 'Hello World' rxeUOuUKx3uVVO6qKQt+jawwRMvglw62D02h3ZTFfb3fz2gL0k29nmmYpe6n09X+LtZFdG4tTp8EOlM1jY75+Z==
Signature Verification Failed
exit status 19
```

### HMAC Content Signatures

```bash
$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 19_aes_encrypt_hmac_sign_content_to_base64.go aes.go helpers.go errors.go 'Hello World'
jz7vcg+B0kjBhTp5TfuV2XgrpA0JOKyAZcB9o+Qp7Jj4cE01W8hFfTLaqs3WFNM89MTfwAn9stUI9KfTRw3UOQ==f+tGhML1HsD0qqfQ+Gi+wNLRCskGtlXi+nhPEne73qM=

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 20_aes_decrypt_hmac_verify_content_from_base64.go aes.go hmac.go helpers.go errors.go 'Hello World' jz7vcg+B0kjBhTp5TfuV2XgrpA0JOKyAZcB9o+Qp7Jj4cE01W8hFfTLaqs3WFNM89MTfwAn9stUI9KfTRw3UOQ==f+tGhML1HsD0qqfQ+Gi+wNLRCskGtlXi+nhPEne73qM=
Signature Verification Succeeded

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 20_aes_decrypt_hmac_verify_content_from_base64.go aes.go hmac.go helpers.go errors.go 'Hello World' jz7vcg+B0kjBhTp5TfuV2XgrpA0JOKyAZcB9o+Qp7Jj4cE01W8hFfTLaqs3WFNM89MTfwAn9stUI9KfTRw3UOQ==f+tGhML1HsD0qqfQ+Gi+wNLRCskGtlXi+nhPEne73qm=
error: content doesn't match
exit status 21

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 20_aes_decrypt_hmac_verify_content_from_base64.go aes.go hmac.go helpers.go errors.go 'Hello World' jz7vcg+B0kjBhTp5TfuV2XgrpA0JOKyAZcB9o+Qp7Jj4cE01W8hFfTLaqs3WFNM89MTfwAn9stUI9KfTRw3UOq==f+tGhML1HsD0qqfQ+Gi+wNLRCskGtlXi+nhPEne73qM=
Signature Verification Failed
exit status 20

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 21_aes_decrypt_hmac_verify_content_from_base64.go aes.go hmac.go helpers.go errors.go jz7vcg+B0kjBhTp5TfuV2XgrpA0JOKyAZcB9o+Qp7Jj4cE01W8hFfTLaqs3WFNM89MTfwAn9stUI9KfTRw3UOQ==f+tGhML1HsD0qqfQ+Gi+wNLRCskGtlXi+nhPEne73qM=
Signature Verification Succeeded

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 21_aes_decrypt_hmac_verify_content_from_base64.go aes.go hmac.go helpers.go errors.go jz7vcg+B0kjBhTp5TfuV2XgrpA0JOKyAZcB9o+Qp7Jj4cE01W8hFfTLaqs3WFNM89MTfwAn9stUI9KfTRw3UOQ==f+tGhML1HsD0qqfQ+Gi+wNLRCskGtlXi+nhPEne73qm=
Signature Verification Failed
exit status 20

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 21_aes_decrypt_hmac_verify_content_from_base64.go aes.go hmac.go helpers.go errors.go jz7vcg+B0kjBhTp5TfuV2XgrpA0JOKyAZcB9o+Qp7Jj4cE01W8hFfTLaqs3WFNM89MTfwAn9stUI9KfTRw3UOq==f+tGhML1HsD0qqfQ+Gi+wNLRCskGtlXi+nhPEne73qM=
Signature Verification Failed
exit status 20
```

### HMAC Wrapper Signatures

```bash
$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 22_aes_encrypt_hmac_sign_wrapper_to_base64.go aes.go hmac.go helpers.go errors.go 'Hello World'
dj82aiViZEVUZr++Q3Zz5Ix5Q7D0G7NX3DjyMm+i9O72/nXGcc6WKuqukg2n+c8TmlkC2t/rjXpLFmq5kNRsyw==tBpwJs9FZ0A+QPWoks7udfz6P0YA8h+vbwlduRcitS4=

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 23_aes_decrypt_hmac_verify_wrapper_from_base64.go aes.go hmac.go helpers.go errors.go dj82aiViZEVUZr++Q3Zz5Ix5Q7D0G7NX3DjyMm+i9O72/nXGcc6WKuqukg2n+c8TmlkC2t/rjXpLFmq5kNRsyw==tBpwJs9FZ0A+QPWoks7udfz6P0YA8h+vbwlduRcitS4=
Signature Verification Succeeded

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 23_aes_decrypt_hmac_verify_wrapper_from_base64.go aes.go hmac.go helpers.go errors.go dj82aiViZEVUZr++Q3Zz5Ix5Q7D0G7NX3DjyMm+i9O72/nXGcc6WKuqukg2n+c8TmlkC2t/rjXpLFmq5kNRsyw==tBpwJs9FZ0A+QPWoks7udfz6P0YA8h+vbwlduRcitS0=
Signature Verification Failed
exit status 20

$ AES_KEY=0123456789012345 HMAC_KEY=0987654321098765 go run 23_aes_decrypt_hmac_verify_wrapper_from_base64.go aes.go hmac.go helpers.go errors.go dj82aiViZEVUZr++Q3Zz5Ix5Q7D0G7NX3DjyMm+i9O72/nXGcc6WKuqukg2n+c8TmlkC2t/rjXpLFmq5kNRsyW==tBpwJs9FZ0A+QPWoks7udfz6P0YA8h+vbwlduRcitS4=
Signature Verification Failed
exit status 20
```

### HMAC signed value chains

```bash
$ HMAC_KEY=0123456789012345 go run 24_chain_list_hmac_sign_to_base64.go helpers.go errors.go Hello World Goodbye World
nDo7/2Jel9aP4RtT/v2vzkqvAuZC5mGjvsNKGGXofqZEMumfFZoIZDqFHLdMqsCQr2n1ujeKyXlvURO/iDAa5w== Hello
RX7rJOTD89VN7Kn/mWo6RTRaQPMkRexAjYH8w27NTG3Vfn9ST2oeguGuraOmHM8KXQLCu08tgaUYnCcy0wH/fQ== World
vt1b+boo+Qnpnx2E/knshgPAHmvXph3dHRacUXtn+bIOS61v2UyPpBFuK/ck5Uz40e7q+mC3kWcT3zWtMzKOXg== Goodbye
wl0TqHTzm5s3+FZYh5R7+rfkFfy4ffjY1LVPUEHO3DlHXOnRbRusEcfiTpPz41QEjvQ6Ywqb3RI3ugHL89DLyg== World

$ HMAC_KEY=0123456789012345 go run 24_chain_list_hmac_sign_to_base64.go helpers.go errors.go Goodbye World Hello World
36Mee0SiQQr4yeJkgBJZfP4eMlRVG9O+FFcwNG6GlfkQOgvBXabzW0P+3ZqJ9qvppJzww2pMXZg4XsfOpaitpA== Goodbye
vuAFbTAZgmuFdo3xpZqY6n7TAq/HNA3PqXrCmgA+SsCMDzlc4GgfPsURx5IFiDY3RIdFoXnivRih/KIHQr+vwQ== World
p18Nb4M1msr7eIv27eIrMDaOUTX4qopGm5nL9T6QS6/PrXMI3at3Wq0vqbrkb9kRNFuWSzgFERyRax7GvBsajw== Hello
wl0TqHTzm5s3+FZYh5R7+rfkFfy4ffjY1LVPUEHO3DlHXOnRbRusEcfiTpPz41QEjvQ6Ywqb3RI3ugHL89DLyg== World

$ HMAC_KEY=0123456789012345 go run 25_chain_list_hmac_verify_from_base64.go hmac.go helpers.go errors.go Hello nDo7/2Jel9aP4RtT/v2vzkqvAuZC5mGjvsNKGGXofqZEMumfFZoIZDqFHLdMqsCQr2n1ujeKyXlvURO/iDAa5w== World RX7rJOTD89VN7Kn/mWo6RTRaQPMkRexAjYH8w27NTG3Vfn9ST2oeguGuraOmHM8KXQLCu08tgaUYnCcy0wH/fQ== Goodbye vt1b+boo+Qnpnx2E/knshgPAHmvXph3dHRacUXtn+bIOS61v2UyPpBFuK/ck5Uz40e7q+mC3kWcT3zWtMzKOXg== World wl0TqHTzm5s3+FZYh5R7+rfkFfy4ffjY1LVPUEHO3DlHXOnRbRusEcfiTpPz41QEjvQ6Ywqb3RI3ugHL89DLyg==
Signature Verification Succeeded

$ HMAC_KEY=0123456789012345 go run 25_chain_list_hmac_verify_from_base64.go hmac.go helpers.go errors.go Hello nDo7/2Jel9aP4RtT/v2vzkqvAuZC5mGjvsNKGGXofqZEMumfFZoIZDqFHLdMqsCQr2n1ujeKyXlvURO/iDAa5w== World RX7rJOTD89VN7Kn/mWo6RTRaQPMkRexAjYH8w27NTG3Vfn9ST2oeguGuraOmHM8KXQLCu08tgaUYnCcy0wH/fQ== Goodbye vt1b+boo+Qnpnx2E/knshgPAHmvXph3dHRacUXtn+bIOS61v2UyPpBFuK/ck5Uz40e7q+mC3kWcT3zWtMzKOXg== World wl0TqHTzm5s3+FZYh5R7+rfkFfy4ffjY1LVPUEHO3DlHXOnRbRusEcfiTpPz41QEjvQ6Ywqb3RI3ugHL89DLyG==
Signature Verification Failed
wl0TqHTzm5s3+FZYh5R7+rfkFfy4ffjY1LVPUEHO3DlHXOnRbRusEcfiTpPz41QEjvQ6Ywqb3RI3ugHL89DLyg== != wl0TqHTzm5s3+FZYh5R7+rfkFfy4ffjY1LVPUEHO3DlHXOnRbRusEcfiTpPz41QEjvQ6Ywqb3RI3ugHL89DLyG==
exit status 20

$ HMAC_KEY=0123456789012345 go run 25_chain_list_hmac_verify_from_base64.go hmac.go helpers.go errors.go Hello nDo7/2Jel9aP4RtT/v2vzkqvAuZC5mGjvsNKGGXofqZEMumfFZoIZDqFHLdMqsCQr2n1ujeKyXlvURO/iDAa5w== World RX7rJOTD89VN7Kn/mWo6RTRaQPMkRexAjYH8w27NTG3Vfn9ST2oeguGuraOmHM8KXQLCu08tgaUYnCcy0wH/fQ== Goodbye vt1b+boo+Qnpnx2E/knshgPAHmvXph3dHRacUXtn+bIOS61v2UyPpBFuK/ck5Uz40e7q+mC3kWcT3zWtMzKOXG== World wl0TqHTzm5s3+FZYh5R7+rfkFfy4ffjY1LVPUEHO3DlHXOnRbRusEcfiTpPz41QEjvQ6Ywqb3RI3ugHL89DLyg==
Signature Verification Failed
vt1b+boo+Qnpnx2E/knshgPAHmvXph3dHRacUXtn+bIOS61v2UyPpBFuK/ck5Uz40e7q+mC3kWcT3zWtMzKOXg== != vt1b+boo+Qnpnx2E/knshgPAHmvXph3dHRacUXtn+bIOS61v2UyPpBFuK/ck5Uz40e7q+mC3kWcT3zWtMzKOXG==
exit status 20

$ HMAC_KEY=0123456789012345 go run 25_chain_list_hmac_verify_from_base64.go hmac.go helpers.go errors.go Hello nDo7/2Jel9aP4RtT/v2vzkqvAuZC5mGjvsNKGGXofqZEMumfFZoIZDqFHLdMqsCQr2n1ujeKyXlvURO/iDAa5w== World RX7rJOTD89VN7Kn/mWo6RTRaQPMkRexAjYH8w27NTG3Vfn9ST2oeguGuraOmHM8KXQLCu08tgaUYnCcy0wH/fq== Goodbye vt1b+boo+Qnpnx2E/knshgPAHmvXph3dHRacUXtn+bIOS61v2UyPpBFuK/ck5Uz40e7q+mC3kWcT3zWtMzKOXg== World wl0TqHTzm5s3+FZYh5R7+rfkFfy4ffjY1LVPUEHO3DlHXOnRbRusEcfiTpPz41QEjvQ6Ywqb3RI3ugHL89DLyg==
Signature Verification Failed
RX7rJOTD89VN7Kn/mWo6RTRaQPMkRexAjYH8w27NTG3Vfn9ST2oeguGuraOmHM8KXQLCu08tgaUYnCcy0wH/fQ== != RX7rJOTD89VN7Kn/mWo6RTRaQPMkRexAjYH8w27NTG3Vfn9ST2oeguGuraOmHM8KXQLCu08tgaUYnCcy0wH/fq==
exit status 20

$ HMAC_KEY=0123456789012345 go run 25_chain_list_hmac_verify_from_base64.go hmac.go helpers.go errors.go Hello nDo7/2Jel9aP4RtT/v2vzkqvAuZC5mGjvsNKGGXofqZEMumfFZoIZDqFHLdMqsCQr2n1ujeKyXlvURO/iDAa5W== World RX7rJOTD89VN7Kn/mWo6RTRaQPMkRexAjYH8w27NTG3Vfn9ST2oeguGuraOmHM8KXQLCu08tgaUYnCcy0wH/fQ== Goodbye vt1b+boo+Qnpnx2E/knshgPAHmvXph3dHRacUXtn+bIOS61v2UyPpBFuK/ck5Uz40e7q+mC3kWcT3zWtMzKOXg== World wl0TqHTzm5s3+FZYh5R7+rfkFfy4ffjY1LVPUEHO3DlHXOnRbRusEcfiTpPz41QEjvQ6Ywqb3RI3ugHL89DLyg==
Signature Verification Failed
nDo7/2Jel9aP4RtT/v2vzkqvAuZC5mGjvsNKGGXofqZEMumfFZoIZDqFHLdMqsCQr2n1ujeKyXlvURO/iDAa5w== != nDo7/2Jel9aP4RtT/v2vzkqvAuZC5mGjvsNKGGXofqZEMumfFZoIZDqFHLdMqsCQr2n1ujeKyXlvURO/iDAa5W==
exit status 20
```

### Merkle Trees

```bash
$ HMAC_KEY=0123456789012345 go run 26_merkle_tree_hmac_sign_to_base64.go helpers.go errors.go
NP0n+gNreXmqI4ATzz5rOcut9+317H1Ab6CgO0KLNJMQpS/vT4GDOhAiRMf+KUvMV+f9sApyc4V4iOng/1RXEw== +
m3M6Y+orFgsXEZ0/HpZJME9cKy+/SpzQ/zvxLujK+CS3xqVX+x0QzIiwYikLswjWUUWgGR6splkAD7z0BLZExg== *
X0gyp/RcB31yImaxxDDo7L0mN83WYc3DtFTETS4FV3CKu7bgLHwolraIgcCE8VU+reP9Vh1tMq+9m1XALg0BNw== -
PJwjHeAcTOaFAswC0nYD6wUaptsYB5wptVbdfBRTz8QRthFsJIqTPUWSGgcbVQWQMqTWgpP2es56jwtJnQLhJQ== x
o/vbZre3SZqFmUqNP+x1dB6zJiYwcJl5pBSEfDHyQ/p6W2caSW4JbjhydykpWrKe2lwYJ6kQ9H1OqX3C7TzPYw== y
IqEsHAfsvx4w/06gqbvQoU0iD+sIUXTtcunF4nyEWBPLUEsBeg7aUyID3Kmy1uutz5To/Mrjb3pu5BqhdLX/rA== 2
npF9VUehurfZWBax5ZUuqKOfE79yg7rPJF0aSQ7u7i6G/POVx95rKJV0qFOYKDbGyMwxJ8BUupcCfnC1p0WGOQ== 7
lCo6zpPpmoEaF9xW7a+ch0dKC+HzlOUGxpVvIorjXIen0H2iphehmZSEhK9LxMbWxE+D64N8/Xvi2cl91DDr2Q== 1

$ HMAC_KEY=0123456789012345 go run 27_merkle_tree_hmac_verify_from_base64.go hmac.go helpers.go errors.go NP0n+gNreXmqI4ATzz5rOcut9+317H1Ab6CgO0KLNJMQpS/vT4GDOhAiRMf+KUvMV+f9sApyc4V4iOng/1RXEw== m3M6Y+orFgsXEZ0/HpZJME9cKy+/SpzQ/zvxLujK+CS3xqVX+x0QzIiwYikLswjWUUWgGR6splkAD7z0BLZExg== X0gyp/RcB31yImaxxDDo7L0mN83WYc3DtFTETS4FV3CKu7bgLHwolraIgcCE8VU+reP9Vh1tMq+9m1XALg0BNw== PJwjHeAcTOaFAswC0nYD6wUaptsYB5wptVbdfBRTz8QRthFsJIqTPUWSGgcbVQWQMqTWgpP2es56jwtJnQLhJQ== o/vbZre3SZqFmUqNP+x1dB6zJiYwcJl5pBSEfDHyQ/p6W2caSW4JbjhydykpWrKe2lwYJ6kQ9H1OqX3C7TzPYw== IqEsHAfsvx4w/06gqbvQoU0iD+sIUXTtcunF4nyEWBPLUEsBeg7aUyID3Kmy1uutz5To/Mrjb3pu5BqhdLX/rA== npF9VUehurfZWBax5ZUuqKOfE79yg7rPJF0aSQ7u7i6G/POVx95rKJV0qFOYKDbGyMwxJ8BUupcCfnC1p0WGOQ== lCo6zpPpmoEaF9xW7a+ch0dKC+HzlOUGxpVvIorjXIen0H2iphehmZSEhK9LxMbWxE+D64N8/Xvi2cl91DDr2Q==
Signature Verification Succeeded

$ HMAC_KEY=0123456789012345 go run 27_merkle_tree_hmac_verify_from_base64.go hmac.go helpers.go errors.go NP0n+gNreXmqI4ATzz5rOcut9+317H1Ab6CgO0KLNJMQpS/vT4GDOhAiRMf+KUvMV+f9sApyc4V4iOng/1RXEw== m3M6Y+orFgsXEZ0/HpZJME9cKy+/SpzQ/zvxLujK+CS3xqVX+x0QzIiwYikLswjWUUWgGR6splkAD7z0BLZExg== X0gyp/RcB31yImaxxDDo7L0mN83WYc3DtFTETS4FV3CKu7bgLHwolraIgcCE8VU+reP9Vh1tMq+9m1XALg0BNw== PJwjHeAcTOaFAswC0nYD6wUaptsYB5wptVbdfBRTz8QRthFsJIqTPUWSGgcbVQWQMqTWgpP2es56jwtJnQLhJQ== o/vbZre3SZqFmUqNP+x1dB6zJiYwcJl5pBSEfDHyQ/p6W2caSW4JbjhydykpWrKe2lwYJ6kQ9H1OqX3C7TzPYw== IqEsHAfsvx4w/06gqbvQoU0iD+sIUXTtcunF4nyEWBPLUEsBeg7aUyID3Kmy1uutz5To/Mrjb3pu5BqhdLX/rA== npF9VUehurfZWBax5ZUuqKOfE79yg7rPJF0aSQ7u7i6G/POVx95rKJV0qFOYKDbGyMwxJ8BUupcCfnC1p0WGOQ== lCo6zpPpmoEaF9xW7a+ch0dKC+HzlOUGxpVvIorjXIen0H2iphehmZSEhK9LxMbWxE+D64N8/Xvi2cl91DDr2q==
Signature Verification Failed
lCo6zpPpmoEaF9xW7a+ch0dKC+HzlOUGxpVvIorjXIen0H2iphehmZSEhK9LxMbWxE+D64N8/Xvi2cl91DDr2Q== != lCo6zpPpmoEaF9xW7a+ch0dKC+HzlOUGxpVvIorjXIen0H2iphehmZSEhK9LxMbWxE+D64N8/Xvi2cl91DDr2q==
exit status 20
```

## Hybrid Cryptography

```
$ go run 28_goroutine_ping_pong.go 1 2 3
B: 1
A: 1
B: 2
A: 2
B: 3
A: 3

$ go run 29_hybrid_crypto_public_key_request.go pem.go rsa.go aes.go helpers.go errors.go test3.pem session_label
Bob received nonce: session_label
Alice received public key: &{31567127335276920051778986890983418137235118195788341168325280010141872825942730386604890414504330424121846232915529846111962553425801730338442721710329105996928572139897441647445680540638619159385964576015724962916475593889066666184228021802959434558601451977786671232463588955484699002568869232321539136212150321079973869708311234964749352015864656432491046294093762734492605713775237154595184621558656835799069556866308339741960587974969797335759107872113116845565849526376848130079712356076286876918399283180445566867123603330810480774172961213531196691266615312515557045027778553516298941059154014833011840984679 65537}

$ go run 30_hybrid_crypto_key_exchange.go pem.go rsa.go aes.go helpers.go errors.go test3.pem session_label 0123456789012345 A B C
Bob received nonce: session_label
Alice received public key: &{31567127335276920051778986890983418137235118195788341168325280010141872825942730386604890414504330424121846232915529846111962553425801730338442721710329105996928572139897441647445680540638619159385964576015724962916475593889066666184228021802959434558601451977786671232463588955484699002568869232321539136212150321079973869708311234964749352015864656432491046294093762734492605713775237154595184621558656835799069556866308339741960587974969797335759107872113116845565849526376848130079712356076286876918399283180445566867123603330810480774172961213531196691266615312515557045027778553516298941059154014833011840984679 65537}
Bob received symmetric key: 0123456789012345

$ go run 31_hybrid_crypto_send_messages.go pem.go rsa.go aes.go helpers.go errors.go test3.pem session_label 0123456789012345 A B C
Bob received nonce: session_label
Bob received symmetric key: 0123456789012345
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

$ go run 32_hybrid_crypto_dual_key_streams.go pem.go rsa.go aes.go helpers.go errors.go test3.pem session_label 0123456789012345 A B C
Server received nonce: session_label
Bob received symmetric key: 0123456789012345
Bob wants to say: A
Alice received symmetric key: HbnT44/p+mKtHhg0eZc2gHEt8NMaMehpdWbl64ZdYOs=
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

$ go run 33_web_service_ping_pong.go helpers.go errors.go 1 2 3 4 5 6
B: 1
A: 1
B: 2
A: 2
B: 3
A: 3
B: 4
A: 4
B: 5
A: 5
B: 6
A: 6
```

## SQLite3 Databases

### Create Database

```bash
$ ls -al test.db
ls: test.db: No such file or directory

$ go run 40_sqlite3_create_database.go helpers.go errors.go test.db
$ ls -al test.db
-rw-r--r--  1 eleanor  admin  0 16 Oct 02:11 test.db

$ go run 41_sqlite3_create_database.go helpers.go errors.go test2.db
$ ls -al *.db
-rw-r--r--  1 eleanor  admin  0 16 Oct 02:11 test.db
-rw-r--r--  1 eleanor  admin  0 16 Oct 13:03 test2.db

$ go run 42_sqlite3_create_database.go helpers.go errors.go test3
$ ls -al *.db
-rw-r--r--  1 eleanor  admin  0 16 Oct 19:26 test.db
-rw-r--r--  1 eleanor  admin  0 16 Oct 19:26 test2.db
-rw-r--r--  1 eleanor  admin  0 16 Oct 20:37 test3.db

$ go run 43_sqlite3_create_database.go helpers.go errors.go test4
$ ls -al *.db
-rw-r--r--  1 eleanor  admin  0 16 Oct 19:26 test.db
-rw-r--r--  1 eleanor  admin  0 16 Oct 19:26 test2.db
-rw-r--r--  1 eleanor  admin  0 16 Oct 20:37 test3.db
-rw-r--r--  1 eleanor  admin  0 16 Oct 21:06 test4.db
```

### Create Table

```bash
$ go run 44_sqlite3_create_table.go helpers.go errors.go test.db
$ go run 44_sqlite3_create_table.go helpers.go errors.go test.db
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
```

### Drop Table, Insert, Select

```bash
$ go run 45_sqlite3_create_table.go helpers.go errors.go test.db
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
```

### Create Table Index

```bash
$ go run 46_sqlite3_create_index.go helpers.go errors.go test.db
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

$ go run 47_sqlite3_create_index.go helpers.go errors.go test.db
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
```

## Encrypted Databases

### Random Field ID

```bash
$ go run 50_table_with_random_ids.go helpers.go errors.go test.db
rows in Account table = 3

$ sqlite3 test.db
SQLite version 3.29.0 2019-07-10 17:32:03
Enter ".help" for usage hints.
sqlite> select * from Account;
VlT/MgPQ9r7oll7X31DdQg==|Gamma|g@someserver.com
Ws0y+Bnk33C/1FWTNaPxDQ==|Beta|b@someserver.com
i5W/q5yjLO/zl8L5b2nIOQ==|Alpha|a@someserver.com

sqlite> .quit
```

### AES Encrypted Fields

```bash
$ AES_KEY=0123456789012345 go run 51_table_with_encrypted_fields_add_records.go aes.go helpers.go errors.go test.db
rows in Account table = 3

$ sqlite3 test.db
SQLite version 3.29.0 2019-07-10 17:32:03
Enter ".help" for usage hints.
sqlite> select * from Account;
E10a+ftltRrE11NazKmqQw==|voEyPVgO0IzwZRr0KVQvnhAdp1xDdjumqndGTg3+PqY=|hz7giZpDSrQj0pyGU+6mA0uoi3ZIhaDplbzDKV7jWRuXLNDZBU+Oc1IrSFydp8lm
UoXYs2XFZn+0NwBPBUck1Q==|fSLD6i/NTwIUoJK6Y+cg8yEAQGo16P8dLPlR0NJgJbU=|/VXUhUvE3AVqabbHtode4omHzP++OLyT2z51A3rRJTbcvA67N36RaJ88WjoltrFs
lQmkb9x1c3RzgnKWjcj0GQ==|08/u0jj7uYG3q0jGNr+STXGa4OJe5ScD1Biig4y99c8=|8NQi5V0Jed4B/MLJpscdoFL9gYezu6TXrq3qFXFUlLa47G+sePF84noglG3UBwo/

sqlite> .quit
```
