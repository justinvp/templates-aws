package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/mq"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := brokerId
		_, err := mq.LookupBroker(ctx, &mq.LookupBrokerArgs{
			BrokerId: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		opt1 := brokerName
		_, err = mq.LookupBroker(ctx, &mq.LookupBrokerArgs{
			BrokerName: &opt1,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

