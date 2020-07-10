import pulumi
import pulumi_aws as aws

example = aws.elasticsearch.Domain("example", elasticsearch_version="2.3")
main = aws.elasticsearch.DomainPolicy("main",
    access_policies=example.arn.apply(lambda arn: f"""{{
    "Version": "2012-10-17",
    "Statement": [
        {{
            "Action": "es:*",
            "Principal": "*",
            "Effect": "Allow",
            "Condition": {{
                "IpAddress": {{"aws:SourceIp": "127.0.0.1/32"}}
            }},
            "Resource": "{arn}/*"
        }}
    ]
}}

"""),
    domain_name=example.domain_name)

