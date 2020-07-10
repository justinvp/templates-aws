import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const s3 = new aws.ec2.VpcEndpoint("s3", {
    serviceName: "com.amazonaws.us-west-2.s3",
    tags: {
        Environment: "test",
    },
    vpcId: aws_vpc_main.id,
});

