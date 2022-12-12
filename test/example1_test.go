package test

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"github.com/brenank/json-schema-to-go-struct-generator/pkg/converter"
	"github.com/stretchr/testify/assert"

	example1 "github.com/brenank/json-schema-to-go-struct-generator/test/generated/example1"
)

//go:generate go run ../cmd/main.go --input ./samples/example1 --output ./generated/example1/model.go

func TestMarshalValidateSuccess(t *testing.T) {
	param := struct {
		Name           string
		Data           string
		ExpectedResult bool
	}{
		Name: "Blue Sky",
		Data: `{
				"id": 1,
				"name": "Unbridled Optimism 2.0",
				"price": 99.99,
				"tags": [ "happy" ] }`,
		ExpectedResult: true,
	}

	prod := &example1.Product{}
	err := json.Unmarshal([]byte(param.Data), &prod)
	assert.Nil(t, err)
	assert.Nil(t, prod.Validate())
}

func TestMarshalValidateFail(t *testing.T) {
	param := struct {
		Name           string
		Data           string
		ExpectedResult bool
	}{
		Name: "Missing Price",
		Data: `{
				"id": 1,
				"name": "Unbridled Optimism 2.0",
				"tags": [ "happy" ] }`,
		ExpectedResult: false,
	}

	prod := &example1.Product{}
	err := json.Unmarshal([]byte(param.Data), &prod)
	assert.Nil(t, err)
	errs := prod.Validate()
	assert.Equal(t, 1, len(errs))
	assert.ErrorIs(t, errs[0], example1.ErrFieldRequired)
}

func TestGenerateExample(t *testing.T) {
	files := []string{
		path.Join(os.Getenv("PWD"), "./samples/example1/example1a.json"),
	}
	err := converter.Convert(files, "models", "./generated/example1/generate-example/all.go", true)
	assert.Nil(t, err)
}
