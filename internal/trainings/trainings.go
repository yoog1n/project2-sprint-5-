package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	data := strings.Split(datastring, ",")

	if len(data) != 3 {
		return fmt.Errorf("неверный формат данных о тренировке")
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}

	duration, err := time.ParseDuration(data[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть положительной")
	}

	t.Steps = steps
	t.TrainingType = data[1]
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var (
		calories     float64
		err          error
		trainingName string
	)

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(
			t.Steps, t.Weight, t.Height, t.Duration,
		)
		trainingName = "Бег"

	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(
			t.Steps, t.Weight, t.Height, t.Duration,
		)
		trainingName = "Ходьба"

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Тип тренировки: %s\n"+"Длительность: %.2f ч.\n"+"Дистанция: %.2f км.\n"+"Скорость: %.2f км/ч\n"+"Сожгли калорий: %.2f\n",
		trainingName, t.Duration.Hours(), distance, speed, calories), nil
}
