package internal

import (
	"sync"
)

func Run(url *string, file *string, threads int) {
	value := make(chan struct{}, threads)
	var wg sync.WaitGroup
	if *url != "" {
		Headermap(url, &wg, false)
	} else if *file != "" {
		FileRead(file, &wg, value)
	}

}
