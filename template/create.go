package template

type CreateTemplate interface {
}

type createTemplate struct {
	Model string
}

func NewCreateTemplate() CreateTemplate {
	return &createTemplate{}
}

func (t *createTemplate) Repository() {

}
func (t *createTemplate) Service() {

}
func (t *createTemplate) Controller() {

}
