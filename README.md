# go JWT AUTH
## 🚀 Características
- Registro y Login de usuarios.

- Contraseñas encriptadas con bcrypt.

- Autenticación basada en JWT almacenado en Cookies (Seguridad mejorada).

- Configuración mediante variables de entorno (.env).

- Middleware de CORS configurado para desarrollo.

## 🛠️ Requisitos
- Go (versión 1.18 o superior)

- PostgreSQL

- Una herramienta para probar la API (Postman, Insomnia o REST Client de VS Code)

## ⚙️ Configuración

1. Clona el repositorio:

```bash
git clone https://github.com/D-aniela/go-jwt-auth.git
cd go-jwt-auth
````

2. Instala las dependencias:

```bash
go mod tidy
````

### 3. Configura las variables de entorno:
Crea un archivo .env en la raíz del proyecto con el contenido que está incluido de ejemplo en:

```.env.example```


## 🏃 Ejecución
Para iniciar el servidor:

```bash
go run main.go
````

### 📌 Endpoints

Método | Endpoint | Descripción | Requiere Cookie
--- | --- | --- | --- |
POST | /user/register | Crea un nuevo usuario | No 
POST | /user/login | Inicia sesión y recibe la cookie JWT | No |
POST | /user/profile | Obtiene los datos del usuario autenticado | Sí |
POST | /user/logout | Cierra la sesión y limpia la cookie | Sí |


## 🔒 Seguridad
Este proyecto utiliza cookies **HTTPOnly**. Esto significa que el token JWT no es accesible mediante JavaScript en el navegador, lo cual mitiga ataques de tipo XSS.

Para que las cookies funcionen en el frontend, asegúrate de que tu cliente (React, Vue, etc.) envíe la opción ```withCredentials: true``` en las peticiones.

## 🧪 Pruebas Rápidas (VS Code)
Si utilizas Visual Studio Code, puedes probar los endpoints sin necesidad de Postman utilizando la extensión [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client).

Abre el archivo ```routes.http```.

Haz clic en Send Request sobre cada endpoint.

Asegúrate de tener habilitado el guardado de cookies en VS Code para que el flujo de Login -> Perfil -> Logout funcione automáticamente:

Ve a ```Settings``` -> Busca ```rest-client.rememberCookies``` -> Actívalo.