package main

import (
	"fmt"
	"log"
	"path"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	holiday "github.com/holiday-jp/holiday_jp-go"
)

var tag string
var state string
var list []string

var sess = session.Must(session.NewSession())
var svc = ec2.New(sess)

func handler(event events.CloudWatchEvent) error {
	schedule := fmt.Sprint(path.Base(event.Resources[0]))

	if schedule == "AutoStart" {
		if holiday.IsHoliday(time.Now()) {
			log.Println("Today is holiday! Skip auto start.")
			return nil
		}
		state = "stopped"
		tag = "AutoStart"
	} else if schedule == "AutoStop" {
		state = "running"
		tag = "AutoStop"
	} else {
		log.Println("Cloudwatch schedule not match.")
		return nil
	}

	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-state-name"),
				Values: []*string{
					aws.String(state),
				},
			},
		},
	}
	resp, err := svc.DescribeInstances(input)
	if err != nil {
		panic(err)
	}

	for _, r := range resp.Reservations {
		for _, i := range r.Instances {
			for _, t := range i.Tags {
				if *t.Key == tag {
					if *t.Value == "ON" {
						list = append(list, *i.InstanceId)
					}
				}
			}
		}
	}

	if schedule == "AutoStart" {
		startInstance(list)
	} else if schedule == "AutoStop" {
		stopInstance(list)
	}

	return nil
}

func startInstance(list []string) {
	for _, i := range list {
		input := &ec2.StartInstancesInput{
			InstanceIds: []*string{
				aws.String(i),
			},
		}

		log.Printf("Start instance (%s).", i)
		_, err := svc.StartInstances(input)
		if err != nil {
			panic(err)
		}
	}
}

func stopInstance(list []string) {
	for _, i := range list {
		input := &ec2.StopInstancesInput{
			InstanceIds: []*string{
				aws.String(i),
			},
		}

		log.Printf("Stop instance (%s).", i)
		_, err := svc.StopInstances(input)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	lambda.Start(handler)
}
