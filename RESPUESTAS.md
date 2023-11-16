1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?
    los archivos con el verbo create crean las tablas y los archivos con add insertan data en esas tablas. una crea y el otro inserta esa es la difirencia.

2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?
    Tendría que ajustar todo los nodos "depends_on" que hagan referencia al nombre antiguo con el nuevo.

3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer? 
    cambiar la variable de entorno al puerto 81, en el archivo .env.

4.- ¿Qué pasa si a la variable de ambiente `BIND_IP` le asignas el valor localhost?
    no se logra conectar a la base de datos. da error connection refuse.

5.- Revisa el archivo `Dockerfile` en la carpeta `movies-api`. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.
    que consta de 2 etapas. despues de buscar información encontre que se hace Multi-stage builds para alivianar el archivo final, se hace el build con la imagen original de go y despues se monta en una más liviana. https://docs.docker.com/language/golang/build-images/#multi-stage-builds