package gestion_descargas

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Descarga representa el registro de una descarga realizada por un usuario
type Descarga struct {
	ID        int    `json:"id"`
	UsuarioID int    `json:"usuario_id"`
	LibroID   int    `json:"libro_id"`
	FechaHora string `json:"fecha_hora"`
}

// archivoDescargas es la ruta donde se persiste el historial en formato JSON
const archivoDescargas = "data/descargas.json"

// limiteDiario define el maximo de descargas permitidas por usuario por dia
const limiteDiario = 5

// CargarDescargas lee el historial de descargas desde el archivo JSON.
// Si el archivo no existe retorna una lista vacia sin error.
func CargarDescargas() ([]Descarga, error) {
	datos, err := os.ReadFile(archivoDescargas)
	if err != nil {
		if os.IsNotExist(err) {
			return []Descarga{}, nil
		}
		return nil, fmt.Errorf("error al leer el archivo de descargas: %w", err)
	}
	var descargas []Descarga
	if err := json.Unmarshal(datos, &descargas); err != nil {
		return nil, fmt.Errorf("error al parsear el archivo de descargas: %w", err)
	}
	return descargas, nil
}

// GuardarDescargas persiste el historial de descargas en el archivo JSON.
func GuardarDescargas(descargas []Descarga) error {
	datos, err := json.MarshalIndent(descargas, "", "  ")
	if err != nil {
		return fmt.Errorf("error al serializar las descargas: %w", err)
	}
	if err := os.WriteFile(archivoDescargas, datos, 0644); err != nil {
		return fmt.Errorf("error al guardar el archivo de descargas: %w", err)
	}
	return nil
}

// VerificarLimite verifica si un usuario puede realizar mas descargas hoy.
// Retorna true si el usuario no ha alcanzado el limite diario.
func VerificarLimite(usuarioID int) bool {
	descargas, err := CargarDescargas()
	if err != nil {
		return false
	}
	// Contar descargas del usuario en el dia de hoy
	hoy := time.Now().Format("2006-01-02")
	count := 0
	for _, d := range descargas {
		if d.UsuarioID == usuarioID && strings.HasPrefix(d.FechaHora, hoy) {
			count++
		}
	}
	return count < limiteDiario
}

// RegistrarDescarga registra una nueva descarga verificando primero
// que el usuario no haya superado el limite diario de descargas.
func RegistrarDescarga(usuarioID, libroID int) error {
	if usuarioID <= 0 {
		return errors.New("el ID de usuario debe ser un numero positivo")
	}
	if libroID <= 0 {
		return errors.New("el ID de libro debe ser un numero positivo")
	}
	// Verificar limite antes de registrar
	if !VerificarLimite(usuarioID) {
		return fmt.Errorf("el usuario %d alcanzo el limite diario de %d descargas", usuarioID, limiteDiario)
	}
	descargas, err := CargarDescargas()
	if err != nil {
		return err
	}
	nuevoID := 1
	if len(descargas) > 0 {
		nuevoID = descargas[len(descargas)-1].ID + 1
	}
	descarga := Descarga{
		ID:        nuevoID,
		UsuarioID: usuarioID,
		LibroID:   libroID,
		FechaHora: time.Now().Format("2006-01-02 15:04:05"),
	}
	descargas = append(descargas, descarga)
	return GuardarDescargas(descargas)
}

// HistorialDescargas retorna todas las descargas realizadas por un usuario
// identificado por su ID.
func HistorialDescargas(usuarioID int) []Descarga {
	if usuarioID <= 0 {
		return []Descarga{}
	}
	descargas, err := CargarDescargas()
	if err != nil {
		return []Descarga{}
	}
	var historial []Descarga
	for _, d := range descargas {
		if d.UsuarioID == usuarioID {
			historial = append(historial, d)
		}
	}
	return historial
}

