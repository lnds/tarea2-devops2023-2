Preguntas

1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?
    r1- Los dos archvivos con la instrucciones Crate son los encargados de crear las tablas Directors y Movies respectivamente.
    r2- Los dos archivos con las instrucciones Add son los encargados de agregar registros a las respectivas tablas

2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?
    r3- Si se cambia el nombre del servicio de base de datos Postgres deja de funcionar la aplicación, porque los otros servicios (Movis-api y Flyway ) pierden la dependencia, ademas el Host tambien se pierde.
    r4- Para que funcione la aplicación se debe corregir la dependencia (depends on) y cambiar la variable de ambiente HOST (server)

3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer? 
    r5- se debe cambiar la variable BIND_PORT en .env

4.- ¿Qué pasa si a la variable de ambiente `BIND_IP` le asignas el valor localhost?
    r6- Significa que el servicio backend solo respondera a solictudes locales y no desde cualquier IP

5.- Revisa el archivo `Dockerfile` en la carpeta `movies-api`. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.
    r7- Se genera la construcción en dos etapas (se utilizan dos imagenes existentes); primera  etapa de construccion de la apliación backend en software Go  y la segunda  etapa despliegala imagen en un contenedor cloud google

Recuerda responder en el archivo `RESPUESTAS.md` y agregar tu nombre en ese archivo.