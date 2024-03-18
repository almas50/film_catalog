package handler

import (
	"FilmCatalog/internal"
	"FilmCatalog/internal/model"
	"FilmCatalog/internal/service"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"strconv"
)

const (
	actorsURL = "/api/actors"
	actorURL  = "/api/actors/{id}"
	filmsURL  = "/api/films"
	filmURL   = "/api/films/{id}"
	userURL   = "/api/sign-up"
)

type Handler struct {
	Service   service.Service
	Validator *validator.Validate
}

func (h *Handler) Register(router *http.ServeMux) {
	router.HandleFunc(http.MethodGet+" "+actorsURL, h.middlewareAuth(h.GetActors()))
	router.HandleFunc(http.MethodPost+" "+actorsURL, h.middlewareAuth(h.CreateActor()))
	router.HandleFunc(http.MethodGet+" "+actorURL, h.GetActor())
	router.HandleFunc(http.MethodPatch+" "+actorURL, h.UpdateActor())
	router.HandleFunc(http.MethodDelete+" "+actorURL, h.DeleteActor())

	router.HandleFunc(http.MethodPost+" "+filmsURL, h.CreateFilm())
	router.HandleFunc(http.MethodGet+" "+filmsURL, h.GetFilms())
	router.HandleFunc(http.MethodGet+" "+filmURL, h.GetFilm())
	router.HandleFunc(http.MethodDelete+" "+filmURL, h.DeleteFilm())
	router.HandleFunc(http.MethodPatch+" "+filmURL, h.UpdateFilm())

	router.HandleFunc(http.MethodPost+" "+userURL, h.CreateUser())
}

// @Summary sign-up
// @Tags user
// @ID sign-up
// @Produce  json
// @Param data body User true "user data"
// @Success 201
// @Failure 400 {string} message
// @Router /api/sign-up [post]
func (h *Handler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var user model.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Println("Invalid data ", err.Error(), user)
			http.Error(w, "Invalid data"+err.Error(), http.StatusBadRequest)
		}
		// Set for all new user role admin for debug
		user.Role = model.RoleAdmin
		err = h.Service.CreateUser(user)
		if err != nil {
			log.Println("can't create user ", err.Error(), user)
			http.Error(w, "can't sign-up"+err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusCreated)
	}
}

// @Summary get all actors
// @Tags actor
// @ID get_actors
// @Security BasicAuth
// @Produce  json
// @Success 200 {object} Actor
// @Failure 400 {string} message
// @Router /api/actors [get]
func (h *Handler) GetActors() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		actors_db, err := h.Service.GetActors()
		if err != nil {
			log.Println("can't get actors ", err.Error())
			http.Error(w, "can't get actors "+err.Error(), http.StatusBadRequest)
		}

		actors, err := json.Marshal(actors_db)
		if err != nil {
			log.Println("can't marshall actors ", err.Error())
			http.Error(w, "can't get actors "+err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(actors)
	}
}

// @Summary create actor
// @Tags actor
// @ID create_actor
// @Security BasicAuth
// @Produce  json
// @Param data body Actor true "actor data"
// @Success 201
// @Failure 400 {string} message
// @Router /api/actors [post]
func (h *Handler) CreateActor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var actor model.Actor
		err := json.NewDecoder(r.Body).Decode(&actor)
		if err != nil {
			log.Println("Invalid data ", err.Error(), actor)
			http.Error(w, "Invalid data"+err.Error(), http.StatusBadRequest)
		}

		err = h.Validator.Struct(actor)
		if err != nil {
			log.Println("Invalid data ", err.Error(), actor)
			http.Error(w, "Invalid data"+err.Error(), http.StatusBadRequest)
		}

		actorId, err := h.Service.CreateActor(actor)
		if err != nil {
			log.Println("can't create actor ", err.Error(), actor)
			http.Error(w, "can't create actor "+err.Error(), http.StatusBadRequest)
		}

		w.Header().Set("Location", fmt.Sprintf("%s/%d", actorsURL, actorId))
		w.WriteHeader(http.StatusCreated)
	}
}

// @Summary get actor by id
// @Tags actor
// @ID get_actor
// @Security BasicAuth
// @Produce  json
// @Param id path string true "actor id"
// @Success 200 {object} Actor
// @Failure 404 {string} message
// @Failure 400 {string} message
// @Router /api/actor [get]
func (h *Handler) GetActor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		actorId, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			log.Println("can't parse id", err.Error(), actorId)
			http.Error(w, "can't parse id "+err.Error(), http.StatusBadRequest)
		}

		actor_db, err := h.Service.GetActor(actorId)
		if err != nil {
			log.Println("can't get actor", err.Error(), actorId)
			http.Error(w, "can't get actor "+err.Error(), http.StatusNotFound)
		}

		actor, err := json.Marshal(actor_db)
		if err != nil {
			log.Println("can't marshall actor ", err.Error(), actor_db)
			http.Error(w, "can't get actor "+err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(actor)
	}
}

// @Summary update actor
// @Tags actor
// @ID update_actor
// @Security BasicAuth
// @Produce  json
// @Param data body Actor true "actor data"
// @Success 204
// @Failure 400 {string} message
// @Router /api/actor [patch]
func (h *Handler) UpdateActor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		actorId, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			log.Println("can't parse id", err.Error(), actorId)
			http.Error(w, "can't parse id "+err.Error(), http.StatusBadRequest)
		}

		var actor model.Actor
		err = json.NewDecoder(r.Body).Decode(&actor)
		if err != nil {
			log.Println("Invalid data ", err.Error(), actor)
			http.Error(w, "Invalid data "+err.Error(), http.StatusBadRequest)
		}

		err = h.Validator.Struct(actor)
		if err != nil {
			log.Println("Invalid data ", err.Error(), actor)
			http.Error(w, "Invalid data"+err.Error(), http.StatusBadRequest)
		}

		actor.ID = uint(actorId)
		err = h.Service.UpdateActor(actor)
		if err != nil {
			log.Println("can't update actor", err.Error(), actor)
			http.Error(w, "can't update actor "+err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// @Summary delete actor
// @Tags actor
// @ID delete_actor
// @Security BasicAuth
// @Produce  json
// @Param id path string true "actor id"
// @Success 204
// @Failure 400 {string} message
// @Router /api/actor [delete]
func (h *Handler) DeleteActor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		actorId, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			log.Println("can't parse id", err.Error(), actorId)
			http.Error(w, "can't parse id "+err.Error(), http.StatusBadRequest)
		}

		err = h.Service.DeleteActor(actorId)
		if err != nil {
			log.Println("can't delete actor", err.Error(), actorId)
			http.Error(w, "can't delete actor "+err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// @Summary get all films
// @Tags film
// @ID get_films
// @Security BasicAuth
// @Produce  json
// @Success 200 {object} Film
// @Failure 400 {string} message
// @Router /api/films [get]
func (h *Handler) GetFilms() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		sortByGet := r.URL.Query().Get("sortBy")
		sortBy := "rating"
		if internal.ValidateSort(sortByGet) {
			sortBy = sortByGet
		}
		// TODO: Validate search
		search := r.URL.Query().Get("search")

		films_db, err := h.Service.GetFilms(sortBy, search)
		if err != nil {
			log.Println("can't get films ", err.Error())
			http.Error(w, "can't get films "+err.Error(), http.StatusBadRequest)
		}

		films, err := json.Marshal(films_db)
		if err != nil {
			log.Println("can't marshall films ", err.Error())
			http.Error(w, "can't get films "+err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(films)
	}
}

// @Summary create film
// @Tags film
// @ID create_film
// @Security BasicAuth
// @Produce  json
// @Param data body CreateFilmDTO true "film data"
// @Success 201
// @Failure 400 {string} message
// @Router /api/films [post]
func (h *Handler) CreateFilm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var film model.CreateFilmDTO
		err := json.NewDecoder(r.Body).Decode(&film)
		if err != nil {
			log.Println("Invalid data ", err.Error(), film)
			http.Error(w, "Invalid data"+err.Error(), http.StatusBadRequest)
		}

		err = h.Validator.Struct(film)
		if err != nil {
			log.Println("Invalid data ", err.Error(), film)
			http.Error(w, "Invalid data"+err.Error(), http.StatusBadRequest)
		}

		res, err := h.Service.CreateFilm(film)
		if err != nil {
			log.Println("can't create film ", err.Error(), film)
			http.Error(w, "can't create film "+err.Error(), http.StatusBadRequest)
		}
		// TODO: fix path location
		w.Header().Set("Location", fmt.Sprintf("%s/%d", filmsURL, res))
		w.WriteHeader(http.StatusCreated)
	}
}

// @Summary update film
// @Tags film
// @ID update_film
// @Security BasicAuth
// @Produce  json
// @Param data body UpdateFilmDTO true "film data"
// @Success 204
// @Failure 400 {string} message
// @Router /api/film [patch]
func (h *Handler) UpdateFilm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var film model.UpdateFilmDTO
		err := json.NewDecoder(r.Body).Decode(&film)
		if err != nil {
			log.Println("Invalid data ", err.Error(), film)
			http.Error(w, "Invalid data"+err.Error(), http.StatusBadRequest)
		}

		err = h.Validator.Struct(film)
		if err != nil {
			log.Println("Invalid data ", err.Error(), film)
			http.Error(w, "Invalid data"+err.Error(), http.StatusBadRequest)
		}

		filmId, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			log.Println("can't parse id", err.Error(), filmId)
			http.Error(w, "can't parse id "+err.Error(), http.StatusBadRequest)
		}

		film.ID = uint(filmId)
		err = h.Service.UpdateFilm(film)
		if err != nil {
			log.Println("can't update film", err.Error(), film)
			http.Error(w, "can't update film "+err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// @Summary get film by id
// @Tags film
// @ID get_film
// @Security BasicAuth
// @Produce  json
// @Param id path string true "film id"
// @Success 200 {object} Film
// @Failure 404 {string} message
// @Failure 400 {string} message
// @Router /api/film [get]
func (h *Handler) GetFilm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		filmId, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			log.Println("can't parse id", err.Error(), filmId)
			http.Error(w, "can't parse id "+err.Error(), http.StatusBadRequest)
		}

		film_db, err := h.Service.GetFilm(filmId)
		if err != nil {
			log.Println("can't get film", err.Error(), film_db)
			http.Error(w, "can't get film "+err.Error(), http.StatusNotFound)
		}

		film, err := json.Marshal(film_db)
		if err != nil {
			log.Println("can't marshall film ", err.Error())
			http.Error(w, "can't get film "+err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(film)
	}
}

// @Summary delete film
// @Tags film
// @ID delete_film
// @Security BasicAuth
// @Produce  json
// @Param id path string true "film id"
// @Success 204
// @Failure 400 {string} message
// @Router /api/film [delete]
func (h *Handler) DeleteFilm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		filmId, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			log.Println("can't parse id", err.Error(), filmId)
			http.Error(w, "can't parse id "+err.Error(), http.StatusBadRequest)
		}

		err = h.Service.DeleteFilm(filmId)
		if err != nil {
			log.Println("can't delete film", err.Error(), filmId)
			http.Error(w, "can't delete film "+err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
