package usecases

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockedInternalObject struct {
	mock.Mock
}

func (m *MockedInternalObject) DoSomething(number int) (bool, error) {

	args := m.Called(number)
	return args.Bool(0), args.Error(1)

}

type RealInternalObject interface {
	DoSomething(number int) (bool, error)
}

type RealObject struct {
	RealInternalObject RealInternalObject
}

func (ro *RealObject) Do(number int) (bool, error) {
	return ro.RealInternalObject.DoSomething(number)
}

type TestSuite struct {
	suite.Suite
	MockedInternalObject *MockedInternalObject
	RealObject           *RealObject
}

func (suite *TestSuite) SetupTest() {
	suite.MockedInternalObject = new(MockedInternalObject)
	suite.RealObject = &RealObject{
		RealInternalObject: suite.MockedInternalObject,
	}
}

func (suite *TestSuite) TestExample() {
	suite.MockedInternalObject.On("DoSomething", 123).Return(true, nil)

	suite.RealObject.Do(123)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
