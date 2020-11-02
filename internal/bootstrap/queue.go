package bootstrap

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/statistico/statistico-odds-warehouse/internal/queue"
	qa "github.com/statistico/statistico-odds-warehouse/internal/queue/aws"
	"github.com/statistico/statistico-odds-warehouse/internal/queue/log"
)

func (c Container) Queue() queue.Queue {
	if c.Config.QueueDriver == "aws" {
		key := c.Config.AwsConfig.Key
		secret := c.Config.AwsConfig.Secret

		sess, err := session.NewSession(&aws.Config{
			Credentials: credentials.NewStaticCredentials(key, secret, ""),
			Region:      aws.String(c.Config.AwsConfig.Region),
		})

		if err != nil {
			panic(err)
		}

		return qa.NewQueue(
			sqs.New(sess),
			c.Logger,
			c.Config.AwsConfig.QueueUrl,
			30,
		)
	}

	if c.Config.QueueDriver == "log" {
		return log.NewQueue(c.Logger)
	}

	panic("Queue driver provided is not supported")
}
