// Package docs GENERATED BY SWAG; DO NOT EDIT
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
        "/api/createApi": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysApi"
                ],
                "summary": "创建基础api",
                "parameters": [
                    {
                        "description": "api路径, api中文描述, api组, 方法",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/system.SysApi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "创建基础api",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/article/type/add": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ArticleType"
                ],
                "summary": "新增博客一级标题",
                "parameters": [
                    {
                        "description": "标题，图标",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/system.SysArticleType"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "新建标题",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/article/type/delete": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ArticleType"
                ],
                "summary": "删除博客一级标题",
                "parameters": [
                    {
                        "description": "标题，图标",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/system.SysArticleType"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "删除标题",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/article/type/edit": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ArticleType"
                ],
                "summary": "修改博客一级标题",
                "parameters": [
                    {
                        "description": "标题，图标",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/system.SysArticleType"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "修改标题",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/article/type/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ArticleType"
                ],
                "summary": "获取博客一级标题",
                "parameters": [
                    {
                        "description": "标题，图标",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/system.SysArticleType"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "修改标题",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/base/captcha": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "生成验证码",
                "responses": {
                    "200": {
                        "description": "生成验证码,返回包括随机数id,base64,验证码长度",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.SysCaptchaResponse"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/base/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户名, 密码, 验证码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回包括用户信息,token,过期时间",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.LoginResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.LoginResponse"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysUser"
                ],
                "summary": "用户注册账号",
                "parameters": [
                    {
                        "description": "用户米，昵称，密码，角色ID",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Register"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户注册账号，返回包括用户信息",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.SysUserResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.SysUserResponse"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.Login": {
            "type": "object",
            "properties": {
                "captcha": {
                    "type": "string"
                },
                "captchaId": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.Register": {
            "type": "object",
            "properties": {
                "authorityIds": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "authorityid": {
                    "type": "integer"
                },
                "enable": {
                    "type": "integer"
                },
                "headerImg": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "response.LoginResponse": {
            "type": "object",
            "properties": {
                "expiresAt": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/system.SysUser"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "response.SysCaptchaResponse": {
            "type": "object",
            "properties": {
                "captchaId": {
                    "type": "string"
                },
                "captchaLength": {
                    "type": "integer"
                },
                "picPath": {
                    "type": "string"
                }
            }
        },
        "response.SysUserResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/system.SysUser"
                }
            }
        },
        "system.SysApi": {
            "type": "object",
            "properties": {
                "apiGroup": {
                    "description": "api组",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "description": {
                    "description": "api中文描述",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "method": {
                    "description": "方法:创建POST(默认)|查看GET|更新PUT|删除DELETE",
                    "type": "string"
                },
                "path": {
                    "description": "api路径",
                    "type": "string"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "system.SysArticleType": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "system.SysAuthority": {
            "type": "object",
            "properties": {
                "authorityId": {
                    "type": "integer"
                },
                "authorityName": {
                    "type": "string"
                },
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysAuthority"
                    }
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "dataAuthorityId": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysAuthority"
                    }
                },
                "default": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "menus": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysBaseMenu"
                    }
                },
                "parentId": {
                    "type": "integer"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "system.SysBaseMenu": {
            "type": "object",
            "properties": {
                "authoritys": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysAuthority"
                    }
                },
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysBaseMenu"
                    }
                },
                "closeTab": {
                    "type": "boolean"
                },
                "component": {
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "defaultMenu": {
                    "type": "boolean"
                },
                "hidden": {
                    "type": "boolean"
                },
                "icon": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "keepAlive": {
                    "type": "boolean"
                },
                "menuBtn": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysBaseMenuBtn"
                    }
                },
                "name": {
                    "type": "string"
                },
                "parameters": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysBaseMenuParameter"
                    }
                },
                "parentId": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "sort": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "system.SysBaseMenuBtn": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "sysBaseMenuId": {
                    "type": "integer"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "system.SysBaseMenuParameter": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "key": {
                    "type": "string"
                },
                "sysBaseMenuID": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "system.SysUser": {
            "type": "object",
            "properties": {
                "authorities": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysAuthority"
                    }
                },
                "authority": {
                    "$ref": "#/definitions/system.SysAuthority"
                },
                "authorityId": {
                    "description": "用户角色ID",
                    "type": "integer"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "enable": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "nickname": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                },
                "username": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "博客系统 Swagger",
	Description:      "Go 语言编程之旅：一起用 Go 做项目",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
