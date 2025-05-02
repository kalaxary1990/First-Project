package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	err1 := errors.New("data format error")
	err2 := errors.New("error of conversion step in pT")
	err3 := errors.New("error of conversion duration in pT")
	err4 := errors.New("stpcount not recognized")
	err5 := errors.New("durcount not recognized")
	dataslise := strings.Split(data, ",")

	if len(dataslise) != 3 {
		return 0, "", 0, err1
	}

	stpcount, err := strconv.Atoi(dataslise[0])
	if err != nil {
		return 0, "", 0, err2
	}
	durcount, err := time.ParseDuration(dataslise[2])
	if err != nil {
		return 0, "", 0, err3
	}

	if stpcount <= 0 {
		return 0, "", 0, err4
	}

	if durcount <= 0 {
		return 0, "", 0, err5
	}

	return stpcount, dataslise[1], durcount, nil

}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	if steps <= 0 || height <= 0 {
		return 0
	}
	stplength := height * stepLengthCoefficient
	return (float64(steps) * stplength) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 || steps <= 0 || height <= 0 {
		return 0
	}
	dist := distance(steps, height)
	return dist / duration.Hours()
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	err1 := errors.New("stpcount not recognized")
	err2 := errors.New("weight not recognized")
	err3 := errors.New("heght not recognized")
	err4 := errors.New("duration not recognized")
	if steps <= 0 {
		return 0, err1
	}
	if weight <= 0 {
		return 0, err2
	}
	if height <= 0 {
		return 0, err3
	}
	if duration <= 0 {
		return 0, err4
	}

	midspeed := meanSpeed(steps, height, duration)

	return (weight * midspeed * duration.Minutes()) / minInH, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	err1 := errors.New("stpcount not recognized")
	err2 := errors.New("weight not recognized")
	err3 := errors.New("heght not recognized")
	err4 := errors.New("duration not recognized")
	if steps <= 0 {
		return 0, err1
	}
	if weight <= 0 {
		return 0, err2
	}
	if height <= 0 {
		return 0, err3
	}
	if duration <= 0 {
		return 0, err4
	}

	midspeed := meanSpeed(steps, height, duration)
	return ((weight * midspeed * duration.Minutes()) / minInH) * walkingCaloriesCoefficient, nil
}
func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию

	stpcount, modtraining, durcount, _ := parseTraining(data)
	walkcalor, err1 := WalkingSpentCalories(stpcount, weight, height, durcount)
	runcalor, err2 := RunningSpentCalories(stpcount, weight, height, durcount)
	err3 := errors.New("неизвестный тип тренировки")
	switch modtraining {

	case "Ходьба":
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			modtraining, durcount.Hours(), distance(stpcount, height), meanSpeed(stpcount, height, durcount), walkcalor), err1

	case "Бег":
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			modtraining, durcount.Hours(), distance(stpcount, height), meanSpeed(stpcount, height, durcount), runcalor), err2

	default:
		return "", err3
	}

}
