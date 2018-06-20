package gopherdb

import (
	"crypto/sha1"
	"math/big"
	"fmt"
)

func  query(key string) error {
	index := hashHelper(key, 4)
	fmt.Println(index)
	return nil
}

func hashHelper(key string, size int) int64 {
	hasher := sha1.New()
	var hashval big.Int
	hash := hasher.Sum([]byte(key))
	hashval.SetBytes(hash)
	var remainder big.Int
	remainder.Mod(&hashval, big.NewInt(int64(size)))
	return remainder.Int64()
}
