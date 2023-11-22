#Respuestas.
#Nombre: Robert Alexis Valenzuela Murillo


1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?
	Los archivos create, crean tablas. Para este caso, crea directors y movies.
	Los archivos Add, insertan un valor dentro de las tablas mencionadas. En este caso, insertan valores de directores y de películas.

2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?
	Probablemente falle, debido a que hay una variable de entorno seteada en el archivo .env con el nombre postgres y también está seteada en las dependencias de los otros servicios declarados (flyway y movies-api).
	Para que tenga éxito, tocaría modificar la variable de entorno a db y la dependencias al mismo nombre.
	
3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer?
	Tendríamos que modificar la variable de entorno BIND_PORT=8000 a BIND_PORT=81

4.- ¿Qué pasa si a la variable de ambiente BIND_IP le asignas el valor localhost?
Si asignas el valor localhost a la variable de entorno BIND_IP, esto limitará la aplicación a escuchar solo en la interfaz de red local de la máquina donde se está ejecutando el contenedor. En otras palabras, la aplicación solo será accesible desde la misma máquina en la que se está ejecutando, y no desde otras máquinas en la red.
	
5.- Revisa el archivo Dockerfile en la carpeta movies-api. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.
	Me llama la atención que hay dos etapas en el archivo, uno de construcción, donde se utiliza la imagen de golang y luego hay una etapa de implementación donde se utiliza la imagen gcr.io/distroless/base-debian11 para el despliegue de la aplicación. También se establece el usuario y grupo "nonroot:nonroot" para ejecutar la app como usuario no privilegiado

#Recuerda responder en el archivo RESPUESTAS.md y agregar tu nombre en ese archivo.