package unicreds

import (
	"encoding/json"
	"testing"

	. "github.com/Versent/unicreds"
	"github.com/stretchr/testify/assert"
)

func TestJsonExpand(t *testing.T) {
	var err error

	originRawJson := []byte(`{
		"k1" : "v1",
		"k2" : "${prefix/a.b.c}",
		"k3" : {
			"k31": "${prefix/a.1}",
			"k32": "${prefix/a.2}",
			"k33": "${other_prefix/a.3}"
		}
  }`)

	expectedRawJson := []byte(`{
		"k1" : "v1",
		"k2" : "a.b.c",
		"k3" : {
			"k31": "a.1",
			"k22": "a.2",
			"k33": "${other_prefix/a.3}"
		}
  }`)

	var expectedJson interface{}
	err = json.Unmarshal(expectedRawJson, &expectedJson)
	assert.Nil(t, err)

	resultRawJson, err := JsonExpand(originRawJson, "prefix", func(key string) (string, error) { return key, nil })
	assert.Nil(t, err)

	var resultJson interface{}
	err = json.Unmarshal(resultRawJson, &resultJson)
	assert.Nil(t, err)

	assert.NotEqual(t, 0, len(resultRawJson))
	assert.Equal(t, expectedJson, resultJson)
}
