# Google Cloud Text-to-Speech API Client Library for Go

[![GoDoc](https://godoc.org/github.com/giefferre/texttospeech?status.svg)][godoc-reference]

Package texttospeech offers a Client to interact with [Google Cloud Text-to-Speech API][tts-docs].

NOTE: this package actually offers the _v1beta1_ version of the API, so it could stop working in case of breaking changes made by Google engineers.

[godoc-reference]: https://godoc.org/github.com/giefferre/texttospeech
[tts-docs]: https://cloud.google.com/text-to-speech/docs/

## Documentation

Documentation is available on [GoDoc][godoc-reference].

## Usage example

The [samples](samples) folder contains a minimal application to help you understand how to use the package.

### Authentication

The package requires valid authentication credentials to access Google Cloud Text-to-Speech API.

- Create a project with the [Google Cloud Console][cloud-console], and enable the [Text-to-Speech API][tts-api].
- From the Cloud Console, create a service account, download its json credentials file, then set the `GOOGLE_APPLICATION_CREDENTIALS` environment variable:

  ```bash
  export GOOGLE_APPLICATION_CREDENTIALS=/path/to/your-project-credentials.json
  ```

For more information about Authentication, please consult [Google's Authentication Overview][adc].

[cloud-console]: https://console.cloud.google.com
[tts-api]: https://console.cloud.google.com/apis/api/texttospeech.googleapis.com/overview?project=_
[adc]: https://cloud.google.com/docs/authentication#developer_workflow

### Run the sample

Before running the example you must first install the Text-to-Speech API client:

```bash
go get -u github.com/giefferre/texttospeech
```

To run the example:

```bash
go run synthesize.go en "Hello world" output.mp3
```

## Reference

This library is inspired by:

- The [Google Cloud Client Libraries for Go][cloud-libraries]
- The [Google Cloud samples repository][cloud-samples]

[cloud-libraries]: https://github.com/GoogleCloudPlatform/google-cloud-go
[cloud-samples]: https://github.com/GoogleCloudPlatform/golang-samples
