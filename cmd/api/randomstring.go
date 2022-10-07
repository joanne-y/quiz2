package main

import "crypto/rand"

type Tools struct{}

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+_#$-!~"

func (t *Tools) generateRandomString(length int) string {
	s := make([]rune, length)
	r := []rune(randomStringSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x := p.Uint64()
		y := uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)

}
