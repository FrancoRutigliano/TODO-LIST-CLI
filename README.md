
# TODO LIST CLI 💙 

Este proyecto consiste en una aplicación de lista de tareas (Todo List) diseñada para funcionar en la interfaz de línea de comandos (CLI). Proporciona una manera simple y eficiente de gestionar tus tareas diarias directamente desde la terminal.



## 🤖Características Principales

**Interfaz de Línea de Comandos Intuitiva**: Accede y gestiona tus tareas de manera rápida y sin esfuerzo desde la terminal.

**Operaciones Básicas**: Agrega nuevas tareas, marca tareas como completadas y elimina tareas de manera sencilla.

**Listado Claro**: Visualiza tus tareas de manera clara y ordenada, con indicadores visuales para las tareas completadas.

**Persistencia de Datos**: Las tareas se almacenan localmente para que puedas retomar tu lista en futuras sesiones.

**Personalización**: Configura y organiza tus tareas de acuerdo con tus preferencias.
## ✨Instalación

1. Instalación:
> - Clona el proyecto.
```bash
    git clone https://github.com/FrancoRutigliano/TODO-LIST-CLI
```
> - Asegurate de tener [Go](https://go.dev/learn/) instalado en tu sistema.

2. **Ejecuta** el siguiente comando para instalar las dependencias.

```bash
 $ go get -u github.com/FrancoRutigliano/TODO-LIST-CLI
```
    
3. Luego compila el programa(***Ejecutable***).

```bash
 $ go build ./cmd/todo
```

## Usos/Ejemplos

***Listar Tareas***
```bash
./todo -list
```

***Añadir Tareas***
```bash
./todo -add Nombre_de_la_tarea
```
> *Para agregar una tarea **no** debes utilizar comillas*.

***Eliminar Tarea***
```bash
./todo -add número_de_tarea_eliminar
```


## Autor

- [@FrancoRutigliano](https://github.com/FrancoRutigliano)


## Contribuciones

¡Las contribuciones son bienvenidas! Si encuentras algún error o tienes ideas para mejorar este proyecto, por favor abre un problema o envía una solicitud de extracción.`.


