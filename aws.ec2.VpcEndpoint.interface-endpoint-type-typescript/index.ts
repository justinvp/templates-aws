import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ec2 = new aws.ec2.VpcEndpoint("ec2", {
    privateDnsEnabled: true,
    securityGroupIds: [aws_security_group_sg1.id],
    serviceName: "com.amazonaws.us-west-2.ec2",
    vpcEndpointType: "Interface",
    vpcId: aws_vpc_main.id,
});

