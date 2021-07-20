package model

type Resource struct {
	Account   string `json:"account"`
	Region    string `json:"region"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	Version   string `json:"version"`
}

type Resources []Resource