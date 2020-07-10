import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.globalaccelerator.EndpointGroup("example", {
    endpointConfigurations: [{
        endpointId: aws_lb_example.arn,
        weight: 100,
    }],
    listenerArn: aws_globalaccelerator_listener_example.id,
});

