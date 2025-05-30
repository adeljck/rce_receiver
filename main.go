package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

var (
	Port string
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		datas, err := c.GetRawData()
		if err != nil {
			return
		}
		if b64, exist := c.GetQuery("b64"); exist && b64 == "1" {
			data, err := decodeBase64(string(datas))
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf(string(data))
		} else {
			fmt.Println(string(datas))
		}
	})
	flag.StringVar(&Port, "p", "8080", "server port")
	flag.Parse()
	log.SetPrefix("[+] ")
	log.Printf("Listening on port %s\n", Port)
	err := r.Run(fmt.Sprintf(":%s", Port))
	if err != nil {
		panic(err)
	}
}

func decodeBase64(encodedStr string) ([]byte, error) {

	cleanStr := strings.Map(func(r rune) rune {
		if strings.ContainsRune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=-_", r) {
			return r
		}
		return -1
	}, encodedStr)

	if data, err := base64.StdEncoding.DecodeString(cleanStr); err == nil {
		return data, nil
	}

	if data, err := base64.RawStdEncoding.DecodeString(cleanStr); err == nil {
		return data, nil
	}

	if data, err := base64.URLEncoding.DecodeString(cleanStr); err == nil {
		return data, nil
	}

	if data, err := base64.RawURLEncoding.DecodeString(cleanStr); err == nil {
		return data, nil
	}

	return nil, fmt.Errorf("Can not decode base64")
}
