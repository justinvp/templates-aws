import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.iam.Role("example", {assumeRolePolicy: JSON.stringify({
    Statement: [{
        Action: "sts:AssumeRole",
        Effect: "Allow",
        Principal: {
            Service: "ec2.amazonaws.com",
        },
    }],
    Version: "2012-10-17",
})});
const example_AmazonEKSWorkerNodePolicy = new aws.iam.RolePolicyAttachment("example-AmazonEKSWorkerNodePolicy", {
    policyArn: "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
    role: example.name,
});
const example_AmazonEKSCNIPolicy = new aws.iam.RolePolicyAttachment("example-AmazonEKSCNIPolicy", {
    policyArn: "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy",
    role: example.name,
});
const example_AmazonEC2ContainerRegistryReadOnly = new aws.iam.RolePolicyAttachment("example-AmazonEC2ContainerRegistryReadOnly", {
    policyArn: "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly",
    role: example.name,
});

