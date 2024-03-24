package repositories

import (
	"banky/user/entities"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type userPostgresRepository struct {
	db *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) UserRepository {
	return &userPostgresRepository{
		db: db,
	}
}

func (r *userPostgresRepository) InsertUser(user *entities.User) (*entities.User, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		log.Errorf("InsertUserData: %v", result.Error)
		return nil, result.Error
	}
	log.Debugf("InsertUserData: %v", result.RowsAffected)

	createdUser, err := r.FindUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (r *userPostgresRepository) FindUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	result := r.db.First(&user, "email = ?", email)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		log.Errorf("FindUserByEmail: %v", result.Error)
		return nil, result.Error
	}
	return &user, nil
}
