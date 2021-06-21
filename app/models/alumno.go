package models

import (
	"fmt"

	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type Alumno struct {
	Id         string `gorm:"primaryKey;"`
	Nombres    string
	Codigo     string
	Matriculas []Matricula
}

func (tab Alumno) ToString() string {
	return tab.Nombres
}

func (tab *Alumno) BeforeCreate(*gorm.DB) error {
	tab.Id = uuid.NewV4().String()
	return nil
}

func (alumno Alumno) FindAll(conn *gorm.DB) ([]Alumno, error) {
	var alumnos []Alumno
	if err := conn.Preload("Matriculas").Find(&alumnos).Error; err != nil {
		return nil, err
	}
	return alumnos, nil
}

func (alumno Alumno) GetAll(conn *gorm.DB) ([]Alumno, error) {
	var alumnos []Alumno
	if err := conn.Find(&alumnos).Error; err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//fmt.Printf("Error: %v", err)
		//return fmt.Errorf("Error: %v", err)
		//continue
		return nil, fmt.Errorf("Error: %v", err)
	}
	return alumnos, nil
}
