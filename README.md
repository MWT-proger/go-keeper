# go-keeper

Представляет собой клиент-серверную систему, позволяющую пользователю надёжно 
и безопасно хранить логины, пароли, бинарные данные и прочую приватную информацию.


## Развертывание проекта

1. Склонируйте репозиторий в любую подходящую директорию на вашем компьютере.

```bash
git clone https://github.com/MWT-proger/go-keeper.git
```


2. Скопируйте шаблон файла с переменным окружения

```bash
  cp deployments/.env.example deployments/.env
```

3. Укажите верные переменные окружения в только что созданный файл [.env](deployments/.env)

*Доступны следующие переменные*
```bash
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=testDB
POSTGRES_PORT=5432
```
4. Запустите БД Postgres следующей командой

```bash
  docker compose -f deployments/docker-compose.yaml --env-file deployments/.env up -d
```

5. Запустите сервер
```
go run ./cmd/server -a "localhost:8000" -d "user=postgres password=postgres host=localhost port=5432 dbname=testDB sslmode=disable"
```

## Документация

#### Aвтоматически собирается и размещается на [mwt-proger.github.io/go-keeper](https://mwt-proger.github.io/go-keeper/)

Ведется и собирается с помощью инструмента [Diplodoc](https://github.com/diplodoc-platform/) и GravityUI

### Сборка актуальной версии для локального запуска

1. Собираем
```bash
gh docs -i ./docs -o ./docs/build
```

2. Запускаем
```bash
open docs/build/index.html 
```

---

[ТЕХНИЧЕСКОЕ ЗАДАНИЕ](docs/specification.md)
