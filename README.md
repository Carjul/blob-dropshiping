
## Configuración del Entorno

### Requisitos Previos

- Go 1.19 o superior
- Docker y Docker Compose
- MongoDB

### Configuración del Proyecto

1. **Clona el Repositorio:**

    ```sh
    git clone https://github.com/tuusuario/myproject.git
    cd myproject
    ```

2. **Instala las Dependencias:**

    ```sh
    go mod tidy
    ```

3. **Configura las Variables de Entorno:**

    Crea un archivo `.env` en el directorio raíz del proyecto y añade las siguientes variables:

    ```sh
    MONGO_URI=mongodb://localhost:27017
    MONGO_DB=mydatabase
    ```

4. **Construye y Ejecuta la Aplicación:**

    ```sh
    go run cmd/myproject/main.go
    ```

### Uso de Docker

1. **Construye la Imagen Docker:**

    ```sh
    docker build -t myproject .
    ```

2. **Inicia los Servicios con Docker Compose:**

    ```sh
    docker-compose up -d
    ```

## Endpoints de la API

### Crear un Blob

- **URL:** `/api/blobs`
- **Método:** `POST`
- **Parámetros:**
  - `title` (string): El título del blob.
  - `content` (string): El contenido del blob.

- **Ejemplo de Solicitud:**

    ```sh
    curl -X POST http://localhost:8080/api/blobs \
    -H "Content-Type: application/json" \
    -d '{"title": "Mi primer blob", "content": "Este es el contenido de mi primer blob."}'
    ```

### Obtener Todos los Blobs

- **URL:** `/api/blobs`
- **Método:** `GET`

- **Ejemplo de Solicitud:**

    ```sh
    curl http://localhost:8080/api/blobs
    ```
