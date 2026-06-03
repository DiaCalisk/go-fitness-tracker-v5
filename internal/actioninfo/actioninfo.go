// Пакет actioninfo реализует вывод общей информации обо всех тренировках и прогулках.
package actioninfo

import (
	"fmt"
	"log"
)

// DataParser это интерфейс для обращения к методам структур Training или DaySteps.
type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

// Info принимает слайс строк с данными о тренировках или прогулках, после чего выводит информацию о них.
func Info(dataset []string, dp DataParser) {
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			log.Println(err)
			continue
		}
		actionString, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(actionString)
	}
}
