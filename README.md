# DNA Mutant Detector API
Esta api está construida sobre AppEngine de google, usando **FireStore** en **Datastore Mode** como base de datos para almacenar todos los checkeos de ADN realizados y los Stats

### Posibles Mejoras
* Agregar Header de Cache (None si no tiene)
* Agregar un Cache de Proxy Reverso Como Varnish para atajar los ADN ya detectados.
* Cambiar la base de datos de FireStore a MongoDB para poder generar stats dinamicamente.
* Separar la capa de Guardado en Base de Datos de la de Calculo agregando RabbitMQ, asi tendriamos mejores tiempos de respuesta y podemos escalar horizontalmente nuestra API
* Test de Integracion

## Endpoint
`https://meli-challenge-261304.appspot.com`

## Metodos

### [GET] / (healthcheck)
```
OK
```
Response Codes:
* 200-OK: Alive
* 500-InternalServerError: error en el servidor

Rate Limit:
`none`


### [POST] /mutant
```
{
	"dna": ["ATGCGAA", "CAGTGCC", "TTATGTT", "AGAAGGG", "CCCCTAA", "TCACTGG", "ATGCGAA"]
}
```

Response Codes:
* 200-OK: mutante
* 403-Forbidden: humano
* 500-InternalServerError: error en el servidor
* 400-BadRequest: Input Mal formado

Rate Limit:
`none`

### [GET] /stats
```
{
    "count_human_dna": 0,
    "count_mutant_dna": 1,
    "ratio": 0
}
```

Response Codes:
* 200-OK
* 500-InternalServerError: error en el servidor

Rate Limit:
`none`
