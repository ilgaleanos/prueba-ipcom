# prueba-ipcom

En el presente repositorio se encuentra la solución a los dos puntos para la prueba.

## Ejecutar el código

Para poder ejecutar el codigo simplemente necesita clonar el repositorio e instalar las dependecnias con el comando
```bash 
go mod download
``` 

y finalmente ejecutar el codigo
```bash 
go run main
``` 

## Primer punto

La solucion está enfocada en solucionar un problema de consumo de api masiva por medio del uso de concurrencia, se consulta en 

```bash 
curl --location --request GET '127.0.0.1:8080/resumen/2019-12-01?dias=15'
``` 

Como nota adicional se han limitado los dias máximos de consulta a 15 días ya que no hay directriz

## Segundo punto

La solucion está enfocada en leer el fichero linea a linea pensando en un gran tamaño del mismo del mismo, de la misma forma se soluciona por medio de la buscqueda de los elementos dentro de la estructura por medio de índices.

```bash 
curl --location --request GET '127.0.0.1:8080/agrupador'
``` 

## Nota
Se ha realizado una integracion de los dos puntos en endpoints como tal, de modo que basta seguir las referencias del main [router] para llegar a cada controlador.
