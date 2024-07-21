package core

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"github.com/spaolacci/murmur3"
	"strings"
)

func isEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func isNotEmpty(s string) bool {
	return !isEmpty(s)
}
func same(a, b string) bool {
	return a == b
}
func diff(a, b string) bool {
	return !same(a, b)
}
func invalid(a interface{}) bool { return a == nil }
func valid(a interface{}) bool   { return a != nil }
func getLink(data string) string {
	// make sure data already been not null
	a := murmur3.Sum64([]byte(data))
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, a)
	if err != nil {
		panic(err)
	}
	b := buf.Bytes()
	return base64.URLEncoding.EncodeToString(b)[:8]
}
func getID(data string) uint64 {
	return murmur3.Sum64([]byte(data))
}
