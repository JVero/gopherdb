package gopherdb

import (
	"crypto/sha1"
	"math/big"
)

func query(key string) error {
	hasher := sha1.New()
	println(big.NewInt(hasher.Sum([]byte(key))))
	return nil
}
