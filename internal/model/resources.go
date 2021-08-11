package model

type Resource struct {
	Name    string `json:"name"`
	Account string `json:"account"`
	Region  string `json:"region"`
	Service string `json:"service"`
	Version string `json:"version"`
}

type Resources []Resource

func (e Resources) Len() int {
	return len(e)
}

func (e Resources) Less(i, j int) bool {
	return e[i].Service > e[j].Service
}

func (e Resources) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}