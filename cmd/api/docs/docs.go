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
        "/admin/createadmin": {
            "post": {
                "description": "admin creation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "Create a new admin from admin panel",
                "operationId": "AdminSignup",
                "parameters": [
                    {
                        "description": "New Admin details",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.AdminLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/admin/login": {
            "post": {
                "description": "Logs in an admin user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "Admin login",
                "parameters": [
                    {
                        "description": "Admin login data",
                        "name": "adminLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.AdminLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/doctor/categories": {
            "get": {
                "description": "List available categories for doctors.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Doctor"
                ],
                "summary": "List Doctor Categories",
                "responses": {
                    "200": {
                        "description": "Categories listed successfully",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    },
                    "400": {
                        "description": "Failed to list categories",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/doctor/login": {
            "post": {
                "description": "Logs in a doctor.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Doctor Authentication"
                ],
                "summary": "Doctor Login",
                "parameters": [
                    {
                        "description": "Doctor Login Request",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.DoctorLogin"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Doctor login successful",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request or login failure",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/doctor/profile": {
            "get": {
                "description": "Get the profile of a doctor.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Doctor"
                ],
                "summary": "Get Doctor Profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Doctor profile retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request or profile retrieval failure",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/doctor/{categoryid}/registration": {
            "post": {
                "description": "Register a doctor.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Doctor Authentication"
                ],
                "summary": "Doctor Registration",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "categoryid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Doctor Registration Data",
                        "name": "registrationData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.DoctorRegistration"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Registration accepted",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request or registration failure",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Logs in a user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "description": "User Login Request",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully logged in",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request or login failure",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "get": {
                "description": "Logs out a user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "Logout user from the app",
                "responses": {
                    "200": {
                        "description": "Logged out successfully",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/user/signup": {
            "post": {
                "description": "Register a new user account with the provided details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "Register a new user account",
                "parameters": [
                    {
                        "description": "User registration request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.UserRegister"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "req.AdminLogin": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "req.DoctorLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "req.DoctorRegistration": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "license_number": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "specialise": {
                    "type": "string"
                }
            }
        },
        "req.UserLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "req.UserRegister": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "res.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {},
                "message": {
                    "type": "string"
                },
                "stastus_code": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
