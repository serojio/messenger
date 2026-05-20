# messenger

## Запуск

### user-service
Для запуска сервиса необходимо перейти в директорию и выполнить команды:

Поднять базу данных
```docker compose up postgres```

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