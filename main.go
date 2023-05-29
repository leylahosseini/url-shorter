package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"time"

	//	_ "github.com/jinzhu/gorm/dialects/mysql"

	"leylahosseini/my-url-shorter/db"
	"leylahosseini/my-url-shorter/db/models"
)

func main() {
	db2 := db.ConnectToMysql()

	//defer db2.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			url := r.FormValue("url")
			code := generateCode()

			db2.Create(&models.Url{URL: url, Code: code})

			fmt.Fprintf(w, "http://localhost:8080/%s", code)
		} else {
			code := r.URL.Path[1:]

			var shortURL models.Url

			db2.Where("code = ?", code).First(&shortURL)

			if shortURL.ID != 0 {
				http.Redirect(w, r, shortURL.URL, http.StatusMovedPermanently)
			} else {
				fmt.Fprintf(w, "Invalid code")
			}
		}
	})

	http.ListenAndServe(":8080", nil)
}

// func InitDB(datasource string) *gorm.DB {
//     user := os.Getenv("DB_USER")
//     pass := os.Getenv("DB_PASSWORD")
//     url := os.Getenv("DB_URL")
//     dbname := os.Getenv("DB_NAME")

//     if datasource == "" {
//         datasource = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, url, dbname)
//     }

//     db, err := gorm.Open("mysql", datasource)

//     if err != nil {
//         fmt.Println("Error Connection Database")
//         panic(err)
//     }

//     fmt.Println("Connected to Database ...")

//     db.AutoMigrate(&ShortURL{})

//     return db
// }

func generateCode() string {
	// generate a unique code for the short URL
	h := sha256.New()
	h.Write([]byte(strconv.FormatInt(time.Now().UnixNano(), 10)))
	code := hex.EncodeToString(h.Sum(nil))[:8]
	return code
}
