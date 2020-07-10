import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.iam.Role("example", {assumeRolePolicy: JSON.stringify({
    Statement: [{
        Action: "sts:AssumeRole",
        Effect: "Allow",
        Principal: {
            Service: "eks-fargate-pods.amazonaws.com",
        },
    }],
    Version: "2012-10-17",
})});
const example_AmazonEKSFargatePodExecutionRolePolicy = new aws.iam.RolePolicyAttachment("example-AmazonEKSFargatePodExecutionRolePolicy", {
    policyArn: "arn:aws:iam::aws:policy/AmazonEKSFargatePodExecutionRolePolicy",
    role: example.name,
});

