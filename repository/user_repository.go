package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
)

type UserRepository interface {
	BaseRepository[entity.UserCredential]
	GetByUsername(username string) (entity.UserCredential, error)
	GetByUsernamePassword(username string, password string) (entity.UserCredential, error)
}

type userRepository struct {
	db *sqlx.DB
}

func (u *userRepository) List() ([]entity.UserCredential, error) {
	var users []entity.UserCredential
	sql := `SELECT id, user_name, is_active FROM user_credential`
	err := u.db.Select(&users, sql)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) Get(id string) (entity.UserCredential, error) {
	var user entity.UserCredential
	sql := `SELECT id, user_name, is_active FROM user_credential WHERE id = $1`
	err := u.db.Get(&user, sql, id)
	if err != nil {
		return entity.UserCredential{}, err
	}
	return user, nil
}

func (u *userRepository) Create(newData entity.UserCredential) error {
	sql := "INSERT INTO user_credential (id, user_name, password, is_active) VALUES (:id, :user_name, :password, :is_active)"
	_, err := u.db.NamedExec(sql, &newData)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) Update(newData entity.UserCredential) error {
	sql := "UPDATE user_credential SET user_name = :user_name, password = :password WHERE id = :id"
	_, err := u.db.NamedExec(sql, &newData)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) Delete(id string) error {
	sql := "DELETE FROM user_credential WHERE id = $1"
	_, err := u.db.Exec(sql, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) GetByUsername(username string) (entity.UserCredential, error) {
	var userCredential entity.UserCredential
	sql := `SELECT id, user_name, is_active FROM user_credential WHERE user_name = $1 and is_active = true`
	err := u.db.Get(&userCredential, sql, username)
	if err != nil {
		return entity.UserCredential{}, err
	}
	return userCredential, nil
}

func (u *userRepository) GetByUsernamePassword(username string, password string) (entity.UserCredential, error) {
	user, err := u.GetByUsername(username)
	if err != nil {
		return entity.UserCredential{}, err
	}
	//pwdCheck := security.CheckPasswordHash(password, user.Password)
	//if !pwdCheck {
	//	return entity.UserCredential{}, fmt.Errorf("password don't match")
	//}
	return user, nil
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}
