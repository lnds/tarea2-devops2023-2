Sebastian Salgado - Daniel Hinojosa

1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?
## R: los archivos CREATE son para poder crear tablas con sus respectivos campos, mientras que los archivos ADD son para poblar dichas tablas y campos.

2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?
## R: Falla al momento de levantar los servicio, dado que estan apuntando al servicio con el nombre postgres.Para solucionarlo se debe cambiar el nombre postgres en los depends donde esta declarado, por el nombre db.

3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer? 

## R: Se deberia cambiar el puerto dentro del archivo .env. Si esta siendo utilizado el puerto se debe utilizar otro.

4.- ¿Qué pasa si a la variable de ambiente `BIND_IP` le asignas el valor localhost?
## R: Si se cambia la variable mencionada, no se lograria comunicar el front con el backend, actualmente permite poder reponder peticiones desde cualquier ip.

5.- Revisa el archivo `Dockerfile` en la carpeta `movies-api`. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este 
## R: Lo que ocurre es que lo primero que esta haciendo es descargar las dependencias de Golang para que al construir la imagen no de error. Tambien esta utilizando USER nonroot:nonroot, practica para evitar vulnerabilidades .