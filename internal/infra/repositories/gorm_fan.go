package repositories

import (
	"github.com/samuelorlato/football-api/internal/domain/entities"
	ports "github.com/samuelorlato/football-api/internal/domain/ports/repositories"
	"github.com/samuelorlato/football-api/internal/integration/persistance/models"
	"gorm.io/gorm"
)

type gormFanRepository struct {
	db *gorm.DB
}

func NewGormFanRepository(db *gorm.DB) ports.FanRepository {
	return &gormFanRepository{
		db,
	}
}

func (f *gormFanRepository) FindByEmail(email string) (*entities.Fan, error) {
	var fanModel models.Fan
	if err := f.db.Where("email = ?", email).First(&fanModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	fan := fanModel.ToEntity()

	return &fan, nil
}

func (f *gormFanRepository) FindByTeam(team string) ([]entities.Fan, error) {
	var fanModels []models.Fan
	if err := f.db.Where("team = ?", team).Find(&fanModels).Error; err != nil {
		return nil, err
	}

	fans := make([]entities.Fan, len(fanModels))
	for i, fanModel := range fanModels {
		fans[i] = fanModel.ToEntity()
	}

	return fans, nil
}

func (f *gormFanRepository) Save(fan entities.Fan) error {
	var fanModel models.Fan
	fanModel.FromEntity(fan)

	return f.db.Create(&fanModel).Error
}
