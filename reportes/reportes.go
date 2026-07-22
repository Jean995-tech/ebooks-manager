package reportes

import (
	"sort"
	"strings"
	"time"

	"github.com/Jean995-tech/ebooks-manager/gestion_descargas"
	"github.com/Jean995-tech/ebooks-manager/gestion_libros"
	"github.com/Jean995-tech/ebooks-manager/gestion_usuarios"
)

type LibroConteo struct {
	Libro     gestion_libros.Libro
	Descargas int
}

type Resumen struct {
	TotalLibros    int
	TotalUsuarios  int
	TotalDescargas int
}

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

func LibrosMasDescargados(n int) []LibroConteo {
	libros := gestion_libros.ListarLibros()
	descargas, _ := gestion_descargas.CargarDescargas()

	conteo := make(map[int]int)
	for _, d := range descargas {
		conteo[d.LibroID]++
	}

	var resultado []LibroConteo
	for _, l := range libros {
		resultado = append(resultado, LibroConteo{
			Libro:     l,
			Descargas: conteo[l.ID],
		})
	}

	sort.Slice(resultado, func(i, j int) bool {
		return resultado[i].Descargas > resultado[j].Descargas
	})

	if n > len(resultado) {
		n = len(resultado)
	}
	return resultado[:n]
}

func UsuariosActivos(dias int) []gestion_usuarios.Usuario {
	usuarios := gestion_usuarios.ListarUsuarios()
	descargas, _ := gestion_descargas.CargarDescargas()

	activos := make(map[int]bool)
	limite := time.Now().AddDate(0, 0, -dias).Format("2006-01-02")

	for _, d := range descargas {
		fecha := strings.Split(d.FechaHora, " ")[0]
		if fecha >= limite {
			activos[d.UsuarioID] = true
		}
	}

	var resultado []gestion_usuarios.Usuario
	for _, u := range usuarios {
		if activos[u.ID] {
			resultado = append(resultado, u)
		}
	}
	return resultado
}
