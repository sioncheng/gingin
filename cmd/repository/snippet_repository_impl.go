package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/sioncheng/gingin/cmd/domain"
)

type SnippetRepositoryImpl struct {
	DB *gorm.DB
}

func (s *SnippetRepositoryImpl) LoadSnippet(id int) domain.Snippet {
	entity := domain.Snippet{}

	s.DB.First(&entity, id)

	return entity
}
