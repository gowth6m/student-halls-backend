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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/hall/all": {
            "get": {
                "description": "Get all halls",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "halls"
                ],
                "summary": "Get all halls",
                "responses": {
                    "200": {
                        "description": "Halls retrieved successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/hall.HallResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/hall/create": {
            "post": {
                "description": "Create a new hall",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "halls"
                ],
                "summary": "Create a new hall",
                "parameters": [
                    {
                        "description": "Hall object to be created",
                        "name": "hall",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hall.CreateHallRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Hall created successfully",
                        "schema": {
                            "$ref": "#/definitions/hall.HallResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/hall/{id}": {
            "get": {
                "description": "Get hall by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "halls"
                ],
                "summary": "Get hall by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hall ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Hall retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/hall.HallResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Hall not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/search": {
            "get": {
                "description": "Search for universities and halls",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search"
                ],
                "summary": "Search for universities and halls",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search query",
                        "name": "query",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Search results retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/search.SearchResult"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/university/all": {
            "get": {
                "description": "Get all universities",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "universities"
                ],
                "summary": "Get all universities",
                "responses": {
                    "200": {
                        "description": "Universities retrieved successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/university.UniversityResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/university/create": {
            "post": {
                "description": "Create a new university",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "universities"
                ],
                "summary": "Create a new university",
                "parameters": [
                    {
                        "description": "University object to be created",
                        "name": "university",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/university.CreateUniversityRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "University created successfully",
                        "schema": {
                            "$ref": "#/definitions/university.UniversityResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/university/{id}": {
            "get": {
                "description": "Get university by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "universities"
                ],
                "summary": "Get university by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "University ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "University retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/university.UniversityResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "University not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/university/{id}/halls": {
            "get": {
                "description": "Get halls by university ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "universities"
                ],
                "summary": "Get halls by university ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "University ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Halls retrieved successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/hall.Hall"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/user/all": {
            "get": {
                "description": "Get all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "Users retrieved successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/user.UserResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/user/create": {
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User object to be created",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User created successfully",
                        "schema": {
                            "$ref": "#/definitions/user.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/user/current": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get current user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get current user",
                "responses": {
                    "200": {
                        "description": "User retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/user.UserResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/user/email/{email}": {
            "get": {
                "description": "Get user by email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/user.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User login credentials",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User logged in successfully",
                        "schema": {
                            "$ref": "#/definitions/user.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Invalid credentials",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/user/{username}": {
            "get": {
                "description": "Get user by username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by username",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User username",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/user.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "hall.CreateHallRequest": {
            "type": "object",
            "required": [
                "coverImage",
                "location",
                "name",
                "postalCode",
                "reviews",
                "universityId"
            ],
            "properties": {
                "coverImage": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "postalCode": {
                    "type": "string"
                },
                "reviews": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "universityId": {
                    "type": "string"
                }
            }
        },
        "hall.Hall": {
            "type": "object",
            "required": [
                "coverImage",
                "id",
                "location",
                "name",
                "postalCode",
                "reviews",
                "university"
            ],
            "properties": {
                "coverImage": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "postalCode": {
                    "type": "string"
                },
                "reviews": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "university": {
                    "type": "string"
                }
            }
        },
        "hall.HallResponse": {
            "type": "object",
            "properties": {
                "coverImage": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "postalCode": {
                    "type": "string"
                },
                "reviews": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "university": {
                    "type": "string"
                }
            }
        },
        "search.SearchResult": {
            "type": "object",
            "properties": {
                "halls": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/hall.Hall"
                    }
                },
                "universities": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/university.University"
                    }
                }
            }
        },
        "university.CreateUniversityRequest": {
            "type": "object",
            "required": [
                "coverImage",
                "halls",
                "location",
                "name",
                "numberOfStudents",
                "website"
            ],
            "properties": {
                "coverImage": {
                    "type": "string"
                },
                "halls": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "numberOfStudents": {
                    "type": "integer"
                },
                "website": {
                    "type": "string"
                }
            }
        },
        "university.University": {
            "type": "object",
            "required": [
                "coverImage",
                "halls",
                "id",
                "location",
                "name",
                "numberOfStudents",
                "website"
            ],
            "properties": {
                "coverImage": {
                    "type": "string"
                },
                "halls": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "numberOfStudents": {
                    "type": "integer"
                },
                "website": {
                    "type": "string"
                }
            }
        },
        "university.UniversityResponse": {
            "type": "object",
            "properties": {
                "coverImage": {
                    "type": "string"
                },
                "halls": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "numberOfStudents": {
                    "type": "integer"
                },
                "website": {
                    "type": "string"
                }
            }
        },
        "user.CreateUserRequest": {
            "type": "object",
            "required": [
                "email",
                "firstName",
                "lastName",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "university": {
                    "type": "string"
                },
                "username": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 3
                },
                "yearOfStudy": {
                    "type": "integer",
                    "maximum": 5,
                    "minimum": 1
                }
            }
        },
        "user.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "user.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/user.UserResponse"
                }
            }
        },
        "user.UserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "university": {
                    "type": "string"
                },
                "userImg": {
                    "type": "string"
                },
                "userType": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                },
                "yearOfStudy": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "api.student-halls.com",
	BasePath:         "/v0",
	Schemes:          []string{"https"},
	Title:            "Student Halls API",
	Description:      "This is the REST API for Student Halls.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
