// Package wallpapercalc предоставляет функции для расчёта количества обоев.
package wallpapercalc

import (
	"fmt"
	"math"
)

// Perimeter вычисляет периметр комнаты.
// Возвращает ошибку, если длина или ширина <= 0.
func Perimeter(length, width float64) (float64, error) {
	if length <= 0 {
		return 0, fmt.Errorf("длина должна быть больше 0")
	}
	if width <= 0 {
		return 0, fmt.Errorf("ширина должна быть больше 0")
	}

	return 2 * (length + width), nil
}

// RollsNeeded вычисляет необходимое количество рулонов.
// perimeter — периметр комнаты
// height — высота стен
// rollWidth — ширина рулона
// rollLength — длина рулона
func RollsNeeded(perimeter, height, rollWidth, rollLength float64) (int, error) {
	if perimeter <= 0 || height <= 0 {
		return 0, fmt.Errorf("периметр и высота должны быть больше 0")
	}
	if rollWidth <= 0 || rollLength <= 0 {
		return 0, fmt.Errorf("размеры рулона должны быть больше 0")
	}

	// Общая площадь стен
	wallArea := perimeter * height

	// Площадь одного рулона
	rollArea := rollWidth * rollLength

	if rollArea <= 0 {
		return 0, fmt.Errorf("площадь рулона не может быть 0")
	}

	// Округляем вверх, чтобы рулонов точно хватило
	rolls := int(math.Ceil(wallArea / rollArea))

	return rolls, nil
}

// AddExtraRolls добавляет дополнительные рулоны (по указателю).
func AddExtraRolls(rolls *int, extra int) error {
	if rolls == nil {
		return fmt.Errorf("указатель на количество рулонов равен nil")
	}
	if extra < 0 {
		return fmt.Errorf("дополнительные рулоны не могут быть отрицательными")
	}

	*rolls += extra
	return nil
}

// FormatWallpaperReport формирует строку отчёта по обоям.
func FormatWallpaperReport(room string, rolls int, rollPrice float64) (string, error) {
	if room == "" {
		return "", fmt.Errorf("название комнаты не может быть пустым")
	}
	if rolls <= 0 {
		return "", fmt.Errorf("количество рулонов должно быть больше 0")
	}
	if rollPrice <= 0 {
		return "", fmt.Errorf("цена рулона должна быть больше 0")
	}

	totalCost := float64(rolls) * rollPrice

	report := fmt.Sprintf(
		"Комната: %s\nНеобходимо рулонов: %d\nЦена за рулон: %.2f руб.\nИтоговая стоимость: %.2f руб.",
		room,
		rolls,
		rollPrice,
		totalCost,
	)

	return report, nil
}
