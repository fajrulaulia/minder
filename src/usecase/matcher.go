package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/fajrulaulia/minder/src/model"
	"github.com/fajrulaulia/minder/src/repository"
	"github.com/fajrulaulia/minder/src/usecase/matcher"
)

type MatcherUsecaseIface interface {
	LikePassAction(ctx context.Context, params *matcher.LikePass) error
	ListPartnerCoupleRecommedation(ctx context.Context, mail string) ([]string, error)
}

type MatcherUsecaseStruct struct {
	MatcherRepo repository.MatcherRepositoryIface
	UserRepo    repository.UserRepositoryIface
}

func NewMatcherUsecase(user repository.MatcherRepositoryIface, UserRepo repository.UserRepositoryIface) MatcherUsecaseIface {
	return &MatcherUsecaseStruct{
		MatcherRepo: user,
		UserRepo:    UserRepo,
	}
}

func (c *MatcherUsecaseStruct) LikePassAction(ctx context.Context, params *matcher.LikePass) error {

	if params.UserEmail == params.EmailTarget {
		return fmt.Errorf("can't pass or like self")
	}

	user1, err := c.UserRepo.GetUserByEmail(ctx, params.UserEmail)
	if err != nil {
		return err
	}
	if user1 == nil {
		return fmt.Errorf("user1 not found")
	}

	user2, err := c.UserRepo.GetUserByEmail(ctx, params.EmailTarget)
	if err != nil {
		return err
	}

	if user2 == nil {
		return fmt.Errorf("user2 not found")
	}

	count, err := c.UserRepo.IsGotLimit(ctx, user1.ID)
	if err != nil {
		return err
	}

	if count >= 2 {
		if user1.SubscribedEndate != nil && user1.SubscribedEndate.After(time.Now()) {
			return fmt.Errorf("you reached limit")
		}
	}

	err = c.MatcherRepo.CreateAction(ctx, model.Matcher{
		User1ID: user1.ID,
		User2ID: user2.ID,
		Action:  params.Action,
	})
	if err != nil {
		return err
	}

	return nil

}

func (c *MatcherUsecaseStruct) ListPartnerCoupleRecommedation(ctx context.Context, mail string) ([]string, error) {

	user, err := c.UserRepo.GetUserByEmail(ctx, mail)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	res, err := c.UserRepo.SelectFreshUser(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
