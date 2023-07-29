package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

const (
	firstJan2023TimestampInNs = 1672531200000000000
	base                      = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	maxRandomNumber           = 1e7
)

// 29 July timestamp = 1690588800000000000
// difference = 1.80576e+16 ==> (62 ^ 9 to 62 ^ 10) ==> 10 digits
// 62 ^ 9 = 1.3537087e+16 ==> 10 ^ 16 ==> 6 June 2023 onwards (10 digits)
// 62 ^ 10 = 2.1834011e+14 ==> 10 ^ 14 ==> till 6 August 2049

// Generates a random 14 digit id.
func NewUniqueId() (string, error) {
	// 10 digits of timestamp (base62Encoded) + 4 digits of random number (base62encoded)
	// we will use crypto/rand for random number generation as it is much safer
	// the id can also be public facing, hence security is key
	b62Timestamp := base62Encode(time.Now().UnixNano() - firstJan2023TimestampInNs)

	b62RandStr, err := getRandomBase62Str(4)
	if err != nil {
		panic("error")
	}

	id := b62Timestamp + b62RandStr
	if len(id) != 14 {
		return id, fmt.Errorf("id should be of length exactly 14. Got %s of length %d", id, len(id))
	}
	return id, nil
}

func getRandomBase62Str(strLength int) (string, error) {
	var randString string
	for i := 0; i < strLength; i++ {
		randBigInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(base))))
		if err != nil {
			return "", err
		}
		randString += string(base[randBigInt.Int64()])
	}
	return randString, nil
}

func base62Encode(num int64) string {
	var b62 string
	for {
		if num == 0 {
			break
		}
		b62 = string(base[num%62]) + b62
		num /= 62
	}
	return b62
}
