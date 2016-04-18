package api

type Image struct {
	Name string `json:"image"`
	Mem  int `json:"mem"`
	Cpus float64 `json:"cpus"`
}

type Meta struct {
	VerifyImageField  Image `json:"verify"`
	BuildImageField   Image `json:"build"`
	TemplateCodeField string `json:"template"`
}

type ServiceDefinition struct {
	Build  Image `json:"build"`
	Verify Image `json:"verify"`
}

type Template struct {
	Type string `json:"type"`
	URI string `json:"uri"`
}

type StackStructure struct {

}

type Stack interface {
	Id() string
	Name() string
	Links() Links
	GetBuildImage() Image
	GetVerifyImage() Image
	GetTemplateCode() string
}

type StackModel struct {
	IDField             string `json:"id"`
	NameField           string `json:"name"`
	LinksField          []Link `json:"links"`
	Services map[string]ServiceDefinition `json:"services"`
	Template Template `json:"template"`
}

func (a StackModel) Id() string {
	return a.IDField
}

func (a StackModel) Name() string {
	return a.NameField
}

func (a StackModel) Links() Links {
	return LinksModel{
		Links: a.LinksField,
	}
}

func (a StackModel) GetBuildImage() Image {
	for _, v := range a.Services {
		if v.Build != (Image{}){
			return v.Build
		}
	}
	return Image{}
}

func (a StackModel) GetVerifyImage() Image {
	for _, v := range a.Services {
		if v.Verify !=  (Image{}){
			return v.Verify
		}
	}
	return Image{}
}

func (a StackModel) GetTemplateCode() string {
	return a.Template.URI
}

type Stacks interface {
	Count() int
	First() Stacks
	Last() Stacks
	Prev() Stacks
	Next() Stacks
	Items() []Stack
}

type StacksModel struct {
	CountField int            `json:"count"`
	SelfField  string         `json:"self"`
	FirstField string         `json:"first"`
	LastField  string         `json:"last"`
	PrevField  string         `json:"prev"`
	NextField  string         `json:"next"`
	ItemsField []StackModel   `json:"items"`
}

func (stacks StacksModel) Count() int {
	return stacks.CountField
}

func (stacks StacksModel) Self() Stacks {
	return nil
}
func (stacks StacksModel) First() Stacks {
	return nil
}
func (stacks StacksModel) Last() Stacks {
	return nil
}
func (stacks StacksModel) Prev() Stacks {
	return nil
}
func (stacks StacksModel) Next() Stacks {
	return nil
}

func (stacks StacksModel) Items() []Stack {
	items := make([]Stack, 0)
	for _, stack := range stacks.ItemsField {
		items = append(items, stack)
	}
	return items
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
