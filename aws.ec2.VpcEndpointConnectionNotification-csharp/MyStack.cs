using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var topic = new Aws.Sns.Topic("topic", new Aws.Sns.TopicArgs
        {
            Policy = @"{
    ""Version"":""2012-10-17"",
    ""Statement"":[{
        ""Effect"": ""Allow"",
        ""Principal"": {
            ""Service"": ""vpce.amazonaws.com""
        },
        ""Action"": ""SNS:Publish"",
        ""Resource"": ""arn:aws:sns:*:*:vpce-notification-topic""
    }]
}

",
        });
        var fooVpcEndpointService = new Aws.Ec2.VpcEndpointService("fooVpcEndpointService", new Aws.Ec2.VpcEndpointServiceArgs
        {
            AcceptanceRequired = false,
            NetworkLoadBalancerArns = 
            {
                aws_lb.Test.Arn,
            },
        });
        var fooVpcEndpointConnectionNotification = new Aws.Ec2.VpcEndpointConnectionNotification("fooVpcEndpointConnectionNotification", new Aws.Ec2.VpcEndpointConnectionNotificationArgs
        {
            ConnectionEvents = 
            {
                "Accept",
                "Reject",
            },
            ConnectionNotificationArn = topic.Arn,
            VpcEndpointServiceId = fooVpcEndpointService.Id,
        });
    }

}

