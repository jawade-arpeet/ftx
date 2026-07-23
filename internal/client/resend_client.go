package client

import "github.com/resend/resend-go/v3"

type ResendClient struct {
	rsnd *resend.Client
}

func newResendClient(apiKey string) *ResendClient {
	return &ResendClient{
		rsnd: resend.NewClient(apiKey),
	}
}
