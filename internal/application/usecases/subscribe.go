package usecases

import (
	"github.com/google/uuid"
	"github.com/samuelorlato/football-api/internal/application/ports"
	"github.com/samuelorlato/football-api/internal/domain/entities"
	ports2 "github.com/samuelorlato/football-api/internal/domain/ports/repositories"
	"github.com/samuelorlato/football-api/pkg/errs"
)

type subscribeUsecase struct {
	fanRepository ports2.FanRepository
}

func NewSubscribeUsecase(fanRepository ports2.FanRepository) ports.SubscribeUsecase {
	return &subscribeUsecase{
		fanRepository,
	}
}

func (s *subscribeUsecase) Execute(request entities.RegisterFanRequest) (*entities.Fan, error) {
	existingFan, err := s.fanRepository.FindByEmail(request.Email)
	if err != nil {
		return nil, errs.NewInternalServerError()
	}
	if existingFan != nil {
		return nil, errs.NewUnprocessableContentError("torcedor com email já cadastrado")
	}

	fanID := uuid.NewString()
	fan := request.ToEntity(fanID)
	err = s.fanRepository.Save(fan)
	if err != nil {
		return nil, errs.NewInternalServerError()
	}

	return &fan, nil
}
