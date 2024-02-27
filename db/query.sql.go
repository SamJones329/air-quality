// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"
)

const createAirStat = `-- name: CreateAirStat :one
INSERT INTO
    air_stats (co2_ppm, temp_tick, humidity_tick)
VALUES
    ($1, $2, $3)
RETURNING
    id, co2_ppm, temp_tick, humidity_tick, created_at
`

type CreateAirStatParams struct {
	Co2Ppm       int32
	TempTick     int32
	HumidityTick int32
}

func (q *Queries) CreateAirStat(ctx context.Context, arg CreateAirStatParams) (AirStat, error) {
	row := q.db.QueryRow(ctx, createAirStat, arg.Co2Ppm, arg.TempTick, arg.HumidityTick)
	var i AirStat
	err := row.Scan(
		&i.ID,
		&i.Co2Ppm,
		&i.TempTick,
		&i.HumidityTick,
		&i.CreatedAt,
	)
	return i, err
}

const getAirStat = `-- name: GetAirStat :one
SELECT
    id, co2_ppm, temp_tick, humidity_tick, created_at
FROM
    air_stats
WHERE
    id = $1
LIMIT
    1
`

func (q *Queries) GetAirStat(ctx context.Context, id int64) (AirStat, error) {
	row := q.db.QueryRow(ctx, getAirStat, id)
	var i AirStat
	err := row.Scan(
		&i.ID,
		&i.Co2Ppm,
		&i.TempTick,
		&i.HumidityTick,
		&i.CreatedAt,
	)
	return i, err
}

const listAirStats = `-- name: ListAirStats :many
SELECT
    id, co2_ppm, temp_tick, humidity_tick, created_at
FROM
    air_stats
ORDER BY
    created_at
`

func (q *Queries) ListAirStats(ctx context.Context) ([]AirStat, error) {
	rows, err := q.db.Query(ctx, listAirStats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AirStat
	for rows.Next() {
		var i AirStat
		if err := rows.Scan(
			&i.ID,
			&i.Co2Ppm,
			&i.TempTick,
			&i.HumidityTick,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
