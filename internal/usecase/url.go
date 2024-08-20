package usecase

import (
	"github.com/jhamiltonjunior/cut-url/internal/domain/entities/url"
	"github.com/jhamiltonjunior/cut-url/internal/domain/repository/url_repository"
)

type URLUseCase struct {
	URLRepository url_repository.Repository
}

func NewURLService(URLRepository url_repository.Repository) *URLUseCase {
	return &URLUseCase{
		URLRepository: URLRepository,
	}
}

func (s *URLUseCase) Create() error {
	entitiesURL := &url.URL{
		OriginalURL:    "Origin",
		DestinationURL: "string",
		UserID:         1,
	}

	_, err := s.URLRepository.CreateULR(entitiesURL)
	if err != nil {
		return err
	}

	return nil
}
