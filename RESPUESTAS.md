Roberto Catalán H.

1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?
  Los archivos con el verbo Create se encargan de crear las tablas necesarias para que la aplicación funcione.
  Los archivos con el verbo Add agregan datos a las tablas creadas con los Create.

2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?
  Si se va a cambiar el nombre al servicio, se deberá actualizar el nombre en las dependencias del archivo .yaml y también se deberá cambiar en las variables de entorno. 

3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer?
  Se tendría que cambiar la variable de ambiente BIND_PORT del archivo .env.

4.- ¿Qué pasa si a la variable de ambiente BIND_IP le asignas el valor localhost?
  Al dejar la variable BIND_PORT con localhost, la API no responde y la aplicación no se logra conectar a la base de datos, por lo tanto, la página no muestra datos.

5.- Revisa el archivo Dockerfile en la carpeta movies-api. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.
  Me llama la atención que se realizan 2 tareas principales, con la ejecución de 2 comandos RUN. También me llama la atención que la segunda tarea usa lo construido en la primera tarea, llamada build-stage.