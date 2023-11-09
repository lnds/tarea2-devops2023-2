# Tarea 3

Gracias al éxito de la aplicación de libros en Flask, les han encargado la construcción de una app web para películas.

Dado que la aplicación tendrá mayor demanda se han cambiado las tecnologías, el backend está escrito en go y se encuentra en la carpeta `movies-api`.

El frontend es una aplicación REACT en Javascript.

# Actividades


- Haz un fork de este repo, y clonalo en tu computador personal
- Completa el archivo docker-compose para ejecutar esta aplicación, y configura el archivo `.env`. Apoyate en los archivos README.md que hay en cada carpeta. 
- Levanta la aplicación con `docker-compose up -d --build` y comprueba que se ejecuta correctamente. Toma un pantallazo de tu navegador y agregalo como evidencia en tu fork repo.
- Agrega un archivo llamado `RESPUESTAS.md` donde debes colocar tu nombre completo y las respuestas a las preguntas que vienen al final. 
- Haz push de tu repo, y luego manda un pull request con tu resultado.
- La entrega de la tarea es individual, pero el trabajo puedes realizarlo con tus compañeros.

# Preguntas

1.- Revisa el contenido del diretorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?
2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?
3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer? 
4.- ¿Qué pasa si a la variable de ambiente `BIND_IP` le asignas el valor localhost?
5.- Revisa el archivo `Dockerfile` en la carpeta `movies-api`. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.

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
