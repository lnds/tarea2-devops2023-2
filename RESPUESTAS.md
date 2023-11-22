# Julio Soto Sánchez
1.- Revisa el contenido del diretorio sql_migrations. ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?

```
En la carpeta sql_migrations, existen 4 archivos de los cuales en su
nombre destaca lo que hará el contenido del archivo, es decir,

el archivo V1__Create posee sentencias de creación de tablas y

el archivo V1__Add posee sentencias Insert a las tablas ya creadas.
```

2.- ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?

Si cambiamos el servicio de postgres a db ocurre el siguiente error.
```
juliosoto@MacBook-Pro-de-Julio tarea2-devops2023-2 % docker-compose up -d --build
service "movies-api" depends on undefined service postgres: invalid compose project
```
Esto es provocado por que en su dependencia, en el archivo docker-compose.yml hace referencia al servicio postgres tanto en el servicio db como en movies-api

Para poder realizar el cambio de postgres a db, debemos además de modificar el servicio, debemos modificar las dependencias de este mismo en el archivo docker-compose.yml
Además debemos modificar la dependencia del servicio movies-api de postgres a db, como se muestra a continuación:

```
depends_on:
      - db
```

3.- Si quisieramos que el servicio movies-api use el puerto 81, ¿Qué cambios habría que hacer?

```
Si quisieramos cambiar el puerto del servicio movies-api, basta con cambiar en el archivo .env
el puerto en la variable de entorno BIND_PORT=81.

Esta modificación es posible, dado que en el archivo de docker-compose.yml utilizamos esta variable
de entorno para levantar la API.

Además en el archivo main.go se hizo necesario realizar una llamada directa a la variable de entorno .env
dado que en ese código el puerto está siendo utilizado directamente en el 8080 como se muestra a continuación:

var port string = "8080"

El código utilizado para que la aplicación utilice la variable de entorno es el siguiente:

import (
	"os"
	"github.com/spf13/cobra"
)

var port string = os.Getenv("BIND_PORT")

```

4.- ¿Qué pasa si a la variable de ambiente BIND_IP le asignas el valor localhost?

```
Al cambiar la variable de entorno BIND_IP no es posible visualizar el contenido del servicio movies-api.

Esto genera el siguiente error visible a través del navegador:

WebSocket connection to 'ws://localhost:3000/ws' failed: The operation couldn’t be completed. Socket is not connected

```

5.- Revisa el archivo Dockerfile en la carpeta movies-api. ¿Qué te llama la atención? Trata de explicar lo que ocurre en este caso.

```
Se puede apreciar en la línea N°1 que a la versión de FROM golang:1.19 AS build-stage
le está asignando un tag llamado bluid-stage,
para que de esta forma sea llamado de aquí en adelante. Además al incluír esta forma de llamado podemos visualizar que se utiliza este "tag"
al momento de utilizar el siguiente commando: COPY --from=build-stage /movies-api /movies-api

```
