package shrd_utils

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const TEST_KEY = "TEST"
const TESTS_KEY = "TESTS"

type testArgs struct {
	Name string
	ID   int
}

func TestCacheKey(t *testing.T) {
	id := uuid.New()
	testArg := testArgs{
		ID:   1,
		Name: "Testing",
	}
	key := BuildCacheKey(TEST_KEY, id.String(), "funcName", testArg, testArg)

	assert.Equal(t, fmt.Sprintf("%s-%s-funcName|Name:%s,ID:%d|Name:%s,ID:%d", TEST_KEY, id, testArg.Name,
		testArg.ID, testArg.Name, testArg.ID), key)
}

func TestBuildPrefixKey(t *testing.T) {
	id := uuid.NewString()
	prefixKey := BuildPrefixKey(TEST_KEY, id, "TESTING")

	assert.Equal(t, fmt.Sprintf("%s-%s-TESTING", TEST_KEY, id), prefixKey)
}
