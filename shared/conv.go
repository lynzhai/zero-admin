package shared

import "github.com/globalsign/mgo/bson"

func BsonToStruct(m bson.M, value interface{}) error {
	bytes, err := bson.Marshal(m)
	if err != nil {
		return err
	}
	err = bson.Unmarshal(bytes, value)
	if err != nil {
		return err
	}
	return nil
}
