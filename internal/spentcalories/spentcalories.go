package spentcalories

import (
	"errors"
	"fmt"
	"log"
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
	if data == "" {
		return 0, "", 0, errors.New("empty line")
	}
	strArr := strings.Split(data, ",")
	if len(strArr) != 3 {
		return 0, "", 0, errors.New("the number of steps is less than or equal to zero")
	}
	stepCount, errStep := strconv.Atoi(strArr[0])
	if errStep != nil {
		return 0, "", 0, errStep
	}
	if stepCount <= 0 {
		return 0, "", 0, errors.New("lost element")
	}
	t, errTime := time.ParseDuration(strArr[2])
	if t <= 0 {
		return 0, "", 0, errors.New("the training's duration is less than or equal to zero")
	}
	if errTime != nil {
		return 0, "", 0, errTime
	}
	action := strArr[1]
	return stepCount, action, t, nil

}

func distance(steps int, height float64) float64 {
	stepLength := stepLengthCoefficient * height
	distanceValue := stepLength * float64(steps)
	distanceKm := distanceValue / mInKm
	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distanceValue := distance(steps, height) / duration.Hours()
	return distanceValue
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	var calories float64
	var returnErr error
	if weight <= 0 {
		return "", errors.New("weight value is incorrect")
	}
	if height <= 0 {
		return "", errors.New("height value is incorrect")
	}
	if data == "" {
		return "", errors.New("empty line")
	}
	steps, activites, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	switch activites {
	case "Бег":
		{
			calories, returnErr = RunningSpentCalories(steps, weight, height, duration)
			if returnErr != nil {
				log.Println(returnErr)
				return "", returnErr
			}
		}
	case "Ходьба":
		{
			calories, returnErr = WalkingSpentCalories(steps, weight, height, duration)
			if returnErr != nil {
				log.Println(returnErr)
				return "", returnErr
			}
		}
	default:
		{
			return "", errors.New("неизвестный тип тренировки")
		}
	}
	speed := meanSpeed(steps, height, duration)
	interval := distance(steps, height)
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activites, duration.Hours(), interval, speed, calories), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps value is incorrect")
	}
	if weight <= 0 {
		return 0, errors.New("weight value is incorrect")
	}
	if height <= 0 {
		return 0, errors.New("height value is incorrect")
	}
	if duration <= 0 {
		return 0, errors.New("duration value is incorrect")
	}
	speed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * speed * durationInMinutes) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps value is incorrect")
	}
	if weight <= 0 {
		return 0, errors.New("weight value is incorrect")
	}
	if height <= 0 {
		return 0, errors.New("height value is incorrect")
	}
	if duration <= 0 {
		return 0, errors.New("duration value is incorrect")
	}
	speed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return ((weight * speed * durationInMinutes) / minInH) * walkingCaloriesCoefficient, nil
}
