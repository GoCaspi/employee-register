package middleware

import (
	"example-project/service/servicefakes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetupService(t *testing.T) {
	fakeDb := servicefakes.FakeDatabaseInterface{}
	actualHandler := SetupService(&fakeDb)
	assert.NotNil(t, actualHandler)
}

func TestSetupEngine(t *testing.T) {
	fakeDb := servicefakes.FakeDatabaseInterface{}
	fakeHandlerFunc1 := SetupService(&fakeDb)
	fakeHandlerFunc2 := SetupService(&fakeDb)
	var fakeHandlerArr []gin.HandlerFunc
	fakeHandlerArr = []gin.HandlerFunc{fakeHandlerFunc1, fakeHandlerFunc2}

	actualHandler := SetupEngine(fakeHandlerArr)
	assert.NotNil(t, actualHandler)

}
