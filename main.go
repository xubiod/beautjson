package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func prg() (err error) {
	if len(os.Args[1:]) != 1 {
		log.Fatalln("missing input filename")
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		return
	}

	n, err := os.Create("result.json")
	if err != nil {
		return
	}

	var result interface{}

	data, err := io.ReadAll(f)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return
	}

	data, err = json.MarshalIndent(result, "", "\t")
	if err != nil {
		return
	}

	_, err = n.Write(data)
	if err != nil {
		return
	}

	err = f.Close()
	if err != nil {
		return
	}

	err = n.Close()
	if err != nil {
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
