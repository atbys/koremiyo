package domain

type Movie struct {
	ID       int
	Title    string
	Rate     float64
	Abstruct string
	FLink    string
	Reviews  []string
}
