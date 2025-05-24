package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

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
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func decodeBase64(encodedStr string) ([]byte, error) {
	// 1. 清理非 Base64 字符
	cleanStr := strings.Map(func(r rune) rune {
		if strings.ContainsRune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=-_", r) {
			return r
		}
		return -1
	}, encodedStr)

	// 2. 尝试用 StdEncoding 解码（带 padding）
	if data, err := base64.StdEncoding.DecodeString(cleanStr); err == nil {
		return data, nil
	}

	// 3. 尝试用 RawStdEncoding 解码（不带 padding）
	if data, err := base64.RawStdEncoding.DecodeString(cleanStr); err == nil {
		return data, nil
	}

	// 4. 尝试用 URLEncoding 解码（支持 -_ 和可选的 padding）
	if data, err := base64.URLEncoding.DecodeString(cleanStr); err == nil {
		return data, nil
	}

	// 5. 尝试用 RawURLEncoding 解码（支持 -_ 且不带 padding）
	if data, err := base64.RawURLEncoding.DecodeString(cleanStr); err == nil {
		return data, nil
	}

	return nil, fmt.Errorf("无法解码 Base64 数据")
}
