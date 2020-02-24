package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf(" Книга %q была написана %s.", a.Title, a.Author)
}

type Book struct {
	Title string

	Author string
}

func (b Book) String() string {

	return fmt.Sprintf("Серия книг %q была создана %s.", b.Title, b.Author)

}

type Stringer interface {
	String() string
}

func main() {
	a := Article{
		Title:  "Эпицентр Удачи",
		Author: "Дмитрием Янковским",
	}
	Print(a)
	b := Book{

		Title:  "S.T.A.L.K.E.R.",
		Author: "всеми, у кого есть фантазия и любит игру S.T.A.L.K.E.R.,\n а так же читал братьев Стругацких",
	}

	Print(b)
}
func Print(s Stringer) {
	fmt.Println(s.String())
}
