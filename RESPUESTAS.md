# Alumno Juan Barraza Zaso

# Preguntas

1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?

    - Los Archivos con el verbo Create son los que contienen las sentencias para crear las tablas, los archivos con el verbo Add contienen las dentencias para insertan datos en las tablas.

2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?

    - El servicio movie-api queda con una dependencia indefinida y no cargara,  Para que cargue el servicio desebe modificar el archivo de variables de ambiente .env variable POSTGRES_SERVER = con el nombre db, tambbien el el servicio movie-apis y flyway parametro depends_on:  

3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer? 
    
    - Se debe modificar el archivo .env variable BIND_PORT = con el valor 81

4.- ¿Qué pasa si a la variable de ambiente `BIND_IP` le asignas el valor localhost?

    - Levantan los servicios pero no se peude acceder al backend

5.- Revisa el archivo `Dockerfile` en la carpeta `movies-api`. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.

    - Archivo debidamente comentado, levanta el container con usuario diferente a root, indica el Sistema operativo que utilizara el servicio. 

# Evidencias, fotos

    - Las imagenes de ejecucion del aplicativo estan en la carpeta Evidencias 