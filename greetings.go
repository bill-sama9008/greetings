package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un saludo para la persona especificada
func Saludo(name string) (string, error) {

	if name == "" {
		return "", errors.New("Nombre vacio")
	}

	// Devuelve un saludo que incluye el nombre en un mensaje
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Saludo(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hola, %v Bienvenido",
		"Que bueno verte, %v",
		"Saludo, %v Encantado de conocerte",
	}

	return formats[rand.Intn(len(formats))]
}
