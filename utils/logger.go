package utils

import (
	"fmt"
	"time"
)

func Log(msg string){
	fmt.Printf("[%s]: %s", time.Now().Format("2006-01-02 15:04:05"), msg)
}

func Version(o bool) string {
	version := "0.0.1"

	if o {
		fmt.Println(version)
	}

	return version
}