package api

import (
	"github.com/sjkyspa/stacks/controller/api/config"
	"github.com/sjkyspa/stacks/controller/api/net"
	"regexp"
)

type Resource interface {
	GetResourceByURI(uri string) (interface{}, error)
}

type ResourceModel struct {
	config  config.Reader
	gateway net.Gateway
}

func (rm ResourceModel) GetResourceByURI(uri string) (interface{}, error) {
	regex, err := regexp.Compile(`/apps/[^/]*?$`)
	if err != nil {
		return "", err
	}
	if regex.MatchString(uri) {
		var app AppModel
		err := rm.gateway.Get(uri, &app)
		app.BuildMapper = NewBuildMapper(rm.config, rm.gateway)
		app.AppMapper = NewAppRepository(rm.config, rm.gateway)
		return app, err
	}

	var model BuildModel
	err = rm.gateway.Get(uri, &model)
	model.BuildMapper = NewBuildMapper(rm.config, rm.gateway)
	model.Resource = NewResource(rm.config, rm.gateway)
	return model, err
}

func NewResource(cf config.Reader, gw net.Gateway) Resource {
	return ResourceModel{
		config: cf,
		gateway: gw,
	}
}
