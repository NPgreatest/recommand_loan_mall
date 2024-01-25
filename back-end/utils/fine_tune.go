package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main.go/global"
	"net/http"
)

const prompt = `
Break down the user's requirements into a line by line of items and features, followed by a score for the importance of that item, like this:
Examples:
Question:We're having friends over tomorrow. Recommend some fruit and snacks for me. My friends loves fruit more, so please recommand more fruit.
Answer:Green Apple:6\nBrazilian banana:6\nCandy:3

Question:I want to make a healthy smoothie. What ingredients should I use?
Answer:Spinach:6\nFrozen berries:7\nGreek yogurt:5\nChia seeds:4\nAlmond milk:5\nBanana:6

Question:I'm planning a weekend camping trip. What food items should I take?
Answer:Canned beans:5\nInstant noodles:6\nBread:4\nPeanut butter:5\nBottled water:7\nTrail mix:3

Now, here is the question, give me answer directly:
Question:
`

// OpenAIResponse represents the structure of the response from OpenAI
type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// CallOpenAI calls the OpenAI API with the given messages
func CallOpenAI(messages string) (string, error) {
	tarMessages := []map[string]string{
		{
			"role":    "system",
			"content": "Refine user needs into itemized content for word vector matching. Predict the importance of this item in the overall demand. mark with number after the :\n",
		},
		{
			"role":    "user",
			"content": prompt + messages,
		},
	}
	payload := map[string]interface{}{
		"model":    global.GVA_CONFIG.Keys.FineTuneModel, // Specify the model
		"messages": tarMessages,                          // The messages to send
	}

	// Marshal the payload into JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Create a new request
	req, err := http.NewRequest("POST", global.GVA_CONFIG.Keys.OpenApiBase+"/chat/completions", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	// Set the necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", global.GVA_CONFIG.Keys.OpenApiKey))

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Unmarshal the response
	var openAIResp OpenAIResponse
	err = json.Unmarshal(body, &openAIResp)
	if err != nil {
		return "", err
	}

	// Return the response text
	if len(openAIResp.Choices) > 0 {
		return openAIResp.Choices[0].Message.Content, nil
	}
	return "", fmt.Errorf("no response from OpenAI")
}
