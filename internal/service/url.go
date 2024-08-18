package service

import (
	"fmt"
	"github.com/jhamiltonjunior/cut-url/internal/domain/entities/url"
	"github.com/jhamiltonjunior/cut-url/internal/domain/repository/url_repository"
)

type URLService struct {
	URLRepository url_repository.Repository
}

func NewURLService(URLRepository url_repository.Repository) *URLService {
	return &URLService{
		URLRepository: URLRepository,
	}
}

func (s *URLService) Create() error {
	entitiesURL := &url.URL{
		OriginalURL:    "Origin",
		DestinationURL: "string",
		UserID:         1,
	}

	fmt.Println("cheguei service")

	_, err := s.URLRepository.CreateULR(entitiesURL)
	if err != nil {
		return err
	}

	return nil
}
