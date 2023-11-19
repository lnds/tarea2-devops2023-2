# Tarea 3 - Gerardo Alonso

Gracias al éxito de la aplicación de libros en Flask, les han encargado la construcción de una app web para películas.

Dado que la aplicación tendrá mayor demanda se han cambiado las tecnologías, el backend está escrito en go y se encuentra en la carpeta `movies-api`.

El frontend es una aplicación REACT en Javascript. Y se encuentra en la carpeta `movies-front`.

# Actividades


- Haz un fork de este repo, y clonalo en tu computador personal
- Completa el archivo docker-compose para ejecutar esta aplicación, y configura el archivo `.env`. Apoyate en los archivos README.md que hay en cada carpeta. 
- Levanta la aplicación con `docker-compose up -d --build` y comprueba que se ejecuta correctamente. Toma un pantallazo de tu navegador y agregalo como evidencia en tu fork repo.
- Agrega un archivo llamado `RESPUESTAS.md` donde debes colocar tu nombre completo y las respuestas a las preguntas que vienen al final. 
- Haz push de tu repo, y luego manda un pull request con tu resultado.
- La entrega de la tarea es individual, pero el trabajo puedes realizarlo con tus compañeros.

# Preguntas

1.- Revisa el contenido del diretorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?

Los archivos CREATE son comandos sql para creación de tablas de base de datos, estos definen su estructura y los tipos de datos que contendran sus registros.
Los archivos con ADD son comandos sql del tipo INSERT, es decir, sobre tablas previamente creadas genera registros o filas dentro de la tabla con la información y tipo de datos que estas esperan.

2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?

Al momento de tratar de levantar el servicio este fallará. El motivo es que hemos generado dependencias hacia este servicio en las sentencias "depends on:". En este ejercicio
- movies-api
- flyway
Para poder mantener funcionando el conjunto se debe cambiar no solo el nombre en el servicio sino también en estas dependencias.

3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer?

Bastaría con solo aplicar el cambio en el archivo .env, dado que se ha utilizando en forma paramétrica en el docker-compose.
Luego, para acceder al servicio habría que ingresar haciendo uso de este nuevo puerto como sigue, http://localhost:81

4.- ¿Qué pasa si a la variable de ambiente `BIND_IP` le asignas el valor localhost?

Perdemos la conexión entre el front (movies-front) y el back (movies-api).
Aparentemente al configurar localhost para el contenedor de docker, este asume que el resto de servicios se encuentra en el mismo entorno (contenedor), sin embargo, como hemos levantado front y
back en diferentes contenedores se debe utilizar una dirección que permita conexiones externas como 0.0.0.0

5.- Revisa el archivo `Dockerfile` en la carpeta `movies-api`. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.

Acorde a lo que pude investigar, lo llamativo de este tipo de ejecución es el uso de go mod download. Este comando permite acelerar la construcción de contenedores docker al saltarse el paso de
primero ejecutar los comandos go get ./... antes de generar la construcción en binario.
Esto se traduce en que todas las dependencias de la imagen son capturadas previo a construir el docker container y en evitar un paso adicional en cada ejecución.

Recuerda responder en el archivo `RESPUESTAS.md` y agregar tu nombre en ese archivo.

# Pauta

- Completa el docker-compose: 3 puntos
- Evidencia (pantallazo del navegador): 0.5 puntos
- Pregunta 1: 0.5 puntos
- Pregunta 2: 0.5 puntos
- Pregunta 3: 0.5 puntos
- Pregunta 4: 0.5 puntos
- Pregunta 5: 0.5 puntos

No se corregirá si el PR no trae el archivo `RESPUESTAS.md` y si no incluye el nombre del alumno.

0.2 puntos extras por detectar errores en el enunciado (deben enviar la corrección como un pull request, si es aceptada la corrección se avisará en el foro del curso).
0.5 puntos extras si se proponen mejoras a los archivos fuentes (go, javascript o los dockerfile provistos).
