package main

import (
	"net/http"
	"fmt"
	"path/filepath"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context){
     
}

func main(){
	router := gin.Default()
	
	router.POST("/", func(c *gin.Context) {
		// single file
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(file.Filename)
		
		dst := filepath.Join("./files/", filepath.Base(file.Filename))
		c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!\n", file.Filename))
	})
	router.Run(":8080")
}
