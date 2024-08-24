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

func (uc *URLUseCase) Create(u url.URL) error {
	entitiesURL := &url.URL{
		OriginalURL:    u.OriginalURL,
		DestinationURL: u.DestinationURL,
		Description:    u.Description,
		UserID:         1,
	}

	_, err := uc.URLRepository.CreateULR(entitiesURL)
	if err != nil {
		return err
	}

	return nil
}

func (uc *URLUseCase) GetAllByUser(id int) ([]*url.URL, error) {
	urls, err := uc.URLRepository.GetAllByUser(id)
	if err != nil {
		return nil, err
	}

	return urls, err
}

func (uc *URLUseCase) GetByName(description string) ([]url.URL, error) {
	u, err := uc.URLRepository.GetByName(description)
	if err != nil {
		return nil, err
	}

	return u, nil
}
