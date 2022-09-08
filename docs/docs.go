// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/cart/": {
            "get": {
                "description": "獲得Cart",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cart"
                ],
                "summary": "Get Cart",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Cart"
                        }
                    }
                }
            },
            "put": {
                "description": "更新Cart",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cart"
                ],
                "summary": "Update Cart",
                "parameters": [
                    {
                        "description": "要更新的cart",
                        "name": "cart",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Cart"
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/api/product/": {
            "get": {
                "description": "獲得products list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Get Product list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Product"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Cart": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.OrderProduct"
                    }
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "model.OrderProduct": {
            "type": "object",
            "properties": {
                "cartID": {
                    "type": "integer"
                },
                "product": {
                    "type": "object",
                    "$ref": "#/definitions/model.Product"
                },
                "productID": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "model.Product": {
            "type": "object",
            "properties": {
                "inventory": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.0.1",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{"http"},
	Title:       "Swagger API",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
