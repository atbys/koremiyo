package user

type UserSession interface {
	Get(interface{}) interface{}
	Set(interface{}, interface{})
	Save() error
	Clear()
}
