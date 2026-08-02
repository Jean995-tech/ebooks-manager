package reportes

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/Jean995-tech/ebooks-manager/gestion_descargas"
	"github.com/Jean995-tech/ebooks-manager/gestion_libros"
	"github.com/Jean995-tech/ebooks-manager/gestion_usuarios"
)

// LibroConteo agrupa un libro con su cantidad total de descargas
type LibroConteo struct {
	Libro     gestion_libros.Libro
	Descargas int
}

// Resumen contiene los totales generales del sistema
type Resumen struct {
	TotalLibros    int
	TotalUsuarios  int
	TotalDescargas int
}

// ResumenGeneral retorna un resumen con los totales actuales del sistema:
// cantidad de libros, usuarios registrados y descargas realizadas.
func ResumenGeneral() Resumen {
	libros := gestion_libros.ListarLibros()
	usuarios := gestion_usuarios.ListarUsuarios()
	descargas, _ := gestion_descargas.CargarDescargas()
	return Resumen{
		TotalLibros:    len(libros),
		TotalUsuarios:  len(usuarios),
		TotalDescargas: len(descargas),
	}
}

// LibrosMasDescargados retorna los n libros con mayor numero de descargas,
// ordenados de mayor a menor. Si n es mayor al total de libros, retorna todos.
func LibrosMasDescargados(n int) ([]LibroConteo, error) {
	if n <= 0 {
		return nil, fmt.Errorf("el numero de libros debe ser mayor a 0")
	}
	libros := gestion_libros.ListarLibros()
	descargas, _ := gestion_descargas.CargarDescargas()

	// Contar descargas por libro usando un mapa ID -> cantidad
	conteo := make(map[int]int)
	for _, d := range descargas {
		conteo[d.LibroID]++
	}

	// Construir la lista de libros con su conteo
	var resultado []LibroConteo
	for _, l := range libros {
		resultado = append(resultado, LibroConteo{
			Libro:     l,
			Descargas: conteo[l.ID],
		})
	}

	// Ordenar de mayor a menor cantidad de descargas
	sort.Slice(resultado, func(i, j int) bool {
		return resultado[i].Descargas > resultado[j].Descargas
	})

	if n > len(resultado) {
		n = len(resultado)
	}
	return resultado[:n], nil
}

// UsuariosActivos retorna los usuarios que han realizado al menos una
// descarga en los ultimos n dias.
func UsuariosActivos(dias int) ([]gestion_usuarios.Usuario, error) {
	if dias <= 0 {
		return nil, fmt.Errorf("el numero de dias debe ser mayor a 0")
	}
	usuarios := gestion_usuarios.ListarUsuarios()
	descargas, _ := gestion_descargas.CargarDescargas()

	// Calcular la fecha limite para considerar un usuario como activo
	activos := make(map[int]bool)
	limite := time.Now().AddDate(0, 0, -dias).Format("2006-01-02")

	for _, d := range descargas {
		fecha := strings.Split(d.FechaHora, " ")[0]
		if fecha >= limite {
			activos[d.UsuarioID] = true
		}
	}

	// Filtrar solo los usuarios que aparecen en el mapa de activos
	var resultado []gestion_usuarios.Usuario
	for _, u := range usuarios {
		if activos[u.ID] {
			resultado = append(resultado, u)
		}
	}
	return resultado, nil
}
