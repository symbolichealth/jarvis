package jarvis

// Jarvis is the main struct that keeps track of chat history with Gemini.
type Jarvis struct {
	gemini  *Gemini
	history []string
}

// Start initializes a new Jarvis instance.
func Start() *Jarvis {
	return &Jarvis{
		gemini:  NewGemini(),
		history: []string{},
	}
}

// Chat sends a message to Gemini and appends the response to the history.
func (j *Jarvis) Chat(message string) (string, error) {
	response, err := j.gemini.Chat(j.history, message)
	if err != nil {
		return "", err
	}
	j.history = append(j.history, message, response)
	return response, nil
}
