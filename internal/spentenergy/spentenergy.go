//Пакет spentenergy содержит функции для расчёта потраченных калорий при ходьбе, беге, а также функции для расчёта средней скорости и дистанции. 
package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)
//WalkingSpentCalories принимает количество шагов, рост и вес пользователя, а так же продолжительность активности и возвращает потраченные калории.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	switch {
	case weight <= 2:
		return 0, errors.New("incorrect weight")
	case height <= 0.5:
		return 0, errors.New("incorrect height")
	case duration <= 0:
		return 0, errors.New("incorrect duration of activity")
	case steps <= 0:
		return 0, errors.New("incorrect number of steps")
	}
	avgSpeed := MeanSpeed(steps, height, duration)
	return (avgSpeed * duration.Minutes() * weight * walkingCaloriesCoefficient) / minInH, nil
}

//RunningSpentCalories принимает количество шагов, рост и вес пользователя, а так же продолжительность активности и возвращает потраченные калории.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	switch {
	case weight <= 2:
		return 0, errors.New("incorrect weight")
	case height <= 0.5:
		return 0, errors.New("incorrect height")
	case duration <= 0:
		return 0, errors.New("incorrect duration of activity")
	case steps <= 0:
		return 0, errors.New("incorrect number of steps")
	}
	avgSpeed := MeanSpeed(steps, height, duration)
	return (avgSpeed * duration.Minutes() * weight) / minInH, nil
}

//MeanSpeed принимает количество шагов, рост пользователя и продолжительность активности и возвращает среднюю скорость.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

//Distance принимает количество шагов и рост пользователя в метрах, а возвращает дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	stepLenght := height * stepLengthCoefficient
	return stepLenght * float64(steps) / mInKm
}
