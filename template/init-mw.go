package template

type InitMWTemplate interface {
}

type initMwTemplate struct {
}

func NewInitMwTemplate() InitMWTemplate {
	return &initMwTemplate{}
}

func (t *initMwTemplate) Middleware() {

}

func (t *initMwTemplate) AuthService() {

}
