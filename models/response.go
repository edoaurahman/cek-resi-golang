package models

import "time"

type Response struct {
	Resi       string    `json:"resi"`
	Expedition string    `json:"expedition"`
	Details    []Details `json:"details"`
}

type Details struct {
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}
