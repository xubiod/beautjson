package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func prg() (err error) {
	var f, n *os.File
	var data []byte
	var result interface{}

	if len(os.Args[1:]) != 1 {
		log.Fatalln("missing input filename")
		return
	}

	if f, err = os.Open(os.Args[1]); err != nil {
		return
	}
	if n, err = os.Create("result.json"); err != nil {
		return
	}
	if data, err = io.ReadAll(f); err != nil {
		return
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return
	}
	if data, err = json.MarshalIndent(result, "", "\t"); err != nil {
		return
	}
	if _, err = n.Write(data); err != nil {
		return
	}
	if err = f.Close(); err != nil {
		return
	}
	if err = n.Close(); err != nil {
		return
	}

	log.Println("complete")
	return
}

func main() {
	if err := prg(); err != nil {
		log.Fatalln(err)
	}
}
