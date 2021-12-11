# wild_hack
wildberries hackathon result

## Начало работы

Чтобы запустить докерфайл

Создать изображение

```bash
  docker build -t docker-gs-wild-hack .
```

Запустить docker контейнер

```bash
  docker run -d -p 8080:8080 <image name>
```

Чтобы поднять все вместе

```bash
  docker-compose up
```

Запустить postgres в докере
```bash
docker run --name kakash_postgres -e POSTGRES_USER=test -e POSTGRES_PASSWORD=test -e POSTGRES_DB=akshakak -d -p 54320:5432 postgres
```