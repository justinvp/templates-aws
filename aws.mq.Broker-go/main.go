package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/mq"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mq.NewBroker(ctx, "example", &mq.BrokerArgs{
			BrokerName: pulumi.String("example"),
			Configuration: &mq.BrokerConfigurationArgs{
				Id:       pulumi.Any(aws_mq_configuration.Test.Id),
				Revision: pulumi.Any(aws_mq_configuration.Test.Latest_revision),
			},
			EngineType:       pulumi.String("ActiveMQ"),
			EngineVersion:    pulumi.String("5.15.0"),
			HostInstanceType: pulumi.String("mq.t2.micro"),
			SecurityGroups: pulumi.StringArray{
				pulumi.Any(aws_security_group.Test.Id),
			},
			Users: mq.BrokerUserArray{
				&mq.BrokerUserArgs{
					Password: pulumi.String("MindTheGap"),
					Username: pulumi.String("ExampleUser"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

