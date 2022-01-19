package db

import (
	"fmt"

	"github.com/tupt0101/student-auth-service/src/datasource/postgres"

	"github.com/tupt0101/student-auth-service/src/domain/access_token"
	"github.com/tupt0101/student-auth-service/src/domain/users"
	"github.com/tupt0101/student-auth-service/src/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT user_id, access_token, expires FROM access_tokens WHERE access_token=$1;"
	queryCreateAccessToken = "UPDATE access_tokens SET access_token=$1, expires=$2 WHERE user_id=$3;"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=$1 WHERE access_token=$2;"
	queryLoginUser         = "SELECT user_id, email, name, created_on FROM users WHERE email=$1 and password=$2;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	stmt, err := postgres.Client.Prepare(queryGetAccessToken)
	if err != nil {
		fmt.Println(err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	var result access_token.AccessToken
	if getErr := row.Scan(&result.UserId, &result.AccessToken, &result.Expires); getErr != nil {
		return nil, errors.NewNotFoundError("no access token found with given id")
	}

	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {
	stmt, err := postgres.Client.Prepare(queryCreateAccessToken)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(at.AccessToken, at.Expires, at.UserId)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	stmt, err := postgres.Client.Prepare(queryCreateAccessToken)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(at.Expires, at.AccessToken)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	stmt, err := postgres.Client.Prepare(queryLoginUser)
	if err != nil {
		fmt.Println(err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	row := stmt.QueryRow(email, password)
	var result users.User
	if getErr := row.Scan(&result.UserId, &result.Email, &result.Name, &result.CreatedOn); getErr != nil {
		return nil, errors.NewNotFoundError("no user found with given email & password")
	}

	return &result, nil
}
