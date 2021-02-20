package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"hub-service/domain"
	"strconv"
	"strings"
)

func NewHubPostgresRepository(db *sql.DB) domain.HubRepository {
	return hubPostgres{
		db: db,
	}
}

type hubPostgres struct {
	db *sql.DB
}

func (u hubPostgres) Create(ctx context.Context, hub domain.Hub) (*domain.Hub, error) {
	var query = `
		INSERT INTO 
			hubs ("name", "geo_location")
		VALUES 
			($1, point($2, $3))
		RETURNING id
	`

	var lastId int64
	err := u.db.QueryRow(query, hub.Name, hub.GeoLocation.Lat, hub.GeoLocation.Long).Scan(&lastId)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return nil, domain.ErrorDuplicated
		}
		return nil, err
	}

	return u.FindById(ctx, lastId)
}

func (u hubPostgres) FindById(ctx context.Context, hubId int64) (*domain.Hub, error) {
	var query = `
		SELECT "id", "name", "geo_location" 
		FROM 
			hubs
		WHERE 
			id = $1
		LIMIT 1 FOR NO KEY UPDATE
	`

	var (
		id          int64
		name        string
		geoLocation string
	)

	err := u.db.QueryRowContext(ctx, query, hubId).Scan(&id, &name, &geoLocation)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrorNotFound
		}
		return nil, errors.New(fmt.Sprintf("error query hub info, details: %v", err.Error()))
	}

	return &domain.Hub{
		Id:          id,
		Name:        name,
		GeoLocation: parsePointToGeoLocation(geoLocation),
	}, nil
}

func (u hubPostgres) Find(ctx context.Context, req domain.FindHubsRequest) ([]*domain.Hub, error) {
	var query = `
		SELECT "id", "name", "geo_location" 
		FROM 
			hubs
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

	if req.Name != "" {
		count++
		params = append(params, req.Name)
		conditions = conditions + ` AND "name"::TEXT LIKE '%' || ` + fmt.Sprintf("$%d", count) + ` || '%'`
	}

	hubs := make([]*domain.Hub, 0)
	rows, err := u.db.QueryContext(ctx, fmt.Sprintf(query, conditions), params...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrorNotFound
		}
		return nil, errors.New(fmt.Sprintf("error query hub info, details: %v", err.Error()))
	}

	defer rows.Close()
	for rows.Next() {
		var (
			id     int64
			name   string
			geoStr string
		)
		rows.Scan(&id, &name, &geoStr)
		hubs = append(hubs, &domain.Hub{
			Id:          id,
			Name:        name,
			GeoLocation: parsePointToGeoLocation(geoStr),
		})
	}

	return hubs, nil
}

func parsePointToGeoLocation(geoStr string) domain.GeoLocation {
	lat := float64(0)
	long := float64(0)
	geoStr = strings.TrimLeft(strings.TrimRight(geoStr, ")"), "(")
	geoS := strings.Split(geoStr, ",")

	if tmp := geoS[0]; tmp != "" {
		lat, _ = strconv.ParseFloat(tmp, 64)
	}

	if tmp := geoS[1]; tmp != "" {
		long, _ = strconv.ParseFloat(tmp, 64)
	}

	return domain.GeoLocation{
		Lat:  lat,
		Long: long,
	}
}
