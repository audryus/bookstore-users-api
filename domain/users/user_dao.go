package users

import (
	"fmt"

	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/logger"

	"gitlab.com/aubayaml/aubayaml-go/bookstore/users-api/datasources/postgres/user_db"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/errors"
)

const (
	queryGetUser = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?"
)

//Get User by id
func (user *User) Get() errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.InternalServerError("database error", err)
	}
	defer stmt.Close()

	if err := stmt.QueryRow(user.ID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error when trying to get user", err)
		return errors.ParseError(err)
	}

	return nil
}

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES (?,?,?,?,?,?)"
)

//Save user
func (user *User) Save() errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.InternalServerError("database error", err)
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		logger.Error("error when trying to INSERT user", err)
		return errors.InternalServerError("database error", err)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to prepare get last inserted ID", err)
		return errors.InternalServerError("database error", err)
	}
	user.ID = userID
	return nil
}

const (
	queryUpdateUser = "UPDATE users SET first_name = ?, last_name =?, email = ?, status = ? WHERE id = ?"
)

//Update user
func (user *User) Update() errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update statement", err)
		return errors.InternalServerError("database error", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.ID)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return errors.InternalServerError("database error", err)
	}
	return nil
}

const (
	queryDeleteUser = "DELETE FROM users WHERE id = ?"
)

//Delete the User
func (user *User) Delete() errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete statement", err)
		return errors.InternalServerError("database error", err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.ID); err != nil {
		logger.Error("error when trying to delete user", err)
		return errors.InternalServerError("database error", err)
	}

	return nil
}

const (
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?"
)

//FindByStatus users
func (user *User) FindByStatus(status string) ([]User, errors.RestErr) {
	stmt, err := user_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find by status statement", err)
		return nil, errors.InternalServerError("database error", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find by status", err)
		return nil, errors.InternalServerError("database error", err)
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var tmp User
		if err := rows.Scan(&tmp.ID, &tmp.FirstName, &tmp.LastName, &tmp.Email, &tmp.DateCreated, &tmp.Status); err != nil {
			logger.Error("error when trying to read data from data base", err)
			return nil, errors.InternalServerError("database error", err)
		}
		results = append(results, tmp)
	}
	if len(results) == 0 {
		return nil, errors.NotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}

const (
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email = ? AND password = ? AND status = ?"
)

//FindByEmailAndPassword users
func (user *User) FindByEmailAndPassword() errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare find by email and password ", err)
		return errors.InternalServerError("database error", err)
	}
	defer stmt.Close()

	fmt.Println(user.Email)
	fmt.Println(user.Password)
	fmt.Println(user.Status)

	if err := stmt.QueryRow(user.Email, user.Password, user.Status).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error when trying to get user by email and password", err)
		return errors.ParseError(err)
	}

	return nil
}
