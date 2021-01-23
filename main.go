package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/google/uuid"
)

var svc *cloudwatchlogs.CloudWatchLogs
var myuuid string

func init() {
	os.Setenv("FOO", "1")
	os.Setenv("BAR", "2")
	os.Setenv("ITYPE", "3")
	myuuid = uuid.New().String()

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: "oseven",
		Config: aws.Config{
			Region: aws.String("eu-central-1"),
		},
	}))

	svc = cloudwatchlogs.New(sess)

	svc.CreateLogStream(&cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  aws.String("test"),
		LogStreamName: aws.String(myuuid),
	})

	// input := &cloudwatchlogs.DescribeLogStreamsInput{
	// 	LogGroupName:        aws.String("test"),
	// 	LogStreamNamePrefix: aws.String(myuuid),
	// }
	// logData, err := svc.DescribeLogStreams(input)
	// fmt.Print(logData)
	// if err != nil {
	// 	panic(err)
	// }
	// nexttoken = logData.LogStreams[0].UploadSequenceToken
	// fmt.Print(nexttoken)
}

type logs struct {
	UUID  string `json:"uuid"`
	Foo   string `json:"foo"`
	Bar   string `json:"bar"`
	Itype string `json:"itype"`
}

func main() {

	foo := os.Getenv("FOO")
	bar := os.Getenv("BAR")
	itype := os.Getenv("ITYPE")

	// // resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	// // if err != nil {
	// // 	panic("PANIC")
	// // }
	// // defer resp.Body.Close()
	// // body, err := ioutil.ReadAll(resp.Body)
	// // fmt.Println(string(body))

	log := &logs{
		UUID:  myuuid,
		Foo:   foo,
		Bar:   bar,
		Itype: itype,
	}

	b, err := json.Marshal(log)
	if err != nil {
		panic(err)
	}

	var nexttoken *string
	for {

		log := &cloudwatchlogs.PutLogEventsInput{
			LogEvents: []*cloudwatchlogs.InputLogEvent{
				{
					Message:   aws.String(string(b)),
					Timestamp: makeTimestamp(),
				},
			},
			LogGroupName:  aws.String("test"),
			LogStreamName: aws.String(myuuid),
		}
		if nexttoken != nil {
			log.SequenceToken = aws.String(*nexttoken)
		}

		data, err := svc.PutLogEvents(log)
		if err != nil {
			panic(err)
		}
		nexttoken = data.NextSequenceToken
		time.Sleep(10 * time.Second)
	}
}

func makeTimestamp() *int64 {
	ts := time.Now().UnixNano() / int64(time.Millisecond)
	return &ts
}
