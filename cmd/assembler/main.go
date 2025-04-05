package main

import (
	"fmt"
	"log"
	"os"

	"p1/pkg/assembler"
	"p1/pkg/assembler/lexer"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Uso: assembler <arquivo.asm>")
	}

	asmFile := os.Args[1]

	// Obtém os tokens do arquivo de assembly.
	tokens := lexer.GetTokens(asmFile)

	// Cria uma nova instância do assembler.
	asmb := assembler.NewAssembler(tokens)

	// Primeira passagem: coleta rótulos e define o PC.
	if err := asmb.FirstPass(); err != nil {
		log.Fatalf("Erro na primeira passagem: %v", err)
	}

	// Segunda passagem: gera o código final, resolvendo operandos.
	if err := asmb.SecondPass(); err != nil {
		log.Fatalf("Erro na segunda passagem: %v", err)
	}

	// Grava o arquivo .mem com o cabeçalho e o código gerado.
	if err := asmb.WriteMEM("io/build/output.mem"); err != nil {
		log.Fatalf("Erro ao escrever o arquivo .mem: %v", err)
	}

	fmt.Println("Arquivo .mem gerado com sucesso!")
}
