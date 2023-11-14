# Respuestas

## Nombre: Helena Alvarado

1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?
- Los archivos con el verbo Create, crean las tablas directors y movies
- Los archivos con el verbo Add, insertan registros a las tablas directors y movies

2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?
- Si cambia el nombre del servicio de postgres a db, se debe cambiar el valor a "depends_on" en el servicio movies-api y en el servicio flyway.


3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer? 

 - Solamente asignar el valor 81 a la variable de entorno BIND_PORT; 
 - A la variable REACT_APP_API_URI se le asignó el valor = 
 "http://localhost:${BIND_PORT}"


4.- ¿Qué pasa si a la variable de ambiente `BIND_IP` le asignas el valor localhost?
- El contenedor movies-api no seria accesible desde fuera del contenedor.


5.- Revisa el archivo `Dockerfile` en la carpeta `movies-api`. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.
- primero se construyó basado en buenas prácticas, comentando adecuadamente.
- en la imagen build-release-stage se genera a partir de build-stage
- Que comparado con el archivo Dockerfile de movies-front este último utiliza la instrucción COPY . .. y Dockerfile de movies-api utiliza COPY ./ ./  

6. - Las imagenes de evidencia se encuentran alojadas en el directorio raíz