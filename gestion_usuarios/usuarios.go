package gestion_usuarios

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Usuario representa un usuario registrado en el sistema
type Usuario struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Email         string `json:"email"`
	FechaRegistro string `json:"fecha_registro"`
}

// archivoUsuarios es la ruta donde se persisten los usuarios en formato JSON
const archivoUsuarios = "data/usuarios.json"

// CargarUsuarios lee la lista de usuarios desde el archivo JSON.
// Si el archivo no existe retorna una lista vacia sin error.
func CargarUsuarios() ([]Usuario, error) {
	datos, err := os.ReadFile(archivoUsuarios)
	if err != nil {
		if os.IsNotExist(err) {
			return []Usuario{}, nil
		}
		return nil, fmt.Errorf("error al leer el archivo de usuarios: %w", err)
	}
	var usuarios []Usuario
	if err := json.Unmarshal(datos, &usuarios); err != nil {
		return nil, fmt.Errorf("error al parsear el archivo de usuarios: %w", err)
	}
	return usuarios, nil
}

// GuardarUsuarios persiste la lista de usuarios en el archivo JSON.
func GuardarUsuarios(usuarios []Usuario) error {
	datos, err := json.MarshalIndent(usuarios, "", "  ")
	if err != nil {
		return fmt.Errorf("error al serializar los usuarios: %w", err)
	}
	if err := os.WriteFile(archivoUsuarios, datos, 0644); err != nil {
		return fmt.Errorf("error al guardar el archivo de usuarios: %w", err)
	}
	return nil
}

// RegistrarUsuario agrega un nuevo usuario al sistema validando que
// el nombre no este vacio y que el email no este duplicado.
func RegistrarUsuario(nombre, email string) error {
	if nombre == "" {
		return errors.New("el nombre no puede estar vacio")
	}
	if email == "" || !strings.Contains(email, "@") {
		return errors.New("el email no es valido")
	}
	usuarios, err := CargarUsuarios()
	if err != nil {
		return err
	}
	// Verificar que el email no este registrado ya
	for _, u := range usuarios {
		if strings.EqualFold(u.Email, email) {
			return fmt.Errorf("ya existe un usuario registrado con el email %s", email)
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

// BuscarUsuario busca un usuario por su email.
// La busqueda no distingue entre mayusculas y minusculas.
func BuscarUsuario(email string) (Usuario, error) {
	if email == "" {
		return Usuario{}, errors.New("el email no puede estar vacio")
	}
	usuarios, err := CargarUsuarios()
	if err != nil {
		return Usuario{}, err
	}
	for _, u := range usuarios {
		if strings.EqualFold(u.Email, email) {
			return u, nil
		}
	}
	return Usuario{}, fmt.Errorf("no se encontro ningun usuario con el email %s", email)
}

// ListarUsuarios retorna todos los usuarios registrados en el sistema.
func ListarUsuarios() []Usuario {
	usuarios, err := CargarUsuarios()
	if err != nil {
		return []Usuario{}
	}
	return usuarios
}
