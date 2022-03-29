### Descripción de los Ejemplos ###

Para ejecutar los ejemplos en la carpeta donde está el código **usar el comando** "go run ."
todos los ejemplos tienen un benchmark utilizando el paquete de testing de golang. para ejecutar **utilizar el comando "go test -bench=."**
 
#### Carpeta - Secuencial ####
Este ejemplo ejecuta secuencialmente una función que verifica si responden los URLs almacenados en un slice. el resultado es desplegado a la consola de salida

El resto de los ejemplos construyen sobre este. Es una adaptación del código de https://github.com/quii/learn-go-with-tests/tree/main/concurrency


#### Carpeta - sync_v1 y v2 ####
Estos ejemplos utilizan el paquete sync, en particular el WaitGroup para sincronizar las gorutinas. El objetivo es ver cómo lanzar gorutinas que **NO** comparten datos.
Se lanza una gorutina para cada url a verificar y se sincroniza con la principal (join) mediante wait.
Este paquete además incluye semáforos Mutex para proteger secciones críticas, pero no se utilizan en este ejemplo.
El primero **lanza como gorutina una función con nombre** y el segundo (v2) **lanza funciones anónimas.**

#### Carpeta - channels_v1 ####
Este ejemplo utiliza otro mecanismo de sincronización de Go, los canales.
En el mismo se lanzan las gorutinas y en lugar de imprimir el resultado en consola lo escriben a un canal, el cual es leído por la gorutina principal para desplegar por la consola de salida