Nombre: Jimena Ester Silva Navarrete

la imagen con el nombre "up-app-movies.png", es un pantallazo que evidencia que logré levantar la app de movies



Preguntas:


1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?

Resp 1:

Los archivos con el verbo Create son los responsables de crear las tablas directors y movies en la base datos.
Los archivos con el verbo Add, en cambio, son los responsables de poblar con datos las tablas, directors y movies, creadas en los archivos con el verbo Create.





2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?

Resp 2:

Al cambiar el nombre del servicio de postgres a db, conlleva también realizar cambios en:
1.- cambiar el nombre de postgres a db en la dependencia de movie-api
2.- cambiar el nombre de postgres a db en la dependecia de flyway
3.- cambiar el valor de la variable POSTGRES_DB, en el archivo .env, de postgres a db




3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer?

Resp 3:
Al cambiar el servicio movies-api para que use el puerto 81, se deben realizar los siguientes cambios en:
1.- modificar el valor de la variable BIND_PORT, en el archivo .env, de 8000 a 81
2.- modificar el valor de la variable REACT_APP_API_URI, en el archivo .env, de http://localhost:8000 a http://localhost:81




4.- ¿Qué pasa si a la variable de ambiente BIND_IP le asignas el valor localhost?

Resp 4:

La variable de entorno BIND_IP se utiliza para especificar la dirección IP a la que un programa o servicio se vincula o escucha.

Cuando se establece el bind IP en 0.0.0.0, significa que el servidor está configurado para escuchar en todas las interfaces de red disponibles en el sistema. En términos simples, el servidor está abierto para aceptar conexiones entrantes desde cualquier dirección IP disponible en la máquina.

El servidor web configurado con el bind IP en 0.0.0.0, escuchando y aceptando conexiones tanto desde la red local (127.0.0.1 o localhost) como desde cualquier otra red externa a la que esté conectado el servidor.

Por tanto, al asignarle el valor de localhost a bind IP, el front no logra conectarse con movies-api, y no responde el back con las solicitudes del front. Debido a que en Docker, los contenedores pueden estar en diferentes redes y, por defecto, se ejecutan aislados unos de otros. Cada contenedor tiene su propia interfaz de red y su propia dirección IP dentro de la red a la que está conectado.

Si estableces la variable de entorno BIND_IP de un contenedor a localhost, este contenedor solo escuchará conexiones que se originen dentro del mismo contenedor. En un entorno Docker donde los contenedores están en diferentes redes, si un contenedor A tiene su BIND_IP configurado como localhost, no podrá recibir conexiones de otros contenedores ya que localhost se refiere al contenedor en sí mismo y no a los otros contenedores de Docker en diferentes redes.





5.- Revisa el archivo Dockerfile en la carpeta movies-api. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.

Resp 5:

Me llama la atención el segundo fragmento del archivo dockerFile, el que corresponde a "Deploy the application binary into a lean image":

FROM gcr.io/distroless/base-debian11 AS build-release-stage: Esta línea indica que se está utilizando la imagen gcr.io/distroless/base-debian11 como base para crear una nueva imagen. Además, se le da un alias o nombre de etapa llamado build-release-stage. Esta imagen está configurada para ser una imagen mínima, sin muchos paquetes o herramientas adicionales, optimizada para ejecutar aplicaciones de manera segura.

WORKDIR /: Esta instrucción establece el directorio de trabajo dentro del contenedor en la raíz del sistema de archivos.

COPY --from=build-stage /movies-api /movies-api: Esta línea copia los archivos desde otra etapa del proceso de construcción (identificada por build-stage) al directorio /movies-api en esta nueva etapa (build-release-stage). Parece que se está copiando un binario o un conjunto de archivos relacionados con la aplicación movies-api desde una etapa previa denominada build-stage.

USER nonroot:nonroot: Cambia el usuario y el grupo del contenedor al usuario nonroot. Esto es una práctica de seguridad común en la que se ejecuta la aplicación como un usuario no privilegiado dentro del contenedor, lo que ayuda a mitigar riesgos de seguridad si la aplicación es comprometida.

En resumen, este fragmento de Dockerfile está creando una nueva imagen basada en gcr.io/distroless/base-debian11, copiando un binario o conjunto de archivos de una etapa anterior y configurando el contenedor para ejecutar la aplicación movies-api como un usuario no privilegiado. Esta estrategia de múltiples etapas es común en Dockerfile para construir imágenes más pequeñas y seguras.