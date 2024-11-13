# WebSocket Repositiry

## Tecnologias
* Gin 
* Gorilla/WebSocket

## To execute
1. Create mod  ```go mod init backend```
2. dependencies ```go mod tidy```
3. Run Server ``` go run ```

## Routes
* ### Ping
    #### Method: GET
    #### ```http://localhost:8080/ping```
    Test route to check if the server is running correctly with Gin
    #### Response
    ```
    {
    pong
    }
    ```
