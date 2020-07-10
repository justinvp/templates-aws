import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new load balancer attachment
const baz = new aws.elb.Attachment("baz", {
    elb: aws_elb_bar.id,
    instance: aws_instance_foo.id,
});

