package infra_connector

import "github.com/svc-bot-mds/terraform-provider-tdh/client/model"

type DataPlanesQuery struct {
	Name string `schema:"dataplaneName,omitempty"`
	model.PageQuery
}

type EligibleDataPlanesQuery struct {
	Provider          string `schema:"provider,omitempty"`
	InfraResourceType string `schema:"infraResourceType,omitempty"`
	OrgId             string `schema:"orgId,omitempty"`
	model.PageQuery
}
