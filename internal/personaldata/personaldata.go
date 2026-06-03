// Пакет personaldata содержит структуру с данными пользователя и функцию с выводом данных о пользователе.
package personaldata

import "fmt"

// Personal содержит данные о пользователе.
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print выводит данные о пользователе.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %0.2f\nРост: %0.2f\n", p.Name, p.Weight, p.Height)
}
