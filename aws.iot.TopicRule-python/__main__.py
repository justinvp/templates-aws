import pulumi
import pulumi_aws as aws

mytopic = aws.sns.Topic("mytopic")
role = aws.iam.Role("role", assume_role_policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "iot.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}

""")
rule = aws.iot.TopicRule("rule",
    description="Example rule",
    enabled=True,
    sns={
        "sns": "RAW",
        "sns": role.arn,
        "sns": mytopic.arn,
    },
    sql="SELECT * FROM 'topic/test'",
    sql_version="2016-03-23")
iam_policy_for_lambda = aws.iam.RolePolicy("iamPolicyForLambda",
    policy=mytopic.arn.apply(lambda arn: f"""{{
  "Version": "2012-10-17",
  "Statement": [
    {{
        "Effect": "Allow",
        "Action": [
            "sns:Publish"
        ],
        "Resource": "{arn}"
    }}
  ]
}}

"""),
    role=role.id)

