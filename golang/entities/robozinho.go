package entities

import "time"

type Robozinho struct {
	Id              uint `gorm:"primaryKey;autoIncrement;not null"`
	Nome            string
	Peso            float64
	Foto            string
	DataCadastro    time.Time
	DataAtualizacao *time.Time
}

func (Robozinho) TableName() string {
	return "Robozinho"
}
