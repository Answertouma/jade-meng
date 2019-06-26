package paris

import (
	"heart/handler"
)

type parisHandler struct {
	handler.Answer
}

func NewParisHandler() *parisHandler {

	han := &parisHandler{}
	han.Love = "Toma"
	han.Sheep = 100
	return han
}
