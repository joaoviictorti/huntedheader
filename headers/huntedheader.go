package headers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

func Headermap(site *string) {
	uri, err := http.Get(*site)
	chaves := map[string]bool{
		"Strict-Transport-Security": false,
		"Content-Security-Policy":   false,
		"X-Frame-Options":           false,
		"X-Content-Type-Options":    false,
		"Referrer-Policy":           false,
		"Permissions-Policy":        false,
	}
	if err != nil {
		Ciano := color.New(color.FgCyan).SprintFunc()
		red := color.New(color.FgRed).SprintFunc()
		fmt.Printf("[%s] [%s] %s", Ciano(time.Now().Format("02/01/2006 15:04:05")), red("CONEXÃO FALHOU"), *site+"\n")
		os.Exit(0)
	}
	for valor := range uri.Header {
		switch valor {
		case "Strict-Transport-Security", "Content-Security-Policy", "X-Frame-Options", "X-Content-Type-Options", "Referrer-Policy", "Permissions-Policy":
			chaves[valor] = true
		default:
			continue
		}
	}
	for cabecalho, value := range chaves {
		var valor bool = value
		if !valor {
			Ciano := color.New(color.FgCyan).SprintFunc()
			red := color.New(color.FgRed).SprintFunc()
			fmt.Printf("[%s] [%s] %s", Ciano(time.Now().Format("02/01/2006 15:04:05")), red(cabecalho), *site+"\n")
		}
	}
}

func Arquivo(arquivo *string) {
	readFile, err := os.Open(*arquivo)
	if err != nil {
		fmt.Print("Erro")
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()
	for _, line := range lines {
		Headermap(&line)
	}
}
