package yandexspeechkitapiv1

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/Badikov/yandexspeechkitapiv1/texttospeech"
	"github.com/Badikov/yandexspeechkitapiv1/speechtotext"
)

type Config struct {
	APP_ENV   string `mapstructure:"APP_ENV"`
	API_KEY   string `mapstructure:"API_KEY"`
	TTS_HTTPS string `mapstructure:"TTS_HTTPS"`
	STT_HTTPS string `mapstructure:"STT_HTTPS"`
	VOICE     string `mapstructure:"VOICE"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	// Read file path
	viper.AddConfigPath(path)
	// set config file and path
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	// watching changes in app.env
	viper.AutomaticEnv()
	// reading the config file
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func YandexSpeechKit(pach_file, some_text string) string {
	// load app.env file data to struct
	config, err := LoadConfig(".")
	// handle errors
	if err != nil {
		log.Fatalf("can't load environment app.env: %v", err)
	}
	
	fmt.Printf(" -----%s----\n", "Reading Environment variables Using Viper package")
	fmt.Printf(" %s = %v \n", "Application_Environment", config.APP_ENV)
	fmt.Printf(" %s = %v \n", "Yandex API key", config.API_KEY)
	fmt.Printf(" %s = %v \n", "Text to speech server addres", config.TTS_HTTPS)
	fmt.Printf(" %s = %v \n", "Speech to text server addres", config.STT_HTTPS)
	fmt.Printf(" %s = %v \n", "Selected voice", config.VOICE)

	if pach_file == "" && some_text != "" {
		pachAudio := texttospeech.TextToSpeech(some_text, config.API_KEY, config.VOICE, config.TTS_HTTPS)
		log.Println(pachAudio)
		return pachAudio
	} else if pach_file != "" && some_text == "" {
		some_text, err = speechtotext.SpeechToText(pach_file, config.API_KEY,config.STT_HTTPS)
		if err != nil {
			log.Fatalf("I can't get text string: %v", err)
			return ""
		}
		log.Println(some_text)
		return some_text
	}
	return ""	
}
