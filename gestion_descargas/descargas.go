package gestion_descargas

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

type Descarga struct {
	ID        int    `json:"id"`
	UsuarioID int    `json:"usuario_id"`
	LibroID   int    `json:"libro_id"`
	FechaHora string `json:"fecha_hora"`
}

const archivoDescargas = "data/descargas.json"
const limiteDiario = 5

func CargarDescargas() ([]Descarga, error) {
	datos, err := os.ReadFile(archivoDescargas)
	if err != nil {
		if os.IsNotExist(err) {
			return []Descarga{}, nil
		}
		return nil, err
	}
	var descargas []Descarga
	err = json.Unmarshal(datos, &descargas)
	return descargas, err
}

func GuardarDescargas(descargas []Descarga) error {
	datos, err := json.MarshalIndent(descargas, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(archivoDescargas, datos, 0644)
}

func VerificarLimite(usuarioID int) bool {
	descargas, err := CargarDescargas()
	if err != nil {
		return false
	}
	hoy := time.Now().Format("2006-01-02")
	count := 0
	for _, d := range descargas {
		if d.UsuarioID == usuarioID && strings.HasPrefix(d.FechaHora, hoy) {
			count++
		}
	}
	return count < limiteDiario
}

func RegistrarDescarga(usuarioID, libroID int) error {
	if !VerificarLimite(usuarioID) {
		return errors.New("limite diario de descargas alcanzado (maximo 5 por dia)")
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

func HistorialDescargas(usuarioID int) []Descarga {
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
