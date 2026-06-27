package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/macbook/lumina-backend-go/internal/domain/entities"
	"github.com/macbook/lumina-backend-go/pkg/database"
	"gorm.io/datatypes"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Warning: .env file not found")
	}

	database.ConnectDB()

	var questions []entities.Question
	if err := database.DB.Find(&questions).Error; err != nil {
		log.Fatal("Failed to fetch questions:", err)
	}

	for _, q := range questions {
		if q.Type != entities.MultipleChoice {
			continue
		}

		var currentOptions interface{}
		if err := json.Unmarshal(q.Options, &currentOptions); err != nil {
			fmt.Printf("Skipping question %s: invalid JSON\n", q.ID)
			continue
		}

		// Check if it's an array
		optsArray, ok := currentOptions.([]interface{})
		if !ok {
			fmt.Printf("Skipping question %s: options is not an array (already converted?)\n", q.ID)
			continue
		}

		if len(optsArray) == 0 {
			continue
		}

		// Convert to map
		newOptions := make(map[string]string)
		keys := []string{"A", "B", "C", "D"}
		
		for i, val := range optsArray {
			if i >= len(keys) {
				break
			}
			text := fmt.Sprintf("%v", val)
			newOptions[keys[i]] = text
			
			if text == q.CorrectAnswer {
				q.CorrectAnswer = keys[i]
			}
		}

		optionsJSON, _ := json.Marshal(newOptions)
		q.Options = datatypes.JSON(optionsJSON)

		if err := database.DB.Save(&q).Error; err != nil {
			fmt.Printf("Failed to update question %s: %v\n", q.ID, err)
		} else {
			fmt.Printf("Updated question %s to A,B,C,D format\n", q.ID)
		}
	}

	fmt.Println("Migration completed.")
}
