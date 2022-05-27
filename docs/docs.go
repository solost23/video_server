// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/category/{user_name}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get user all category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Class"
                ],
                "summary": "get user all category",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "add category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Class"
                ],
                "summary": "create_class",
                "parameters": [
                    {
                        "description": "类别",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Class"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/category/{user_name}/{class_id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Class"
                ],
                "summary": "update_class",
                "parameters": [
                    {
                        "description": "类别",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Class"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/comment/{video_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get comment by video id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "get_comment_by_video_id",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "add comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "create comment",
                "parameters": [
                    {
                        "description": "评论",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Comment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/comment/{video_id}/{comment_id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "delete comment",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "user login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "login",
                "parameters": [
                    {
                        "description": "用户",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "add user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "register",
                "parameters": [
                    {
                        "description": "用户",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/role": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get all roleAuth",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "get all roleAuth",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "add roleAuth",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "add roleAuth",
                "parameters": [
                    {
                        "description": "角色",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CasbinModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete roleAuth",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "delete role",
                "parameters": [
                    {
                        "description": "角色",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CasbinModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/role/{role_name}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get roleAuth",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "get roleAuth",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get all user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get_all_user_info",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/user/{user_name}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get_user_info",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "update_user_info",
                "parameters": [
                    {
                        "description": "用户",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/video": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get all video",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Video"
                ],
                "summary": "get_all_video",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/video/{user_name}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get video by user_name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Video"
                ],
                "summary": "get_video_by_userName",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/video/{user_name}/{class_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get video by user_name and class_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Video"
                ],
                "summary": "get_video_by_userName_and_classID",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "add video",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Video"
                ],
                "summary": "add video",
                "parameters": [
                    {
                        "description": "视频",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Video"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/video/{user_name}/{class_id}/{video_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get video",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Video"
                ],
                "summary": "get video",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete video",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Video"
                ],
                "summary": "delete video",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CasbinModel": {
            "type": "object",
            "properties": {
                "method": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "ptype": {
                    "type": "string"
                },
                "role_name": {
                    "type": "string"
                }
            }
        },
        "model.Class": {
            "type": "object",
            "properties": {
                "createTime": {
                    "description": "DeleteStatus string ` + "`" + `gorm:\"type:enum('DELETE_STATUS_NORMAL','DELETE_STATUS_DEL');default:DELETE_STATUS_NORMAL\"` + "`" + `",
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "introduce": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updateTime": {
                    "type": "integer"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "model.Comment": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "createTime": {
                    "description": "DeleteStatus string ` + "`" + `gorm:\"delete_status;type:enum('DELETE_STATUS_NORMAL','DELETE_STATUS_DEL');default:DELETE_STATUS_NORMAL\"` + "`" + `",
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "isThumb": {
                    "type": "string"
                },
                "parentId": {
                    "type": "string"
                },
                "updateTime": {
                    "type": "integer"
                },
                "videoID": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "commentCount": {
                    "type": "integer"
                },
                "createTime": {
                    "description": "DeleteStatus string ` + "`" + `gorm:\"delete_status;type:enum('DELETE_STATUS_NORMAL','DELETE_STATUS_DEL');default:DELETE_STATUS_NORMAL\"` + "`" + `",
                    "type": "integer"
                },
                "fansCount": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "introduce": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updateTime": {
                    "type": "integer"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "model.Video": {
            "type": "object",
            "properties": {
                "classID": {
                    "type": "string"
                },
                "commentCount": {
                    "type": "integer"
                },
                "createTime": {
                    "type": "integer"
                },
                "deleteStatus": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "introduce": {
                    "type": "string"
                },
                "thumbCount": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updateTime": {
                    "type": "integer"
                },
                "userID": {
                    "type": "string"
                },
                "videoUrl": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "video_server Swagger",
	Description:      "this is a video server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
