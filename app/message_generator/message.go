package messagegenerator

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

type MessageJsonRepresentation struct {
	Message Message `json:"message"`
}

type Message struct {
	Fields []Field `json:"fields"`
}

type Field struct {
	Name  string    `json:"name"`
	Type  FieldType `json:"type"`
	Rules Rule      `json:"rules"`
}

type Rule struct {
	Format     FieldFormat `json:"format"`
	Min        interface{} `json:"min"`
	Max        interface{} `json:"max"`
	StartsWith string      `json"startsWith"`
	EndsWith   string      `json"endsWith"`
}
