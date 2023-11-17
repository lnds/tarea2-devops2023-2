# Rodrigo Antonio Sáez Zuñiga

# Preguntas

1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?

Los archivos create son para crear las tablas con las estructuras definidas dentro de cada archivo create, mientras que los archivos add son para poblar con los datos dentro de los archivos las tablas creadas.

2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?

Si solo se cambia el nombre del servicio, tendremos errores ya que hay dependencias de otros servicios para que se ejecuten como lo son el movies-api y flyway. Otro cambio que debemos hacer para que se ejecute de forma correcta es modificar la variable POSTGRES_SERVER para que el backend movies-api se pueda conectar a la base de datos

3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer?

Para modificar el puerto a 81 implica cambiar el puerto de salida en el servicio movies-api y cambiar en el .env la variable BIND_PORT para que siga funcionando.

4.- ¿Qué pasa si a la variable de ambiente `BIND_IP` le asignas el valor localhost?

Si bien la aplicación no genera errores en su creación y se puede levantar, el frontend no se logra conectar con el backend de forma correcta por lo que no retorna la información necesaria de la base de datos.

5.- Revisa el archivo `Dockerfile` en la carpeta `movies-api`. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.

Dentro del dockerfile se esta haciendo multi-stage-build desde una imagen con golang a una con base debian, por lo general esto se utiliza para crear imagenes mas livianas y dejar solo los componentes que necesito de cada una.
También se puede ver que se deja un usuario final que no sea root para que tenga permisos limitados dentro del contenedor.