package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	err1 := errors.New("data format error")
	err2 := errors.New("error of conversion step in pP")
	err3 := errors.New("stpcount not recognized")
	err4 := errors.New("error of conversion duration in pP")
	err5 := errors.New("durcount not recognized")
	dataslise := strings.Split(data, ",")
	if len(dataslise) != 2 {
		return 0, 0, err1
	}

	stpcount, err := strconv.Atoi(dataslise[0])
	if err != nil {

		return 0, 0, err2
	}
	if stpcount <= 0 {
		return 0, 0, err3
	}

	durcount, err := time.ParseDuration(dataslise[1])
	if err != nil {
		return 0, 0, err4
	}
	if durcount <= 0 {
		return 0, 0, err5
	}

	return stpcount, durcount, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функци

	stpcount, durcount, err1 := parsePackage(data)

	if err1 != nil {
		log.Println(err1)
		return ""
	}

	distant := (float64(stpcount) * stepLength) / float64(mInKm)

	calorcount, _ := spentcalories.WalkingSpentCalories(stpcount, weight, height, durcount)

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		stpcount, distant, calorcount)

}
