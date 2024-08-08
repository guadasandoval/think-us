// 3. Suma a cero – 3.5 puntos
// Dado un arreglo de enteros, elimina todos los nodos consecutivos cuya suma sea cero y
// devuelve los nodos restantes. Un arreglo vacío también puede ser un resultado válido. Si se
// recibe un valor nulo, devuelve un arreglo vacío.
// Ejemplo:
// Entrada: [3, 4, -7, 5, -6, 2, 5, -1, 8]
// Salida: [5, 8]

package main

import (
	"fmt"
)

func sumaACero(slice []int) []int {
	if slice == nil {
		return []int{}
	}
	acumuladorSuma := 0
	sumaMap := make(map[int]int) //guardo la suma y la posicion
	sumaMap[0] = -1 // Esto es para manejar el caso donde la suma acumulada sea 0 

	for i := 0; i < len(slice); i++ {
		
		acumuladorSuma += slice[i]

		if pos, ok := sumaMap[acumuladorSuma]; ok {
			slice = append(slice[:pos+1], slice[i+1:]...)
			i = pos
			acumuladorSuma = 0
			sumaMap = make(map[int]int)
			sumaMap[0] = -1
		} else {
			sumaMap[acumuladorSuma] = i
		}
	}

	return slice
}

func main() {
	slice := []int{3, 4, -7, 5, -6, 2, 5, -1, 8}
	result := sumaACero(slice)
	fmt.Println(result) 
}
