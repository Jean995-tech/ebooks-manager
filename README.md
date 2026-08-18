<div align="center">

# 📚 ebooks-manager

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/Licencia-Académica-blue?style=for-the-badge)
![Status](https://img.shields.io/badge/Estado-Completo-green?style=for-the-badge)
![UIDE](https://img.shields.io/badge/UIDE-Ciberseguridad-8B0000?style=for-the-badge)

**Sistema de Gestión de Libros Electrónicos**  
Desarrollado en Go · Programación Funcional · API REST · Sin dependencias externas

*Programación Orientada a Objetos — Jean Pierre Males Cedeño — UIDE 2026*

</div>

---

## 🎯 ¿Qué hace este sistema?

Administra una biblioteca digital desde la consola y a través de una API REST. Puedes agregar libros, registrar usuarios, controlar descargas y generar reportes. Los datos se guardan en archivos JSON locales entre sesiones.

---

## 📋 Avances del Proyecto

### ✅ Autónomo 1 — Planeación y estructura base
- Estructura del proyecto en 4 módulos independientes
- Sistema CRUD completo para libros
- Registro y consulta de usuarios
- Control de descargas con límite diario
- Persistencia en archivos JSON
- Menú interactivo en consola

### ✅ Autónomo 2 — Interfaces, encapsulación y manejo de errores
- Interfaces definidas para cada módulo
- Comentarios en todas las funciones
- Manejo de errores mejorado con mensajes descriptivos
- Validaciones de datos
- Encapsulación documentada

### ✅ Proyecto Final — Servicios Web REST
- 8 endpoints REST implementados
- Serialización JSON en todas las respuestas
- Servidor web en puerto 8080
- Compatible con cualquier cliente HTTP

---

## 🗂️ Módulos

```
📦 ebooks-manager
 ┣ 📖 gestion_libros     → Catálogo completo de libros (CRUD)
 ┣ 👤 gestion_usuarios   → Registro y consulta de usuarios
 ┣ ⬇️  gestion_descargas → Historial con límite diario por usuario
 ┣ 📊 reportes           → Estadísticas y resúmenes del sistema
 ┣ 🔌 interfaces         → Contratos de cada módulo
 ┗ 🌐 server             → API REST con 8 servicios web
```

---

## 🏗️ Estructura del Proyecto

```
ebooks-manager/
├── main.go                      # Punto de entrada y menú principal
├── server.go                    # Servidor web con 8 servicios REST
├── go.mod                       # Configuración del módulo Go
├── README.md
├── data/
│   ├── libros.json              # Catálogo persistente de libros
│   ├── usuarios.json            # Usuarios registrados
│   └── descargas.json           # Historial de descargas
├── gestion_libros/
│   └── libros.go
├── gestion_usuarios/
│   └── usuarios.go
├── gestion_descargas/
│   └── descargas.go
├── reportes/
│   └── reportes.go
└── interfaces/
    └── interfaces.go
```

---

## 🌐 API REST — 8 Servicios Web

| # | Método | Endpoint | Descripción |
|---|---|---|---|
| 1 | GET | `/libros` | Listar todos los libros |
| 2 | POST | `/libros` | Agregar un libro nuevo |
| 3 | DELETE | `/libros?id=1` | Eliminar un libro por ID |
| 4 | GET | `/libros/buscar?q=Harry` | Buscar libros por título o autor |
| 5 | GET | `/usuarios` | Listar todos los usuarios |
| 6 | POST | `/usuarios` | Registrar un usuario nuevo |
| 7 | POST | `/descargas` | Registrar una descarga |
| 8 | GET | `/descargas/historial?usuario_id=1` | Ver historial de descargas |

### Ejemplos de respuesta JSON:

**GET /libros:**
```json
[
  {
    "id": 1,
    "titulo": "Harry Potter",
    "autor": "J.K. Rowling",
    "genero": "Fantasia",
    "formato": "PDF",
    "anio": 1997
  }
]
```

**GET /reportes/resumen:**
```json
{
  "TotalLibros": 5,
  "TotalUsuarios": 3,
  "TotalDescargas": 4
}
```

---

## ⚙️ Requisitos

- Go 1.21 o superior → [descargar Go](https://go.dev/dl/)
- Git → [descargar Git](https://git-scm.com/download/win)
- Sin dependencias externas

---

## 🚀 Instalación y Uso

```bash
# 1. Clonar el repositorio
git clone https://github.com/Jean995-tech/ebooks-manager.git

# 2. Entrar al directorio
cd ebooks-manager

# 3. Ejecutar el sistema
go run .
```

### Iniciar el servidor web:
```
Selecciona la opcion 5 en el menu principal
Servidor corriendo en http://localhost:8080
```

---

## ✅ Funcionalidades

| Módulo | Funciones |
|---|---|
| 📖 Libros | Agregar, buscar, listar, eliminar con validaciones |
| 👤 Usuarios | Registrar con validación de email, buscar, listar |
| ⬇️ Descargas | Registrar, ver historial, verificar límite diario |
| 📊 Reportes | Libros más descargados, usuarios activos, resumen |
| 🌐 API REST | 8 endpoints con serialización JSON |

---

## 📦 Paquetes Utilizados

| Paquete | Uso |
|---|---|
| `fmt` | Entrada/salida en consola |
| `bufio` | Lectura de input del usuario |
| `os` | Interacción con archivos del sistema |
| `strings` | Manipulación de cadenas |
| `strconv` | Conversión entre tipos de datos |
| `time` | Registro de fecha y hora |
| `encoding/json` | Serialización JSON |
| `net/http` | Servidor web y API REST |
| `errors` | Manejo de errores descriptivos |
| `sort` | Ordenamiento de reportes |

---

## 🔧 Control de Versiones

```bash
git add .
git commit -m "descripcion del cambio"
git push
```

---

<div align="center">

**Jean Pierre Males Cedeño**  
Ingeniería en Ciberseguridad · UIDE · 2026

</div>
