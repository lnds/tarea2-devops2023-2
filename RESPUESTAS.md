## Respuestas
# Alumno Eugenio Bravo G.
1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?
    Que en los del verbo create, estan creando las tablas y en el otro estan haciendo insersiones en las tablas
2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?
    Cambiar el nombre del servicio postgres a db en el archivo docker-compose.yml debería funcionar siempre se actualicen todas las referencias en el archivo para reflejar este cambio, admas se debe cambiar todas las instancias de depends_on, links, y cualquier otro lugar donde el servicio postgres esté referenciado, asegurandose de que los otros servicios estén configurados para usar el nuevo nombre en sus conexiones
3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer?
    Se debe cambiar el puerto del servicio movies-api a 81, deberías actualizar la sección correspondiente en el archivo docker-compose.yml y en .env
4.- ¿Qué pasa si a la variable de ambiente BIND_IP le asignas el valor localhost?
    Cambiar la variable de entorno BIND_IP a localhost significa que el servicio movies-api solo escuchará conexiones en la interfaz de loopback. Esto puede ser problemático si esperas que tu aplicación sea accesible desde fuera del contenedor. 
5.- Revisa el archivo Dockerfile en la carpeta movies-api. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.
    que primero se copia el package y luego el resto de archivos