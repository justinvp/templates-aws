import pulumi
import pulumi_aws as aws

web = aws.ec2.PlacementGroup("web", strategy="cluster")

