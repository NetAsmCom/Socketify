package main

import (
	"encoding/json"
	"os"
)

type response struct {
	ID  int
	Msg string
}

func main() {
	b, _ := json.Marshal(response{123, "Socketify!"})
	l := make([]byte, 4)
	getNativeEndian().PutUint32(l, uint32(len(b)))
	os.Stdout.Write(l)
	os.Stdout.WriteString(string(b) + "\n")
}
