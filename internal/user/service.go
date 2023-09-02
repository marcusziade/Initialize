package user

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// Model, which corresponds to the "users" database table.
type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

// Interface that provides user methods.
type Service interface {
	Create(u *User) error
	Get(id int) (*User, error)
	GetAll() ([]*User, error)
	Delete(id int) error
	Update(u *User) error
}

// Creates a user service with necessary dependencies.
func NewService(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}

type service struct {
	db *gorm.DB
}

// Creates a new user.
func (s *service) Create(u *User) error {
	return s.db.Create(u).Error
}

// Gets a user by ID.
func (s *service) Get(id int) (*User, error) {
	var u User
	if err := s.db.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

// Gets all users.
func (s *service) GetAll() ([]*User, error) {
	var us []*User
	if err := s.db.Find(&us).Error; err != nil {
		return nil, err
	}
	return us, nil
}

// Deletes a user.
func (s *service) Delete(id int) error {
	var u User
	if err := s.db.First(&u, id).Error; err != nil {
		return err
	}
	if err := s.db.Delete(&u).Error; err != nil {
		return err
	}
	return nil
}

// Updates a user.
func (s *service) Update(u *User) error {
	return s.db.Save(u).Error
}