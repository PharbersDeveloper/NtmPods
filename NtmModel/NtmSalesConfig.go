package NtmModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

// SalesConfig Info
type SalesConfig struct {
	ID         		string        `json:"-"`
	Id_        		bson.ObjectId `json:"-" bson:"_id"`
	ScenarioId 		string        `json:"-" bson:"scenario-id"`
	DestConfigID     		string        `json:"-" bson:"dest-config-id"`
	GoodsConfigID			string 	  	  `json:"-" bson:"goods-config-id"`
	DestConfig 		*DestConfig   `json:"-"`
	GoodsConfig 	*GoodsConfig  `json:"-"`
	AccessStatus   	string  `json:"access-status" bson:"access-status"`
	LastYearSales 	float64 `json:"last-year-sales" bson:"last-year-sales"`
	Potential     	float64 `json:"potential" bson:"potential"`
	//ReportID		string	`json"-" bson:"report-id"` //预留字段
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c SalesConfig) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *SalesConfig) SetID(id string) error {
	c.ID = id
	return nil
}


// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u SalesConfig) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "destConfigs",
			Name: "destConfig",
		},
		{
			Type: "goodsConfigs",
			Name: "goodsConfig",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u SalesConfig) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}
	if u.DestConfigID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.DestConfigID,
			Type: "destConfigs",
			Name: "destConfig",
		})
	}

	if u.GoodsConfigID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.GoodsConfigID,
			Type: "goodsConfigs",
			Name: "goodsConfig",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u SalesConfig) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if u.DestConfigID != "" && u.DestConfig != nil {
		result = append(result, u.DestConfig)
	}

	if u.GoodsConfigID != "" && u.GoodsConfig != nil {
		result = append(result, u.GoodsConfig)
	}

	return result
}

func (u *SalesConfig) SetToOneReferenceID(name, ID string) error {
	if name == "DestConfig" {
		u.DestConfigID = ID
		return nil
	}
	if name == "goodsConfig" {
		u.GoodsConfigID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (u *SalesConfig) GetConditionsBsonM(parameters map[string][]string) bson.M {
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