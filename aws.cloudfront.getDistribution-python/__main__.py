import pulumi
import pulumi_aws as aws

test = aws.cloudfront.get_distribution(id="EDFDVBD632BHDS5")

