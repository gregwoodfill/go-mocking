package method

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestMain(m *testing.M) {
	realDdbClient := ddbClient
	m.Run()
	ddbClient = realDdbClient
}

type MockDDBClient struct {
	mock.Mock
	// implement a common interface that is also the declared type in the production file
	// any methods called without a definition will panic
	dynamodbiface.DynamoDBAPI
}

// override the function you need to mock
func (m *MockDDBClient) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	args := m.Called(input)
	var err error
	if args.Error(1) != nil {
		err = args.Error(1)
	} else {
		err = nil
	}
	return args.Get(0).(*dynamodb.GetItemOutput), err
}

func TestReadDynamoDBItem(t *testing.T) {
	// set the package level ddb client to the mock one
	mockDdbClient := &MockDDBClient{}
	ddbClient = mockDdbClient
	key := "a key"
	expectedInput := &dynamodb.GetItemInput{Key: map[string]*dynamodb.AttributeValue{
		"pk": {S: &key},
	}}
	outputValue := "the value"
	output := &dynamodb.GetItemOutput{Item: map[string]*dynamodb.AttributeValue{
		"result": {S: &outputValue},
	}}
	mockDdbClient.On("GetItem", expectedInput).Return(output, nil)

	result, err := ReadDynamoDBItem(key)

	t.Run("it calls GetItem on the ddb client", func(t *testing.T) {
		mockDdbClient.AssertExpectations(t)
	})

	t.Run("it returns GetItemOutput from calling GetItem on ddbclient", func(t *testing.T) {
		assert.Equal(t, output, result)
	})

	t.Run("it does not return an error", func(t *testing.T) {
		assert.Nil(t, err)
	})
}
