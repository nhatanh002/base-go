package application

import (
	"base-go/application/cats"
)

type App struct {
	Cats cats.CatsInteractor
}

func NewApp(cats cats.CatsInteractor) *App {
	return &App{
		Cats: cats,
	}
}
