package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//criacao de um arquivo
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//size, err := f.WriteString("Agora contém dados no arquivo txt")
	size, err := f.Write([]byte("Escrevendo dados no arquivo"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", size)
	f.Close()

	//leitura de um arquivo
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Arquivo aberto com sucesso!\n")
	fmt.Println(string(arquivo))

	newFile, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(newFile)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Print(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
