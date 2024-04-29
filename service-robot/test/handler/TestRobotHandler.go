package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	Handler "robot-fleet-monitoring/service-robot/handler"
	"robot-fleet-monitoring/service-robot/usecase"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TestSuite represents the test suite for the RobotHandler
type TestSuite struct {
	suite.Suite
	usecase *usecase.RobotUsecase
}

func (suite *TestSuite) TestRobotHandler_GetAll(t *testing.T) {

	// Setup
	RobotHandler := Handler.NewRobotHandler(suite.usecase)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/Robots", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test
	Robots := RobotHandler.GetAll(c)

	// Assert
	assert.NotNil(t, Robots)
	assert.NotEmpty(t, Robots)
}

func (suite *TestSuite) TestRobotHandler_GetById(t *testing.T) {

	// Setup
	RobotHandler := Handler.NewRobotHandler(suite.usecase)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/Robots/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test
	Robots := RobotHandler.GetById(c)

	// Assert
	assert.NotNil(t, Robots)
	assert.NotEmpty(t, Robots)
}

func (suite *TestSuite) TestRobotHandler_Create(t *testing.T) {

	// Setup
	RobotHandler := Handler.NewRobotHandler(suite.usecase)

	e := echo.New()
	requestBody := []byte(`{"name": "Milo", "price": 10000}`)
	req := httptest.NewRequest(http.MethodPost, "/Robots", bytes.NewBuffer(requestBody))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test
	Robots := RobotHandler.Create(c)

	// Assert
	assert.NotNil(t, Robots)
	assert.NotEmpty(t, Robots)
}

func (suite *TestSuite) TestRobotHandler_Update(t *testing.T) {

	// Setup
	RobotHandler := Handler.NewRobotHandler(suite.usecase)

	e := echo.New()
	requestBody := []byte(`{"name": "Milo", "price": 10000}`)
	req := httptest.NewRequest(http.MethodPost, "/Robots/1", bytes.NewBuffer(requestBody))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test
	Robots := RobotHandler.Update(c)

	// Assert
	assert.NotNil(t, Robots)
	assert.NotEmpty(t, Robots)
}

func (suite *TestSuite) TestRobotHandler_Delete(t *testing.T) {

	// Setup
	RobotHandler := Handler.NewRobotHandler(suite.usecase)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/Robots/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test
	Robots := RobotHandler.Delete(c)

	// Assert
	assert.NotNil(t, Robots)
	assert.NotEmpty(t, Robots)
}
