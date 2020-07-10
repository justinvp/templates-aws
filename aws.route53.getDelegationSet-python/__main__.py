import pulumi
import pulumi_aws as aws

dset = aws.route53.get_delegation_set(id="MQWGHCBFAKEID")

