package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"hub-service/domain"
	"strings"
)

func NewTeamPostgresRepository(db *sql.DB) domain.TeamRepository {
	return teamPostgres{
		db: db,
	}
}

type teamPostgres struct {
	db *sql.DB
}

func (u teamPostgres) Create(ctx context.Context, team domain.Team) (*domain.Team, error) {
	var query = `
		INSERT INTO 
			teams ("name", "type", "hub_id")
		VALUES 
			($1, $2, $3)
		RETURNING id
	`

	var lastId int64
	err := u.db.QueryRow(query, team.Name, team.Type, team.HubId).Scan(&lastId)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return nil, domain.ErrorDuplicated
		}
		return nil, err
	}

	return u.FindById(ctx, lastId)
}

func (u teamPostgres) FindById(ctx context.Context, teamId int64) (*domain.Team, error) {
	var query = `
		SELECT "id", "name", "type", "hub_id" 
		FROM 
			teams
		WHERE 
			id = $1
		LIMIT 1 FOR NO KEY UPDATE
	`

	var (
		id       int64
		name     string
		teamType string
		hubId    int64
	)

	err := u.db.QueryRowContext(ctx, query, teamId).Scan(&id, &name, &teamType, &hubId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrorNotFound
		}
		return nil, errors.New(fmt.Sprintf("error query team info, details: %v", err.Error()))
	}

	return &domain.Team{
		Id:    id,
		Name:  name,
		Type:  domain.TeamType(teamType),
		HubId: hubId,
	}, nil
}

func (u teamPostgres) AddUsers(ctx context.Context, teamId int64, userIds []int64) error {
	var query = `
		INSERT INTO 
			teams_users ("team_id", "user_id")
		VALUES 
			%s
	`

	vals := make([]string, 0)
	params := make([]interface{}, 0)
	count := 0

	for _, userId := range userIds {
		vals = append(vals, fmt.Sprintf("($%d, $%d)", count+1, count+2))
		params = append(params, teamId, userId)
		count = count + 2
	}

	if len(vals) == 0 {
		return nil
	}

	query = fmt.Sprintf(query, strings.Join(vals, ","))

	_, err := u.db.Exec(query, params...)
	if err != nil {
		return err
	}

	return nil
}

func (u teamPostgres) Find(ctx context.Context, req domain.FindTeamsRequest) ([]*domain.Team, error) {
	var query = `
		SELECT "id", "name", "type", "hub_id"
		FROM 
			teams
		WHERE 
		%s
	`
	count := 0
	conditions := "1=1"
	params := make([]interface{}, 0)

	if len(req.Ids) > 0 {
		inParams := make([]string, 0)
		for _, id := range req.Ids {
			count++
			inParams = append(inParams, fmt.Sprintf("$%d", count))
			params = append(params, id)
		}
		conditions = conditions + " AND id IN (" + strings.Join(inParams, ",") + ")"
	}

	if len(req.Types) > 0 {
		inParams := make([]string, 0)
		for _, id := range req.Types {
			count++
			inParams = append(inParams, fmt.Sprintf("$%d", count))
			params = append(params, id)
		}
		conditions = conditions + ` AND "type" IN (` + strings.Join(inParams, ",") + ")"
	}

	if req.HubId > 0 {
		count++
		params = append(params, req.HubId)
		conditions = conditions + " AND hub_id = " + fmt.Sprintf("$%d", count)
	}

	if req.Name != "" {
		count++
		params = append(params, req.Name)
		conditions = conditions + ` AND "name"::TEXT LIKE '%' || ` + fmt.Sprintf("$%d", count) + ` || '%'`
	}

	teams := make([]*domain.Team, 0)
	rows, err := u.db.QueryContext(ctx, fmt.Sprintf(query, conditions), params...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrorNotFound
		}
		return nil, errors.New(fmt.Sprintf("error query user info, details: %v", err.Error()))
	}

	defer rows.Close()
	for rows.Next() {
		var (
			id       int64
			name     string
			teamType string
			hubId    int64
		)
		rows.Scan(&id, &name, &teamType, &hubId)
		teams = append(teams, &domain.Team{
			Id:    id,
			Name:  name,
			Type:  domain.TeamType(teamType),
			HubId: hubId,
		})
	}

	return teams, nil
}
