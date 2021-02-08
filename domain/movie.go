package domain

type Movie struct {
	Id       int
	Title    string
	Rate     float32
	Abstruct string
	FLink    string
	Reviews  []string
}
