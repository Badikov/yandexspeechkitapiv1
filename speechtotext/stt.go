package speechtotext

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// максимальный размер файла — 1 МБ;
// максимальная длительность — 30 секунд;
// максимальное количество аудиоканалов — 1.
func SpeechToText(file_pach string, apiKey string, stt_https_url string) (string, error) {
	audioData, err := os.ReadFile(file_pach)
	if err != nil {
		log.Fatalf("I can't read the audio file: %v", err)
	}

	req, err := http.NewRequest("POST", stt_https_url, bytes.NewReader(audioData))
	if err != nil {
		log.Fatalf("I can't create request: %v", err)
	}
	req.Header.Add("Authorization", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("I can't send request: %v", err)
	}
	defer res.Body.Close()

	log.Println("\nResponse Status", res.Status)
	log.Println("\nResponse Headers", res.Header)

	var r struct {
		Result string `json:"result,omitempty"`
	}

	dec := json.NewDecoder(res.Body)

	if err := dec.Decode(&r); err != nil {
		return "", err
	}

	return r.Result, nil
}