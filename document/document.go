package document

import "encoding/json"

type Document struct {
	id     uint
	fields map[string]interface{}
}

func (self *Document) CreateDocument(id uint) *Document {
	self.id = id
	self.fields = make(map[string]interface{})
	return self
}

func (self *Document) DeleteMemoryDocument() *Document {
	self.fields = nil
	self = nil
	return self
}

func (self *Document) GetDocument() *Document {
	return self
}

func (self *Document) AddField(name string, value interface{}) *Document {
	self.fields[name] = value
	return self
}

func (self *Document) ToJson() string {

	sdata, err := json.Marshal(self.fields)

	if err != nil {
		panic("Error ao converter para json")
	}

	return string(sdata)
}

func (self *Document) ToMap(jsonString string) map[string]interface{} {

	json.Unmarshal([]byte(jsonString), &self.fields)

	delete(self.fields, "_id.$oid")

	return self.fields
}

func (self *Document) GetMap() map[string]interface{} {
	return self.fields
}

func (self *Document) GetID() uint {
	return self.id
}
