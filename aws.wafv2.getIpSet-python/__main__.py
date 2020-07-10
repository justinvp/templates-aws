import pulumi
import pulumi_aws as aws

example = aws.wafv2.get_ip_set(name="some-ip-set",
    scope="REGIONAL")

