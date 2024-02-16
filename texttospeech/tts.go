package texttospeech

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/oklog/ulid/v2"
)


func TextToSpeech(some_text string, apiKey string, voice string,http_addres string) (string)  {
	var Url *url.URL
	Url, err := url.Parse(http_addres)
	if err != nil {
		log.Fatalf("can't parse https url: %v", err)
	}

	params := url.Values{}
	params.Add("text", some_text)
	params.Add("lang", "ru-RU")
	params.Add("voice", voice)
	Url.RawQuery = params.Encode()

	req, err := http.NewRequest("POST", Url.String(), nil)
	if err != nil {
		log.Fatalf("can't create request: %v", err)
	}

	req.Header.Add("Authorization", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("can't get response: %v", err)
	}
	defer res.Body.Close()
	//hier mey be errors
	log.Println("\nResponse Status", res.Status)
	log.Println("\nResponse Headers", res.Header)

	bodyData, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("can't read byte data: %v", err)
	}
	
	tmpfileName := ulid.Make().String() + ".ogg"

	abs,err := filepath.Abs("./tmp/" + tmpfileName)
	if err != nil {
		log.Fatalf("can't get .ogg file pach: %v", err)
	}
	// https://habr.com/ru/articles/118898/
	tmpfile, err := os.Create(abs)
	if err != nil {
		log.Fatalf("can't create audio.ogg file: %v", err)
	}

	defer tmpfile.Close()
	tmpfile.Write(bodyData)

	log.Println(abs)
	return abs
}


