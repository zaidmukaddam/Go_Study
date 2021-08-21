package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/reujab/wallpaper"
)

type Image struct {
	Urls struct {
		Full string `json:"full"`
	} `json:"urls"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// fmt.Println("Current wallpaper:", background)
	secretKey := os.Getenv("UNSPLASH_ACCESS_KEY")

	// fmt.Printf("Results: %v\n", secretKey)
	resp, err := http.Get("https://api.unsplash.com/photos/random?client_id=" + secretKey + "&featured=true&orientation=landscape&query=landscape")
	if err != nil {
		fmt.Print(err)
		panic(err.Error())
	}
	b := resp.Body
	defer b.Close()

	image := Image{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
		panic(err.Error())
	}
	err = json.Unmarshal(body, &image)
	if err != nil {
		fmt.Print(err)
		panic(err.Error())

	}
	// fmt.Printf("Results: %v\n", image.Urls.Full)

	wallpaper.SetFromURL(image.Urls.Full)
	os.Exit(0)
}
