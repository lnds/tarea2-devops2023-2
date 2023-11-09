Nombre: Sebastian Ignacio Jara Castro

La evidencia solicitada en el tercer punto de la seccion 'Actividades' del enunciado se encuentra en el archivo 'evidencia.jpg'.

1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?

Los primeros crean tablas en la base de datos y los segundos añaden registros a las tablas


2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?

Ese cambio de nombre del servicio provocaría que deberia actualizarse "depends_on" en el servicio flyway y la variable de entorno POSTGRES_SERVER con el valor del nuevo nombre 'bd'


3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer?

- Cambiar el puerto en  REACT_APP_API_URI, quedando "http://localhost:81"
- Configurar variable de entorno BIND_PORT=81
- Cambiar ports de servicio "movies-api" a "81:81"

4.- ¿Qué pasa si a la variable de ambiente BIND_IP le asignas el valor localhost?

La api no seria accesible desde fuera del contenedor.

5.- Revisa el archivo Dockerfile en la carpeta movies-api. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.

Me llama la atención que se muestran dos etapas diferenciadas, una donde se preparan las dependencias y se compila la aplicacion, y la segunda donde se arma una imagen que contiene la aplicacion usando una imagen de poco tamaño como las del tipo "distroless". Otra cosa que me llama la atencion es la linea "USER nonroot:nonroot", que no conocia su uso y veo que ayuda a mejorar la seguridad.