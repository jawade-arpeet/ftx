package repository

import "ftx/internal/client"

type Repository struct{}

func New(client *client.Client) *Repository {
	return &Repository{}
}
