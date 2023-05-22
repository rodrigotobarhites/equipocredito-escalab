# Instalación de la aplicación Docker Compose

## Requisitos

- Docker: asegúrate de tener Docker instalado en tu máquina. Si no lo tienes, puedes descargarlo [aquí](https://www.docker.com/products/docker-desktop).
- Docker Compose: necesitarás Docker Compose para orquestar tus servicios. Si no lo tienes, puedes seguir las instrucciones de instalación [aquí](https://docs.docker.com/compose/install/). Asegúrate de instalar al menos la versión 3.8.

## Descargando el archivo docker-compose.yaml

El archivo `docker-compose.yaml` se encuentra disponible en un repositorio git público. Para clonar el repositorio y obtener el archivo, ejecuta el siguiente comando:

```bash
git clone https://github.com/rodrigotobarhites/equipocredito-escalab.git
```




## Configuración y uso
Navega hasta el directorio del repositorio que acabas de clonar:

```bash
cd  equipocredito-escalab/02.\ PARTE\ 1\ -\ Docker\ Compose/
```


Asegúrate de que el archivo docker-compose.yaml esté en el directorio. Puedes verificarlo ejecutando:

```bash
ls
```
Deberías ver el archivo docker-compose.yaml en la lista.

Inicia los servicios con Docker Compose. Ejecuta el siguiente comando:

```bash
docker-compose up -d
```

Esto iniciará tus servicios en segundo plano. Para ver los logs, puedes ejecutar:

```bash
docker-compose logs -f
```

Para detener los servicios, puedes ejecutar:

```bash
docker-compose down
```

Esto detendrá y eliminará los contenedores, pero mantendrá tus datos.

Si quieres detener los servicios y eliminar los datos, puedes ejecutar:

```bash
docker-compose down -v
```

Esto eliminará los volúmenes definidos en tu archivo docker-compose.yaml.

## Servicios
Este docker-compose contiene dos servicios:

web: este es tu servicio de aplicación que se expone en el puerto 80 (localhost). Se conecta a la base de datos mediante la red webnet.
db: este es tu servicio de base de datos que utiliza una imagen Postgres y almacena los datos en un volumen Docker llamado db_data.
Ambos servicios se reinician automáticamente si se detienen inesperadamente.

## Version compatible

Si por alguna razón no es compatible con la version de Docker Compose, se debe descargar e instalar lo siguiente.

```bash
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
docker-compose --version
```

Esto debería imprimir algo como: docker-compose version 1.29.2, build 5becea4c.


Adjuntamos un video con la ejecución del proyecto:
https://youtu.be/10JKsie5BDk

