package interfaces

import (
	"github.com/Jean995-tech/ebooks-manager/gestion_libros"
	"github.com/Jean995-tech/ebooks-manager/gestion_usuarios"
	"github.com/Jean995-tech/ebooks-manager/gestion_descargas"
)

// GestorLibros define las operaciones basicas que debe tener
// cualquier gestor de libros en el sistema
type GestorLibros interface {
	// Agregar agrega un nuevo libro al catalogo
	Agregar(titulo, autor, genero, formato string, anio int) error
	// Listar retorna todos los libros disponibles
	Listar() []gestion_libros.Libro
	// Buscar busca libros por titulo o autor
	Buscar(query string) []gestion_libros.Libro
	// Eliminar elimina un libro por su ID
	Eliminar(id int) error
}

// GestorUsuarios define las operaciones basicas para
// administrar usuarios del sistema
type GestorUsuarios interface {
	// Registrar agrega un nuevo usuario validando email unico
	Registrar(nombre, email string) error
	// Listar retorna todos los usuarios registrados
	Listar() []gestion_usuarios.Usuario
	// Buscar busca un usuario por su email
	Buscar(email string) (gestion_usuarios.Usuario, error)
}

// GestorDescargas define las operaciones para controlar
// las descargas de libros por usuario
type GestorDescargas interface {
	// Registrar registra una nueva descarga verificando el limite diario
	Registrar(usuarioID, libroID int) error
	// Historial retorna todas las descargas de un usuario
	Historial(usuarioID int) []gestion_descargas.Descarga
	// VerificarLimite verifica si el usuario puede seguir descargando
	VerificarLimite(usuarioID int) bool
}