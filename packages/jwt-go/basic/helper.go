package main

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}

// Set error message in Error struct
func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}
