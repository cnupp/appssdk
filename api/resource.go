package api

import "github.com/sjkyspa/stacks/controller/api/config"

type Resource interface {
	GetResourceByURI(uri string) (interface{}, error)
}

type ResourceModel struct {
	config config.ConfigRepository
}

func (rm ResourceModel) GetResourceByURI(uri string) (interface{}, error) {
	return "", nil
}

func NewResource(cf config.ConfigRepository) Resource {
	return ResourceModel{
		config: cf,
	}
}
