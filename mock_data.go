package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/brianvoe/gofakeit/v6"
)

type GeneratedImage struct {
	ID         string  `json:"id"`
	Title      string  `json:"title"`
	Resolution string  `json:"resolution"`
	Model      string  `json:"model"`
	Prompt     string  `json:"prompt"`
	URL        string  `json:"url"`
	Author     *Author `json:"author"`
	Likes      int     `json:"likes"`
	Views      int     `json:"views"`
	CreatedAt  string  `json:"created_at"`
}

type Author struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	ImagesNum int    `json:"images_num"`
	CreatedAt string `json:"created_at"`
}

func generateData(num int) []GeneratedImage {
	gofakeit.Seed(0)

	var images []GeneratedImage

	for i := 0; i < num; i++ {
		image := GeneratedImage{
			ID:         gofakeit.UUID(),
			Title:      gofakeit.Word(),
			Resolution: fmt.Sprintf("%dx%d", gofakeit.Number(800, 1920), gofakeit.Number(600, 1080)),
			Model:      gofakeit.FirstName(),
			Prompt:     gofakeit.Sentence(5),
			URL:        gofakeit.ImageURL(800, 600),
			Author: &Author{
				ID:        gofakeit.UUID(),
				Username:  gofakeit.Username(),
				ImagesNum: gofakeit.Number(1, 100),
				CreatedAt: gofakeit.Date().Format("2006-01-02"),
			},
			Likes:     gofakeit.Number(0, 1000),
			Views:     gofakeit.Number(0, 10000),
			CreatedAt: gofakeit.Date().Format("2006-01-02 15:04:05"),
		}

		images = append(images, image)
	}

	return images
}

func saveData(images []GeneratedImage) error {
	err := os.MkdirAll("data", 0755)
	if err != nil {
		return err
	}

	filename := "data/data.json"

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(images)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func loadData(images *[]GeneratedImage) error {
	filename := "data/data.json"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &images)
	if err != nil {
		return err
	}

	return nil
}
