package main

import "fmt"

type Hash uint

const (
	MD4         Hash = 1 + iota // import golang.org/x/crypto/md4
	MD5                         // import crypto/md5
	SHA1                        // import crypto/sha1
	SHA224                      // import crypto/sha256
	SHA256                      // import crypto/sha256
	SHA384                      // import crypto/sha512
	SHA512                      // import crypto/sha512
	MD5SHA1                     // no implementation; MD5+SHA1 used for TLS RSA
	RIPEMD160                   // import golang.org/x/crypto/ripemd160
	SHA3_224                    // import golang.org/x/crypto/sha3
	SHA3_256                    // import golang.org/x/crypto/sha3
	SHA3_384                    // import golang.org/x/crypto/sha3
	SHA3_512                    // import golang.org/x/crypto/sha3
	SHA512_224                  // import crypto/sha512
	SHA512_256                  // import crypto/sha512
	BLAKE2s_256                 // import golang.org/x/crypto/blake2s
	BLAKE2b_256                 // import golang.org/x/crypto/blake2b
	BLAKE2b_384                 // import golang.org/x/crypto/blake2b
	BLAKE2b_512                 // import golang.org/x/crypto/blake2b

)

func main() {
	fmt.Println("MD4: ", MD4)
	fmt.Println("MD5: ", MD5)
	fmt.Println("SHA1: ", SHA1)
	fmt.Println("SHA224: ", SHA224)
	fmt.Println("SHA256: ", SHA256)
	fmt.Println("SHA384: ", SHA384)
	fmt.Println("SHA512: ", SHA512)
	fmt.Println("MD5SHA1: ", MD5SHA1)
	fmt.Println("RIPEMD160: ", RIPEMD160)
	fmt.Println("SHA3_224: ", SHA3_224)
	fmt.Println("SHA3_256: ", SHA3_256)
	fmt.Println("SHA3_384: ", SHA3_384)
	fmt.Println("SHA3_512: ", SHA3_512)
	fmt.Println("SHA512_224: ", SHA512_224)
	fmt.Println("SHA512_256: ", SHA512_256)
	fmt.Println("BLAKE2s_256: ", BLAKE2s_256)
	fmt.Println("BLAKE2b_256: ", BLAKE2b_256)
	fmt.Println("BLAKE2b_384: ", BLAKE2b_384)
	fmt.Println("BLAKE2b_512: ", BLAKE2b_512)
}

/*
MD4:  1
MD5:  2
SHA1:  3
SHA224:  4
SHA256:  5
SHA384:  6
SHA512:  7
MD5SHA1:  8
RIPEMD160:  9
SHA3_224:  10
SHA3_256:  11
SHA3_384:  12
SHA3_512:  13
SHA512_224:  14
SHA512_256:  15
BLAKE2s_256:  16
BLAKE2b_256:  17
BLAKE2b_384:  18
BLAKE2b_512:  19
*/
