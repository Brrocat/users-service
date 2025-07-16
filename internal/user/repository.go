package user

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetUserByID(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetAllUsers() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) UpdateUserByID(id uint, user *User) (*User, error) {
	var existing User
	if err := r.db.First(&existing, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&existing).Updates(user).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

func (r *Repository) DeleteUserByID(id uint) error {
	return r.db.Delete(&User{}, id).Error
}
