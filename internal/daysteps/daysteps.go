package daysteps

import (
	"errors"
	"fmt"
	"log"
	"spentcalories"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	if data == "" {
		return 0, 0, errors.New("empty line")
	}
	strArr := strings.Split(data, ",")
	if len(strArr) != 2 {
		return 0, 0, errors.New("missed one or a few params")
	}
	stepCount, errStep := strconv.Atoi(strArr[0])
	if errStep != nil {
		return 0, 0, errStep
	}
	if stepCount <= 0 {
		return 0, 0, errors.New("lost element")
	}
	t, errTime := time.ParseDuration(strArr[1])
	if t <= 0 {
		return 0, 0, errors.New("the training's duration is less than or equal to zero")
	}
	if errTime != nil {
		return 0, 0, errTime
	}
	return stepCount, t, nil
}

func DayActionInfo(data string, weight, height float64) string {
	stepCount, walkDuration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if stepCount <= 0 {
		log.Println("steeps coun less or equal zero")
		return ""
	}
	distance := stepLength * float64(stepCount)
	distance /= mInKm
	lostCalories, errCalories := spentcalories.WalkingSpentCalories(stepCount, weight, height, walkDuration)
	if errCalories != nil {
		log.Println(errCalories)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", stepCount, distance, lostCalories)
}
