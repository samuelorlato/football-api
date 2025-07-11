package repositories

import (
	"github.com/samuelorlato/football-api/internal/domain/entities"
	ports "github.com/samuelorlato/football-api/internal/domain/ports/repositories"
	"github.com/samuelorlato/football-api/internal/integration/persistance/models"
	"gorm.io/gorm"
)

type gormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) ports.UserRepository {
	return &gormUserRepository{
		db,
	}
}

func (u *gormUserRepository) FindByID(ID string) (*entities.User, error) {
	var userModel models.User
	if err := u.db.Where("id = ?", ID).First(&userModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	user := userModel.ToEntity()

	return &user, nil
}

func (u *gormUserRepository) FindByUsername(username string) (*entities.User, error) {
	var userModel models.User
	if err := u.db.Where("name = ?", username).First(&userModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	user := userModel.ToEntity()

	return &user, nil
}

func (u *gormUserRepository) Save(user entities.User) error {
	var userModel models.User
	userModel.FromEntity(user)

	return u.db.Create(&userModel).Error
}
