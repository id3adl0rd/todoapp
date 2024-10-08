### VERBA-group тест-кейс to-do list

## Задача: Разработка REST API Для управления задачами (TO-DO list)

В рамках задачи был разработан REST API с применением gorilla/mux, gorm, zerolog, godotenv, как основных инструментов.
В качестве базы данных был использован Postgresql.

# Почему был выбран ORM?
На самом деле ORM крайне специфическая вещь и ее стоит обходить стороной при разработке чего-то крупного и 
высоконагруженного, так как в среднем ORM съедают довольно прилично перфоманса. Лучше использовать классические 
SQL-запросы. Да, в них можно ошибится или написать что-то с сверх тяжелое, особенно используя Join-ы, но у всего есть
цена. Единственные плюсы ORM-a, это возможность быстро сменить оперируемую БД посредством смены драйвера, а также
наличие нативной защиты от SQL-инъекций.

# Как запустить проект?
Все необходимые конфигурации файлы уже сформированы, ничего не требует вашего вмешательства.
Сначала необходимо запустить Docker и исполнить команду в терминале, находясь в директории проекта:
```shell
docker compose up -d
```
После того, как Docker поднимет под postgresql, исполняем следующую команду:
```go
go run main.go
```

# Прочее
В ТЗ не были описано требование в написание тест-кейсов, покрывающих функционал проекта.
Тест endpoint-ов производились вручную, посредством Postman-a.