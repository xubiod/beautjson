package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args[1:]) != 1 {
		log.Fatalln("missing input filename")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	n, err := os.Create("result.json")
	if err != nil {
		log.Fatalln(err)
	}

	var result interface{}

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatalln(err)
	}

	data, err = json.MarshalIndent(result, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}

	n.Write(data)

	err = f.Close()
	if err != nil {
		log.Fatalln(err)
	}

	err = n.Close()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("complete")
}
