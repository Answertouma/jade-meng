package main

import (
	"fmt"
	"love"
	"context"
)

type WoquHandler struct {

}

func NewWoquHandler() WoquHandler  {
	return WoquHandler{}
}

func (w WoquHandler) GetId(ctx context.Context) (int32, error)  {
	fmt.Println("getId")
	return 1,nil
}

func (w WoquHandler) GetStruct(ctx context.Context) (*love.LoveStruct, error) {
	fmt.Println("getstruct")
	return &love.LoveStruct{1,"woqu"}, nil
}

