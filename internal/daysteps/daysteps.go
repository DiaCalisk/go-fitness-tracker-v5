//Пакет daysteps содержит функционал по разбору строки с данными о дневной активности и формирования строки с информацией о ней.
package daysteps

import (
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"

	"time"
	"strings"
	"errors"
	"strconv"
	"fmt"
)

//DaySteps содержит количество шагов, длительность тренировки, имя, вес и рост пользователя.  
type DaySteps struct {
	personaldata.Personal
	Steps int
	Duration time.Duration
}

//Parse парсит строку с данными и записывает данные в соответствующие поля структуры DaySteps.
func (ds *DaySteps) Parse(datastring string) (err error) {

	splitedString := strings.Split(datastring, ",")

	if len(splitedString) != 2 {
		return errors.New("incorrect amount of transferred data")
	}

	steps, err := strconv.Atoi(splitedString[0])

	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("incorrect number of steps")
	}

	ds.Steps = steps

	duration, err := time.ParseDuration(splitedString[1])

	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("incorrect duration of activity")
	}

	ds.Duration = duration

	return nil
}

//ActionInfo формирует и возвращает строку с данными о тренировке. 
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)

	if err != nil {
		return "", err
	}

	formatedString := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %0.2f км.\nВы сожгли %0.2f ккал.\n", ds.Steps, distance, spentCalories )
	return formatedString, nil
}
