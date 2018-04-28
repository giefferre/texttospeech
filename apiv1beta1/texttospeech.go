// Package texttospeech offers a Client to interact with
// Google Cloud Text-to-Speech API.
package texttospeech

import (
	"context"

	gax "github.com/googleapis/gax-go"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1beta1"
	"google.golang.org/grpc"
)

// Client is a client for interacting with Google Cloud Text-to-Speech API.
type Client struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	client texttospeechpb.TextToSpeechClient
}

// NewClient creates a new texttospeech client.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &Client{
		conn:   conn,
		client: texttospeechpb.NewTextToSpeechClient(conn),
	}

	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *Client) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.conn.Close()
}

// ListVoices returns a list of Voice supported for synthesis.
func (c *Client) ListVoices(ctx context.Context, req *texttospeechpb.ListVoicesRequest, opts ...gax.CallOption) (*texttospeechpb.ListVoicesResponse, error) {
	var resp *texttospeechpb.ListVoicesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ListVoices(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SynthesizeSpeech synthesizes speech synchronously: receive results after
// all text input has been processed.
func (c *Client) SynthesizeSpeech(ctx context.Context, req *texttospeechpb.SynthesizeSpeechRequest, opts ...gax.CallOption) (*texttospeechpb.SynthesizeSpeechResponse, error) {
	var resp *texttospeechpb.SynthesizeSpeechResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.SynthesizeSpeech(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// defaultClientOptions reports the default set of client options to use with this package.
func defaultClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("texttospeech.googleapis.com:443"),
		option.WithScopes(defaultAuthScopes()...),
	}
}

// defaultAuthScopes reports the default set of authentication scopes to use with this package.
func defaultAuthScopes() []string {
	return []string{
		"https://www.googleapis.com/auth/cloud-platform",
	}
}
