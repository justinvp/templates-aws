import pulumi
import pulumi_aws as aws

rule = aws.cfg.Rule("rule", source={
    "owner": "AWS",
    "sourceIdentifier": "S3_BUCKET_VERSIONING_ENABLED",
},
opts=ResourceOptions(depends_on=["aws_config_configuration_recorder.foo"]))
role = aws.iam.Role("role", assume_role_policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "config.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}

""")
foo = aws.cfg.Recorder("foo", role_arn=role.arn)
role_policy = aws.iam.RolePolicy("rolePolicy",
    policy="""{
  "Version": "2012-10-17",
  "Statement": [
  	{
  		"Action": "config:Put*",
  		"Effect": "Allow",
  		"Resource": "*"

  	}
  ]
}

""",
    role=role.id)

