package main

import "github.com/vedantwankhade/snippetbox/internal/models"

type templateData struct {
    Snippet *models.Snippet
    Snippets []*models.Snippet
}
