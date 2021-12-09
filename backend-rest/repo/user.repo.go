package repo

import (
	"github.com/kabi175/alumini-app-backend/domain/models"
	"gorm.io/gorm"
)

type defaultUserRepo struct {
	db *gorm.DB
}

func NewDefaultUserRepo(db *gorm.DB) models.UserRepo {
	db.AutoMigrate(&models.User{})
	return &defaultUserRepo{
		db: db,
	}
}

func (r *defaultUserRepo) Create(user *models.User) (*models.User, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *defaultUserRepo) DeleteByID(userID uint) error {
	result := r.db.Where("ID = ?", userID).Delete(&models.User{})
	return result.Error
}

func (r *defaultUserRepo) FindByID(userID uint) (*models.User, error) {
	user := &models.User{}
	result := r.db.Where(&models.User{ID: userID}).Find(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *defaultUserRepo) FindByMail(email string) (*models.User, error) {
	user := &models.User{}
	result := r.db.Where(&models.User{MailID: email}).Find(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *defaultUserRepo) FindByDesignation(designation string) ([]*models.User, error) {
	users := []*models.User{}
	result := r.db.Where(&models.User{Designation: designation}).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *defaultUserRepo) FindByBatch(batch uint) ([]*models.User, error) {
	users := []*models.User{}
	result := r.db.Where(&models.User{Batch: batch}).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *defaultUserRepo) UpdatePassword(userID uint, passwordHash string, randomHash string) error {
	result := r.db.Model(&models.User{ID: userID}).Updates(&models.User{PasswordHash: passwordHash})
	return result.Error
}

func (r *defaultUserRepo) UpdateVerify(userID uint, verification bool) error {
	result := r.db.Model(&models.User{ID: userID}).Updates(&models.User{Verified: verification})
	return result.Error
}

func (r *defaultUserRepo) UpdateActive(userID uint, isActive bool) error {
	result := r.db.Model(&models.User{ID: userID}).Updates(&models.User{Active: isActive})
	return result.Error
}

func (r *defaultUserRepo) UpdateName(userID uint, name string) error {
	result := r.db.Model(&models.User{ID: userID}).Updates(&models.User{Name: name})
	return result.Error
}
