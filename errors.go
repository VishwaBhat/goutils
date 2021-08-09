package main

import "errors"

var (
	//ErrNoElements indicates when no elements are present in the container
	ErrNoElements = errors.New("no elements the container")
)
