Laboratorio 1 - Sistemas Distribuidos


Integrantes:

    > Felix Bastias // 201904558-k
    > Maximiliano Tapia // 202073552-2


Consideraciones:
    > El programa asume que cada capitan al realizar la solicitud de X suministro, solicita siempre una cantidad de 7


Instrucciones de ejecucion:
- Se parte ejecutando el lado del servidor (dentro de carpeta server) con los siguientes comandos en el orden en que se leen:
    > go mod init test
    > docker build . -t test-main:latest
    > docker run -p 3000:8080 test-main:latest

    Donde test y test-main es un nombre escogido por nosotros para probar la correcta ejecucion de los comandos

- Una vez listo eso, se sigue con correr el lado del cliente (dentro de carpeta client, puede hacerse en una terminal paralela a la que se ejecuta el lado del servidor) con el comando:
    > go run client.go
