Tags: [Bases de Datos](../Indexes/Bases%20de%20Datos.md)
Date: 2024-08-20 13:02
Status: #complete 

# Que son las bases de datos?

Las bases de datos son un conjunto de datos persistentes que pueden ser manipulados por un DBMS(Database Manage System).

## DBMS
[DBMS](../Indexes/DBMS.md)

Un DBMS es lo que a menudo se confunde con bases de datos. Un DBMS son aquellos softwares que proporcionan el manejo de los datos en una base de datos(por ejemplo. MySQL, MongoDB, SQLite, PostgreSQL).

No debemos confundir los DBMS con bases de datos, ya que los DBMS no son bases de datos. Los DBMS operan las bases de datos(por decirlo de alguna forma).

## Bases de Datos Relacionales

Las bases de datos relacionales son un tipo de bases de datos. Este tipo se caracteriza por su forma de organizar datos en filas y columnas, que en conjunto forman una tabla.

*Por que Relacionales?*

Se les conoce de esa manera por que esas bases de datos tienden a formar relaciones con la informacion que se encuentran en las tablas.

Las bases de datos Relacionales operan con un lenguaje de alto nivel llamado *SQL*. Con este lenguaje es posible manipular los datos y acceder a ellos sin la necesidad de crear algoritmos complejos para la busqueda de los datos.

Por ejemplo, si queremos seleccionar todos los datos de nuestra tabla "alumnos".
```
SELECT * FROM alumnos
```

A diferencia de un algoritmo donde se tendria que hacer uso de condicionales y bucles, SQL facilita el proceso con una simple consulta.
