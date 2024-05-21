package repository

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"github.com/fajrulaulia/minder/config"
	"github.com/fajrulaulia/minder/src/model"
)

type UserRepositoryIface interface {
	CreateUser(ctx context.Context, payload model.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	SelectFreshUser(ctx context.Context, id int) ([]string, error)
	IsGotLimit(ctx context.Context, id int) (int, error)
}

type UserRepositoryStruct struct {
	Config *config.Config
}

func NewUserRepository(c *config.Config) UserRepositoryIface {
	return &UserRepositoryStruct{
		Config: c,
	}
}

func (c *UserRepositoryStruct) CreateUser(ctx context.Context, payload model.User) error {

	db, err := c.Config.Db.MySQL().PrepareContext(ctx, "INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(payload.Username, payload.Email, payload.Password)
	if err != nil {
		return err
	}

	return nil
}

func (c *UserRepositoryStruct) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {

	var user model.User
	var subscribedEndDate sql.NullTime
	err := c.Config.Db.MySQL().QueryRowContext(ctx, "SELECT id, username, email, password, subscribed_enddate FROM users WHERE email=? ", email).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &subscribedEndDate)

	if err != nil {
		log.Println("GetUserByEmail error", err)
		if strings.Contains(err.Error(), "no rows") {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (c *UserRepositoryStruct) SelectFreshUser(ctx context.Context, id int) ([]string, error) {

	rows, err := c.Config.Db.MySQL().QueryContext(ctx, `
		SELECT 
		username
		FROM
		users  u
		WHERE
		u.id NOT IN (
			SELECT 
				DISTINCT user2_id
			FROM
				minder_db.matches
			WHERE
				DATE(created_at) = CURDATE() AND user1_id=?
		)
		AND id <> ?;
    `, id, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []string
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, user.Username)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (c *UserRepositoryStruct) IsGotLimit(ctx context.Context, id int) (int, error) {

	db := c.Config.Db.MySQL()
	var count int

	err := db.QueryRowContext(ctx, `
	SELECT  count(*)
	FROM minder_db.matches
	WHERE DATE(created_at) = CURDATE() AND user1_id =?;
    `, id).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
