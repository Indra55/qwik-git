package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func hashObjects(data []byte) (string, []byte) {
	header:= fmt.Sprintf("blob %d\x00",len(data))
	buf:=append([]byte(header),data...)
	h:= sha1.New()
	h.Write(buf)
	return hex.EncodeToString(h.Sum(nil)),buf
}