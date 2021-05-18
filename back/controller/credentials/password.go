package credentials

import (
	"crypto/sha256"
	"fmt"
)

type password struct {
	Hash string
}

func (p *password) Verify(authHeader string) bool {
	fmt.Println(p)
	return authHeader == p.Hash
}

func ByPassword(pass string) Auth {
	return &password{
		Hash: fmt.Sprintf("%x", sha256.Sum256([]byte(pass))),
	}
}
