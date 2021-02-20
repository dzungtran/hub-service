package usecase

import (
	"context"
	"errors"
	"fmt"
	"hub-service/domain"
	"hub-service/pkg/core/utils"
)

type (
	TeamAddUsersUseCase interface {
		Execute(context.Context, TeamAddUsersInput) error
	}

	// Input data
	TeamAddUsersInput struct {
		UserIds []int64 `json:"user_ids" validate:"required"`
		TeamId  int64
	}

	teamAddUsersInteractor struct {
		teamRepo domain.TeamRepository
		userRepo domain.UserRepository
	}
)

func NewTeamAddUsersUseCase(teamRepo domain.TeamRepository, userRepo domain.UserRepository) TeamAddUsersUseCase {
	return teamAddUsersInteractor{
		teamRepo: teamRepo,
		userRepo: userRepo,
	}
}

// Execute create User with dependencies
func (i teamAddUsersInteractor) Execute(ctx context.Context, input TeamAddUsersInput) error {

	err := i.validateUserIds(ctx, input.UserIds)
	if err != nil {
		return err
	}

	team, err := i.teamRepo.FindById(ctx, input.TeamId)
	if err != nil {
		if err == domain.ErrorNotFound {
			return errors.New("Team not found")
		}
		return err
	}

	err = i.teamRepo.AddUsers(ctx, team.Id, input.UserIds)
	if err != nil {
		return err
	}

	return nil
}

func (i teamAddUsersInteractor) validateUserIds(ctx context.Context, userIds []int64) error {
	users, err := i.userRepo.Find(ctx, domain.FindUsersRequest{Ids: userIds})
	if err != nil {
		return err
	}

	if len(userIds) > 0 && len(users) == 0 {
		return errors.New("users not found")
	}

	validIds := make([]int64, 0)
	invalidIds := make([]int64, 0)
	for _, u := range users {
		validIds = append(validIds, u.Id)
	}

	for _, id := range userIds {
		if !utils.IsInt64SliceContains(validIds, id) {
			invalidIds = append(invalidIds, id)
		}
	}

	if len(invalidIds) > 0 {
		return errors.New(fmt.Sprintf("User Ids %v does not exists", invalidIds))
	}

	return nil
}
