package main

import (
	"math/rand"
	"net/http"
	"fmt"
	"log"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"flag"
	"time"
	"os"
)

func randomString(length int) string {
    rand.Seed(time.Now().UnixNano())

    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    result := make([]byte, length)

    for i := range result {
        result[i] = charset[rand.Intn(len(charset))]
    }

    return string(result)
}

func main(){
	router := gin.Default()

	domain := flag.String("domain", "Domain", "Website Link")

	flag.Parse()
	
	router.POST("/", func(c *gin.Context) {
		// single file
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println(file.Filename)
		random_string := randomString(6)
		dst := filepath.Join("./files/", random_string, filepath.Base(file.Filename))
		c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("%s/%s/%s\n", *domain, random_string ,file.Filename))
	})

	router.LoadHTMLFiles("index.html")
	
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	files := http.FileServer(http.FS(os.DirFS("./files")))
	
	router.GET("/:random_string/:name", gin.WrapH(files))
	
	router.Run(":8080")
}
