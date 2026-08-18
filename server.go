package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jean995-tech/ebooks-manager/gestion_descargas"
	"github.com/Jean995-tech/ebooks-manager/gestion_libros"
	"github.com/Jean995-tech/ebooks-manager/gestion_usuarios"
	"github.com/Jean995-tech/ebooks-manager/reportes"
)

// iniciarServidor arranca el servidor web en el puerto 8080
func iniciarServidor() {
	// Rutas de libros
	http.HandleFunc("/libros", manejarLibros)
	http.HandleFunc("/libros/buscar", buscarLibroHandler)

	// Rutas de usuarios
	http.HandleFunc("/usuarios", manejarUsuarios)

	// Rutas de descargas
	http.HandleFunc("/descargas", manejarDescargas)
	http.HandleFunc("/descargas/historial", historialDescargasHandler)
	http.HandleFunc("/descargas/limite", verificarLimiteHandler)

	// Ruta de reportes
	http.HandleFunc("/reportes/resumen", resumenHandler)
	http.HandleFunc("/reportes/top", topLibrosHandler)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// responderJSON envia una respuesta en formato JSON
func responderJSON(w http.ResponseWriter, datos interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(datos)
}

// manejarLibros maneja GET y POST para /libros
func manejarLibros(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Servicio 1: Listar todos los libros
		libros := gestion_libros.ListarLibros()
		responderJSON(w, libros, http.StatusOK)

	case http.MethodPost:
		// Servicio 2: Agregar un libro nuevo
		var libro gestion_libros.Libro
		if err := json.NewDecoder(r.Body).Decode(&libro); err != nil {
			responderJSON(w, map[string]string{"error": "datos invalidos"}, http.StatusBadRequest)
			return
		}
		err := gestion_libros.AgregarLibro(libro.Titulo, libro.Autor, libro.Genero, libro.Formato, libro.Anio)
		if err != nil {
			responderJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
			return
		}
		responderJSON(w, map[string]string{"mensaje": "libro agregado exitosamente"}, http.StatusCreated)

	case http.MethodDelete:
		// Servicio 3: Eliminar un libro por ID
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			responderJSON(w, map[string]string{"error": "ID invalido"}, http.StatusBadRequest)
			return
		}
		err = gestion_libros.EliminarLibro(id)
		if err != nil {
			responderJSON(w, map[string]string{"error": err.Error()}, http.StatusNotFound)
			return
		}
		responderJSON(w, map[string]string{"mensaje": "libro eliminado exitosamente"}, http.StatusOK)

	default:
		responderJSON(w, map[string]string{"error": "metodo no permitido"}, http.StatusMethodNotAllowed)
	}
}

// buscarLibroHandler maneja GET /libros/buscar?q=query
// Servicio 4: Buscar libros por titulo o autor
func buscarLibroHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		responderJSON(w, map[string]string{"error": "metodo no permitido"}, http.StatusMethodNotAllowed)
		return
	}
	query := r.URL.Query().Get("q")
	resultados := gestion_libros.BuscarLibro(query)
	responderJSON(w, resultados, http.StatusOK)
}

// manejarUsuarios maneja GET y POST para /usuarios
func manejarUsuarios(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Servicio 5: Listar todos los usuarios
		usuarios := gestion_usuarios.ListarUsuarios()
		responderJSON(w, usuarios, http.StatusOK)

	case http.MethodPost:
		// Servicio 6: Registrar un usuario nuevo
		var usuario gestion_usuarios.Usuario
		if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
			responderJSON(w, map[string]string{"error": "datos invalidos"}, http.StatusBadRequest)
			return
		}
		err := gestion_usuarios.RegistrarUsuario(usuario.Nombre, usuario.Email)
		if err != nil {
			responderJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
			return
		}
		responderJSON(w, map[string]string{"mensaje": "usuario registrado exitosamente"}, http.StatusCreated)

	default:
		responderJSON(w, map[string]string{"error": "metodo no permitido"}, http.StatusMethodNotAllowed)
	}
}

// manejarDescargas maneja GET y POST para /descargas
func manejarDescargas(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Servicio 7: Registrar una descarga
		var descarga gestion_descargas.Descarga
		if err := json.NewDecoder(r.Body).Decode(&descarga); err != nil {
			responderJSON(w, map[string]string{"error": "datos invalidos"}, http.StatusBadRequest)
			return
		}
		err := gestion_descargas.RegistrarDescarga(descarga.UsuarioID, descarga.LibroID)
		if err != nil {
			responderJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
			return
		}
		responderJSON(w, map[string]string{"mensaje": "descarga registrada exitosamente"}, http.StatusCreated)

	default:
		responderJSON(w, map[string]string{"error": "metodo no permitido"}, http.StatusMethodNotAllowed)
	}
}

// historialDescargasHandler maneja GET /descargas/historial?usuario_id=1
// Servicio 8: Ver historial de descargas de un usuario
func historialDescargasHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		responderJSON(w, map[string]string{"error": "metodo no permitido"}, http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Query().Get("usuario_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		responderJSON(w, map[string]string{"error": "ID de usuario invalido"}, http.StatusBadRequest)
		return
	}
	historial := gestion_descargas.HistorialDescargas(id)
	responderJSON(w, historial, http.StatusOK)
}

// verificarLimiteHandler maneja GET /descargas/limite?usuario_id=1
func verificarLimiteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		responderJSON(w, map[string]string{"error": "metodo no permitido"}, http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Query().Get("usuario_id")
	id, _ := strconv.Atoi(idStr)
	puede := gestion_descargas.VerificarLimite(id)
	responderJSON(w, map[string]bool{"puede_descargar": puede}, http.StatusOK)
}

// resumenHandler maneja GET /reportes/resumen
func resumenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		responderJSON(w, map[string]string{"error": "metodo no permitido"}, http.StatusMethodNotAllowed)
		return
	}
	resumen := reportes.ResumenGeneral()
	responderJSON(w, resumen, http.StatusOK)
}

// topLibrosHandler maneja GET /reportes/top?n=3
func topLibrosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		responderJSON(w, map[string]string{"error": "metodo no permitido"}, http.StatusMethodNotAllowed)
		return
	}
	nStr := r.URL.Query().Get("n")
	n, err := strconv.Atoi(nStr)
	if err != nil || n <= 0 {
		n = 5
	}
	top, err := reportes.LibrosMasDescargados(n)
	if err != nil {
		responderJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}
	responderJSON(w, top, http.StatusOK)
}