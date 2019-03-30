package NtmModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

// Salesreport Info
type Salesreport struct {
	ID         		string        `json:"-"`
	Id_        		bson.ObjectId `json:"-" bson:"_id"`
	ScenarioId 		string        `json:"-" bson:"scenario-id"`
	HospitalSalesReportIDs			[]string	`json:"-" bson:"hospital-sales-report-ids"`
	RepresentativeSalesReportIDs	[]string  	`json:"-" bson:"representative-sales-report-ids"`
	ProductSalesReportIDs			[]string	`json:"-" bson:"product-sales-report-ids"`

	HospitalSalesReport 		[]*Hospitalsalesreport			`json:"-"`
	RepresentativeSalesReport	[]*Representativesalesreport	`json:"-"`
	ProductSalesReport			[]*Productsalesreport			`json:"-"`

	Time 						float64 `json:"time" bson:"time"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Salesreport) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Salesreport) SetID(id string) error {
	c.ID = id
	return nil
}


// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u Salesreport) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "hospitalsalesreports",
			Name: "hospitalsalesreports",
		},
		{
			Type: "representativesalesreports",
			Name: "representativesalesreports",
		},
		{
			Type: "productsalesreports",
			Name: "productsalesreports",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u Salesreport) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	for _, kID := range u.HospitalSalesReportIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "hospitalsalesreports",
			Name: "hospitalsalesreports",
		})
	}


	for _, kID := range u.RepresentativeSalesReportIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "representativesalesreports",
			Name: "representativesalesreports",
		})
	}

	for _, kID := range u.ProductSalesReportIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "productsalesreports",
			Name: "productsalesreports",
		})
	}
	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u Salesreport) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	for key := range u.HospitalSalesReport {
		result = append(result, u.HospitalSalesReport[key])
	}

	for key := range u.RepresentativeSalesReport {
		result = append(result, u.RepresentativeSalesReport[key])
	}

	for key := range u.ProductSalesReport {
		result = append(result, u.ProductSalesReport[key])
	}

	return result
}

func (u *Salesreport) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "hospitalsalesreports" {
		u.HospitalSalesReportIDs = IDs
		return nil
	} else if name == "representativesalesreports" {
		u.RepresentativeSalesReportIDs = IDs
		return nil
	} else if name == "productsalesreports" {
		u.ProductSalesReportIDs = IDs
		return nil
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (u *Salesreport) AddToManyIDs(name string, IDs []string) error {
	if name == "hospitalsalesreports" {
		u.HospitalSalesReportIDs = append(u.HospitalSalesReportIDs, IDs...)
		return nil
	} else if name == "representativesalesreports" {
		u.RepresentativeSalesReportIDs = append(u.RepresentativeSalesReportIDs, IDs...)
		return nil
	} else if name == "productsalesreports" {
		u.ProductSalesReportIDs = append(u.ProductSalesReportIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}


func (u *Salesreport) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "ids":
			r := make(map[string]interface{})
			var ids []bson.ObjectId
			for i := 0; i < len(v); i++ {
				ids = append(ids, bson.ObjectIdHex(v[i]))
			}
			r["$in"] = ids
			rst["_id"] = r
		case "scenario-id":
			rst[k] = v[0]
		}
	}

	return rst
}