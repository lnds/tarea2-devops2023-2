1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?
RESPUESTA: La diferencia es la siguiente:
           - Los arichivos con verbo Create son los responsables de crear las tablas en cambio,
           - Los archivos con el verbo Add son los encargados de insertar registros a las tablas creadas.     

2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db?
RESPUESTA: Los servicios movies-api y flyway pierden la dependencia con el Servicio postgres y como resultado la aplicación no leerá la Base de Datos.

¿Qué otros cambios tendrías que hacer?
RESPUESTA: Al reemplazar el nombre se deben realizar los siguientes cambios:
           - Reemplazar la sección depends_on del servcio movies-api.
           - Reemplazar la sección depends_on del servcio flyway.

3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer?
RESPUESTA: Modificar la variable de ambiente BIND_PORT en el archivo .env

4.- ¿Qué pasa si a la variable de ambiente BIND_IP le asignas el valor localhost?
RESPUESTA: El contenedor no podría entregar respuesta al exterior.

5.- Revisa el archivo Dockerfile en la carpeta movies-api. 
¿Qué te llama la atención? 
RESPUESTA: El uso de buenas prácticas en la construcción de imágenes de contenedores y el empleo de imágenes  ligeras y seguras para reducir el tamaño del contenedor y mejorar la seguridad.

Trata de explicar lo que ocurre en este caso.
RESPUESTA: Este archivo se utiliza para construir y empaquetar una aplicación escrita en Go en un contenedor.

El primer bloque de codigo "FROM golang:1.19 AS build-stage":
- Utiliza la imagen oficial de Go 1.19 como base para la construcción de la aplicación. 
- Luego, establece el directorio de trabajo en "/app".
- Se copian los archivos go.mod y go.sum al directorio de trabajo.
- Se ejecuta go mod download para descargar las dependencias del módulo Go.
- Se copian todos los demás archivos del contexto de construcción al directorio de trabajo.
- Se ejecuta CGO_ENABLED=0 GOOS=linux go build -o /movies-api para compilar la aplicación Go. 

El segundo bloque de codigo "FROM gcr.io/distroless/base-debian11 AS build-release-stage"
- Utiliza una imagen basada en Debian 11. 
- Establece el directorio de trabajo en la raíz ("/").
- Copia el ejecutable compilado desde la etapa de construcción (--from=build-stage) al directorio raíz del contenedor (/movies-api).
