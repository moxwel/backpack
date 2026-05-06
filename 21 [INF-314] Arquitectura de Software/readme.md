# Arquitectura de Software

En este curso se estudió los fundamentos de la arquitectura de software, incluyendo patrones de diseño, principios SOLID, y arquitecturas como MVC, microservicios, y serverless. Se hizo uso de herramientas como Docker para contenerizar aplicaciones y Kubernetes para orquestar despliegues de aplicaciones en un entorno de producción simulado.

---

Los archivos del proyecto se encuentran en el siguiente repositorio:

- [utfsm-arquisw-tareafinal](https://github.com/moxwel/utfsm-arquisw-tareafinal)\
_**Microservicio** en **FastAPI** para la gestión de canales y miembros, con persistencia en MongoDB y mensajería de eventos vía RabbitMQ._

- [utfsm-arquisw-tareafinal-gateway](https://github.com/PipeM113/utfsm-arquisw-tareafinal-gateway)\
_**API Gateway** en **FastAPI** para enrutar peticiones a los microservicios del sistema de chat._

- [utfsm-arquisw-tareafinal-front](https://github.com/moxwel/utfsm-arquisw-tareafinal-front)\
_**Frontend** en **React** para el sistema de chat, consumiendo los microservicios desplegados en Kubernetes a través del API Gateway._

## Descripción

Este proyecto es un microservicio en FastAPI para la gestión de canales y miembros para un sistema de chat (clon de Discord) ficticio. El servicio cuenta con persistencia de datos en MongoDB y mensajería de eventos vía RabbitMQ.

Este microservicio formaba parte de un sistema más grande, desplegado en Kubernetes, con otros microservicios para autenticación, gestión de usuarios, etc...
