WallpaperCalc

Описание проекта:
WallpaperCalc — это консольное приложение на языке Go для расчёта количества рулонов обоев и общей стоимости покупки.

Проект выполнен в рамках лабораторной работы (Variant 2).

Университет: Astana International University (AIU)
Автор: Damir Ravil
Год: 2026


Цель работы:
Реализовать пакет с функциями для расчёта:
- Периметра комнаты
- Количества рулонов
- Добавления запасных рулонов
- Формирования итогового отчёта

Также продемонстрировать:
- Работу с модулями Go
- Использование внешних библиотек
- Обработку ошибок
- Работу с указателями


Структура проекта:

lab4-variant2/
|
|-- go.mod
|-- go.sum
|-- main.go
|
|-- pkg/
    |-- wallpapercalc/
        |-- wallpapercalc.go


Реализованные функции:

1. Perimeter(length, width float64)
Вычисляет периметр:
P = 2 × (length + width)

2. RollsNeeded(perimeter, height, rollWidth, rollLength float64)
S = perimeter × height
rollArea = rollWidth × rollLength
rolls = ceil(S / rollArea)

3. AddExtraRolls(rolls *int, extra int)
Добавляет дополнительные рулоны через указатель.

4. FormatWallpaperReport(room string, rolls int, rollPrice float64)
Формирует текстовый отчёт.


Используемые библиотеки:
- github.com/google/uuid
- github.com/fatih/color
- github.com/dustin/go-humanize


Установка:

go mod init github.com/DamirRavil/lab4-variant2
go get github.com/google/uuid
go get github.com/fatih/color
go get github.com/dustin/go-humanize
go mod tidy
go run main.go


Вывод:
В ходе выполнения лабораторной работы был разработан модульный проект на языке Go с использованием внешних библиотек. Были закреплены навыки работы с функциями, указателями и обработкой ошибок.

Проект соответствует требованиям варианта 2.
