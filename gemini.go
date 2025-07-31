package jarvis

import (
	"context"
	"errors"
	"os"

	genai "github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// Gemini provides methods to interact with the Google Gemini API.
type Gemini struct {
	client *genai.Client
}

// NewGemini creates a new Gemini client using the GEMINI_API_KEY environment variable.
func NewGemini() *Gemini {
	apiKey := os.Getenv("GEMINI_API_KEY")
	ctx := context.Background()
	var client *genai.Client
	if apiKey != "" {
		c, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
		if err == nil {
			client = c
		}
	}
	return &Gemini{client: client}
}

// Chat sends a message along with the previous history to the Gemini API.
func (g *Gemini) Chat(history []string, message string) (string, error) {
	if g.client == nil {
		return "", errors.New("gemini client not initialized")
	}

	ctx := context.Background()
	cs := g.client.GenerativeModel("gemini-2.5-flash").StartChat()

	// Convert existing history to chat messages.
	for i, h := range history {
		role := "user"
		if i%2 == 1 {
			role = "model"
		}
		cs.History = append(cs.History, &genai.Content{Role: role, Parts: []genai.Part{genai.Text(h)}})
	}

	resp, err := cs.SendMessage(ctx, genai.Text(message))
	if err != nil {
		return "", err
	}
	if len(resp.Candidates) == 0 {
		return "", errors.New("no response from gemini")
	}
	if len(resp.Candidates[0].Content.Parts) == 0 {
		return "", errors.New("empty response from gemini")
	}
	if txt, ok := resp.Candidates[0].Content.Parts[0].(genai.Text); ok {
		return string(txt), nil
	}
	return "", errors.New("unexpected response from gemini")
}
