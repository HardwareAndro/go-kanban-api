package config

import (
	"github.com/go-playground/validator/v10"
	"log"
)

type GoAppTools struct {
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
	Validate    *validator.Validate
}
