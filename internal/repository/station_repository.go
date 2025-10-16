package repository

import (
	"database/sql"
	"tariff-api/internal/model"
)

type StationRepository struct {
	db *sql.DB
}

func NewStationRepository(db *sql.DB) *StationRepository {
	return &StationRepository{db: db}
}

func (repo *StationRepository) GetStations() ([]model.Station, error) {
	rows, err := repo.db.Query(`
		SELECT ID, CODE, NAME, PARAGRAPH
		FROM stations
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stations := make([]model.Station, 0)
	for rows.Next() {
		var station model.Station
		if err := rows.Scan(&station.ID, &station.CODE, &station.NAME, &station.PARAGRAPH); err != nil {
			return nil, err
		}
		stations = append(stations, station)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stations, nil
}
