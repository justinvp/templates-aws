import pulumi
import json
import pulumi_aws as aws

example = aws.iam.Role("example", assume_role_policy=json.dumps({
    "Statement": [{
        "Action": "sts:AssumeRole",
        "Effect": "Allow",
        "Principal": {
            "Service": "ec2.amazonaws.com",
        },
    }],
    "Version": "2012-10-17",
}))
example__amazon_eks_worker_node_policy = aws.iam.RolePolicyAttachment("example-AmazonEKSWorkerNodePolicy",
    policy_arn="arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
    role=example.name)
example__amazon_ekscni_policy = aws.iam.RolePolicyAttachment("example-AmazonEKSCNIPolicy",
    policy_arn="arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy",
    role=example.name)
example__amazon_ec2_container_registry_read_only = aws.iam.RolePolicyAttachment("example-AmazonEC2ContainerRegistryReadOnly",
    policy_arn="arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly",
    role=example.name)

