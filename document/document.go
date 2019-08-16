package document

import "encoding/json"

type Document struct {
	id     uint
	fields map[string]interface{}
}

func (doc *Document) CreateDocument(id uint) *Document {
	doc.id = id
	doc.fields = make(map[string]interface{})
	doc.fields["_key"] = id
	return doc
}

func (doc *Document) AddField(name string, value interface{}) *Document {
	doc.fields[name] = value
	return doc
}

func (doc *Document) ToJson() string {

	sdata, err := json.Marshal(doc.fields)

	if err != nil {
		panic("Error ao converter para json")
	}

	return string(sdata)
}

func (doc *Document) ToMap(jsonString string) map[string]interface{} {

	json.Unmarshal([]byte(jsonString), doc.fields)

	return doc.fields
}

func (doc *Document) GetMap() map[string]interface{} {
	return doc.fields
}

func (doc *Document) GetID() uint {
	return doc.id
}
