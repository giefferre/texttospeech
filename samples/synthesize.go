// Command synthesize sends input text the Google Text-to-Speech API
// and converts it to an MP3 audio file.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	texttospeech "github.com/giefferre/texttospeech/apiv1beta1"
	"golang.org/x/net/context"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1beta1"
)

const usage = `Usage: synthesize <language> <text> <outputfilename>
Language must be a BCP-47 language tag.
Output will be an MP3 encoded audio file.
More information about BCP-47 here https://www.rfc-editor.org/rfc/bcp/bcp47.txt.
`

func main() {
	flag.Parse()
	if len(os.Args) < 4 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}

	// Checks for available voices with the given language code.
	languageCode := os.Args[1]
	resp, err := listVoices(languageCode)
	if err != nil {
		log.Fatal(err)
	}

	if len(resp.GetVoices()) == 0 {
		log.Fatalf("No voices available for language code '%s'. Please try again with a different one.", languageCode)
	}

	// Synthesizes the input text to audio bytes.
	text := os.Args[2]
	voice := resp.GetVoices()[0]
	out, err := synthesizeSpeech(text, voice)
	if err != nil {
		log.Fatal(err)
	}

	// Saves audio bytes to output file.
	fileName := os.Args[3]
	err = ioutil.WriteFile(fileName, out.GetAudioContent(), 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Done!")
}

func listVoices(languageCode string) (*texttospeechpb.ListVoicesResponse, error) {
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	resp, err := client.ListVoices(ctx, &texttospeechpb.ListVoicesRequest{
		LanguageCode: languageCode,
	})
	return resp, err
}

func synthesizeSpeech(text string, voice *texttospeechpb.Voice) (*texttospeechpb.SynthesizeSpeechResponse, error) {
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	resp, err := client.SynthesizeSpeech(ctx, &texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{
				Text: text,
			},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: voice.LanguageCodes[0],
			Name:         voice.Name,
			SsmlGender:   voice.SsmlGender,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	})
	return resp, err
}
