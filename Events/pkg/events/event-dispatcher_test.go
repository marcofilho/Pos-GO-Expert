package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type EventTest struct {
	Name    string
	Payload interface{}
}

func (e *EventTest) GetName() string {
	return e.Name
}

func (e *EventTest) GetPayload() interface{} {
	return e.Payload
}

func (e *EventTest) GetDateTime() time.Time {
	return time.Now()
}

type EventTestHandler struct {
	ID int
}

func (e *EventTestHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event      EventTest
	event2     EventTest
	handler    EventTestHandler
	handler2   EventTestHandler
	handler3   EventTestHandler
	dispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.dispatcher = NewEventDispatcher()

	suite.handler = EventTestHandler{
		ID: 1,
	}

	suite.handler2 = EventTestHandler{
		ID: 2,
	}

	suite.handler3 = EventTestHandler{
		ID: 3,
	}

	suite.event = EventTest{
		Name:    "test",
		Payload: "test",
	}

	suite.event2 = EventTest{
		Name:    "test2",
		Payload: "test2",
	}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.dispatcher.handlers[suite.event.GetName()]))

	assert.Equal(suite.T(), &suite.handler, suite.dispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.dispatcher.handlers[suite.event.GetName()][1])

}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WhenHandlerAlreadyRegistered() {
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.dispatcher.handlers[suite.event.GetName()]))

	suite.dispatcher.Clear()
	suite.Equal(0, len(suite.dispatcher.handlers))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.dispatcher.handlers[suite.event.GetName()]))

	exists := suite.dispatcher.Has(suite.event.GetName(), &suite.handler)
	suite.True(exists)

	exists = suite.dispatcher.Has(suite.event.GetName(), &suite.handler2)
	suite.True(exists)

	exists = suite.dispatcher.Has(suite.event.GetName(), &suite.handler3)
	suite.False(exists)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_GetHandlers() {
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.dispatcher.handlers[suite.event.GetName()]))

	suite.dispatcher.Remove(suite.event.GetName(), &suite.handler)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))
	assert.Equal(suite.T(), &suite.handler2, suite.dispatcher.handlers[suite.event.GetName()][0])

	suite.dispatcher.Remove(suite.event.GetName(), &suite.handler2)
	suite.Equal(0, len(suite.dispatcher.handlers[suite.event.GetName()]))

	suite.dispatcher.Remove(suite.event.GetName(), &suite.handler3)
	suite.Equal(0, len(suite.dispatcher.handlers[suite.event.GetName()]))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eventHandler := &MockHandler{}
	eventHandler.On("Handle", &suite.event)

	eventHandler2 := &MockHandler{}
	eventHandler2.On("Handle", &suite.event)

	suite.dispatcher.Register(suite.event.GetName(), eventHandler)
	suite.dispatcher.Register(suite.event.GetName(), eventHandler2)

	suite.dispatcher.Dispatch(&suite.event)

	eventHandler.AssertExpectations(suite.T())
	eventHandler.AssertNumberOfCalls(suite.T(), "Handle", 1)

	eventHandler2.AssertExpectations(suite.T())
	eventHandler2.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func TestSuit(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
