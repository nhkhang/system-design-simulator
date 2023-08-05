package constants

type ResponseStatus int

const (
	Success      ResponseStatus = 200
	Unauthorized ResponseStatus = 401
	NotFound     ResponseStatus = 404
)
