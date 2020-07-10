import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const model = new aws.sagemaker.Model("m", {
    executionRoleArn: aws_iam_role_foo.arn,
    primaryContainer: {
        image: "174872318107.dkr.ecr.us-west-2.amazonaws.com/kmeans:1",
    },
});
const assumeRole = pulumi.output(aws.iam.getPolicyDocument({
    statements: [{
        actions: ["sts:AssumeRole"],
        principals: [{
            identifiers: ["sagemaker.amazonaws.com"],
            type: "Service",
        }],
    }],
}, { async: true }));
const role = new aws.iam.Role("r", {
    assumeRolePolicy: assumeRole.json,
});

