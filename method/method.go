package method

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"log"
)

var sess = session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
}))

// we want to mock ddbclient
var ddbClient dynamodbiface.DynamoDBAPI = dynamodb.New(sess)

func ReadDynamoDBItem(key string) (*dynamodb.GetItemOutput, error) {
	item := dynamodb.GetItemInput{Key: map[string]*dynamodb.AttributeValue{
		"pk": &dynamodb.AttributeValue{S: &key},
	}}
	output, err := ddbClient.GetItem(&item)
	if err != nil {
		log.Println(err)
	}
	return output, err
}
