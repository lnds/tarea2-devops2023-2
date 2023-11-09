# Respuestas
# Jaime Norambuena Sagredo

## 1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?
Los que tienen el verbo Create hacen referencia a creación de tablas y los que tienen el verbo Add son inserciones de datos.

## 2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?
Si lo cambio, debería cambiar la dependencia (depends_on) del servicio flyway y además la asignación de POSTGRES_SERVER en el archivo .env

## 3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer?
Habría que modificar el puerto en REACT_APP_API_URI y BIND_PORT en el archivo .env y el ports en el servicio movies-api.

## 4.- ¿Qué pasa si a la variable de ambiente BIND_IP le asignas el valor localhost?
Arroja error: request returned Bad Gateway for API route and version 

## 5.- Revisa el archivo Dockerfile en la carpeta movies-api. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.
En realidad casi todo es nuevo para mi, pero me llamó la atención el build-stage, por como copia y descarga las dependencias (COPY/RUN) y utilizar el USER nonroot:nonroot como buena práctica.