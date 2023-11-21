1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?
    Los archivos con el verbo Create tienen instrucciones DDL para crear la respectiva tabla en la base de datos, y los archivos con el verbo Add tienen instrucciones DML para agregar registros a la respectiva tabla.

2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? 
     se genera un error porque no se encuentra la dependencia en los servicios donde se está utilizando.
       "ERROR: Service 'movies-api' depends on service 'postgres' which is undefined."
       "ERROR: Service 'flyway' depends on service 'postgres' which is undefined." 
        
    ¿Qué otros cambios tendrías que hacer?
     modificar el valor de la dependencia en todos los servicios donde se estaba haciendo referencia a "postgres" cambiandolo por "db"
       depends_on:
         - db
     Para este caso se debe modificar en los servicios movies-api y flyway
     tambien se debe modificar el archivo .env 
       POSTGRES_SERVER=db

3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer?
    Modificar el archivo .env
      BIND_PORT=81
      REACT_APP_API_URI=http://localhost:81
    Modificar el archivo docker-compose.yml cambiando el ports en servicio movies-api
         ports:
          - "${BIND_PORT}:${BIND_PORT}"

4.- ¿Qué pasa si a la variable de ambiente BIND_IP le asignas el valor localhost?
     El front end no trae datos porque no encuentra la api.

5.- Revisa el archivo Dockerfile en la carpeta movies-api. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.
     
    # especifica la imagen base con la que será construida la docker image en este caso la verion 20.5 de nodeJS basada en la distribucion Alpine Linux 3.17
    FROM node:20.5-alpine3.17

    #establece /usr/app como directorio de trabajo en el cual van a ser ejecutados los comandos 
    WORKDIR /usr/app
    #copia el archivo package.json del directorio actual al directorio /usr/app
    COPY package.json .

    # Instala npm en el contenedor
    RUN npm install --quiet

    # copia todos los archivos y directorios del directorio actual al directorio /usr/app del contenedor
    COPY . ..
---------------------------
TAREA 2 - FRANKLIN CRUCES
---------------------------
