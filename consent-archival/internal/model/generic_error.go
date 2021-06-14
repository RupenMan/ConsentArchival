package model

type genericError struct {
	StatusCode int
	Body errorBody 
}