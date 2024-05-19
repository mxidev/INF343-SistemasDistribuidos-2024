Félix Bastías       201904558-k
Maximiliano Tapia   202073552-2

Para ejecutar se debe estar en una terminal en la dirección de las carpetas director, namenode, datanodes y el doshbank segun la MV que se le asignó. Luego, ocupar los comandos de make:

    make para buildear el docker.
    make run para iniciar el programa.

Para ejecutar los mercenarios, se debe ocupar el comando:
    go run jugador.go (en el caso del jugador)
    go run mercenarioX.go (donde X es un número del 1 al 7 y son los bots)

Ejemplo:
    Para NameNode, ubicado en la MV dist077: Entrar al directorio INF343-SistemasDistribuidos-2024/Laboratorio 4/namenode. Luego ejecutar los comandos make de arriba.
    Asi para todas las demás entidades según su MV asignada.

MV Asignadas
- NameNode, DoshBank -> MV dist077

- Director, Mercenarios y DataNode1 -> MV dist078

- DataNode2 -> MV dist079

- DataNode3 -> MV dist080


CONSIDERACIONES
- Se asume que en una misma MV se dockeriza Director y Mercenarios.
- Se intentó ocupar el comando make en las MV, pero se debía hacer un sudo y no teniamos las claves. Es por esto que se dockerizó de manera local como si fuera cada MV y no se presentaron problemas.
- Se recomienda ocupar make en todas las entidades antes de ejecutar los mercenarios.
- Debido a limitaciones de tiempo, no se pudo implementar todas las funcionalidades que se pedia por enunciado. Pero si se realizo la dockerización, makefiles, y se implementó el diseño de pisos junto con las entidades. Además, se hizo un "esquema" de conexiones para probar la comunicación.
