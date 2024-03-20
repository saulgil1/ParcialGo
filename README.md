# ParcialGo

Functionality of the module

The purpose of this module is to return the computation time of instances using various methods, from Single Threaded, Multi Threaded on a variable, Multi Threaded with a collection, and finally Multi Threaded implementing semaphores. This is to better visualize the differences between each method as well as to analyze the results to determine which solution to implement in different situations depending on the required purpose.

How to use it

First, select an instance to compute, this instance should be a text file with 3 columns. The first column would be a number A, the second column would be a number B, and finally, the third column would represent the equation to perform with A and B, for example:

320 240 -

Then, proceed to process the file using:

archivoProcesado := gestorArchivos.CargarArchivo(instancia)

Where the text document is processed, it is recommended to store it in a variable since this processed file will be sent as a parameter to the following functions.

Finally, it is sent through:

concurrencia.SingleThreaded(archivoProcesado) -- ST
concurrencia.DataRaceUnaVaraible(archivoProcesado) -- MT on a variable
concurrencia.DataRaceColeccion(archivoProcesado) -- MT on a collection
concurrencia.MultiThreadedSemaforo(arhivoProcesado) -- MT using semaphores
