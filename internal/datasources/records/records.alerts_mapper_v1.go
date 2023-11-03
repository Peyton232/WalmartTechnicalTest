package records

import (
	V1Domains "github.com/peyton232/walmart/internal/business/domains/v1"
)

func (u *Alerts) ToV1Domain() V1Domains.AlertDomain {
	return V1Domains.AlertDomain{
		AlertID:     u.AlertID,
		ServiceID:   u.ServiceID,
		ServiceName: u.ServiceName,
		Model:       u.Model,
		AlertType:   u.AlertType,
		AlertTS:     u.AlertTS,
		Severity:    u.Severity,
		TeamSlack:   u.TeamSlack,
		CreatedAt:   u.CreatedAt,
	}
}

func FromUsersV1Domain(u *V1Domains.AlertDomain) Alerts {
	return Alerts{
		AlertID:     u.AlertID,
		ServiceID:   u.ServiceID,
		ServiceName: u.ServiceName,
		Model:       u.Model,
		AlertType:   u.AlertType,
		AlertTS:     u.AlertTS,
		Severity:    u.Severity,
		TeamSlack:   u.TeamSlack,
		CreatedAt:   u.CreatedAt,
	}
}

func ToArrayOfUsersV1Domain(u *[]Alerts) []V1Domains.AlertDomain {
	var result []V1Domains.AlertDomain

	for _, val := range *u {
		result = append(result, val.ToV1Domain())
	}

	return result
}
