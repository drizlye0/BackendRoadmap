Tags: [Bases de Datos](../Indexes/Bases%20de%20Datos.md)
[SQL](../Indexes/SQL.md) Date: 2024-08-20 13:09 Status: #complete

# Problem

Query all columns for all American cities in the CITY table with populations
larger than 100000. The CountryCode for America is USA.

The CITY table is described as follows
![Table](https://s3.amazonaws.com/hr-challenge-images/8137/1449729804-f21d187d0f-CITY.jpg)

## Solution

```
SELECT *
FROM CITY
WHERE POPULATION >= 100000 AND COUNTRYCODE="USA"
```
