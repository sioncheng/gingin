package repository

import (
	"github.com/sioncheng/gingin/cmd/domain"
)

type SnippetRepository interface {
	LoadSnippet(id int) domain.Snippet
}
