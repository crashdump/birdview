package model

import "time"

type EOL struct {
	Type    string    `json:"type"`
	Version string    `json:"version"`
	EOL     time.Time `json:"eol"`
}

type EOLs []EOL
