# Saludos en Go

Este paquete proporciona una forma simple de obtener saludos personalizados en Go.

## Instalación
Ejecuta el siguiente comando para instalar el paquete
```bash
go get -u git github.com:bill-sama9008/greetings
```

## Uso
Aqui tienes un ejemplo de cómo utilizar el paquete en tu código:

```go
package main

import (
	"fmt"
	"log"

	"github.com/bill-sama9008/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Andres", "Bill", "Elons"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// message, err := greetings.Saludo("Andres")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println(messages)
}
```