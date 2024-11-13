# WebSocket Repositiry
## Tecnologías
* Gin
* Gorilla/WebSocket

## Para ejecutar
1. Crear mod  ```go mod init backend```
2. Dependencias ```go mod tidy```
3. Ejecutar servidor ``` go run ```

## Rutas
* ### Ping
    #### Método: GET
    #### ```http://localhost:8080/ping```
    Ruta de prueba para verificar si el servidor está funcionando correctamente con Gin
    #### Respuesta
    ```
    {
    pong
    }
    ```

# Mensaje

Este endpoint permite enviar un mensaje mediante una solicitud HTTP POST.

## Cuerpo de la solicitud

La solicitud debe incluir los siguientes parámetros en el cuerpo de la solicitud en formato raw:

- `Username` (string): El nombre de usuario del remitente.
    
- `Contenido` (string): El contenido del mensaje.
    

Ejemplo:

``` json
{
    "Username": "Manuel",
    "Contenido": "Wow que bacán la mensajería"
}

 ```

## Respuesta

La respuesta a la solicitud contendrá los detalles del mensaje enviado.

Este endpoint permite al usuario agregar un nuevo mensaje.

#### Cuerpo de la solicitud

- `Username`: (texto) El nombre de usuario del remitente.
    
- `Contenido`: (texto) El contenido del mensaje.
    

#### Respuesta

La respuesta contendrá el estado de la adición del mensaje.
