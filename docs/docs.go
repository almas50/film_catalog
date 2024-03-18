// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/actor": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "actor"
                ],
                "summary": "get actor by id",
                "operationId": "get_actor",
                "parameters": [
                    {
                        "type": "string",
                        "description": "actor id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/FilmCatalog.Actor"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "actor"
                ],
                "summary": "delete actor",
                "operationId": "delete_actor",
                "parameters": [
                    {
                        "type": "string",
                        "description": "actor id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "actor"
                ],
                "summary": "update actor",
                "operationId": "update_actor",
                "parameters": [
                    {
                        "description": "actor data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/FilmCatalog.Actor"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/actors": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "actor"
                ],
                "summary": "get all actors",
                "operationId": "get_actors",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/FilmCatalog.Actor"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "actor"
                ],
                "summary": "create actor",
                "operationId": "create_actor",
                "parameters": [
                    {
                        "description": "actor data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/FilmCatalog.Actor"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/film": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "film"
                ],
                "summary": "get film by id",
                "operationId": "get_film",
                "parameters": [
                    {
                        "type": "string",
                        "description": "film id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/FilmCatalog.Film"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "film"
                ],
                "summary": "delete film",
                "operationId": "delete_film",
                "parameters": [
                    {
                        "type": "string",
                        "description": "film id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "film"
                ],
                "summary": "update film",
                "operationId": "update_film",
                "parameters": [
                    {
                        "description": "film data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/FilmCatalog.UpdateFilmDTO"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/films": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "film"
                ],
                "summary": "get all films",
                "operationId": "get_films",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/FilmCatalog.Film"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "film"
                ],
                "summary": "create film",
                "operationId": "create_film",
                "parameters": [
                    {
                        "description": "film data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/FilmCatalog.CreateFilmDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/sign-up": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "sign-up",
                "operationId": "sign-up",
                "parameters": [
                    {
                        "description": "user data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/FilmCatalog.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "FilmCatalog.Actor": {
            "type": "object",
            "properties": {
                "birthday": {
                    "type": "string"
                },
                "films": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/FilmCatalog.Film"
                    }
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "FilmCatalog.CreateFilmDTO": {
            "type": "object",
            "properties": {
                "actors": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "description": {
                    "type": "string",
                    "maxLength": 1000
                },
                "name": {
                    "description": "sort and search",
                    "type": "string",
                    "maxLength": 150,
                    "minLength": 1
                },
                "rating": {
                    "description": "sort(default)",
                    "type": "integer",
                    "maximum": 10,
                    "minimum": 0
                },
                "release": {
                    "description": "sort",
                    "type": "string"
                }
            }
        },
        "FilmCatalog.Film": {
            "type": "object",
            "properties": {
                "actors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/FilmCatalog.Actor"
                    }
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "description": "sort and search",
                    "type": "string"
                },
                "rating": {
                    "description": "sort(default)",
                    "type": "integer"
                },
                "release": {
                    "description": "sort",
                    "type": "string"
                }
            }
        },
        "FilmCatalog.UpdateFilmDTO": {
            "type": "object",
            "properties": {
                "actors": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                },
                "release": {
                    "type": "string"
                }
            }
        },
        "FilmCatalog.User": {
            "type": "object"
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger FilmCatalog API",
	Description:      "Catalog of films and actors",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}