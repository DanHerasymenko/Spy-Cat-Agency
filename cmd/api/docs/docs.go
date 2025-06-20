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
        "/api/cats/create": {
            "post": {
                "description": "Create a new spy cat",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cats"
                ],
                "summary": "Create a new spy cat",
                "parameters": [
                    {
                        "description": "CreateCat request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CatCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Cat"
                        }
                    },
                    "400": {
                        "description": "Bad request",
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
        "/api/cats/list": {
            "get": {
                "description": "Retrieve a full list of all registered spy cats in the agency",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cats"
                ],
                "summary": "List all spy cats",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Cat"
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
        "/api/cats/{id}": {
            "get": {
                "description": "Retrieve a spy cat by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cats"
                ],
                "summary": "Get a spy cat",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Cat ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Cat"
                        }
                    },
                    "400": {
                        "description": "Bad request",
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
            },
            "delete": {
                "description": "Remove a spy cat by ID",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Cats"
                ],
                "summary": "Delete a spy cat",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Cat ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
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
        "/api/cats/{id}/salary": {
            "put": {
                "description": "Update salary of a spy cat by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cats"
                ],
                "summary": "Update cat salary",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Cat ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update salary request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CatUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Cat"
                        }
                    },
                    "400": {
                        "description": "Bad request",
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
        "/api/missions": {
            "get": {
                "description": "Get all created missions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Missions"
                ],
                "summary": "List all missions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Mission"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a mission with 1–3 targets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Missions"
                ],
                "summary": "Create a new mission",
                "parameters": [
                    {
                        "description": "Mission create body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.MissionCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Mission"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/missions/targets/{id}": {
            "put": {
                "description": "Update notes or completion status for a target",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Missions"
                ],
                "summary": "Update target",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Target ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Target update body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TargetUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Target"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a target by ID (only if target is not completed)",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Missions"
                ],
                "summary": "Delete target",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Target ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/missions/{id}": {
            "get": {
                "description": "Retrieve mission details, including assigned cat and targets",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Missions"
                ],
                "summary": "Get mission by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Mission ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Mission"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a mission by ID (only if not assigned to a cat)",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Missions"
                ],
                "summary": "Delete a mission",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Mission ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/missions/{id}/assign": {
            "post": {
                "description": "Assign a cat to a mission (1 cat per mission)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Missions"
                ],
                "summary": "Assign cat to mission",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Mission ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Cat assign body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CatAssign"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/missions/{id}/targets": {
            "post": {
                "description": "Add a target to an existing mission (only if mission is not completed)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Missions"
                ],
                "summary": "Add target to mission",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Mission ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Target create body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TargetCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Target"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Cat": {
            "type": "object",
            "properties": {
                "breed": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "salary": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                },
                "years_experience": {
                    "type": "integer"
                }
            }
        },
        "model.CatAssign": {
            "type": "object",
            "required": [
                "cat_id"
            ],
            "properties": {
                "cat_id": {
                    "type": "integer"
                }
            }
        },
        "model.CatCreate": {
            "type": "object",
            "required": [
                "breed",
                "name",
                "salary",
                "years_experience"
            ],
            "properties": {
                "breed": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "salary": {
                    "type": "number",
                    "minimum": 0
                },
                "years_experience": {
                    "type": "integer",
                    "minimum": 0
                }
            }
        },
        "model.CatUpdate": {
            "type": "object",
            "required": [
                "salary"
            ],
            "properties": {
                "salary": {
                    "type": "number",
                    "minimum": 0
                }
            }
        },
        "model.Mission": {
            "type": "object",
            "properties": {
                "cat": {
                    "$ref": "#/definitions/model.Cat"
                },
                "cat_id": {
                    "type": "integer"
                },
                "completed": {
                    "type": "boolean"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "targets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Target"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.MissionCreate": {
            "type": "object",
            "required": [
                "cat_id",
                "name",
                "targets"
            ],
            "properties": {
                "cat_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "targets": {
                    "type": "array",
                    "maxItems": 3,
                    "minItems": 1,
                    "items": {
                        "$ref": "#/definitions/model.TargetCreate"
                    }
                }
            }
        },
        "model.Target": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean"
                },
                "country": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "mission_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "notes": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.TargetCreate": {
            "type": "object",
            "required": [
                "country",
                "name"
            ],
            "properties": {
                "country": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "notes": {
                    "type": "string"
                }
            }
        },
        "model.TargetUpdate": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean"
                },
                "notes": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "SpyCat Agency API",
	Description:      "A spy cat management system API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
