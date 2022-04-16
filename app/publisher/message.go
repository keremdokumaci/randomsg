package publisher

type FieldType string
type FieldFormat string

const (
	String FieldType = "string"
	Int              = "int"
	Float            = "float"
	Time             = "time"
)

const (
	UUID FieldFormat = "uuid"
)

type Message struct {
	Fields []Field
}

type Field struct {
	Name  string    `json:"name"`
	Type  FieldType `json:"type"`
	Rules []Rule
}

type Rule struct {
	Format             FieldFormat `json:"format"`
	GreaterThan        interface{} `json:"gt"`
	GreaterOrEqualThan interface{} `json:"gte"`
}
