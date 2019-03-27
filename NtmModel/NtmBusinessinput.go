package NtmModel

import "gopkg.in/mgo.v2/bson"

type Businessinput struct {
	ID               string        `json:"-"`
	Id_              bson.ObjectId `json:"-" bson:"_id"`
	DestConfigId     string        `json:"dest-config-id" bson:"dest-config-id"`
	ResourceConfigId string        `json:"resource-config-id" bson:"resource-config-id"`
	SalesTarget      float64       `json:"sales-target" bson:"sales-target"`
	Budget           float64       `json:"budget" bson:"budget"`
	MeetingPlaces    float64       `json:"meeting-places" bson:"meeting-places"`
	VisitTime        float64       `json:"visit-time" bson:"visit-time"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Businessinput) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Businessinput) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *Businessinput) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	r := make(map[string]interface{})
	var ids []bson.ObjectId
	for k, v := range parameters {
		switch k {
		case "ids":
			for i := 0; i < len(v); i++ {
				ids = append(ids, bson.ObjectIdHex(v[i]))
			}
			r["$in"] = ids
			rst["_id"] = r
		case "dest-config-id":
			rst[k] = v[0]
		case "resource-config-id":
			rst[k] = v[0]
		}
	}
	return rst
}
