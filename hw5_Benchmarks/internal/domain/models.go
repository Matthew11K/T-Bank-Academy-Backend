package domain

type Data struct {
	ID       int               `json:"id" msgpack:"id" gob:"id"`
	Name     string            `json:"name" msgpack:"name" gob:"name"`
	Tags     map[string]string `json:"tags" msgpack:"tags" gob:"tags"`
	IsActive bool              `json:"is_active" msgpack:"is_active" gob:"is_active"`
	Score    float64           `json:"score" msgpack:"score" gob:"score"`
}
