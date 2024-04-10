package models

import "time"

type Cliente struct {
	Id        uint       `json:"id" gorm:"primaryKey"`
	Nombre    string     `json:"nombre"`
	Apellido  string     `json:"apellido"`
	Fecha     *time.Time `json:"fecha"`
	Cuit      string     `json:"cuit"`
	Domicilio string     `json:"domicilio"`
	Telefono  string     `json:"telefono"`
	Mail      string     `json:"mail"`
}
