Tags: [Bases de Datos](../Indexes/Bases%20de%20Datos.md)
Date: 2024-08-20 13:04
Status: #progress 

# Que son las Relaciones?
Las relaciones son las conexiones que pueden llegar a existir entre las *Entidades*, estas relaciones representan como la entidades estan conectadas entre si y como interactuan en el modelo *ER*.

Por ejemplo, entidades como *Estudiantes* y *Escuelas* estan relacionadas, ya que los Estudiantes van a a las Escuelas y las Escuelas tienen estudiantes. Otro ejemplo podria ser con *Autores* y *Libros*, los autores escriben los libros y los libros son escritos por autores.

## Cardinalidad
La *Cardinalidad* es el termino para identificar el como nuestras entidades estan relacionadas. La cardinalidad basicamente se trata del tipo de relacion que tienen nuestras entidades.
- __Relacion uno a uno(1:1)__: Este tipo de relacion sucede cuando una instancia de la entidad esta relacionada con una sola instancia de otra entidad.
- __Relacion uno a muchos(1:N)__: Sucede cuando una instancia de la entidad esta relacionada con multiples instancias de otra entidad.
- __Relacion muchos a uno(N:1)__: Sucede al reves que 1:N, multiples instancias de una Entidad estan relacionadas con una sola instancia de otra Entidad.
- __Relacion muchos a muchos(N:N)__: Es cuando multiples instancias de una entidad estan relacionadas con otras multiples instancias de otra entidad.

# Referencias
[modelo-entidad-relacion](https://keepcoding.io/blog/modelo-entidad-relacion/)
[Cardinalidad en Base de Datos](https://www.youtube.com/watch?v=f5ZB05OWNCM&list=PLE19WVM4rff2rESwl-_MC9dxGT_tIeQcO&index=4)
