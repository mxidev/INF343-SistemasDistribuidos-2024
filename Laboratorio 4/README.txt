Félix Bastías       201904558-k
Maximiliano Tapia   202073552-2

Para ejecutar se debe estar en una terminal en la dirección de las carpetas director, namenode, datanodes y el doshbank segun la MV que se le asigno. Luego, ocupar los comandos de make:

make para buildear el docker.
make run para iniciar el programa.

Ejemplo:
    Para NameNode, ubicado en la MV dist077: Entrar al directorio INF343-SistemasDistribuidos-2024/Laboratorio 4/namenode. Luego ejecutar los comandos make de arriba.
    Asi para todas las demas entidades segun su MV asignada.

MV Asignadas
- NameNode, DoshBank -> MV dist077

- Director, Mercenarios y DataNode1 -> MV dist078

- DataNode2 -> MV dist079

- DataNode3 -> MV dist080


CONSIDERACIONES
- Se asume que en una misma MV se dockeriza Director y Mercenarios
- Debido a limitaciones de tiempo, no se pudo implementar todas las funcionalidades que se pedia por enunciado. Pero si se realizo la dockerizacion, makefiles, y se implementaron el diseño de pisos junto con las entidades. Ademas, se hizo un "esquema" de conexiones para probar la comunicacion.
