package main

import (
	"fmt"
	"io/ioutil"

	"github.com/saulgil1/ParcialGo/concurrencia"
	"github.com/saulgil1/ParcialGo/gestorArchivos"
)

func main() {
	fmt.Println("Experimentación ST vs MT")

	// Directorio que contiene los archivos
	directorio := "./InstanciasConcurrencia/"

	// Obtener la lista de archivos en el directorio
	archivos, err := ioutil.ReadDir(directorio)
	if err != nil {
		fmt.Println("Error al leer el directorio:", err)
		return
	}

	// Iterar sobre los archivos
	for _, archivo := range archivos {
		// Ignorar si no es un archivo regular
		if !archivo.Mode().IsRegular() {
			continue
		}

		// Obtener el nombre del archivo
		nombreArchivo := archivo.Name()

		// Imprimir el nombre del archivo antes de procesarlo
		fmt.Println("Procesando archivo:", nombreArchivo)

		// Construir la ruta completa al archivo
		rutaArchivo := directorio + nombreArchivo

		fmt.Println("-> Ejecutando versión de un solo hilo para el archivo:", nombreArchivo)
		coleccionOperaciones := gestorArchivos.CargarArchivo(rutaArchivo)

		// Versión SingleThreaded
		concurrencia.SingleThreaded(coleccionOperaciones)

		/*
			///Versión MultiThreaded - Data Race Una Variable
			concurrencia.DataRaceUnaVaraible(coleccionOperaciones)
		*/

		/*
			//Versión MultiThreaded - Data Race coleccion reservada
			concurrencia.DataRaceColeccion(coleccionOperaciones)
		*/

		/*
			//Versión MultiThreaded - Semaforo
			concurrencia.MultiThreadedSemaforos(coleccionOperaciones)
		*/
		fmt.Println()
	}

}
