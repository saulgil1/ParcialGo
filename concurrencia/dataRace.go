package concurrencia

import (
	"comparativosConcurrencia/calculadora"
	"fmt"
	"sync"
	"time"
)

func SingleThreaded(coleccionOperaciones []calculadora.Operacion) {

	tiempoInicial := time.Now()

	for i := 0; i < len(coleccionOperaciones); i++ {
		coleccionOperaciones[i].Operar()
	}

	var sum float32
	sum = 0
	for i := 0; i < len(coleccionOperaciones); i++ {
		sum += coleccionOperaciones[i].Resultado

	}

	tiempoFinal := time.Since(tiempoInicial)
	fmt.Println("Suma de resultados -->", sum)
	fmt.Println("Tiempo de computo -->", tiempoFinal)
}

func DataRaceUnaVaraible(coleccionOperaciones []calculadora.Operacion) {

	tiempoInicial := time.Now()
	var wg sync.WaitGroup
	var sumaTotal float32

	for _, operacion := range coleccionOperaciones {
		wg.Add(1)

		// Función para sumar
		go func(op calculadora.Operacion) {
			op.Operar()
			sumaTotal += op.Resultado
			wg.Done()
		}(operacion)
	}

	wg.Wait()
	tiempoFinal := time.Since(tiempoInicial)
	fmt.Println("Suma de resultados con posible pérdida por concurrencia:", sumaTotal)
	fmt.Println("Tiempo de computo -->", tiempoFinal)
}

func DataRaceColeccion(coleccionOperaciones []calculadora.Operacion) {

	tiempoInicial := time.Now()

	//Mecanismo de control de rutinas go
	var wg sync.WaitGroup
	slice := make([]float32, len(coleccionOperaciones))

	for i, operacion := range coleccionOperaciones {
		wg.Add(1)

		// Funcion para sumar
		go func(op calculadora.Operacion, j int) {
			op.Operar()
			slice[j] = op.Resultado
			wg.Done()
		}(operacion, i)
	}
	wg.Wait()

	sumaResultados := float32(0)
	for _, resultado := range slice {
		sumaResultados += resultado
	}
	tiempoFinal := time.Since(tiempoInicial)
	fmt.Println("Suma de los resultados:", sumaResultados)
	fmt.Println("Tiempo de computo -->", tiempoFinal)
}

func MultiThreadedSemaforos(coleccionOperaciones []calculadora.Operacion) {

	tiempoInicial := time.Now()

	var wg sync.WaitGroup
	var mutex sync.Mutex
	var sumaTotal float32

	for _, operacion := range coleccionOperaciones {
		wg.Add(1)

		// Función para sumar
		go func(op calculadora.Operacion) {
			op.Operar()

			mutex.Lock()
			sumaTotal += op.Resultado
			mutex.Unlock()

			wg.Done()
		}(operacion)
	}

	wg.Wait()
	tiempoFinal := time.Since(tiempoInicial)
	fmt.Println("Suma de resultados con posible pérdida por concurrencia:", sumaTotal)
	fmt.Println("Tiempo de computo -->", tiempoFinal)
}

/* Ejemplos de uso de la funcion en el main

// Versión SingleThreaded
concurrencia.SingleThreaded(coleccionOperaciones)

///Versión MultiThreaded - Data Race Una Variable
concurrencia.DataRaceUnaVaraible(coleccionOperaciones)

//Versión MultiThreaded - Data Race coleccion reservada
concurrencia.DataRaceColeccion(coleccionOperaciones)

//Versión MultiThreaded - Semaforo
concurrencia.MultiThreadedSemaforos(coleccionOperaciones)

*/
