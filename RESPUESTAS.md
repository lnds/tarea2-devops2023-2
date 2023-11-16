# Preguntas

# 1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?

## Los archivos create, dan instrucciones para crear nuevas tablas, describen sus columnas y características, mientras que los archivos add, dan instrucciones de agregar elementos (filas) a las tablas. 

# 2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?

## Si solo se cambia el nombre de postgres a db, la aplicación no puede acceder a la base de datos, el despliegue de directores y movies se ve en blanco. 
## Los otros elementos que se deberían cambiar están en el docker-compose.yml, y corresponden al nombre del contenedor postgres => db, a la dependencia de los contenedores flyway y movies-api, que también deben cambiar de postgres a db.

# 3.- Si quisiéramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer? 

## Para que la aplicacion funcine en su conjunto es necesario modificar: 
## - Puertos en movies-api (81:81)
## - REACT_APP_API_URI='http://localhost:81' en .env
## - BIND_PORT=81 en .env

# 4.- ¿Qué pasa si a la variable de ambiente `BIND_IP` le asignas el valor localhost?

## 

# 5.- Revisa el archivo `Dockerfile` en la carpeta `movies-api`. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.

## 

