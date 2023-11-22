# RESPUESTAS

# Autor: Rodrigo Muñoz Soto

## 1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?

R: Los archivos que contienen Create son utilizados para sentenciar una tarea del tipo DDL, en este caso, definición/creación de tablas. Por otra parte, los arhivos que contienen Add son utilizados para sentenciar tareas del tipo DML, en este caso, para poblar (insert into) una tabla creada antereiormente.

## 2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?

R: Si se cambia únicamente el nombre del servicio postgres a db, al levantar la app esta no correrá, ya que dicho cambio de nombre de servicio también impacta a servicios que son dependientes de este, como movies-api y flyway. Para levantar y correr la app correctamente, también se debe realizar la modificación de nombre en los servicios declarados en depends_on de movies-api y flywau, y en el nombre de la variable de entorno POSTGRES_SERVER en el archivo de configuración .env

## 3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer?

R: Se tendría que hacer la modificación en el archivo de configuración .env para la variable BIND_PORT. Como los servicios que utilizan este puerto está configurado de forma parametrizada, no es necesario realizar cambios más allá del propio valor de la varible en el archivo .env

## 4.- ¿Qué pasa si a la variable de ambiente `BIND_IP` le asignas el valor localhost?

R: fallará la comunicación entre el servicio de front y la api

## 5.- Revisa el archivo `Dockerfile` en la carpeta `movies-api`. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.

R: Me llama la atención que utilice dos imagenes distintas para construir el container (golang y debian11), lo que nos muestra una gran flexibilidad a la hora de poder construir nuestros contenedores.
Además, me parece que el uso de un usuario nonroot es muy adecuado para evitar potenciales vulnerabilidades

## Recuerda responder en el archivo `RESPUESTAS.md` y agregar tu nombre en ese archivo.
