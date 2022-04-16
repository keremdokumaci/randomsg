package messagegenerator

import (
	"encoding/json"
	"math"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/keremdokumaci/random-message-generator/app/helper"
)

type generator struct {
	MessageJsonRepresentation
}

func NewMessageGenerator(content string) generator {
	messageGenerator := generator{}
	err := json.Unmarshal([]byte(content), &messageGenerator)
	if err != nil {
		helper.ErrorText(err.Error())
	}

	return messageGenerator
}

func (g generator) GenerateMessage() string {
	message := make(map[string]interface{})
	for _, field := range g.Message.Fields {
		message[field.Name] = g.generateField(field)
	}

	marshaledMessage, err := json.Marshal(message)
	if err != nil {
		helper.ErrorText(err.Error())
	}

	return string(marshaledMessage)
}

func (g generator) generateField(field Field) interface{} {
	var generatedField interface{}

	switch field.Type {
	case String:
		generatedField = g.stringFieldValue(field.Rules)
		break
	case Int:
		generatedField = g.intFieldValue(field.Rules)
		break
	case Time:
		generatedField = g.timeFieldValue(field.Rules)
		break
	default:
		helper.ErrorText("Couldn't find the type.")
	}

	return generatedField
}

func (g generator) stringFieldValue(rules Rule) string {
	var value string
	if rules.Format == UUID {
		value = uuid.NewString()
	}

	return value
}

func (g generator) intFieldValue(rules Rule) int {
	var value int
	rand.Seed(time.Now().UnixNano())
	if rules.Min != nil && rules.Max != nil {
		max := int(rules.Max.(float64))
		min := int(rules.Min.(float64))
		value = rand.Intn(max-min) + min
	} else if rules.Min != nil && rules.Max == nil {
		min := int(rules.Min.(float64))
		value = rand.Intn(math.MaxInt-min) + min
	} else if rules.Min == nil && rules.Max != nil {
		max := int(rules.Max.(float64))
		value = rand.Intn(max-math.MinInt) + math.MinInt
	} else {
		value = rand.Int()
	}

	return value
}

func (g generator) timeFieldValue(rules Rule) time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
