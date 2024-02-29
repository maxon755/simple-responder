package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.Int("port", 8080, "The port to start an application")
	responseStatus := flag.Int("status", 200, "The status to respond with")
	responseBody := flag.String("body", "OK", "The body to respond with")
	contentFilePath := flag.String("body-file", "", "The path to file with response content")
	responseDelay := flag.Int("delay", 0, "The delay before response. Usefull for timeout emulation")

	flag.Parse()

	if *contentFilePath != "" {
		responseContentFromFile, err := readFileContent(*contentFilePath)
		if err != nil {
			fmt.Println("Could not read file: " + err.Error())
			return
		}

		*responseBody = responseContentFromFile
	}

	r := gin.Default()
	r.Any("/*proxyPath", func(c *gin.Context) {
		time.Sleep(time.Duration(*responseDelay) * time.Second)

		c.String(*responseStatus, *responseBody)
	})
	fmt.Println("Time to make a request!")
	r.Run(fmt.Sprintf("0.0.0.0:%d", *port))
}

func readFileContent(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
