// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Space VPX WebSite Backend Service",
    "title": "Space VPX Backend Service",
    "version": "development"
  },
  "paths": {
    "/auth/register": {
      "post": {
        "tags": [
          "Auth"
        ],
        "summary": "Register User",
        "operationId": "RegisterUser",
        "parameters": [
          {
            "description": "Register User Body",
            "name": "RegisterUser",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/RegisterUser"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Register Response",
            "schema": {
              "$ref": "#/definitions/LoginResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "403": {
            "description": "Forbidden",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "422": {
            "description": "Unprocessable Entity",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "LoginResponse": {
      "type": "object",
      "properties": {
        "access_token": {
          "type": "string"
        },
        "refresh_token": {
          "type": "string"
        }
      }
    },
    "RegisterUser": {
      "type": "object",
      "required": [
        "email",
        "name",
        "password",
        "again_password"
      ],
      "properties": {
        "again_password": {
          "type": "string"
        },
        "email": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Space VPX WebSite Backend Service",
    "title": "Space VPX Backend Service",
    "version": "development"
  },
  "paths": {
    "/auth/register": {
      "post": {
        "tags": [
          "Auth"
        ],
        "summary": "Register User",
        "operationId": "RegisterUser",
        "parameters": [
          {
            "description": "Register User Body",
            "name": "RegisterUser",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/RegisterUser"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Register Response",
            "schema": {
              "$ref": "#/definitions/LoginResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "403": {
            "description": "Forbidden",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "422": {
            "description": "Unprocessable Entity",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "LoginResponse": {
      "type": "object",
      "properties": {
        "access_token": {
          "type": "string"
        },
        "refresh_token": {
          "type": "string"
        }
      }
    },
    "RegisterUser": {
      "type": "object",
      "required": [
        "email",
        "name",
        "password",
        "again_password"
      ],
      "properties": {
        "again_password": {
          "type": "string"
        },
        "email": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
}
