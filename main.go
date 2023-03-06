package main

import (
	"log"
)

func main() {
	var images []GeneratedImage
	err := loadData(&images)
	if err != nil {
		log.Println("Failed to load data from file, generating new data...")
		images = generateData(10)
		err := saveData(images)
		if err != nil {
			log.Fatal(err)
		}
	}
}
