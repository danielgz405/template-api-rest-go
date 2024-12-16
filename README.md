# 🌐 Go Backend Template with REST APIs and WebSockets

![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

Bienvenido a **Go Backend Template**, una plantilla diseñada para desarrolladores que buscan construir aplicaciones backend modernas con **APIs RESTful** y soporte para **WebSockets bidireccionales**. 🚀 Este proyecto está pensado para ser un punto de partida flexible y adaptable para tus propios proyectos.

---

## ✨ Características

- **Plantilla Genérica**: Diseñada para ajustarse a diferentes tipos de proyectos backend.
- **RESTful APIs**: Incluye ejemplos básicos de rutas CRUD.
- **WebSockets Bidireccionales**: Implementación base con códigos de eventos que puedes personalizar.
- **MongoDB Integrado**: Soporte nativo para bases de datos MongoDB.
- **Código Limpio y Escalable**: Utiliza buenas prácticas y un diseño modular.

---

## 📚 Tabla de Contenidos

1. [Introducción](#-introducción)
2. [Requisitos](#-requisitos)
3. [Estructura del Proyecto](#-estructura-del-proyecto)
4. [Configuración](#-configuración)
9. [Licencia](#-licencia)

---

## 💡 Introducción

Esta plantilla sirve como base para proyectos backend con **Go** que necesiten tanto APIs REST como WebSockets. Los códigos de eventos para WebSockets son **placeholders**, lo que significa que puedes reemplazarlos según las necesidades de tu proyecto. 

Al ser una plantilla genérica, se adapta a diferentes casos de uso, como:

- Servicios de mensajería.
- Sistemas de notificaciones en tiempo real.
- Aplicaciones CRUD robustas.

---

## 🔧 Requisitos

Antes de empezar, asegúrate de tener lo siguiente instalado:

- [Go 1.20+](https://go.dev/dl/)
- [MongoDB](https://www.mongodb.com/)
- [Git](https://git-scm.com/)
- **(Opcional)** Docker para ejecución y despliegue.

---

## 📂 Estructura del Proyecto

La estructura de este proyecto está diseñada para facilitar la organización y escalabilidad:

go-backend-template/

├── Database/         # Conexión y configuración de MongoDB

├── Handles/          # Controladores HTTP para rutas REST

├── Middleware/       # Middlewares como autenticación y logs

├── Modules/          # Componentes y módulos reutilizables

├── Pages/            # Plantillas HTML (si las hay)

├── Repository/       # Interacción directa con MongoDB

├── Responses/        # Estructuras para respuestas JSON

├── Server/           # Configuración y ejecución del servidor

├── Estructures/      # Modelos de datos y estructuras compartidas

├── Websockets/       # Implementación y manejo de WebSockets

└── main.go           # Punto de entrada principal

---

## ⚙️ Configuración

1. **Clona el repositorio**:
   ```bash
   git clone https://github.com/tu-usuario/go-backend-template.git
   cd go-backend-template
2. **Instala las dependencias**:
   ```bash
   go mod tidy
3. **Configura tus variables de entorno en un archivo .env**:
   ```env
    PORT=5050
    JWT_SECRET=owo
    DB_URI=mongodb://localhost:27017/
    DB_URI_TEST=mongodb://localhost:27017/
    TESTING_MODE=true
4. **Ejecuta el servidor**:
   ```bash
   go run main.go
   ```

## 📄 Licencia
Este proyecto está licenciado bajo la MIT License. ¡Siéntete libre de usar esta plantilla en tus proyectos!