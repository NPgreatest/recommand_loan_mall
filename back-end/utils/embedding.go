package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main.go/global"
	"net/http"
)

type EmbeddingRequest struct {
	Input string `json:"input"`
	Model string `json:"model"`
}
type EmbeddingResponse struct {
	Data []struct {
		Embedding []float32 `json:"embedding"`
	} `json:"data"`
}

func GetEmbedding(input string) ([]float32, error) {
	requestBody, err := json.Marshal(EmbeddingRequest{Input: input,
		Model: "text-embedding-ada-002"})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", global.GVA_CONFIG.Keys.OpenApiBase+"/embeddings", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+global.GVA_CONFIG.Keys.OpenApiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var embeddingResp EmbeddingResponse
	err = json.Unmarshal(body, &embeddingResp)
	if err != nil {
		return nil, err
	}
	if len(embeddingResp.Data) > 0 {
		return embeddingResp.Data[0].Embedding, nil
	}

	return nil, fmt.Errorf("no embedding data found")
}
