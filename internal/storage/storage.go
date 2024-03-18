package storage

import "FilmCatalog/internal/model"

type Storage interface {
	CreateActor(model.Actor) (int64, error)
	GetActors() ([]model.Actor, error)
	GetActor(int) (model.Actor, error)
	CreateFilm(model.CreateFilmDTO) (int64, error)
	GetFilms(string, string) ([]model.Film, error)
	GetFilm(int) (model.Film, error)
	DeleteActor(int) error
	DeleteFilm(int) error
	UpdateActor(model.Actor) error
	UpdateFilm(model.UpdateFilmDTO) error

	CreateUser(model.User) error
	GetUser(string, string) (model.User, error)
}
