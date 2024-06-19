package models

import (
	"fmt"
)


type Articulo struct {
	ID int `json:"id"`
	Titulo string `json:"titulo"`
	Contenido string `json:"contenido"`
	Autor string `json:"autor"`
}