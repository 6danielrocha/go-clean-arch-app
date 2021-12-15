package usecases

import (
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/interfaces"
)

// A UserRepository belong to the inteface layer
type UserRepository struct {
	SQLHandler interfaces.SQLHandler
}

// FindAll is returns the number of entities.
func (ur *UserRepository) FindAll() (users domain.Users, err error) {
	const query = `
		SELECT
			id,
			name
		FROM
			users
	`
	rows, err := ur.SQLHandler.Query(query)

	if err != nil {
		return
	}

	//before row use
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err = rows.Scan(&id, &name); err != nil {
			return
		}
		user := domain.User{
			ID:   id,
			Name: name,
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}

// FindByID is returns the entity identified by the given id.
func (ur *UserRepository) FindByID(userID int) (user domain.User, err error) {

	const query = `
		SELECT
			id,
			name
		FROM
			users
		WHERE
			id = $1
	`
	row, err := ur.SQLHandler.Query(query, userID)

	if err != nil {
		return
	}

	//before row use
	defer row.Close()

	var id int
	var name string

	row.Next()
	if err = row.Scan(&id, &name); err != nil {
		return
	}

	user.ID = id
	user.Name = name

	return
}
