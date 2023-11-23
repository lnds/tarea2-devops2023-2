# RESPUESTAS: 
# Erikson Fuenzalida Diaz

1.- Revisa el contenido del directorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?
R.- Los archivos CREATE se usan para crear las tablas necesarias que usaremos en el proyecto (tablas: directors y movies), y se complementan con los archivos ADD para hacer un insert de la data que necesita cada una de las tablas creadas.

2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?
R.- Si cambiamos el servicio de postgres a db deberiamos cambiar o actualizar las referencias para que funcione,  ya que los demas archivos o servicios que usen como referencia postgres necesitan tambien este nombre para funcionar correctamente.

3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer?
R.- Si hicieramos este cambio deberiamos hacer cambios en el archivo docker-compose.yml y en .env en las variables como BIND_PORT y REACT_APP_API_URI.

4.- ¿Qué pasa si a la variable de ambiente BIND_IP le asignas el valor localhost?
R.- Si realizamos esto tendriamos problemas ya que el front de la api no estaria accesible dede fuera del contenedor y solo funcionaria internamente.

5.- Revisa el archivo Dockerfile en la carpeta movies-api. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.
R.- En el dockerfile se esta haciendo con multi-stage-build desde una imagen golang una con dedian, esto permite que la creacion de imagenes se mas liviana en el proceso. Ademas el uso de USER nonroot permite evitar vulnerabilidades en la ejecucion.