import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ssm.Association("example", {
    targets: [{
        key: "InstanceIds",
        values: [aws_instance_example.id],
    }],
});

