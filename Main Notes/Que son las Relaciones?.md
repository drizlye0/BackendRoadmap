Tags: [Bases de Datos](../Indexes/Bases%20de%20Datos.md) Date: 2024-08-20 13:04
Status: #progress

# Que son las Relaciones?

Las relaciones son las conexiones que pueden llegar a existir entre las
_Entidades_, estas relaciones representan como la entidades estan conectadas
entre si y como interactuan en el modelo _ER_.

Por ejemplo, entidades como _Estudiantes_ y _Escuelas_ estan relacionadas, ya
que los Estudiantes van a a las Escuelas y las Escuelas tienen estudiantes. Otro
ejemplo podria ser con _Autores_ y _Libros_, los autores escriben los libros y
los libros son escritos por autores.

## Cardinalidad

La _Cardinalidad_ es el termino para identificar el como nuestras entidades
estan relacionadas. La cardinalidad basicamente se trata del tipo de relacion
que tienen nuestras entidades.

- **Relacion uno a uno(1:1)**: Este tipo de relacion sucede cuando una instancia
  de la entidad esta relacionada con una sola instancia de otra entidad.
- **Relacion uno a muchos(1:N)**: Sucede cuando una instancia de la entidad esta
  relacionada con multiples instancias de otra entidad.
- **Relacion muchos a uno(N:1)**: Sucede al reves que 1:N, multiples instancias
  de una Entidad estan relacionadas con una sola instancia de otra Entidad.
- **Relacion muchos a muchos(N:N)**: Es cuando multiples instancias de una
  entidad estan relacionadas con otras multiples instancias de otra entidad.

# Referencias

[modelo-entidad-relacion](https://keepcoding.io/blog/modelo-entidad-relacion/)
[Cardinalidad en Base de Datos](https://www.youtube.com/watch?v=f5ZB05OWNCM&list=PLE19WVM4rff2rESwl-_MC9dxGT_tIeQcO&index=4)
