package gestion_usuarios

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Usuario struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Email         string `json:"email"`
	FechaRegistro string `json:"fecha_registro"`
}

const archivoUsuarios = "data/usuarios.json"

func CargarUsuarios() ([]Usuario, error) {
	datos, err := os.ReadFile(archivoUsuarios)
	if err != nil {
		if os.IsNotExist(err) {
			return []Usuario{}, nil
		}
		return nil, err
	}
	var usuarios []Usuario
	err = json.Unmarshal(datos, &usuarios)
	return usuarios, err
}

func GuardarUsuarios(usuarios []Usuario) error {
	datos, err := json.MarshalIndent(usuarios, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(archivoUsuarios, datos, 0644)
}

func RegistrarUsuario(nombre, email string) error {
	usuarios, err := CargarUsuarios()
	if err != nil {
		return err
	}
	for _, u := range usuarios {
		if u.Email == email {
			return errors.New("ya existe un usuario con ese email")
		}
	}
	nuevoID := 1
	if len(usuarios) > 0 {
		nuevoID = usuarios[len(usuarios)-1].ID + 1
	}
	usuario := Usuario{
		ID:            nuevoID,
		Nombre:        nombre,
		Email:         email,
		FechaRegistro: time.Now().Format("2006-01-02"),
	}
	usuarios = append(usuarios, usuario)
	return GuardarUsuarios(usuarios)
}

func BuscarUsuario(email string) (Usuario, error) {
	usuarios, err := CargarUsuarios()
	if err != nil {
		return Usuario{}, err
	}
	for _, u := range usuarios {
		if u.Email == email {
			return u, nil
		}
	}
	return Usuario{}, errors.New("usuario no encontrado")
}

func ListarUsuarios() []Usuario {
	usuarios, err := CargarUsuarios()
	if err != nil {
		return []Usuario{}
	}
	return usuarios
}
