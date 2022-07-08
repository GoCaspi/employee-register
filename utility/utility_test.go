package utility

import (
	"crypto/sha256"
	"example-project/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestGetBearerToken(t *testing.T) {

	uuid := uuid.New()
	uuidString := uuid.String()

	responseRecoder := httptest.NewRecorder()
	fakeContest, _ := gin.CreateTestContext(responseRecoder)
	fakeContest.Request = httptest.NewRequest("GET", "http://localhost:9090/token", nil)
	fakeContest.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", uuidString))

	actualToken := GetBearerToken(fakeContest)
	assert.Equal(t, uuidString, actualToken)
}

func TestHashUsernameAndPassword(t *testing.T) {
	mockAuth := model.Auth{Password: "pa55word", Username: "Hacker"}

	hashedPassword := sha256.Sum256([]byte(mockAuth.Password))
	hashedUsername := sha256.Sum256([]byte(mockAuth.Username))

	actualU, actualP := HashUsernameAndPassword(mockAuth).Username, HashUsernameAndPassword(mockAuth).Password
	assert.Equal(t, actualP, hashedPassword)
	assert.Equal(t, actualU, hashedUsername)
}

func TestHashEmployees(t *testing.T) {

	mockEmps := []model.EmployeePayload{
		model.EmployeePayload{ID: "1", FirstName: "Tom", LastName: "Riddle", Email: "tomsnake.com", Auth: model.Auth{"volde", "mort"}},
		model.EmployeePayload{ID: "2", FirstName: "Tom", LastName: "Riddle", Email: "tomsnake.com", Auth: model.Auth{"volde", "mort"}},
		model.EmployeePayload{ID: "3", FirstName: "Tom", LastName: "Riddle", Email: "tomsnake.com", Auth: model.Auth{"volde", "mort"}},
	}
	var expectedArr []model.Employee
	for i := 0; i < len(mockEmps); i++ {
		apender := model.Employee{ID: fmt.Sprint(i + 1), FirstName: "Tom", LastName: "Riddle", Email: "tomsnake.com", Auth: HashUsernameAndPassword(mockEmps[i].Auth)}
		expectedArr = append(expectedArr, apender)
	}

	actual := HashEmployees(mockEmps)
	assert.Equal(t, expectedArr, actual)

}
