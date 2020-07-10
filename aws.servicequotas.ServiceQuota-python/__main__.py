import pulumi
import pulumi_aws as aws

example = aws.servicequotas.ServiceQuota("example",
    quota_code="L-F678F1CE",
    service_code="vpc",
    value=75)

