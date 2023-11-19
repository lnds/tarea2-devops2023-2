# Preguntas y Respuestas:
## Matías Mayol

# 1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?

## Los archivos create, dan instrucciones para crear nuevas tablas, describen sus columnas y características, mientras que los archivos add, dan instrucciones de agregar elementos (filas) a las tablas. 

# 2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?

## Si solo se cambia el nombre de postgres a db, la aplicación no puede acceder a la base de datos, el despliegue de directores y películas se ve en blanco. 
## Los otros elementos que se deberían cambiar están en el docker-compose.yml, y corresponden al nombre del contenedor postgres => db, a la dependencia de los contenedores flyway y movies-api, que también deben cambiar de postgres a db.

# 3.- Si quisiéramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer? 

## Para que la aplicación funcione en su conjunto es necesario modificar: 
## - Puertos en movies-api (81:81)
## - REACT_APP_API_URI='http://localhost:81' en .env
## - BIND_PORT=81 en .env

# 4.- ¿Qué pasa si a la variable de ambiente `BIND_IP` le asignas el valor localhost?

## Se ve el front de la aplicación, pero no se observan las peliculas, por lo que aparentemente no se esta teniendo acceso a la base de datos. 

# 5.- Revisa el archivo `Dockerfile` en la carpeta `movies-api`. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.

## Me llama la atención el uso de dos imágenes, una para construir la aplicación y otra para montarla. 

## A continuación se explica que ocurre en cada linea de código. 

### Linea 1:
### FROM golang:1.19 AS build-stage
### Esta linea define la imagen base y su versión para construir el ambiente

### Linea 2:
### WORKDIR /app
### Define el directorio dentro del contenedor a /app, para trabajar dentro de este. 

### Linea 3:
### COPY go.mod go.sum ./
### copia go.mod y go.sum a /app.

### Linea 4:
### RUN go mod download
### descarga las dependencias del modulo Go basado en los elementos del paso anterior. 

### Linea 5:
### COPY ./ ./
### copia el resto de la aplicación a /app.

### Linea 6:
### RUN CGO_ENABLED=0 GOOS=linux go build -o /movies-api
### Construye la aplicación Go. CGO_ENABLED=0 se asegura de que el binario este enlazado, independiente del ambiente anfitrión. GOOS=linux define el sistema operativo. El binario resultante es movies-api y esta en el directorio raíz (root, /) del contenedor.

### Linea 8:
### FROM gcr.io/distroless/base-debian11 AS build-release-stage
### define la imagen base para montar la aplicación. 

### Linea 9:
### WORKDIR /:
### Fija el directorio de trabajo a / (raíz, root).

### Linea 10:
### COPY --from=build-stage /movies-api /movies-api:
### Copia el binario de movies-api desde la etapa de construcción a la raíz (/, root, del directorio de trabajo actual).

### Linea 11:
### USER nonroot:nonroot:
### Cambia el usuario a no raíz, esto es una medida de seguridad para quitarle privilegios al usuario de la aplicación, reduciendo vulnerabilidades. 



























