# HamsterGen

## Описание

Простой в использовании генератор промо-кодов Hamster Kombat для получения ключей. В данный момент поддерживаются следующие игры:
- Bike Ride 3D (`BIKE`)
- Train Miner (`TRAIN`)
- Chain Cube 2048 (`CUBE`)
- My Clone Army (`CLONE`)
- Merge Away (`MERGE`)
- Twerk Race (`TWERK`)
- Polysphere (`POLY`)
- Mow and Trim (`TRIM`)
- Mud Race (`RACE`)
- Cafe Dash (`CAFE`)

## Использование

### Сборка и запуск

Компиляция исполняемого файла: `go build`

Запуск из исходного кода: `go run .`

### Параметры запуска

```
--app string
  Название игры (доступные варианты: BIKE | CUBE | CLONE | TRAIN | MERGE | TWERK | POLY | TRIM | RACE | CAFE)
--client-seed int
  Случайное значение для генерации клиентского ID
--token string
  Токен клиента (если доступен, иначе будет получен автоматически)
```

<details>
  <summary>Пример генерации промо-кода</summary>
Чтобы сгенерировать код для игры BIKE без указания токена, выполните следующую команду:

`go run . -app BIKE -client-seed 177013`
</details>

## Отказ от ответственности

HamsterGen предоставляется "как есть" без каких-либо гарантий. Авторы проекта не несут ответственности за любой ущерб, возникший в результате использования данного программного обеспечения.