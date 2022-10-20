package repositories

import (
	"golang-api/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User, ID int) (models.User, error)
	DeleteUser(user models.User, ID int) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User

	err := r.db.First(&user, ID).Error
	return user, err
}

// Using Create method
func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

// update
func (r *repository) UpdateUser(user models.User, ID int) (models.User, error) {
	err := r.db.Model(&user).Where("id=?", ID).Updates(&user).Error // Using Save method

	return user, err
}

// delete
func (r *repository) DeleteUser(user models.User, ID int) (models.User, error) {
	err := r.db.Delete(&user).Error // Using Delete method

	return user, err
}
