nombre completo : Francisca Catalina Valdés Barrera

1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?

    Los archivos que contienen el verbo "Create" son los encargados de crear las tablas , mientras que aquellos archivos que usan el verbo "Add" se refieren a la insercion de datos en las tablas que se han creado.

2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? 

     sucede lo siguiente :

     C:\Users\fran9\Desktop\tarea2-devops2023-2>docker compose down
     service "movies-api" depends on undefined service postgres: invalid compose project

    C:\Users\fran9\Desktop\tarea2-devops2023-2>docker-compose up -d --build
    service "flyway" depends on undefined service postgres: invalid compose project

    -El mensaje anterior nos indica que en el archivo docker-compose.yml, el servicio llamado "flyway" tiene una dependencia declarada hacia otro servicio llamado "postgres",
     pero el servicio "postgres" no está definido.

    -movies-api tambien depende del  servicio "postgres"
 
    ¿Qué otros cambios tendrías que hacer?

     Modificar las siguientes secciones del archivo docker-compose.yml
  movies-api:
    depends_on:
      - db

  flyway:
    image: flyway/flyway:latest
    container_name: migration_flyway
    volumes:
      - ./sql_migrations:/flyway/sql
    command: ["-url=jdbc:postgresql://${POSTGRES_SERVER}:${POSTGRES_PORT}/${POSTGRES_DB}", "-user=${POSTGRES_USER}", "-password=${POSTGRES_PASSWORD}", "migrate"]
    depends_on:
      - db

  y por último modificar el archivo .env

   POSTGRES_SERVER=db

3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer?

Para hacer que movies-api  use el puerto 81, debemos realizar los siguientes cambios:

- cambiar el valor de la variable BIND_PORT, en el archivo .env de   8000 a 81
- cambiar el valor de la variable REACT_APP_API_URI, en el archivo .env de http://localhost:8000 a http://localhost:81

y correr los siguientes comandos 

C:\Users\fran9\Desktop\tarea2-devops2023-2>docker compose down
C:\Users\fran9\Desktop\tarea2-devops2023-2>docker-compose up -d --build

4.- ¿Qué pasa si a la variable de ambiente BIND_IP le asignas el valor localhost?
    Modifique el valor en el archivo .env de esta forma 
    BIND_IP=localhost
    ejecute los comandos 
    C:\Users\fran9\Desktop\tarea2-devops2023-2>docker compose down
    C:\Users\fran9\Desktop\tarea2-devops2023-2>docker-compose up -d --build
   
   movies_front levanta pero ya no se conecta con la api , esto por que al cambiar el BIND_IP  a localhost deja a la api solo accesible dentro del propio contenedor
   No será posible acceder al servicio desde afuera como esta intentando la aplicacion movies api
  

5.- Revisa el archivo Dockerfile en la carpeta movies-api. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.

   Me llamo la atención el uso de (USER nonroot:nonroot) es una práctica de seguridad recomendada para reducir el impacto potencial de posibles vulnerabilidades.
   la copia de archivos se realiza de manera selectiva y en dos fases. Primero, se copian los archivos go.mod y go.sum para descargar las dependencias y aprovechar la caché de Docker. 
   luego, se copian todos los archivos restantes. Esta estrategia optimiza la construcción al minimizar el impacto en las capas de caché.


se adjuntan las siguientes imagenes 
evidencia-movies-front.PNG
evidencia-directors.PNG

