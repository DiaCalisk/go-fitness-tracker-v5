//Пакет trainings содержит функционал по разбору строки с данными о тренировках и формирования строки с информацией о них.
package trainings

import 	(
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"time"
	"errors"
	"strconv"
	"strings"
	"fmt"
)

//Training содержит количество шагов, тип тренировки, длительность тренировки, имя, вес и рост пользователя.  
type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
	personaldata.Personal
}

//Parse парсит строку с данными и записывает данные в соответствующие поля структуры Training.
func (t *Training) Parse(datastring string) (err error) {

	splitedString := strings.Split(datastring, ",")

	if len(splitedString) != 3 {
		return errors.New("incorrect amount of transferred data")
	}

	steps, err := strconv.Atoi(splitedString[0])

	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("incorrect number of steps")
	}

	t.Steps = steps

	t.TrainingType = splitedString[1]

	duration, err := time.ParseDuration(splitedString[2])

	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("incorrect duration of activity")
	}

	t.Duration = duration

	return nil
}


//ActionInfo формирует и возвращает строку с данными о тренировке, исходя из того, какой тип тренировки был передан. 
func (t Training) ActionInfo() (string, error) {
	var spentCalories float64
	var err error
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	switch t.TrainingType {
	case "Бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "Ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "", errors.New("unknown type of activity")
	}

	if err != nil {
		return "", err
	}

	formatedString := fmt.Sprintf("Тип тренировки: %s\nДлительность: %0.2f ч.\nДистанция: %0.2f км.\nСкорость: %0.2f км/ч\nСожгли калорий: %0.2f\n", t.TrainingType, t.Duration.Hours(), distance, avgSpeed, spentCalories)
	return formatedString, nil
}
