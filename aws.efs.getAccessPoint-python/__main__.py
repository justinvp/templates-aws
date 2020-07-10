import pulumi
import pulumi_aws as aws

test = aws.efs.get_access_point(access_point_id="fsap-12345678")

