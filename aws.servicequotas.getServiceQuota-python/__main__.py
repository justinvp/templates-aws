import pulumi
import pulumi_aws as aws

by_quota_code = aws.servicequotas.get_service_quota(quota_code="L-F678F1CE",
    service_code="vpc")
by_quota_name = aws.servicequotas.get_service_quota(quota_name="VPCs per Region",
    service_code="vpc")

