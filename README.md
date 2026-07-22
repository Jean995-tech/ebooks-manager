<div align="center">

# 📚 ebooks-manager

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/Licencia-Académica-blue?style=for-the-badge)
![Status](https://img.shields.io/badge/Estado-En%20Desarrollo-yellow?style=for-the-badge)
![UIDE](https://img.shields.io/badge/UIDE-Ciberseguridad-8B0000?style=for-the-badge)

**Sistema de Gestión de Libros Electrónicos**  
Desarrollado en Go · Programación Funcional · Sin dependencias externas

*Programación Orientada a Objetos — Jean Pierre Males Cedeño — UIDE 2026*

</div>

---

## 🎯 ¿Qué hace este sistema?

Administra una biblioteca digital desde la consola. Puedes agregar libros, registrar usuarios, controlar descargas y generar reportes, todo sin instalar nada más que Go. Los datos se guardan en archivos JSON locales entre sesiones.

---

## 🗂️ Módulos

```
📦 ebooks-manager
 ┣ 📖 gestion_libros     → Catálogo completo de libros (CRUD)
 ┣ 👤 gestion_usuarios   → Registro y consulta de usuarios
 ┣ ⬇️  gestion_descargas → Historial con límite diario por usuario
 ┗ 📊 reportes           → Estadísticas y resúmenes del sistema
```

---

## 🏗️ Estructura del Proyecto

```
ebooks-manager/
├── main.go                  # Punto de entrada y menú principal
├── go.mod                   # Configuración del módulo Go
├── README.md
├── data/
│   ├── libros.json          # Catálogo persistente de libros
│   ├── usuarios.json        # Usuarios registrados
│   └── descargas.json       # Historial de descargas
├── gestion_libros/
│   └── libros.go
├── gestion_usuarios/
│   └── usuarios.go
├── gestion_descargas/
│   └── descargas.go
└── reportes/
    └── reportes.go
```

---

## ⚙️ Requisitos

- Go 1.21 o superior → [descargar Go](https://go.dev/dl/)
- Sin dependencias externas (solo librería estándar de Go)

---

## 🚀 Instalación y Uso

```bash
# 1. Clonar el repositorio
git clone https://github.com/Jean995-tech/ebooks-manager.git

# 2. Entrar al directorio
cd ebooks-manager

# 3. Ejecutar el sistema
go run main.go
```

---

## ✅ Funcionalidades

| Módulo | Funciones disponibles |
|---|---|
| 📖 Libros | Agregar, buscar, listar, eliminar |
| 👤 Usuarios | Registrar, buscar, listar |
| ⬇️ Descargas | Registrar descarga, ver historial, verificar límite diario |
| 📊 Reportes | Libros más descargados, usuarios activos, resumen general |

---

## 📦 Paquetes Utilizados

| Paquete | Uso |
|---|---|
| `fmt` | Entrada/salida en consola |
| `bufio` | Lectura de input del usuario |
| `os` | Interacción con archivos del sistema |
| `strings` | Manipulación y búsqueda de cadenas |
| `strconv` | Conversión entre tipos de datos |
| `time` | Registro de fecha y hora en descargas |
| `encoding/json` | Persistencia de datos en archivos JSON |

> Todos los paquetes son de la **librería estándar de Go**. No se requiere `go get` ni instalación adicional.

---

## 🛡️ Alcance del Proyecto (Etapa 1)

**Incluido:**
- ✅ CRUD completo del catálogo de libros
- ✅ Registro y consulta de usuarios
- ✅ Control de descargas con límite diario (5 por usuario)
- ✅ Persistencia en archivos JSON
- ✅ Reportes básicos del sistema
- ✅ Interfaz de consola con menú interactivo

**Fuera del alcance:**
- ❌ Autenticación con contraseña
- ❌ Interfaz gráfica o web
- ❌ Base de datos relacional
- ❌ Sistema de pagos

---

<div align="center">

**Jean Pierre Males Cedeño**  
Ingeniería en Ciberseguridad · UIDE · 2026

</div>
