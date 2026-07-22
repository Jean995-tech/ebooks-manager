package gestion_libros

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

type Libro struct {
	ID      int    `json:"id"`
	Titulo  string `json:"titulo"`
	Autor   string `json:"autor"`
	Genero  string `json:"genero"`
	Formato string `json:"formato"`
	Anio    int    `json:"anio"`
}

const archivoLibros = "data/libros.json"

func CargarLibros() ([]Libro, error) {
	datos, err := os.ReadFile(archivoLibros)
	if err != nil {
		if os.IsNotExist(err) {
			return []Libro{}, nil
		}
		return nil, err
	}
	var libros []Libro
	err = json.Unmarshal(datos, &libros)
	return libros, err
}

func GuardarLibros(libros []Libro) error {
	datos, err := json.MarshalIndent(libros, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(archivoLibros, datos, 0644)
}

func AgregarLibro(titulo, autor, genero, formato string, anio int) error {
	libros, err := CargarLibros()
	if err != nil {
		return err
	}
	nuevoID := 1
	if len(libros) > 0 {
		nuevoID = libros[len(libros)-1].ID + 1
	}
	libro := Libro{
		ID:      nuevoID,
		Titulo:  titulo,
		Autor:   autor,
		Genero:  genero,
		Formato: formato,
		Anio:    anio,
	}
	libros = append(libros, libro)
	return GuardarLibros(libros)
}

func BuscarLibro(query string) []Libro {
	libros, err := CargarLibros()
	if err != nil {
		return []Libro{}
	}
	query = strings.ToLower(query)
	var resultados []Libro
	for _, l := range libros {
		if strings.Contains(strings.ToLower(l.Titulo), query) ||
			strings.Contains(strings.ToLower(l.Autor), query) {
			resultados = append(resultados, l)
		}
	}
	return resultados
}

func ListarLibros() []Libro {
	libros, err := CargarLibros()
	if err != nil {
		return []Libro{}
	}
	return libros
}

func EliminarLibro(id int) error {
	libros, err := CargarLibros()
	if err != nil {
		return err
	}
	for i, l := range libros {
		if l.ID == id {
			libros = append(libros[:i], libros[i+1:]...)
			return GuardarLibros(libros)
		}
	}
	return errors.New("libro no encontrado")
}
