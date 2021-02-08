package domain

type Movie struct {
	Id       int
	Title    string
	Rate     float64
	Abstruct string
	FLink    string
	Reviews  []string
}
