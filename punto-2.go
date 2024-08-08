// 2. Encripta tu mensaje – 3.5 puntos
// Has sido encargado de desarrollar una nueva forma de encriptar comunicaciones.
// Básicamente, cada vocal del mensaje de entrada deberá ser precedida por otra cadena,
// llamada la clave. La función recibirá dos parámetros de cadena: el primero será la clave y el
// segundo, el mensaje. La función devolverá una cadena.
// Ejemplo:
// Clave: dcj
// Mensaje: I love prOgrAmming!
// Esto debería devolver: dcjI ldcjovdcje prdcjOgrdcjAmmdcjing!
// No se tomarán en cuenta las vocales con acentos.
// Cuando se reciba una cadena vacía, debería devolver una cadena vacía. Si el mensaje es
// nulo o vacío, devolver una cadena vacía. Si la clave es nula o vacía, entonces utiliza
// DCJ como valor predeterminado.

package main

import (
	"fmt"
	"strings"
)

func encriptarMensaje(key, mensaje string) string {
	
	if len(mensaje) == 0 {
		return ""
	}

	if len(key) == 0{
		key = "DCJ"
	}

	vocales := "aeiouAEIOU"
	
	var result strings.Builder

	for _, char := range mensaje {
		if strings.ContainsRune(vocales, char) {
			result.WriteString(key)  //concateno la clave a la vocal
		}
		result.WriteRune(char)
	}

	return result.String()
}

func main() {
	key := "dcj"
	message := "I love prOgrAmming!"
	encryptedMessage := encriptarMensaje(key, message)
	fmt.Println(encryptedMessage)
}
