import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const my_instance = new aws.opsworks.Instance("my-instance", {
    instanceType: "t2.micro",
    layerIds: [aws_opsworks_custom_layer_my_layer.id],
    os: "Amazon Linux 2015.09",
    stackId: aws_opsworks_stack_main.id,
    state: "stopped",
});

