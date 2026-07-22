package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Jean995-tech/ebooks-manager/gestion_descargas"
	"github.com/Jean995-tech/ebooks-manager/gestion_libros"
	"github.com/Jean995-tech/ebooks-manager/gestion_usuarios"
	"github.com/Jean995-tech/ebooks-manager/reportes"
)

var scanner = bufio.NewScanner(os.Stdin)

func leer(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func leerInt(prompt string) int {
	val, _ := strconv.Atoi(leer(prompt))
	return val
}

func main() {
	os.MkdirAll("data", 0755)

	for {
		fmt.Println("\n========================================")
		fmt.Println("   SISTEMA DE GESTION DE EBOOKS")
		fmt.Println("========================================")
		fmt.Println("1. Gestion de Libros")
		fmt.Println("2. Gestion de Usuarios")
		fmt.Println("3. Gestion de Descargas")
		fmt.Println("4. Reportes")
		fmt.Println("0. Salir")
		fmt.Println("----------------------------------------")

		opcion := leer("Selecciona una opcion: ")

		switch opcion {
		case "1":
			menuLibros()
		case "2":
			menuUsuarios()
		case "3":
			menuDescargas()
		case "4":
			menuReportes()
		case "0":
			fmt.Println("Hasta luego!")
			return
		default:
			fmt.Println("Opcion no valida.")
		}
	}
}

func menuLibros() {
	for {
		fmt.Println("\n--- GESTION DE LIBROS ---")
		fmt.Println("1. Agregar libro")
		fmt.Println("2. Listar libros")
		fmt.Println("3. Buscar libro")
		fmt.Println("4. Eliminar libro")
		fmt.Println("0. Volver")

		op := leer("Opcion: ")
		switch op {
		case "1":
			titulo := leer("Titulo: ")
			autor := leer("Autor: ")
			genero := leer("Genero: ")
			formato := leer("Formato (PDF/EPUB/MOBI): ")
			anio := leerInt("Anio: ")
			err := gestion_libros.AgregarLibro(titulo, autor, genero, formato, anio)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Libro agregado exitosamente.")
			}
		case "2":
			libros := gestion_libros.ListarLibros()
			if len(libros) == 0 {
				fmt.Println("No hay libros en el catalogo.")
			}
			for _, l := range libros {
				fmt.Printf("[%d] %s - %s (%d) [%s]\n", l.ID, l.Titulo, l.Autor, l.Anio, l.Formato)
			}
		case "3":
			query := leer("Buscar por titulo o autor: ")
			resultados := gestion_libros.BuscarLibro(query)
			if len(resultados) == 0 {
				fmt.Println("No se encontraron resultados.")
			}
			for _, l := range resultados {
				fmt.Printf("[%d] %s - %s\n", l.ID, l.Titulo, l.Autor)
			}
		case "4":
			id := leerInt("ID del libro a eliminar: ")
			err := gestion_libros.EliminarLibro(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Libro eliminado.")
			}
		case "0":
			return
		}
	}
}

func menuUsuarios() {
	for {
		fmt.Println("\n--- GESTION DE USUARIOS ---")
		fmt.Println("1. Registrar usuario")
		fmt.Println("2. Listar usuarios")
		fmt.Println("3. Buscar usuario por email")
		fmt.Println("0. Volver")

		op := leer("Opcion: ")
		switch op {
		case "1":
			nombre := leer("Nombre: ")
			email := leer("Email: ")
			err := gestion_usuarios.RegistrarUsuario(nombre, email)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Usuario registrado exitosamente.")
			}
		case "2":
			usuarios := gestion_usuarios.ListarUsuarios()
			if len(usuarios) == 0 {
				fmt.Println("No hay usuarios registrados.")
			}
			for _, u := range usuarios {
				fmt.Printf("[%d] %s - %s (registro: %s)\n", u.ID, u.Nombre, u.Email, u.FechaRegistro)
			}
		case "3":
			email := leer("Email: ")
			u, err := gestion_usuarios.BuscarUsuario(email)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Encontrado: [%d] %s - %s\n", u.ID, u.Nombre, u.Email)
			}
		case "0":
			return
		}
	}
}

func menuDescargas() {
	for {
		fmt.Println("\n--- GESTION DE DESCARGAS ---")
		fmt.Println("1. Registrar descarga")
		fmt.Println("2. Ver historial de un usuario")
		fmt.Println("3. Verificar limite diario")
		fmt.Println("0. Volver")

		op := leer("Opcion: ")
		switch op {
		case "1":
			usuarioID := leerInt("ID de usuario: ")
			libroID := leerInt("ID de libro: ")
			err := gestion_descargas.RegistrarDescarga(usuarioID, libroID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Descarga registrada exitosamente.")
			}
		case "2":
			usuarioID := leerInt("ID de usuario: ")
			historial := gestion_descargas.HistorialDescargas(usuarioID)
			if len(historial) == 0 {
				fmt.Println("No hay descargas para este usuario.")
			}
			for _, d := range historial {
				fmt.Printf("Libro ID: %d | Fecha: %s\n", d.LibroID, d.FechaHora)
			}
		case "3":
			usuarioID := leerInt("ID de usuario: ")
			if gestion_descargas.VerificarLimite(usuarioID) {
				fmt.Println("El usuario puede seguir descargando hoy.")
			} else {
				fmt.Println("El usuario alcanzo el limite diario de 5 descargas.")
			}
		case "0":
			return
		}
	}
}

func menuReportes() {
	for {
		fmt.Println("\n--- REPORTES ---")
		fmt.Println("1. Resumen general")
		fmt.Println("2. Libros mas descargados")
		fmt.Println("3. Usuarios activos")
		fmt.Println("0. Volver")

		op := leer("Opcion: ")
		switch op {
		case "1":
			r := reportes.ResumenGeneral()
			fmt.Printf("\nTotal libros:     %d\n", r.TotalLibros)
			fmt.Printf("Total usuarios:   %d\n", r.TotalUsuarios)
			fmt.Printf("Total descargas:  %d\n", r.TotalDescargas)
		case "2":
			n := leerInt("Cuantos libros mostrar: ")
			top := reportes.LibrosMasDescargados(n)
			for i, lc := range top {
				fmt.Printf("%d. %s - %d descargas\n", i+1, lc.Libro.Titulo, lc.Descargas)
			}
		case "3":
			dias := leerInt("Activos en los ultimos cuantos dias: ")
			activos := reportes.UsuariosActivos(dias)
			if len(activos) == 0 {
				fmt.Println("No hay usuarios activos en ese periodo.")
			}
			for _, u := range activos {
				fmt.Printf("[%d] %s - %s\n", u.ID, u.Nombre, u.Email)
			}
		case "0":
			return
		}
	}
}
