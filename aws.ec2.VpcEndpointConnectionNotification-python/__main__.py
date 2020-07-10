import pulumi
import pulumi_aws as aws

topic = aws.sns.Topic("topic", policy="""{
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

""")
foo_vpc_endpoint_service = aws.ec2.VpcEndpointService("fooVpcEndpointService",
    acceptance_required=False,
    network_load_balancer_arns=[aws_lb["test"]["arn"]])
foo_vpc_endpoint_connection_notification = aws.ec2.VpcEndpointConnectionNotification("fooVpcEndpointConnectionNotification",
    connection_events=[
        "Accept",
        "Reject",
    ],
    connection_notification_arn=topic.arn,
    vpc_endpoint_service_id=foo_vpc_endpoint_service.id)

