// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://sumelms.com/docs/terms",
        "contact": {
            "name": "LMS Support",
            "url": "https://sumelms.com/support",
            "email": "support@sumelms.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/courses": {
            "get": {
                "description": "List a new courses",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "List courses",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoints.ListCoursesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "description": "Create a new course",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "Create course",
                "parameters": [
                    {
                        "description": "Add Course",
                        "name": "course",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/endpoints.CreateCourseRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoints.CreateCourseResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/courses/{uuid}": {
            "get": {
                "description": "Find a new course",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "Find course",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Course UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoints.FindCourseResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "put": {
                "description": "Update a course",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "Update course",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Course UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Course",
                        "name": "course",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/endpoints.UpdateCourseRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoints.UpdateCourseResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "description": "Delete a new course",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "Delete course",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Course UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoints.DeleteCourseResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/matrices": {
            "get": {
                "description": "List a new matrices",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "matrices"
                ],
                "summary": "List matrices",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "course search by uuid",
                        "name": "course_uuid",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoints.ListMatricesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "description": "Create a new matrix",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "matrices"
                ],
                "summary": "Create matrix",
                "parameters": [
                    {
                        "description": "Add Matrix",
                        "name": "matrices",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/endpoints.CreateMatrixRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoints.CreateMatrixResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/matrices/{uuid}": {
            "get": {
                "description": "Find a new matrix",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "matrices"
                ],
                "summary": "Find matrix",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Matrix UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoints.FindMatrixResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "put": {
                "description": "Update a matrix",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "matrices"
                ],
                "summary": "Update matrix",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Matrix UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Matrix",
                        "name": "matrix",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/endpoints.UpdateMatrixRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoints.UpdateMatrixResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "description": "Delete a new matrix",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "matrices"
                ],
                "summary": "Delete matrix",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Matrix UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoints.DeleteMatrixResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/subscriptions": {
            "get": {
                "description": "List a new subscriptions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscriptions"
                ],
                "summary": "List subscriptions",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "course search by uuid",
                        "name": "course_uuid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "user search by uuid",
                        "name": "user_uuid",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoints.ListSubscriptionsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "description": "Create a new subscription",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscriptions"
                ],
                "summary": "Create subscription",
                "parameters": [
                    {
                        "description": "Add Subscription",
                        "name": "subscription",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/endpoints.CreateSubscriptionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoints.CreateSubscriptionResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/subscriptions/{uuid}": {
            "get": {
                "description": "Find a new subscription",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscriptions"
                ],
                "summary": "Find subscription",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Subscription UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoints.FindSubscriptionResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "put": {
                "description": "Update a subscription",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscriptions"
                ],
                "summary": "Update subscription",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Subscription UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Subscription",
                        "name": "subscription",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/endpoints.UpdateSubscriptionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoints.UpdateSubscriptionResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "description": "Delete a new subscription",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscriptions"
                ],
                "summary": "Delete subscription",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Subscription UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoints.DeleteSubscriptionResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "endpoints.CourseResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "excerpt": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "image_cover": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "underline": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "endpoints.CreateCourseRequest": {
            "type": "object",
            "required": [
                "code",
                "description",
                "excerpt",
                "name",
                "underline"
            ],
            "properties": {
                "code": {
                    "type": "string",
                    "maxLength": 15
                },
                "description": {
                    "type": "string",
                    "maxLength": 255
                },
                "excerpt": {
                    "type": "string",
                    "maxLength": 140
                },
                "image": {
                    "type": "string"
                },
                "image_cover": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 100
                },
                "underline": {
                    "type": "string",
                    "maxLength": 100
                }
            }
        },
        "endpoints.CreateCourseResponse": {
            "type": "object",
            "properties": {
                "course": {
                    "$ref": "#/definitions/endpoints.CourseResponse"
                }
            }
        },
        "endpoints.CreateMatrixRequest": {
            "type": "object",
            "required": [
                "course_uuid",
                "name"
            ],
            "properties": {
                "code": {
                    "type": "string",
                    "maxLength": 45
                },
                "course_uuid": {
                    "type": "string"
                },
                "description": {
                    "type": "string",
                    "maxLength": 255
                },
                "name": {
                    "type": "string",
                    "maxLength": 100
                }
            }
        },
        "endpoints.CreateMatrixResponse": {
            "type": "object",
            "properties": {
                "matrix": {
                    "$ref": "#/definitions/endpoints.MatrixResponse"
                }
            }
        },
        "endpoints.CreateSubscriptionRequest": {
            "type": "object",
            "required": [
                "course_uuid",
                "role",
                "user_uuid"
            ],
            "properties": {
                "course_uuid": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "string"
                },
                "matrix_uuid": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "user_uuid": {
                    "type": "string"
                }
            }
        },
        "endpoints.CreateSubscriptionResponse": {
            "type": "object",
            "properties": {
                "subscription": {
                    "$ref": "#/definitions/endpoints.SubscriptionResponse"
                }
            }
        },
        "endpoints.DeleteCourseResponse": {
            "type": "object",
            "properties": {
                "course": {
                    "$ref": "#/definitions/endpoints.DeletedCourseResponse"
                }
            }
        },
        "endpoints.DeleteMatrixResponse": {
            "type": "object",
            "properties": {
                "matrix": {
                    "$ref": "#/definitions/endpoints.DeletedMatrixResponse"
                }
            }
        },
        "endpoints.DeleteSubscriptionResponse": {
            "type": "object",
            "properties": {
                "subscription": {
                    "$ref": "#/definitions/endpoints.DeletedSubscriptionResponse"
                }
            }
        },
        "endpoints.DeletedCourseResponse": {
            "type": "object",
            "properties": {
                "deleted_at": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "endpoints.DeletedMatrixResponse": {
            "type": "object",
            "properties": {
                "deleted_at": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "endpoints.DeletedSubscriptionResponse": {
            "type": "object",
            "properties": {
                "deleted_at": {
                    "type": "string"
                },
                "reason": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "endpoints.FindCourseResponse": {
            "type": "object",
            "properties": {
                "course": {
                    "$ref": "#/definitions/endpoints.CourseResponse"
                }
            }
        },
        "endpoints.FindMatrixResponse": {
            "type": "object",
            "properties": {
                "matrix": {
                    "$ref": "#/definitions/endpoints.MatrixResponse"
                }
            }
        },
        "endpoints.FindSubscriptionResponse": {
            "type": "object",
            "properties": {
                "subscription": {
                    "$ref": "#/definitions/endpoints.SubscriptionResponse"
                }
            }
        },
        "endpoints.ListCoursesResponse": {
            "type": "object",
            "properties": {
                "courses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/endpoints.CourseResponse"
                    }
                }
            }
        },
        "endpoints.ListMatricesResponse": {
            "type": "object",
            "properties": {
                "matrices": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/endpoints.MatrixResponse"
                    }
                }
            }
        },
        "endpoints.ListSubscriptionsResponse": {
            "type": "object",
            "properties": {
                "subscriptions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/endpoints.SubscriptionsResponse"
                    }
                }
            }
        },
        "endpoints.MatrixResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "course_uuid": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "endpoints.SubscriptionCourseResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "endpoints.SubscriptionMatrixResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "endpoints.SubscriptionResponse": {
            "type": "object",
            "properties": {
                "course_uuid": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "string"
                },
                "matrix_uuid": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_uuid": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "endpoints.SubscriptionsResponse": {
            "type": "object",
            "properties": {
                "course": {
                    "$ref": "#/definitions/endpoints.SubscriptionCourseResponse"
                },
                "course_uuid": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "string"
                },
                "matrix": {
                    "$ref": "#/definitions/endpoints.SubscriptionMatrixResponse"
                },
                "matrix_uuid": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_uuid": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "endpoints.UpdateCourseRequest": {
            "type": "object",
            "required": [
                "code",
                "description",
                "excerpt",
                "name",
                "underline",
                "uuid"
            ],
            "properties": {
                "code": {
                    "type": "string",
                    "maxLength": 100
                },
                "description": {
                    "type": "string",
                    "maxLength": 255
                },
                "excerpt": {
                    "type": "string",
                    "maxLength": 140
                },
                "image": {
                    "type": "string"
                },
                "image_cover": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 100
                },
                "underline": {
                    "type": "string",
                    "maxLength": 100
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "endpoints.UpdateCourseResponse": {
            "type": "object",
            "properties": {
                "course": {
                    "$ref": "#/definitions/endpoints.CourseResponse"
                }
            }
        },
        "endpoints.UpdateMatrixRequest": {
            "type": "object",
            "required": [
                "name",
                "uuid"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "description": {
                    "type": "string",
                    "maxLength": 255
                },
                "name": {
                    "type": "string",
                    "maxLength": 100
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "endpoints.UpdateMatrixResponse": {
            "type": "object",
            "properties": {
                "matrix": {
                    "$ref": "#/definitions/endpoints.MatrixResponse"
                }
            }
        },
        "endpoints.UpdateSubscriptionRequest": {
            "type": "object",
            "required": [
                "role",
                "uuid"
            ],
            "properties": {
                "expires_at": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "endpoints.UpdateSubscriptionResponse": {
            "type": "object",
            "properties": {
                "subscription": {
                    "$ref": "#/definitions/endpoints.SubscriptionResponse"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Description for what is this security definition being used",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        },
        "OAuth2AccessCode": {
            "type": "oauth2",
            "flow": "accessCode",
            "authorizationUrl": "https://sso.sumelms.com/oauth/authorize",
            "tokenUrl": "https://sso.sumelms.com/oauth/token",
            "scopes": {
                "admin": "Grants read and write access to administrative information"
            }
        },
        "OAuth2Application": {
            "type": "oauth2",
            "flow": "application",
            "tokenUrl": "https://sso.sumelms.com/oauth/token",
            "scopes": {
                "admin": "Grants read and write access to administrative information",
                "write": "Grants write access"
            }
        },
        "OAuth2Implicit": {
            "type": "oauth2",
            "flow": "implicit",
            "authorizationUrl": "https://sso.sumelms.com/oauth/authorize",
            "scopes": {
                "admin": "Grants read and write access to administrative information",
                "write": "Grants write access"
            }
        },
        "OAuth2Password": {
            "type": "oauth2",
            "flow": "password",
            "tokenUrl": "https://sso.sumelms.com/oauth/token",
            "scopes": {
                "admin": "Grants read and write access to administrative information",
                "read": "Grants read access",
                "write": "Grants write access"
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Sumé LMS Course API",
	Description:      "This is the Sumé LMS API for Course Microservice",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
