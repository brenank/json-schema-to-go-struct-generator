package test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	_ "github.com/azarc-io/json-schema-to-go-struct-generator/test/generated/aliases"
)

//go:generate go run ../cmd/main.go --input ./samples/aliases --output ./generated/aliases/model.go

func TestAliases(t *testing.T) {
	pkg := GetPackageStructs("github.com/azarc-io/json-schema-to-go-struct-generator/test/generated/aliases")

	assert.NotNil(t, pkg)
	assert.True(t, pkg.HasField("Foo1_Foo2_Foo3"))
	assert.True(t, pkg.HasField("Foo1"))
	assert.True(t, pkg.HasField("Foo2"))
	assert.True(t, pkg.HasField("Foo3"))
}
