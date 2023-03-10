package patient

import (
	"github.com/graphql-go/graphql"
	pt_schemas "github.com/life-entify/patient/v1/graph/schemas"
	person_schemas "github.com/life-entify/person/v1/graph/schemas"
)

func CreatePatientWithMD(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Create Patient With Metadata",
		Type:        pt_schemas.PatientType,
		Args: graphql.FieldConfigArgument{
			"oldId": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"person_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"next_of_kins": &graphql.ArgumentConfig{
				Type: graphql.NewList(pt_schemas.NextOfKinMetaInputType),
			},
		},
		Resolve: resolver,
	}
}
func CreatePatient(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Create Patient",
		Type:        pt_schemas.PatientType,
		Args: graphql.FieldConfigArgument{
			"oldId": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"profile": &graphql.ArgumentConfig{
				Type: person_schemas.ProfileInputType,
			},
			"next_of_kins": &graphql.ArgumentConfig{
				Type: graphql.NewList(pt_schemas.NextOfKinInputType),
			},
		},
		Resolve: resolver,
	}
}
func UpdatePatientProfile(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Update Patient",
		Type:        pt_schemas.PatientType,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			//don't add profile or person details
			"patient": &graphql.ArgumentConfig{
				Type: pt_schemas.PatientInputType,
			},
			"person_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"profile": &graphql.ArgumentConfig{
				Type: person_schemas.ProfileInputType,
			},
		},
		Resolve: resolver,
	}
}
