package users

type IBusinessLogic interface {
	GetUser() string
}

type BusinessLogic struct {
}

func (bl *BusinessLogic) GetUser() string {
	return "Rohit"
}

func NewBusinessLogic() *BusinessLogic {
	return &BusinessLogic{}
}
