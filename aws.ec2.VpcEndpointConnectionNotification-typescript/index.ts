import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const topic = new aws.sns.Topic("topic", {
    policy: `{
    "Version":"2012-10-17",
    "Statement":[{
        "Effect": "Allow",
        "Principal": {
            "Service": "vpce.amazonaws.com"
        },
        "Action": "SNS:Publish",
        "Resource": "arn:aws:sns:*:*:vpce-notification-topic"
    }]
}
`,
});
const fooVpcEndpointService = new aws.ec2.VpcEndpointService("foo", {
    acceptanceRequired: false,
    networkLoadBalancerArns: [aws_lb_test.arn],
});
const fooVpcEndpointConnectionNotification = new aws.ec2.VpcEndpointConnectionNotification("foo", {
    connectionEvents: [
        "Accept",
        "Reject",
    ],
    connectionNotificationArn: topic.arn,
    vpcEndpointServiceId: fooVpcEndpointService.id,
});

