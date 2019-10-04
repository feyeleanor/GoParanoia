# GoParanoia
Code from the workshop "Adventures in Paranoia with Go and SQLite"


## AES Encryption and Decryption

$ AES_KEY="0123456789012345" go run 01_encrypt_aes.go "Hello World"

[189 66 213 32 192 240 163 198 89 18 112 190 150 234 133 248 10 214 159 253 11 245 129 197 68 225 13 23 141 226 137 82]

$ AES_KEY="0123456789012345" go run 02_decrypt_aes.go 189 66 213 32 192 240 163 198 89 18 112 190 150 234 133 248 10 214 159 253 11 245 129 197 68 225 13 23 141 226 137 82

Hello World
  		
$ PEM_KEY="0123456789012345" go run 03_generate_rsa_keypair.go 1024

[45 45 45 45 45 66 69 71 73 78 32 82 83 65 32 80 82 73 86 65 84 69 32 75 69 89 45 45 45 45 45 10 80 114 111 99 45 84 121 112 101 58 32 52 44 69 78 67 82 89 80 84 69 68 10 68 69 75 45 73 110 102 111 58 32 65 69 83 45 50 53 54 45 67 66 67 44 51 49 100 48 54 50 101 101 54 54 54 99 52 52 57 52 54 49 56 97 50 55 55 97 55 49 102 57 53 52 97 57 10 10 87 74 98 112 85 105 83 69 73 100 113 66 100 48 117 81 71 73 107 113 75 51 83 53 81 68 102 106 56 53 115 103 74 67 83 47 89 86 57 56 102 74 102 105 69 104 106 98 49 85 111 71 68 48 49 80 104 104 113 53 80 66 87 50 10 55 65 83 88 76 82 103 121 68 66 79 55 86 79 67 102 115 87 83 119 90 108 114 119 43 50 113 84 51 52 117 97 73 110 43 68 110 112 57 105 57 114 113 51 74 55 51 115 43 84 78 109 82 78 108 81 101 113 98 113 106 107 53 121 10 116 65 100 79 115 119 43 47 108 53 105 113 112 105 51 88 110 109 116 121 79 70 67 88 89 101 104 98 86 101 99 107 77 79 99 101 105 84 114 85 97 76 89 84 52 83 65 82 84 120 68 112 116 108 86 111 65 114 104 87 87 112 84 78 10 101 111 49 99 112 113 67 82 53 116 53 73 73 70 105 90 66 66 67 109 99 75 102 76 82 74 118 80 113 106 54 113 53 53 113 99 97 116 106 47 100 43 115 51 55 71 72 74 109 48 71 66 80 76 66 98 110 80 80 78 48 99 81 80 10 118 114 119 49 54 120 120 100 117 113 67 98 116 73 67 116 109 70 57 84 121 57 100 99 74 101 85 65 90 75 118 73 103 84 110 47 115 78 82 52 88 108 72 53 76 75 122 113 117 70 57 88 43 73 104 80 79 89 90 111 53 115 43 119 10 69 100 54 119 49 112 118 110 102 81 75 100 106 57 117 65 84 80 109 87 111 116 49 100 119 115 47 82 54 84 103 112 98 51 120 43 120 90 117 78 87 51 103 81 98 97 115 109 67 113 104 114 110 87 104 48 47 110 79 71 113 57 49 99 10 121 89 110 86 73 100 78 75 105 120 87 88 121 105 84 110 103 70 66 75 106 76 80 118 57 43 79 98 102 50 90 47 48 53 111 117 52 111 90 55 108 73 65 68 102 119 108 72 67 79 67 70 78 81 86 86 89 72 47 98 43 99 86 106 10 120 110 98 65 65 80 90 109 76 100 85 118 76 82 76 54 98 87 111 111 56 71 116 115 77 98 101 120 108 66 54 43 110 43 122 121 80 105 120 106 86 78 69 115 114 56 108 56 82 119 48 82 118 120 108 48 112 47 49 69 113 101 51 73 10 104 43 86 102 106 116 100 100 66 87 69 69 122 55 117 87 54 67 118 70 115 55 84 53 116 79 65 98 121 43 66 70 73 113 73 113 57 51 120 88 52 70 106 113 104 69 69 112 86 109 115 112 119 84 71 90 106 67 52 82 115 69 81 57 10 105 84 68 117 53 79 83 109 79 109 71 106 86 112 89 100 65 68 48 53 118 102 120 47 106 120 77 73 50 76 83 113 105 49 70 116 85 76 122 116 88 109 67 90 87 100 88 75 98 81 88 76 110 98 48 98 80 99 86 86 49 53 101 73 10 102 82 104 82 115 113 86 111 101 86 68 52 76 74 52 56 83 122 97 103 79 71 113 121 88 68 101 100 85 106 79 77 57 117 79 119 49 83 99 86 103 67 75 89 121 43 122 105 72 81 77 110 104 82 81 86 77 68 65 67 106 81 86 74 10 74 105 106 47 88 89 82 99 113 66 97 72 81 89 82 99 53 120 80 49 121 52 117 90 72 102 99 107 54 98 85 69 122 70 90 117 105 88 49 120 106 102 70 121 106 117 120 100 76 54 104 69 75 109 121 111 66 57 87 48 97 107 104 98 10 109 108 122 116 79 108 99 78 54 52 75 73 56 50 72 120 66 117 88 89 113 112 73 107 104 50 81 114 43 74 65 104 112 108 110 85 79 105 90 73 84 113 89 61 10 45 45 45 45 45 69 78 68 32 82 83 65 32 80 82 73 86 65 84 69 32 75 69 89 45 45 45 45 45 10]

$ PEM_KEY="0123456789012345" go run 04_generate_rsa_keypair.go test.pem 1024

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

$ go run 04_generate_rsa_keypair.go test2.pem 1024

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

$ PEM_KEY="0123456789012345" go run 05_extract_rsa_pubkey.go test.pem

$ cat test.pem.pub
```
-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAOJeXdQw2dRkYxPwAgGfLh8cDKlMW7I1yksd4sqYOGb9TF4UYg3IOm9Q
+gLhn8WWt72ua7pcEy2MPDUTXUB6WohhwiP+b384W2M63KlBBWh23S5z4mwOfRU4
IHO/qjMQT6j2o1zbGd7Fie0wiujQREegta5jnb67zt2OdDxAQpeNAgMBAAE=
-----END RSA PUBLIC KEY-----
```

$ go run 05_extract_rsa_pubkey.go test2.pem

$ cat test2.pem.pub 
```
-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAMloI+V5GjOo3129h+QeVU1uqbsMAxqK2lFe4i0l4YvavfdUbw0XrE8i
CJLaMMVVFV4pcpnKvLa67STDykbl2Dzg7+0rbei62WzXf6aYyjMBED2BbitZW8/O
XKhvWaZ89ahzdgiW9Eo/9M5gnav6zj2eKsl7RP/bROcOwl0HVZDZAgMBAAE=
-----END RSA PUBLIC KEY-----
```

$ go run 06_rsa_encrypt_data_pkcs1.go test2.pem.pub "Hello World"

[154 5 176 158 66 131 236 171 170 64 127 249 90 20 20 240 207 92 178 103 174 71 185 70 79 39 109 114 109 115 175 150 146 173 48 90 196 196 21 210 127 192 74 111 34 111 63 162 196 120 60 234 40 8 88 33 147 201 21 72 217 162 68 31 215 203 63 17 84 104 154 44 180 110 245 90 118 30 225 193 223 143 106 175 217 47 9 41 255 252 235 221 61 97 2 111 77 208 212 64 151 148 26 171 163 127 78 154 193 197 159 117 87 60 96 255 65 92 195 193 171 85 38 83 146 220 197 78]

$ go run 07_rsa_decrypt_data_pkcs1.go test2.pem 154 5 176 158 66 131 236 171 170 64 127 249 90 20 20 240 207 92 178 103 174 71 185 70 79 39 109 114 109 115 175 150 146 173 48 90 196 196 21 210 127 192 74 111 34 111 63 162 196 120 60 234 40 8 88 33 147 201 21 72 217 162 68 31 215 203 63 17 84 104 154 44 180 110 245 90 118 30 225 193 223 143 106 175 217 47 9 41 255 252 235 221 61 97 2 111 77 208 212 64 151 148 26 171 163 127 78 154 193 197 159 117 87 60 96 255 65 92 195 193 171 85 38 83 146 220 197 78

Hello World

$ go run 06_rsa_encrypt_data_pkcs1.go test.pem.pub "Hello World"

[12 137 31 32 242 3 70 242 75 74 238 176 150 100 253 128 178 8 69 158 153 77 5 125 223 45 66 124 206 212 231 115 175 193 33 140 82 202 80 209 31 176 215 146 28 166 106 232 240 212 106 105 139 90 175 120 156 207 210 254 122 196 136 201 153 220 129 242 203 227 51 167 238 17 231 68 65 83 77 213 72 143 190 136 116 174 139 9 197 51 147 186 25 175 173 136 195 142 23 86 93 117 47 106 59 93 14 99 66 234 86 145 12 226 254 209 164 155 110 204 243 154 138 83 244 129 42 204]

$ PEM_KEY="0123456789012345" go run 07_rsa_decrypt_data_pkcs1.go test.pem 12 137 31 32 242 3 70 242 75 74 238 176 150 100 253 128 178 8 69 158 153 77 5 125 223 45 66 124 206 212 231 115 175 193 33 140 82 202 80 209 31 176 215 146 28 166 106 232 240 212 106 105 139 90 175 120 156 207 210 254 122 196 136 201 153 220 129 242 203 227 51 167 238 17 231 68 65 83 77 213 72 143 190 136 116 174 139 9 197 51 147 186 25 175 173 136 195 142 23 86 93 117 47 106 59 93 14 99 66 234 86 145 12 226 254 209 164 155 110 204 243 154 138 83 244 129 42 204

Hello World

$ go run 08_rsa_sign_data_pkcs1.go test2.pem "Hello World"
[69 26 19 178 176 16 88 201 10 20 183 191 61 242 15 128 38 254 22 70 164 144 190 182 102 191 150 240 190 116 211 44 246 4 84 58 92 165 123 114 167 220 152 161 9 90 80 52 118 97 78 252 118 193 125 238 67 84 78 69 19 175 206 165 6 254 205 139 247 29 123 56 20 146 254 108 199 196 213 62 240 142 176 182 205 1 155 228 223 94 190 174 128 89 77 73 21 6 230 56 83 204 234 217 199 176 91 14 175 151 124 159 104 127 198 30 186 8 3 108 87 245 182 92 112 73 108 189]

$ go run 09_rsa_verify_data_pkcs1.go test2.pem.pub "Hello World" 69 26 19 178 176 16 88 201 10 20 183 191 61 242 15 128 38 254 22 70 164 144 190 182 102 191 150 240 190 116 211 44 246 4 84 58 92 165 123 114 167 220 152 161 9 90 80 52 118 97 78 252 118 193 125 238 67 84 78 69 19 175 206 165 6 254 205 139 247 29 123 56 20 146 254 108 199 196 213 62 240 142 176 182 205 1 155 228 223 94 190 174 128 89 77 73 21 6 230 56 83 204 234 217 199 176 91 14 175 151 124 159 104 127 198 30 186 8 3 108 87 245 182 92 112 73 108 189

Signature Verification Succeeded

$ go run 09_rsa_verify_data_pkcs1.go test2.pem.pub "Hello World" 69 26 19 178 176 16 88 201 10 20 183 191 61 242 15 128 38 254 22 70 164 144 190 182 102 191 150 240 190 116 211 44 246 4 84 58 92 165 123 114 167 220 152 161 9 90 80 52 118 97 78 252 118 193 125 238 67 84 78 69 19 175 206 165 6 254 205 139 247 29 123 56 20 146 254 108 199 196 213 62 240 142 176 182 205 1 155 228 223 94 190 174 128 89 77 73 21 6 230 56 83 204 234 217 199 176 91 14 175 151 124 159 104 127 198 30 186 8 3 108 87 245 182 92 112 73 108

Signature Verification Failed

exit status 1

$ go run 10_rsa_encrypt_data_oaep.go test3.pem.pub "test" "Hello World"

[160 101 217 128 189 116 224 8 27 109 115 36 109 149 57 68 21 132 234 204 51 2 154 56 222 101 227 182 188 102 184 173 83 226 198 252 180 179 204 214 154 244 38 112 231 216 161 33 88 158 75 134 88 236 54 183 66 129 246 61 246 213 136 124 73 148 100 5 150 136 253 227 147 231 48 42 118 55 89 224 110 53 8 182 231 124 117 225 224 222 246 18 174 189 184 146 212 169 61 81 79 238 80 186 193 45 238 251 182 23 52 15 57 30 232 11 20 58 9 86 151 149 127 9 39 15 189 8 223 54 246 230 17 144 73 66 68 213 72 115 29 149 42 170 226 225 145 26 231 175 61 248 227 140 229 79 133 218 251 37 134 147 74 47 49 67 92 117 185 97 235 121 39 34 145 38 34 7 59 158 151 53 203 233 217 194 251 42 11 109 13 42 189 91 62 174 64 183 181 83 249 216 43 155 129 208 157 152 103 236 81 14 179 9 233 247 154 12 85 130 14 124 85 8 250 84 0 25 192 208 233 173 121 88 146 159 80 124 65 129 104 232 151 136 159 212 211 236 250 135 246 11 119 27 220 48]

$ go run 11_rsa_decrypt_data_oaep.go test3.pem "test" 160 101 217 128 189 116 224 8 27 109 115 36 109 149 57 68 21 132 234 204 51 2 154 56 222 101 227 182 188 102 184 173 83 226 198 252 180 179 204 214 154 244 38 112 231 216 161 33 88 158 75 134 88 236 54 183 66 129 246 61 246 213 136 124 73 148 100 5 150 136 253 227 147 231 48 42 118 55 89 224 110 53 8 182 231 124 117 225 224 222 246 18 174 189 184 146 212 169 61 81 79 238 80 186 193 45 238 251 182 23 52 15 57 30 232 11 20 58 9 86 151 149 127 9 39 15 189 8 223 54 246 230 17 144 73 66 68 213 72 115 29 149 42 170 226 225 145 26 231 175 61 248 227 140 229 79 133 218 251 37 134 147 74 47 49 67 92 117 185 97 235 121 39 34 145 38 34 7 59 158 151 53 203 233 217 194 251 42 11 109 13 42 189 91 62 174 64 183 181 83 249 216 43 155 129 208 157 152 103 236 81 14 179 9 233 247 154 12 85 130 14 124 85 8 250 84 0 25 192 208 233 173 121 88 146 159 80 124 65 129 104 232 151 136 159 212 211 236 250 135 246 11 119 27 220 48

Hello World
	