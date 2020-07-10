import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const lb = new aws.ec2.Eip("lb", {
    instance: aws_instance_web.id,
    vpc: true,
});

