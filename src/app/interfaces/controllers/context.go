package controllers

type Context interface {
	Param(key string) string
	BindJSON(v interface{}) error
	JSON(code int, obj interface{})
}
