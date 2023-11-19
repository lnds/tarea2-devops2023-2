# Respuesta de Cristian Luttgue


# Preguntas

1.- Revisa el contenido del diretorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?

La diferencia es que uno de los verbos realiza inserción en la base de datos (Add), a diferencia de create, el cual crea las tablas para realizar la inserción de datos.

2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?

Falla al momento de levantar nuevamente, debido a que postgres está siendo apuntado en las dependencias de los otros servicios. Para solucionar el problema, es necesario cambiar en los depends donde estaba declarado postgres por db.

3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer? 

Debería cambiar el puerto en el .env a 81. Con esto ya sería posible visualizar la api en el puerto 81, siempre y cuando no se encuentre ocupado el puerto de la máquina host.

4.- ¿Qué pasa si a la variable de ambiente `BIND_IP` le asignas el valor localhost?

El front no logra conectar con el backend, esto se debe a que el 0.0.0.0 permitia responder peticiones desde cualquier ip.

5.- Revisa el archivo `Dockerfile` en la carpeta `movies-api`. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.

Me llama la atención que primero se realiza la operación de descarga de Golang antes que la del sistema operativo desde un registry de Google, esto se puede deber a que este lenguaje fue desarrollado por Google y está publicado con todo lo necesario para levantar en el container en forma sencilla.

