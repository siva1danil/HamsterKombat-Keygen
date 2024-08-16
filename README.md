# HamsterGen

## Описание

Простой в использовании генератор промо-кодов Hamster Kombat для получения ключей. В данный момент поддерживаются следующие игры:
- Bike Ride 3D (`BIKE`)
- Train Miner (`TRAIN`)
- Chain Cube 2048 (`CUBE`)
- My Clone Army (`CLONE`)
- Merge Away (`MERGE`)
- Twerk Race (`TWERK`)

## Использование

Компиляция: `go build`

Запуск без компиляции: `go run .`

Параметры:

```
  -app string
        App name (one of BIKE|CUBE|CLONE|TRAIN|MERGE|TWERK)
  -client-seed int
        Random seed used when generating client ID
  -token string
        Client token (if available)
```

<details>
  <summary>Пример генерации кода</summary>
  
  Для генерации `BIKE` кода можно использовать следующую команду:
  
  `go run . -app BIKE -client-seed 177013`
</details>

## Отказ от ответственности

HamsterGen предоставляется "как есть" без каких-либо гарантий. Авторы проекта не несут ответственности за любой ущерб, возникший в результате использования данного программного обеспечения.