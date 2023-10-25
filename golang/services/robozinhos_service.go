package services

import (
	"db.sampes.puc/entities"
	"db.sampes.puc/repositories"
)

func CreateRobozinho(input entities.Robozinho, raw bool) error {
	if raw {
		return repositories.CreateRobozinhoRaw(input)
	}
	return repositories.CreateRobozinho(input)
}

func UpdateRobozinho(input entities.Robozinho, raw bool) error {
	if raw {
		return repositories.UpdateRobozinhoRaw(input)
	}
	return repositories.UpdateRobozinho(input)
}

func DeleteRobozinho(id uint, raw bool) error {
	if raw {
		return repositories.DeleteRobozinhoRaw(id)
	}
	return repositories.DeleteRobozinho(id)
}

func GetRobozinho(id uint, raw bool) (*entities.Robozinho, error) {
	if raw {
		return repositories.GetRobozinhoRaw(id)
	}
	return repositories.GetRobozinho(id)
}

func GetAllRobozinhos(raw bool) ([]entities.Robozinho, error) {
	if raw {
		return repositories.GetAllRobozinhosRaw()
	}
	return repositories.GetAllRobozinhos()
}
