package service

import (
	"FilmCatalog/internal/model"
	"FilmCatalog/internal/storage"
)

var _ Service = &service{}

type service struct {
	Storage storage.Storage
}

func NewService(Storage storage.Storage) (Service, error) {
	return &service{
		Storage,
	}, nil
}

type Service interface {
	CreateActor(actor model.Actor) (int64, error)
	GetActors() ([]model.Actor, error)
	GetActor(int) (model.Actor, error)
	CreateFilm(dto model.CreateFilmDTO) (int64, error)
	GetFilms(string, string) ([]model.Film, error)
	GetFilm(int) (model.Film, error)
	DeleteActor(int) error
	DeleteFilm(int) error
	UpdateActor(model.Actor) error
	UpdateFilm(dto model.UpdateFilmDTO) error

	CreateUser(model.User) error
	GetUser(string, string) (model.User, error)
}

func (s service) GetUser(username string, password string) (model.User, error) {
	user, err := s.Storage.GetUser(username, password)
	return user, err
}

func (s service) CreateUser(user model.User) error {
	err := s.Storage.CreateUser(user)
	return err
}

func (s service) UpdateFilm(dto model.UpdateFilmDTO) error {
	err := s.Storage.UpdateFilm(dto)
	return err
}

func (s service) CreateActor(actor model.Actor) (actorId int64, err error) {
	actorId, err = s.Storage.CreateActor(actor)
	return actorId, err
}

func (s service) GetActor(id int) (actor model.Actor, err error) {
	actor, err = s.Storage.GetActor(id)
	return actor, err
}

func (s service) UpdateActor(actor model.Actor) (err error) {
	err = s.Storage.UpdateActor(actor)
	return err
}

func (s service) CreateFilm(film model.CreateFilmDTO) (filmId int64, err error) {
	filmId, err = s.Storage.CreateFilm(film)
	return filmId, err
}

func (s service) GetActors() ([]model.Actor, error) {
	actors, err := s.Storage.GetActors()
	return actors, err
}

func (s service) GetFilms(sortBy string, search string) (films []model.Film, err error) {
	films, err = s.Storage.GetFilms(sortBy, search)
	return films, err
}

func (s service) GetFilm(id int) (film model.Film, err error) {
	film, err = s.Storage.GetFilm(id)
	return film, err
}

func (s service) DeleteActor(id int) (err error) {
	err = s.Storage.DeleteActor(id)
	return err
}

func (s service) DeleteFilm(id int) (err error) {
	err = s.Storage.DeleteFilm(id)
	return err
}
