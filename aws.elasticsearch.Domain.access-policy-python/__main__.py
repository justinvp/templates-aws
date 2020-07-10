import pulumi
import pulumi_aws as aws

config = pulumi.Config()
domain = config.get("domain")
if domain is None:
    domain = "tf-test"
current_region = aws.get_region()
current_caller_identity = aws.get_caller_identity()
example = aws.elasticsearch.Domain("example", access_policies=f"""{{
  "Version": "2012-10-17",
  "Statement": [
    {{
      "Action": "es:*",
      "Principal": "*",
      "Effect": "Allow",
      "Resource": "arn:aws:es:{current_region.name}:{current_caller_identity.account_id}:domain/{domain}/*",
      "Condition": {{
        "IpAddress": {{"aws:SourceIp": ["66.193.100.22/32"]}}
      }}
    }}
  ]
}}

""")

