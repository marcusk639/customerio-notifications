package repositories

import (
	"github.com/customerio/fullstack-backend-homework/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	Repository
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{Repository: Repository{db: db}}
}

// GetUserByEmail retrieves a user by email
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT id, email FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
