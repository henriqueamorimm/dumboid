package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

func generateDumbo(length int) string {
	characters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	rand.Seed(time.Now().UnixNano())
	uuid := make([]byte, length)
	for i := 0; i < length; i++ {
		uuid[i] = characters[rand.Intn(len(characters))]
	}
	return string(uuid)
}

func main() {
	miniUUID := generateDumbo(250) //Tamanho do DumboID gerado
	fmt.Fprintf(os.Stdout, "Generated DumboID : %s\n", miniUUID)
}
