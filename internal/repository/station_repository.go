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
		SELECT id, tr_start, tr_end, dist_tr
		FROM tarif_tr
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stations := make([]model.Station, 0)
	for rows.Next() {
		var station model.Station
		if err := rows.Scan(&station.ID, &station.TrStartCode, &station.TrEndCode, &station.Distance); err != nil {
			return nil, err
		}
		stations = append(stations, station)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stations, nil
}
