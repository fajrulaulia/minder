package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/fajrulaulia/minder/config"
	"github.com/fajrulaulia/minder/src/model"
)

type MatcherRepositoryIface interface {
	CreateAction(ctx context.Context, payload model.Matcher) error
}

type MatcherRepositoryStruct struct {
	Config *config.Config
}

func NewMatcherRepository(c *config.Config) MatcherRepositoryIface {
	return &MatcherRepositoryStruct{
		Config: c,
	}
}

func (c *MatcherRepositoryStruct) CreateAction(ctx context.Context, payload model.Matcher) error {

	db := c.Config.Db.MySQL()
	var count int

	err := db.QueryRowContext(ctx, `
        SELECT COUNT(*)
        FROM minder_db.matches
        WHERE (user1_id = ? AND user2_id = ?) 
        AND DATE(created_at) = CURDATE()
    `, payload.User1ID, payload.User2ID).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("you already like/pass this user")
	}
	stmt, err := db.PrepareContext(ctx, `
        INSERT INTO minder_db.matches
        (user1_id, user2_id, action, created_at)
        VALUES (?, ?, ?, CURRENT_TIMESTAMP)
    `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, payload.User1ID, payload.User2ID, payload.Action)
	if err != nil {
		log.Println("CreateAction err", err)
		return err
	}

	return nil
}
