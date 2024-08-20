Tags: [[Bases de Datos]] [[DBMS]]
Date: 2024-08-20 13:06
Status: #complete 

# Configurar PostgreSQL

Para poder configurar PostgreSQL en linux, debemos tener un poco de conocimiento en la terminal del linux.

En mi caso estare usando archlinux para configurar PostgreSQL
## Pasos
1. Despues de haber instalado PostgreSQL para nuestra distribucion(en mi caso es archlinux) ejecutar los siguientes comandos en la terminal.
   ```
   $ sudo -iu postgres
   [postgres]$ initdb -D /var/lib/postgres/data 
	```
	Con esto tenemos iniciado el database cluster	
2. Ahora tenemos que iniciar el servicio `postgresql` en el systemd
   ```
   $ sudo systemctl start postgresql   
   $ sudo systemctl enable postgresql
	```
	Con esto ya estaria PostgreSQL iniciado en nuestro sistema
3. Ahora tendremos que crear un usuario para nuestras bases de datos
   ```
   $ sudo -u postgres createuser --createdb --pwprompt nombredeusuario   
   // despues creamos una db que tenga nuestro usuario como owner
   $ sudo -u postgres createdb --owner=nombredeusuario nombredeusuario
	```
4. Con esto ya podriamos estar usando el usuario creado para hacer nuestras consultas
   ```
   $ psql --host localhost --user nombredeusuario   
	```
