package gestion_libros

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Libro representa un libro electronico en el catalogo del sistema
type Libro struct {
	ID      int    `json:"id"`
	Titulo  string `json:"titulo"`
	Autor   string `json:"autor"`
	Genero  string `json:"genero"`
	Formato string `json:"formato"`
	Anio    int    `json:"anio"`
}

// archivoLibros es la ruta donde se persiste el catalogo en formato JSON
const archivoLibros = "data/libros.json"

// CargarLibros lee el catalogo de libros desde el archivo JSON.
// Si el archivo no existe retorna una lista vacia sin error.
func CargarLibros() ([]Libro, error) {
	datos, err := os.ReadFile(archivoLibros)
	if err != nil {
		if os.IsNotExist(err) {
			return []Libro{}, nil
		}
		return nil, fmt.Errorf("error al leer el archivo de libros: %w", err)
	}
	var libros []Libro
	if err := json.Unmarshal(datos, &libros); err != nil {
		return nil, fmt.Errorf("error al parsear el archivo de libros: %w", err)
	}
	return libros, nil
}

// GuardarLibros persiste la lista de libros en el archivo JSON.
func GuardarLibros(libros []Libro) error {
	datos, err := json.MarshalIndent(libros, "", "  ")
	if err != nil {
		return fmt.Errorf("error al serializar los libros: %w", err)
	}
	if err := os.WriteFile(archivoLibros, datos, 0644); err != nil {
		return fmt.Errorf("error al guardar el archivo de libros: %w", err)
	}
	return nil
}

// AgregarLibro agrega un nuevo libro al catalogo validando que los
// campos obligatorios no esten vacios antes de persistir.
func AgregarLibro(titulo, autor, genero, formato string, anio int) error {
	if titulo == "" || autor == "" {
		return errors.New("el titulo y el autor son campos obligatorios")
	}
	if anio < 1000 || anio > 2100 {
		return errors.New("el anio debe estar entre 1000 y 2100")
	}
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

// BuscarLibro busca libros por titulo o autor usando coincidencia parcial
// sin distinguir entre mayusculas y minusculas.
func BuscarLibro(query string) []Libro {
	if query == "" {
		return []Libro{}
	}
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

// ListarLibros retorna todos los libros del catalogo ordenados por ID.
func ListarLibros() []Libro {
	libros, err := CargarLibros()
	if err != nil {
		return []Libro{}
	}
	return libros
}

// EliminarLibro elimina un libro del catalogo por su ID.
// Retorna error si el ID no existe en el catalogo.
func EliminarLibro(id int) error {
	if id <= 0 {
		return errors.New("el ID debe ser un numero positivo")
	}
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
	return fmt.Errorf("no se encontro ningun libro con ID %d", id)
}