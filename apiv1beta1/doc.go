/*
Package texttospeech offers a Client to interact with Google Cloud Text-to-Speech API.

Client initialization

	// Retrieve context
	ctx := context.Background()

	// Init client and check for errors
	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

ListVoices method

Returns a list of Voice supported for synthesis.

	resp, err := client.ListVoices(ctx, &texttospeechpb.ListVoicesRequest{
		LanguageCode: "en",
	})
	if err != nil {
		log.Fatal(err)
	}
	// do something with response

SynthesizeSpeech

Synthesizes speech synchronously: receive results after all text input has been processed.

	resp, err := client.SynthesizeSpeech(ctx, &texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{
				Text: "Hello world",
			},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-US",
			Name:         "en-US-Wavenet-A",
			SsmlGender:   "MALE",
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	// do something with response
*/
package texttospeech
