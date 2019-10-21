# Trabajo  <img src="https://upload.wikimedia.org/wikipedia/commons/f/fc/UPC_logo_transparente.png" width="40" height="40"/>

El trabajo final consta en construir una aplicación distribuida aplicando machine learning y blockchain para la solución de un problema en la industria peruana. Aplicando el correcto diseño de aplicaciones concurrentes. Se debe usar el lenguaje Go y el verificador de modelos Spin y Github.

![Alt Text](https://s3.amazonaws.com/dev.assets.neo4j.com/wp-content/uploads/20180109025607/import-bitcoin-blockchain-neo4j-graph-database.gif)


## Blockchain 
La cadena de bloques, más conocida por el término en inglés blockchain, es un registro único, consensuado y distribuido en varios nodos de una red. En el caso de las criptomonedas, podemos pensarlo como el libro contable donde se registra cada una de las transacciones.

Su funcionamiento puede resultar complejo de entender si profundizamos en los detalles internos de su implementación, pero la idea básica es sencilla de seguir.

En cada bloque se almacena:

- Una cantidad de registros o transacciones válidas,
- Información referente a ese bloque,
- Su vinculación con el bloque anterior y el bloque siguiente a través del hash de cada bloque ─un código único que sería como   la huella digital del bloque.

Por lo tanto, **cada bloque tiene un lugar específico e inamovible dentro de la cadena**, ya que cada bloque contiene información del hash del bloque anterior. La cadena completa se guarda en cada nodo de la red que conforma la blockchain, por lo que **se almacena una copia exacta de la cadena en todos los participantes de la red**.

## ¿Por qué blockchain es tan segura?

Al ser una tecnología distribuida, donde cada nodo de la red almacena una copia exacta de la cadena, se garantiza la disponibilidad de la información en todo momento. En caso de que un atacante quisiera provocar una denegación de servicio, debería anular todos los nodos de la red, ya que basta con que al menos uno esté operativo para que la información esté disponible.

![Alt Text](https://www.cognierblock.com/img/bi-admin.gif)



