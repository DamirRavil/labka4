package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"github.com/google/uuid"

	"github.com/DamirRavil/labka4/pkg/wallpapercalc"
)

func main() {

	length := 5.0 // длина комнаты (м)
	width := 4.0  // ширина комнаты (м)
	height := 2.7 // высота стен (м)

	rollWidth := 1.06  // ширина рулона (м)
	rollLength := 10.0 // длина рулона (м)

	rollPrice := 3500.0
	extraRolls := 1 // запасной рулон

	roomID := uuid.New().String()

	// 1. Периметр
	perimeter, err := wallpapercalc.Perimeter(length, width)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// 2. Количество рулонов
	rolls, err := wallpapercalc.RollsNeeded(perimeter, height, rollWidth, rollLength)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// 3. Добавляем запас
	err = wallpapercalc.AddExtraRolls(&rolls, extraRolls)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// 4. Формируем отчёт
	report, err := wallpapercalc.FormatWallpaperReport(roomID, rolls, rollPrice)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	totalCost := float64(rolls) * rollPrice

	// Цветной вывод
	color.Cyan("=== ОТЧЁТ ПО ОБОЯМ ===")
	fmt.Println(report)

	// Форматированный вывод
	fmt.Printf("\nОбщая стоимость: %.2f руб.\n", totalCost)

	// humanize
	fmt.Println("Стоимость с разделителями:", humanize.Commaf(totalCost))
}
