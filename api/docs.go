// Package api GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package api

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
        "/api/v1/distro/count": {
            "get": {
                "description": "List all users",
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "integer"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/distro/list": {
            "get": {
                "description": "List all users",
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/users/": {
            "post": {
                "description": "Create a New User",
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dbmodels.LinuxUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dbmodels.LinuxUser"
                        }
                    }
                }
            }
        },
        "/api/v1/users/list": {
            "get": {
                "description": "List all users",
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dbmodels.LinuxUser"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{id}": {
            "get": {
                "description": "Retrieves a User by ID",
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dbmodels.LinuxUser"
                        }
                    }
                }
            },
            "put": {
                "description": "Create a New User",
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dbmodels.LinuxUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dbmodels.LinuxUser"
                        }
                    }
                }
            },
            "delete": {
                "description": "Retrieves a User by ID",
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dbmodels.LinuxUser"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dbmodels.LinuxUser": {
            "type": "object",
            "properties": {
                "linuxDistro": {
                    "type": "string"
                },
                "linuxUserID": {
                    "type": "integer"
                },
                "linuxUserName": {
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
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Golug Demo Code",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
