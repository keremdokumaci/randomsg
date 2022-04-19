package messagegenerator

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/keremdokumaci/randomsg/app/helper"
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
	sampleMessage := make(map[string]interface{})
	sampleMessage = g.Message.Sample

	for field, value := range g.Message.Sample {
		fieldType := reflect.TypeOf(value).Kind()
		fmt.Println(field, value, fieldType)
		if isNumeric(value) {
			sampleMessage[field] = g.numericFieldValue(Rule{})
		} else if fieldType == reflect.String {
			sampleMessage[field] = g.stringFieldValue(Rule{})
		} else if fieldType == reflect.Bool {
			sampleMessage[field] = g.boolFieldValue()
		} else {
			fmt.Println("Complex Type") // todo:
		}
	}

	marshaledMessage, err := json.Marshal(sampleMessage)
	if err != nil {
		helper.ErrorText(err.Error())
	}
	return string(marshaledMessage)
}

func (g generator) stringFieldValue(rules Rule) string {
	var value string = randStringRunes(rand.Intn(100)) // todo: 100 -> as rule (length range vs)

	if rules.Format == UUID {
		value = uuid.NewString()
	}
	if rules.StartsWith != "" {
		len := len(rules.StartsWith)
		value = value[len:]
		value = rules.StartsWith + value
	}
	return value
}

func (g generator) numericFieldValue(rules Rule) int {
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

func (g generator) boolFieldValue() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func isNumeric(val interface{}) bool {
	stringValue := fmt.Sprintf("%v", val)

	dotFound := false
	for _, v := range stringValue {
		if v == '.' {
			if dotFound {
				return false
			}
			dotFound = true
		} else if v < '0' || v > '9' {
			return false
		}
	}
	return true
}
