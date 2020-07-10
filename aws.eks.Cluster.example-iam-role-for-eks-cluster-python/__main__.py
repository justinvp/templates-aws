import pulumi
import pulumi_aws as aws

example = aws.iam.Role("example", assume_role_policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "eks.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}

""")
example__amazon_eks_cluster_policy = aws.iam.RolePolicyAttachment("example-AmazonEKSClusterPolicy",
    policy_arn="arn:aws:iam::aws:policy/AmazonEKSClusterPolicy",
    role=example.name)
example__amazon_eks_service_policy = aws.iam.RolePolicyAttachment("example-AmazonEKSServicePolicy",
    policy_arn="arn:aws:iam::aws:policy/AmazonEKSServicePolicy",
    role=example.name)

