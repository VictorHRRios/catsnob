                  888                               888      
                  888                               888      
                  888                               888      
 .d8888b  8888b.  888888 .d8888b  88888b.   .d88b.  88888b.  
d88P"        "88b 888    88K      888 "88b d88""88b 888 "88b 
888      .d888888 888    "Y8888b. 888  888 888  888 888  888 
Y88b.    888  888 Y88b.       X88 888  888 Y88..88P 888 d88P 
 "Y8888P "Y888888  "Y888  88888P' 888  888  "Y88P"  88888P"  
                                                             

Para usar este proyecto se ocupara una base de datos postgres


1. Instala postgres

Para macOS con brew
```
brew install postgresql@15
```

Para linux/wsl
```
sudo apt update
sudo apt install postgresql postgresql-contrib
```

2. Para linux
Actualiza la contraseña de postgres, puedes usar algo simple como postgres
```
sudo passwd postgres
```
3. Inicializa el server de postgres
Mac: `brew services start postgresql@15`
Linux: `sudo service postgresql start`

4. Para conectarte el servidor puedes usar cualquier cliente, se da un ejemplo de psql ya que es el default, pero puedes
usar uno que tenga GUI como PGAdmin

Entra a la shell de psql:
```sudo -u postgres psql```

5. Crea una nueva base de datos
```CREATE DATABASE catsnob```

6. Conectate a la base de datos
```\c catsnob```

7. Cambia la contraseña de la base de datos(solo linux)
```ALTER USER postgres PASSWORD 'postgres';```

8. Listo! Puedes checar que todo funcione haciendo queries, por ejemplo:
```SELECT version();```


