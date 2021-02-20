package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"hub-service/domain"
	"strings"
)

func NewUserPostgresRepository(db *sql.DB) domain.UserRepository {
	return userPostgres{
		db: db,
	}
}

type userPostgres struct {
	db *sql.DB
}

func (u userPostgres) Create(ctx context.Context, user domain.User) (*domain.User, error) {
	var query = `
		INSERT INTO 
			users ("role", "email")
		VALUES 
			($1, $2)
		RETURNING id
	`

	var lastId int64
	err := u.db.QueryRow(query, user.Role, user.Email).Scan(&lastId)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return nil, domain.ErrorDuplicated
		}
		return nil, err
	}

	return u.FindById(ctx, lastId)
}

func (u userPostgres) FindById(ctx context.Context, userId int64) (*domain.User, error) {
	var query = `
		SELECT "id", "role", "email" 
		FROM 
			users
		WHERE 
			id = $1
		LIMIT 1 FOR NO KEY UPDATE
	`

	var (
		id    int64
		role  string
		email string
	)

	err := u.db.QueryRowContext(ctx, query, userId).Scan(&id, &role, &email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrorNotFound
		}
		return nil, errors.New(fmt.Sprintf("error query user info, details: %v", err.Error()))
	}

	return &domain.User{
		Id:    id,
		Role:  domain.UserRoleType(role),
		Email: email,
	}, nil
}

func (u userPostgres) Find(ctx context.Context, req domain.FindUsersRequest) ([]*domain.User, error) {
	var query = `
		SELECT u."id", u."role", u."email" 
		FROM 
			users AS u %s
		WHERE 
		%s
	`

	count := 0
	conditions := "1=1"
	joinStr := ""
	params := make([]interface{}, 0)

	if len(req.Ids) > 0 {
		inParams := make([]string, 0)
		for _, id := range req.Ids {
			count++
			inParams = append(inParams, fmt.Sprintf("$%d", count))
			params = append(params, id)
		}

		conditions = conditions + " AND u.id IN (" + strings.Join(inParams, ",") + ")"
	}

	if req.TeamId > 0 {
		count++
		params = append(params, req.TeamId)
		joinStr = " INNER JOIN teams_users AS tu ON tu.team_id = " + fmt.Sprintf("$%d", count) + " AND u.id = tu.user_id"
	}

	users := make([]*domain.User, 0)
	rows, err := u.db.QueryContext(ctx, fmt.Sprintf(query, joinStr, conditions), params...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrorNotFound
		}
		return nil, errors.New(fmt.Sprintf("error query user info, details: %v", err.Error()))
	}

	defer rows.Close()
	for rows.Next() {
		var (
			id    int64
			role  string
			email string
		)
		rows.Scan(&id, &role, &email)
		users = append(users, &domain.User{
			Id:    id,
			Role:  domain.UserRoleType(role),
			Email: email,
		})
	}

	return users, nil
}
