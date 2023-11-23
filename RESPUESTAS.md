## Alumno
Alejandro Montoya

## 1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?
R: los archivos con el verbo create, contienen query para crear tablas en la base de datos y los archivos con el verbo add tienes querys que insertan datos en tablas.

## 2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?
R: al cambiar el nombre del servicio, debemos cambiar el nombre en todos los parametros que esten usando el servicio por nombre, por ejemplo conexiones a la base de datos.
en este ejercicio se debe actualizar la variable de entorno POSTGRES_SERVER, las dependencias al servicio (movies-api y flyway) y la aplicacion funcionara correctamente.

## 3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer?
R: Se debe actualizar la variable de entorno BIND_PORT con el nuevo puerto y tambien la url en la variable REACT_APP_API_URI para que utilice el nuevo puerto (http://localhost:81).

## 4.- ¿Qué pasa si a la variable de ambiente BIND_IP le asignas el valor localhost?
R: En este caso al modificar la variable deja de funcionar el front, ya que no todo esta en el mismo container. Al modificar la variable como localhost esta solo permitiendo lo interno (rechaza las conexiones externas).

## 5.- Revisa el archivo Dockerfile en la carpeta movies-api. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.
R: Me llama la atencion lo simple y preciso que esta realiza la creacion, compilacion y despliegue de la api. En un par de lineas esta todo lo necesario (tambien al final aplica una buena practica de seguridad).
