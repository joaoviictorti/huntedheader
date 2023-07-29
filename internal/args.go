package internal

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func Banner() string {
	bannerheader := `
    __                __           ____                   __         
   / /_  __  ______  / /____  ____/ / /_  ___  ____ _____/ /__  _____
  / __ \/ / / / __ \/ __/ _ \/ __  / __ \/ _ \/ __ / __  / _ \/ ___/
 / / / / /_/ / / / / /_/  __/ /_/ / / / /  __/ /_/ / /_/ /  __/ /    
/_/ /_/\__,_/_/ /_/\__/\___/\__,_/_/ /_/\___/\__,_/\__,_/\___/_/                                                           
	`
	return bannerheader
}

func Argumentos() {
	parser := argparse.NewParser("Hunteheader", Banner())
	url := parser.String("u", "url", &argparse.Options{Required: false, Help: "Insert url", Default: ""})
	file := parser.String("f", "file", &argparse.Options{Required: false, Help: "Insert file url", Default: ""})
	threads := parser.Int("t", "threads", &argparse.Options{Required: false, Help: "Insert threads", Default: 5})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
		return
	}
	Run(url, file, *threads)

}
