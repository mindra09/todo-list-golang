// internal/service/ai_service.go
package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type AIService interface {
	PredictCategory(file string) (string, error)
}

type AIClient struct{}

func NewAIClient() *AIClient {
	return &AIClient{}
}

func (a *AIClient) PredictCategory(file string) (string, error) {
	body, _ := json.Marshal(map[string]string{
		"file": file,
	})

	req, _ := http.NewRequest(
		"POST",
		"https://ai.example.com/predict-category",
		bytes.NewBuffer(body),
	)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var res struct {
		Data struct {
			Category string `json:"category"`
		} `json:"data"`
	}

	json.NewDecoder(resp.Body).Decode(&res)
	return res.Data.Category, nil
}
