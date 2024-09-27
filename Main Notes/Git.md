Tags: [Version Control System](../Indexes/Version%20Control%20System.md) Date:
2024-08-20 13:12 Status: #complete

# Git

Git es un software para el control de versiones de nuestro codigo. Fue
desarollado por Linus Torvalds(si, el creador del kernel linux) y hasta el dia
de hoy, se ha convertido en un estandar que todo desarrollador tiene que
aprender si desea trabajar con otros desarrolladores.

A dia de hoy, git es usado por las grandes empresas de software para llevar a
cabo sus proyectos, ademas de eso, git puede ser usado no solamente para
proyectos de software, si no para practicamente cualquier proyecto en el cual
necesesites tener un seguimiento de cambios en los archivos.

Con git, podemos trabajar con repositorios que estan localmente en nuestro
dispositivo, git hace un siguimiento de cambios de nuestros archivos por si
queremos volver a un punto en especifico de nuestro codigo. Para poder usar git
hay que usar la consola de comandos(powershell en windows, bash para linux y
macos).

## Comandos De Git

- **git init**: Es el comando para poder iniciar un repositorio en la carpeta
  actual donde nos encontremos.
- **git add**: Con este comando puedes los archivos que quieres que git tenga un
  siguimiento de cambios, si un archivo no se agrega con este comando, git no lo
  tomara en cuenta y no existira en el repositorio.
- **git status**: Muestra el estado actual de nuestro repositorio, aqui podemos
  ver cuantos archivos tienen cambios o cuales no se han agregado, etc.
- **git commit _archivo_ -m ' '**: Este es el comando para poder guardar nuestro
  cambios de nuestros archivos en el repositorio.
- **git branch**: Muestra las ramas actuales que hay en nuestro repositorio.
- **git checkout**: Permite cambiarnos de rama.
- **git merge**: Si estamos trabajando con ramas, este comando nos servira para
  subir nuestros cambios a la rama principal u otra.
- **git clone**: Clona un repositorio que este en algun servicio de hostings
  para repositorios en nuestra maquina local.
- **git remote add**: Nos permite agregar un repositorio remoto.
- **git fetch**: Hace un Fetch a nuestro repositorio remoto.
- **git pull**: Extrae los cambios de una rama remota y los refleja en nuestro
  repositorio.
- **git push**: Sube todos nuestros cambios al repositorio remoto.
