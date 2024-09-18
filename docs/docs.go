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
        "/forcast/get/{lat}/{long}": {
            "get": {
                "description": "Accepts latitude and longitude coordinates then returns a short forcast+ temperature characterization",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "forcast"
                ],
                "summary": "Fetch a short forcast",
                "parameters": [
                    {
                        "type": "number",
                        "default": 39.7456,
                        "description": "latitude",
                        "name": "lat",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "number",
                        "default": -97.0892,
                        "description": "longitude",
                        "name": "long",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/data.Forcast"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/data.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/data.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/data.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "data.Forcast": {
            "type": "object",
            "properties": {
                "characterization": {
                    "type": "string"
                },
                "shortForcast": {
                    "type": "string"
                }
            }
        },
        "data.HTTPError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
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
	Title:            "National Weather Service API",
	Description:      "This is a simple API designed to call the official National Weather Service API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
