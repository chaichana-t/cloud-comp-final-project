package timestream

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/timestreamwrite"
	"strconv"
	"time"
)

func getRecord(name, density string, timestamp int64) *timestreamwrite.Record {
	return &timestreamwrite.Record{
		Dimensions: []*timestreamwrite.Dimension{
			{
				Name:  aws.String("name"),
				Value: aws.String(name),
			},
		},
		MeasureName:      aws.String("customer"),
		MeasureValue:     aws.String(density),
		MeasureValueType: aws.String("BIGINT"),
		Time:             aws.String(strconv.FormatInt(timestamp, 10)),
		TimeUnit:         aws.String("SECONDS"),
	}
}

func WriteSnapshot(snapshot map[string]string) {
	timestamp := time.Now().Unix()

	var records []*timestreamwrite.Record

	for name, density := range snapshot {
		records = append(records, getRecord(name, density, timestamp))
	}

	input := &timestreamwrite.WriteRecordsInput{
		DatabaseName: aws.String("TimestreamCloudFinalProject"),
		TableName:    aws.String("History"),
		Records: records,
	}

	_, err := writeSvc.WriteRecords(input)

	if err != nil {
		fmt.Println(err)
	}
}
