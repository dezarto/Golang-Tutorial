package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Channel struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Field1      string `json:"field1"`
	Field2      string `json:"field2"`
	Field3      string `json:"field3"`
	Field4      string `json:"field4"`
	Field5      string `json:"field5"`
	Field6      string `json:"field6"`
	Field7      string `json:"field7"`
	Field8      string `json:"field8"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	LastEntryID int    `json:"last_entry_id"`
}

type Feed struct {
	CreatedAt string `json:"created_at"`
	EntryID   int    `json:"entry_id"`
	Field1    string `json:"field1"`
	Field2    string `json:"field2"`
	Field3    string `json:"field3"`
	Field4    string `json:"field4"`
	Field5    string `json:"field5"`
	Field6    string `json:"field6"`
	Field7    string `json:"field7"`
	Field8    string `json:"field8"`
}

type APIResponse struct {
	Channel Channel `json:"channel"`
	Feeds   []Feed  `json:"feeds"`
}

func fetchAPIData(url string) (APIResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return APIResponse{}, fmt.Errorf("API'den veri çekerken hata oluştu: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return APIResponse{}, fmt.Errorf("API'den beklenmeyen durum kodu: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return APIResponse{}, fmt.Errorf("Veri okunurken hata oluştu: %v", err)
	}

	var data APIResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return APIResponse{}, fmt.Errorf("JSON çözülürken hata oluştu: %v", err)
	}

	return data, nil
}

func main() {
	app := fiber.New()

	app.Get("/api/data", func(c *fiber.Ctx) error {
		apiURL := "https://api.thingspeak.com/channels/2578919/feeds.json?api_key=18J81HFG0751ONGI&results=20" // API URL'si

		data, err := fetchAPIData(apiURL)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(data)
	})

	app.Listen(":3000")
}
