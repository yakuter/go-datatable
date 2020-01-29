package main

import (
	"encoding/json"
	"fmt"
	"go-datatable/backend/inc"
	"go-datatable/backend/model"
	"go-datatable/backend/src"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

func main() {
	// Create DB connection
	db = inc.InitDB()
	defer db.Close()

	fillWithData()

	router := gin.Default()
	router.Use(src.CORS())

	router.GET("/posts", src.GetPosts)

	router.Run(":8081")
}

func fillWithData() {
	db = inc.GetDB()
	var posts []model.Post

	db.Find(&posts)

	if len(posts) > 0 {
		return
	}

	// JSON dosyasını açalım.
	jsonDosya, err := os.Open("posts.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonDosya.Close()

	// Dosya içeriğini okuyalım.
	icerik, err := ioutil.ReadAll(jsonDosya)
	if err != nil {
		fmt.Println(err)
	}

	// Dosyadaki veriyi ayrıştırıp değişkenimize aktaralım.
	json.Unmarshal(icerik, &posts)

	// JSON verisindeki isimleri ekrana yazdıralım.
	for _, post := range posts {
		db.FirstOrCreate(&post, post)
	}
}
