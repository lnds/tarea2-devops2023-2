# Respuestas
## Nombre Alumno: Sebastián Urbano

### P1: Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?

Los archivos con el verbo Create permiten crear tablas, mientras que los verbos Add permiten agregar nuevas filas o registros a la tabla de la base de datos.

### P2: ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?

Al realizar el cambio del nombre del servicio de postgres a db, se tendría que cambiar en el depends_on de flyway además de cambiarlo en el .env en el POSTGRES_SERVER, dado que ahora deberían tomar el valor de db.

### P3: Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer? 

Si queremos que el servicio moves-api ocupe el puerto 81, lo que tendríamos que hacer sería lo siguiente:
- Cambiar la variables de entorno llamada `BIND_PORT` a 81, lo cual cambiaría el port de la api.
- Cambiar la url de conexión que vamos a ocupar en el frontend, que sería cambiar la variable `REACT_APP_API_URI` a `http://localhost:8000`. Lo otro que se podría haber realizado aquí para evitar realizar este paso, es haber configurado esta variable en dos partes, una que entregue localhost y la otra que sea la parte del puerto de la api, donde pudieramos reemplazar por la variable que cambiamos en el paso anterior `BIND_PORT`.

### P4: ¿Qué pasa si a la variable de ambiente `BIND_IP` le asignas el valor localhost?

Si dejaramos la variable `BIND_IP` como localhost, no podríamos acceder a la API desde fuera del contenedor, dado que recibiría sólo tráfico dentro del contenedor y no se podría acceder desde alguna aplicación o navegador fuera del contenedor.

### P5: Revisa el archivo `Dockerfile` en la carpeta `movies-api`. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.

Revisando el contenedor, puedo ver que se tienen varias etapas dentro del contenedor, en este caso 2, una que se llama `build-stage`, en la cual se construye la imagen, para después pasar a la segunda etapa que es `build-release-stage`, en la cual se levanta una imagen más pequeña para que se ejecute. Lo otro interesante es con respecto al comando `USER nonroot:nonroot`, lo que permite ejecutar el contenedor como `nonroot`, de manera que no se ejecute la aplicación como `root` en el contenedor.