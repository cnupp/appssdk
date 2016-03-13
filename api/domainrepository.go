package api
import (
	"github.com/sjkyspa/stacks/controller/api/config"
	"github.com/sjkyspa/stacks/controller/api/net"
	"encoding/json"
	"fmt"
)

//go:generate counterfeiter -o fakes/fake_domain_repository.go . DomainRepository
type DomainRepository interface {
	Create(params DomainParams) (createdDomain Domain, apiErr error)
	GetDomain(name string) (Domain, error)
	GetDomains() (Domains, error)
	Delete(id string) (apiErr error)
}

type DefaultDomainRepository struct {
	config  config.Reader
	gateway net.Gateway
}

func NewDomainRepository(config config.Reader, gateway net.Gateway) DomainRepository {
	return DefaultDomainRepository{config: config, gateway: gateway}
}

func (cc DefaultDomainRepository) Create(params DomainParams) (createdDomain Domain, apiErr error) {
	data, err := json.Marshal(params)
	if err != nil {
		apiErr = fmt.Errorf("Can not serilize the data")
		return
	}

	res, err := cc.gateway.Request("POST", "/domains", data)
	if err != nil {
		apiErr = err
		return
	}

	location := res.Header.Get("Location")

	var domainModel DomainModel
	apiErr = cc.gateway.Get(location, &domainModel)
	if apiErr != nil {
		return
	}
	createdDomain = domainModel

	return
}

func (cc DefaultDomainRepository) GetDomain(name string) (app Domain, apiErr error) {
	var remoteDomain DomainModel
	apiErr = cc.gateway.Get(fmt.Sprintf("/domains/%s", name), &remoteDomain)
	if apiErr != nil {
		return
	}
	app = remoteDomain
	return
}


func (cc DefaultDomainRepository) GetDomains() (domains Domains, apiErr error) {
	var remoteDomains DomainsModel
	apiErr = cc.gateway.Get(fmt.Sprintf("/domains"), &remoteDomains)
	if apiErr != nil {
		return
	}
	remoteDomains.DomainMapper = cc
	domains = remoteDomains
	return
}


func (cc DefaultDomainRepository) Delete(id string) (apiErr error) {
	apiErr = cc.gateway.Delete(fmt.Sprintf("/domains/%s", id), "")
	if apiErr != nil {
		return
	}
	return
}
