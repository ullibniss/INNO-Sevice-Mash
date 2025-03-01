package group

import "inno-lb-master/internal/endpoint"

type Group struct {
	Alias     string               `json:"alias"`
	Endpoints []*endpoint.Endpoint `json:"endpoints"`
}
