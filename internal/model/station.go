package model

type Station struct {
	ID          int     `json:"id"`
	TrStartCode int     `json:"tr_start"`
	TrEndCode   int     `json:"tr_end"`
	Distance    float64 `json:"dist_tr"`
}
