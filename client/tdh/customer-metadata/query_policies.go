package customer_metadata

import "github.com/svc-bot-mds/terraform-provider-tdh/client/model"

type PoliciesQuery struct {
	Id         string   `schema:"id,omitempty"`
	Type       string   `schema:"serviceType"`
	Names      []string `schema:"name,omitempty"`
	ResourceId string   `schema:"resourceId,omitempty"`
	model.PageQuery
}