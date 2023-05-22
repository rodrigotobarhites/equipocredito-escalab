# Instalación de la aplicación Kubernetes

## Pre-requisitos

1. Kubernetes CLI (kubectl)
2. Git
3. Acceso a un cluster de Kubernetes

Si no tienes `kubectl` instalado, puedes seguir las instrucciones [aquí](https://kubernetes.io/docs/tasks/tools/install-kubectl/).

Para verificar que tienes `kubectl` instalado y configurado correctamente, puedes ejecutar el siguiente comando:

```shell
kubectl version
```


## Instalación

Primero, clona el repositorio de git que contiene los archivos yaml. Reemplaza <url-del-repositorio> con la URL del repositorio:

```shell
git clone https://github.com/rodrigotobarhites/equipocredito-escalab.git
cd equipocredito-escalab/03.\ PARTE\ 2\ -\ Kubernetes/
```

Una vez que hayas clonado el repositorio y estés en el directorio correcto, puedes empezar a aplicar los archivos yaml en el orden indicado.

Los comandos son los siguientes:

```shell
kubectl apply -f 01-namespace.yaml
kubectl apply -f 02-frontend-configmap.yaml
kubectl apply -f 03-frontend-deployment.yaml
kubectl apply -f 04-backend-configmap.yaml
kubectl apply -f 05-backend-secrets.yaml
kubectl apply -f 06-backend-deployment.yaml
kubectl apply -f 07-frontend-service.yaml
kubectl apply -f 08-backend-service.yaml
kubectl apply -f 09-loadbalancer.yaml
```

Estos comandos crearán los componentes necesarios para la aplicación en tu cluster de Kubernetes.

## Verificación

Para verificar que los servicios y deployments se han creado correctamente, puedes usar los siguientes comandos:

```shell
kubectl get services
kubectl get deployments
```

Al visualizar service pueden ejecutar el sitio desde un balanceador de carga relacionado segun la IP otorgada o bien generar un port-forward hacia el puerto 8080.

```shell
kubectl port-forward -n equipocredito svc/frontend 8080:8080
```

Estos comandos mostrarán una lista de todos los servicios y deployments en tu cluster. Deberías ver los servicios y deployments que acabas de crear listados aquí.

Adjuntamos un video con la ejecución del proyecto:
http://youtube.com/...

