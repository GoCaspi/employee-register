package cache

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddToCacheMap(t *testing.T) {

	var MyCacheMap = NewCacheMap{}
	uuid := uuid.New()
	uuidString := uuid.String()

	actual := AddToCacheMap("1", uuidString, MyCacheMap)
	assert.Equal(t, actual, MyCacheMap)
}

func TestRemoveFromCacheMap(t *testing.T) {
	var MyCacheMap = NewCacheMap{}

	var tests = []struct {
		Map             NewCacheMap
		idHasTokenInMap bool
	}{
		{MyCacheMap, false},
		{MyCacheMap, true},
	}

	for _, tt := range tests {
		if !tt.idHasTokenInMap {
			actual := RemoveFromCacheMap("1", MyCacheMap)
			assert.Equal(t, actual, MyCacheMap)
		} else {
			uuid := uuid.New()
			uuidString := uuid.String()
			newMap := AddToCacheMap("1", uuidString, MyCacheMap)
			actual := RemoveFromCacheMap("1", newMap)
			assert.Equal(t, actual, MyCacheMap)
		}
	}

}

func TestGiveIdToToken(t *testing.T) {
	var MyCacheMap = NewCacheMap{}
	var fakeToken = "abc123"
	var tests = []struct {
		Map          NewCacheMap
		tokenIsInMap bool
		token        string
	}{
		{MyCacheMap, false, fakeToken},
		{MyCacheMap, true, fakeToken},
	}
	for _, tt := range tests {
		if tt.tokenIsInMap {
			AddToCacheMap("1", tt.token, tt.Map)
			actual := GiveIdToToken(tt.token, tt.Map)
			assert.Equal(t, "1", actual)
		} else {
			RemoveFromCacheMap("1", tt.Map)
			actual := GiveIdToToken(tt.token, tt.Map)
			assert.Equal(t, "", actual)
		}
	}

}

func TestIdIsInMap(t *testing.T) {
	var MyCacheMap = NewCacheMap{}
	var fakeToken = "abc123"
	var tests = []struct {
		idInMap bool
		id      string
		Map     NewCacheMap
	}{
		{false, "1", MyCacheMap},
		{true, "1", MyCacheMap},
	}
	for _, tt := range tests {
		if !tt.idInMap {
			actual := IdIsInMap(tt.id, tt.Map)
			assert.Equal(t, false, actual)
		} else {
			AddToCacheMap(tt.id, fakeToken, tt.Map)
			actual := IdIsInMap(tt.id, tt.Map)
			assert.Equal(t, true, actual)
		}
	}
}

func TestTokenIsInMap(t *testing.T) {
	var MyCacheMap = NewCacheMap{}
	//	var myMap map[string]string
	var fakeToken = "abc123"

	var tests = []struct {
		isInMap bool
		Map     map[string]string
	}{
		{true, MyCacheMap},
		{false, MyCacheMap},
	}

	for _, tt := range tests {
		if tt.isInMap {
			AddToCacheMap("1", fakeToken, tt.Map)
			actual := TokenIsInMap(fakeToken, tt.Map)
			assert.Equal(t, true, actual)
		} else {
			RemoveFromCacheMap("1", tt.Map)
			actual := TokenIsInMap(fakeToken, tt.Map)
			assert.Equal(t, false, actual)
		}
	}
}
