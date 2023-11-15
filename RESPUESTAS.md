1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?
R. Las sentencias con el verbo CREATE son para agregar objetos a la base de datos (en el caso de los scripts, tablas). Las sentencias por el verbo ADD son para agregar data a las tablas.

2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?
R. Cambiar el valor de la variable de entorno POSTGRES_SERVER a db ya que nos conectamos al alias del cotenedor el cual está dado por el nombre del contenedor.

3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer? 
R. Cambiar la variable BIND_PORT a 81 y el puerto expuesto a 81.

4.- ¿Qué pasa si a la variable de ambiente `BIND_IP` le asignas el valor localhost?
R. Perdemos comunicación con el API ya que BIND_IP pasaría a tener la ip local y no la ip del servidor.

5.- Revisa el archivo `Dockerfile` en la carpeta `movies-api`. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.
R. Llama la atención que, a diferencia del dockerfile de movies-front, tiene 2 secciones. La primera para compilar la app y la segunda para publicarla.