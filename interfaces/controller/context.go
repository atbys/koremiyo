package controller

type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
	HTML(int, string, interface{})
	String(int, string, ...interface{})
}
