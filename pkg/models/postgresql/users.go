package postgresql

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
	"tadeobennett.net/quotation/pkg/models"
	// "errors"
	// "tadeobennett.net/quotation/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12) //iterates the hash 12  times
	if err != nil {
		return err
	}

	s := `INSERT INTO users(name, email, password_hash, activated)
	VALUES ($1, $2, $3, true)`

	// INSERT INTO users(name, email, password_hash, activated) VALUES ('Tadeo', 'tadeo@gmail.com', 'tadeo2002', true);

	_, err = m.DB.Exec(s, name, email, hashedPassword) //just returns the row and error

	if err != nil {
		switch {
		case err.Error() == `pq: duplicated key value violates unique constraint "users_email_key"`:
			return models.ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// func (m *UserModel) Get(id int) (*models.User, error){
// 	return 0, nil
// }
