import pulumi
import json
import pulumi_aws as aws

example = aws.iam.Role("example", assume_role_policy=json.dumps({
    "Statement": [{
        "Action": "sts:AssumeRole",
        "Effect": "Allow",
        "Principal": {
            "Service": "eks-fargate-pods.amazonaws.com",
        },
    }],
    "Version": "2012-10-17",
}))
example__amazon_eks_fargate_pod_execution_role_policy = aws.iam.RolePolicyAttachment("example-AmazonEKSFargatePodExecutionRolePolicy",
    policy_arn="arn:aws:iam::aws:policy/AmazonEKSFargatePodExecutionRolePolicy",
    role=example.name)

