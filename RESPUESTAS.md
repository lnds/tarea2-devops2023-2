# Ariel Veliz Guerrero
1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?

  Los archivos con verbo create corresponde a la creacion de tablas con sus respectivos atributos, y los verbos Add corresponden a la insercion de registros en las tablas correspondientes.

2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?

  Si se cambia el nombre del servicio, debemos verificar y cambiar su nombre tambien en todos los demas servicios que dependan de el, ademas modificar la variable de entorno POSGREST_DB para poder conectarnos a la base de datos.

3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer?

  Primero deberiamos cambiar nuestra variable de entorno BIND_PORT con el puerto de salida que queramos utilizar, ademas verificar en nuestro servicio decladado en el docker-compose para que los puertos esten asociados a puerto correspondiente, finalmente verificar en nuestra api que el puerto tambien este seteado con nuestra variable o coincida con el puerto que estamos especificando (81). ya que en el archivo main.go el puerto se encuentra en duro ``var port string = "8080"``

4.- ¿Qué pasa si a la variable de ambiente BIND_IP le asignas el valor localhost?

  Al levantar el contenedor, en la consola se muestra que que no se puede conectar a la API, imagino por que localhost no es aceptado como una IP valida. 

5.- Revisa el archivo Dockerfile en la carpeta movies-api. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.
  
  Podemos identificar que la api esta construida en base a una imagen de goland con su respectivo tag, ademas se puede ver que contiene dos pasos, el cual genera el deploy en base a la copia al buil-stage.

Recuerda responder en el archivo RESPUESTAS.md y agregar tu nombre en ese archivo.