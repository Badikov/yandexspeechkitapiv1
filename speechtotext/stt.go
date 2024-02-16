package speechtotext

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
)

// максимальный размер файла — 1 МБ;
// максимальная длительность — 30 секунд;
// максимальное количество аудиоканалов — 1.
func SpeechToText(file_pach string, apiKey string, stt_https_url string) (string) {
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

	body, err := io.ReadAll(res.Body) // response body is []byte
	if err != nil {
		log.Fatalf("I can't get response: %v", err)
	}

	log.Println(string(body))

	// str := string(body[:])

	return string(body)
}