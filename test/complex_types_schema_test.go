package test

import (
	"os"
	"path"
	"testing"

	"github.com/brenank/json-schema-to-go-struct-generator/pkg/converter"
	_ "github.com/brenank/json-schema-to-go-struct-generator/test/generated/complex-types"
	"github.com/stretchr/testify/assert"
)

//go:generate go run ../cmd/main.go --input ./samples/complex-types --output ./generated/complex-types/models.go

func TestDoesNotContainDuplicateStructs(t *testing.T) {
	pkg := GetPackageStructs("github.com/brenank/json-schema-to-go-struct-generator/test/generated/complex-types")

	assert.NotNil(t, pkg)
	assert.True(t, pkg.HasField("Bar1"))
	assert.True(t, pkg.HasField("Bar2"))
	assert.True(t, pkg.HasField("Bar10"))
	assert.True(t, pkg.HasField("Bar11"))
	assert.True(t, pkg.HasField("Bar12"))

	assert.True(t, pkg.HasField("Foo"))
	assert.True(t, pkg.HasFieldWithPrefix("Foo_"))

	assert.True(t, pkg.HasField("Person"))
	assert.True(t, pkg.HasFieldWithPrefix("Person_"))
}

func TestGenerate1(t *testing.T) {
	files := []string{path.Join(os.Getenv("PWD"), "./samples/complex-types/duplicate-structs-for-single-type-schema.json")}
	err := converter.Convert(files, "models", "./generated/complex-types/generate1/duplicate-structs.go", true)
	assert.Nil(t, err)
}

func TestGenerate2(t *testing.T) {
	files := []string{
		path.Join(os.Getenv("PWD"), "./samples/complex-types/duplicate-structs-without-creating-duplicate-models-schema.json"),
	}
	err := converter.Convert(files, "models", "./generated/complex-types/generate2/test2.go", true)
	assert.Nil(t, err)
}
