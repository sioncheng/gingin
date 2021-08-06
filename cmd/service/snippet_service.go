package service

import (
	"github.com/sioncheng/gingin/cmd/domain"
	"github.com/sioncheng/gingin/cmd/repository"
)

type SnippetService struct {
	Repository repository.SnippetRepository
}

func (s *SnippetService) LoadSnippet(id int) domain.Snippet {
	return s.Repository.LoadSnippet(id)
}
