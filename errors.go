package main

import "errors"

func SetError(msg string) error {
	return errors.New(msg)
}