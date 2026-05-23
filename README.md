# messenger

## Запуск

### user-service
Для запуска сервиса необходимо перейти в директорию и выполнить команды:

```cd services/user-service```

Поднять базу данных
```docker compose up -d postgres```

Выполнить миграции
```docker compose up migrate```
(Проверить таблицу с помощью psql)
```docker exec -it user-service-postgres-1 psql -U postgres -d msgr_users```

Запустить сервис на localhost
```go run cmd/app/main.go```


(Для обновления/удаления миграций)
Apply up
```docker compose run migrate up```

Rollback one migration
```docker compose run migrate down 1```

Rollback all
```docker compose run migrate down```

Current version
```docker compose run migrate version```

Force version
```docker compose run migrate force 1```

(Аннигиляция)
```docker compose down --remove-orphans```

### post-service
Для запуска сервиса необходимо перейти в директорию и выполнить команды:

```cd services/post-service```

Поднять базу данных
```docker compose up -d postgres```

Выполнить миграции
```docker compose up migrate```
(Проверить таблицу с помощью psql)
```docker exec -it post-service-postgres-1 psql -U postgres -d msgr_posts```

Запустить сервис на localhost
```go run cmd/app/main.go```


(Для обновления/удаления миграций)
Apply up
```docker compose run migrate up```

Rollback one migration
```docker compose run migrate down 1```

Rollback all
```docker compose run migrate down```

Current version
```docker compose run migrate version```

Force version
```docker compose run migrate force 1```

(Аннигиляция)
```docker compose down --remove-orphans```