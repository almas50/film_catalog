package postgres

import (
	"FilmCatalog/internal/model"
	"FilmCatalog/internal/storage"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var _ storage.Storage = &db{}

type db struct {
	Db *gorm.DB
}

func NewStorage() (storage.Storage, error) {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	postgres_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	postgres_db.AutoMigrate(&model.Actor{})
	postgres_db.AutoMigrate(&model.Film{})
	postgres_db.AutoMigrate(&model.User{})

	return &db{
		Db: postgres_db,
	}, nil
}

func (s *db) GetUser(username string, password string) (model.User, error) {
	var user model.User
	user = model.User{Name: username, Password: password}
	res := s.Db.Where("name = ? AND password = ?", username, password).Find(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) || res.RowsAffected == 0 {
		return user, fmt.Errorf("not found")
	}
	return user, res.Error
}

func (s *db) CreateUser(user model.User) error {
	res := s.Db.Create(&user)
	return res.Error
}

func (s *db) UpdateFilm(filmdto model.UpdateFilmDTO) error {
	var actors []model.Actor
	for _, actor := range filmdto.Actors {
		var a model.Actor
		s.Db.First(&a, actor)
		actors = append(actors, a)
	}
	f := model.Film{ID: filmdto.ID, Name: filmdto.Name, Description: filmdto.Description, Release: filmdto.Release, Rating: filmdto.Rating}
	res := s.Db.Model(&f).Updates(f)
	s.Db.Model(&f).Association("Actors").Replace(actors)

	return res.Error
}

func (s *db) CreateActor(actor model.Actor) (int64, error) {
	res := s.Db.Create(&actor)
	return res.RowsAffected, res.Error
}

func (s *db) GetActors() ([]model.Actor, error) {
	var actors []model.Actor
	s.Db.Preload("Films").Find(&actors)
	return actors, nil
}

func (s *db) GetActor(id int) (model.Actor, error) {
	var actor model.Actor
	s.Db.Preload("Films").First(&actor, id)
	return actor, nil
}

func (s *db) UpdateActor(actor model.Actor) error {
	s.Db.Model(&actor).Updates(model.Actor{Name: actor.Name, Gender: actor.Gender, Birthday: actor.Birthday})
	return nil
}

func (s *db) GetFilms(sortBy string, search string) ([]model.Film, error) {
	var films []model.Film
	s.Db.Preload("Actors").Order(sortBy+" desc").Find(&films, "name LIKE ?", "%"+search+"%")
	return films, nil
}

func (s *db) GetFilm(id int) (model.Film, error) {
	var film model.Film
	s.Db.Preload("Actors").First(&film, id)
	return film, nil
}

func (s *db) CreateFilm(film model.CreateFilmDTO) (int64, error) {
	var actors []model.Actor
	for _, actor := range film.Actors {
		var a model.Actor
		s.Db.First(&a, actor)
		actors = append(actors, a)
	}
	f := model.Film{Name: film.Name, Description: film.Description, Release: film.Release, Rating: film.Rating, Actors: actors}
	res := s.Db.Create(&f)
	return res.RowsAffected, res.Error
}

func (s *db) DeleteActor(id int) (err error) {
	s.Db.Delete(&model.Actor{}, id)
	return nil
}

func (s *db) DeleteFilm(id int) (err error) {
	s.Db.Model(&model.Film{ID: uint(id)}).Association("Actors").Clear()
	s.Db.Delete(&model.Film{}, id)
	return nil
}
