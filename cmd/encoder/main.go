package main

import ( 
	"log"
	"os"
	"p1/pkg/encoder"
) 

func main() { 
	if len(os.Args) < 2 { 
		log.Fatal("Uso: encoder <arquivo.mem>") 
	}

	memFile := os.Args[1]
	encoder.RunBinary(memFile)
}