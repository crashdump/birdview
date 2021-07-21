package model

type Resource struct {
	Name      string `json:"name"`
	Account   string `json:"account"`
	Region    string `json:"region"`
	Service   string `json:"service"`
	Version   string `json:"version"`
}

type Resources []Resource



