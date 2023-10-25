package repositories

import (
	"db.sampes.puc/entities"
	"errors"
	"gorm.io/gorm"
	"time"
)

func CreateRobozinho(input entities.Robozinho) error {
	input.DataCadastro = time.Now()
	return getDbInstance().Create(&input).Error
}

func UpdateRobozinho(input entities.Robozinho) error {
	time := time.Now()
	input.DataAtualizacao = &time
	return getDbInstance().Where(&entities.Robozinho{Id: input.Id}).Save(&input).Error
}

func DeleteRobozinho(id uint) error {
	return getDbInstance().Where(&entities.Robozinho{Id: id}).Delete(&entities.Robozinho{}).Error
}

func GetRobozinho(id uint) (*entities.Robozinho, error) {
	var output entities.Robozinho
	res := getDbInstance().Where(&entities.Robozinho{Id: id}).First(&output)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, res.Error
		}
	}
	return &output, nil
}

func GetAllRobozinhos() ([]entities.Robozinho, error) {
	var output []entities.Robozinho
	res := getDbInstance().Limit(10).Offset(0).Order("Id desc").Find(&output)
	if res.Error != nil {
		return nil, res.Error
	}
	return output, nil
}

func CreateRobozinhoRaw(input entities.Robozinho) error {
	return getDbInstance().Exec(`INSERT INTO Robozinho (Nome, Peso, Foto, data_cadastro) VALUES (?, ?, ?, ?)`, input.Nome, input.Peso, input.Foto, time.Now()).Error
}

func UpdateRobozinhoRaw(input entities.Robozinho) error {
	return getDbInstance().Exec("UPDATE Robozinho SET Nome = ?, Peso = ?, Foto = ?, data_atualizacao = ? WHERE Id = ?", input.Nome, input.Peso, input.Foto, time.Now(), input.Id).Error
}

func DeleteRobozinhoRaw(id uint) error {
	return getDbInstance().Exec(`DELETE FROM Robozinho WHERE Id = ?`, id).Error
}

func GetRobozinhoRaw(id uint) (*entities.Robozinho, error) {
	var output entities.Robozinho
	res := getDbInstance().Raw(`SELECT * FROM Robozinho WHERE Id = ? for update`, id).First(&output)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, res.Error
		}
	}
	return &output, nil
}

func GetAllRobozinhosRaw() ([]entities.Robozinho, error) {
	var output []entities.Robozinho
	res := getDbInstance().Raw(`SELECT * FROM Robozinho ORDER BY Id DESC LIMIT 10 OFFSET 0`).Scan(&output)
	if res.Error != nil {
		return nil, res.Error
	}
	return output, nil
}
