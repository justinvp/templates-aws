package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		topic, err := sns.NewTopic(ctx, "topic", &sns.TopicArgs{
			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\":\"2012-10-17\",\n", "    \"Statement\":[{\n", "        \"Effect\": \"Allow\",\n", "        \"Principal\": {\n", "            \"Service\": \"vpce.amazonaws.com\"\n", "        },\n", "        \"Action\": \"SNS:Publish\",\n", "        \"Resource\": \"arn:aws:sns:*:*:vpce-notification-topic\"\n", "    }]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		fooVpcEndpointService, err := ec2.NewVpcEndpointService(ctx, "fooVpcEndpointService", &ec2.VpcEndpointServiceArgs{
			AcceptanceRequired: pulumi.Bool(false),
			NetworkLoadBalancerArns: pulumi.StringArray{
				pulumi.Any(aws_lb.Test.Arn),
			},
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcEndpointConnectionNotification(ctx, "fooVpcEndpointConnectionNotification", &ec2.VpcEndpointConnectionNotificationArgs{
			ConnectionEvents: pulumi.StringArray{
				pulumi.String("Accept"),
				pulumi.String("Reject"),
			},
			ConnectionNotificationArn: topic.Arn,
			VpcEndpointServiceId:      fooVpcEndpointService.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

