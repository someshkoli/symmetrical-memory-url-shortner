package shortner

import (
	"math/rand"
	"net"
	"strconv"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func NewPath(randRune string) string {
	return "/g/" + randRune
}

func NewURL(host string, port int, path string) string {
	return net.JoinHostPort(host, strconv.Itoa(port)) + path
}
