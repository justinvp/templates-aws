import pulumi
import pulumi_aws as aws

parent = aws.route53.HealthCheck("parent",
    child_health_threshold=1,
    child_healthchecks=[aws_route53_health_check["child"]["id"]],
    tags={
        "Name": "tf-test-calculated-health-check",
    },
    type="CALCULATED")

