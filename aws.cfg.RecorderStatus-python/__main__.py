import pulumi
import pulumi_aws as aws

foo_recorder_status = aws.cfg.RecorderStatus("fooRecorderStatus", is_enabled=True,
opts=ResourceOptions(depends_on=["aws_config_delivery_channel.foo"]))
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
role_policy_attachment = aws.iam.RolePolicyAttachment("rolePolicyAttachment",
    policy_arn="arn:aws:iam::aws:policy/service-role/AWSConfigRole",
    role=role.name)
bucket = aws.s3.Bucket("bucket")
foo_delivery_channel = aws.cfg.DeliveryChannel("fooDeliveryChannel", s3_bucket_name=bucket.bucket)
foo_recorder = aws.cfg.Recorder("fooRecorder", role_arn=role.arn)
role_policy = aws.iam.RolePolicy("rolePolicy",
    policy=pulumi.Output.all(bucket.arn, bucket.arn).apply(lambda bucketArn, bucketArn1: f"""{{
  "Version": "2012-10-17",
  "Statement": [
    {{
      "Action": [
        "s3:*"
      ],
      "Effect": "Allow",
      "Resource": [
        "{bucket_arn}",
        "{bucket_arn1}/*"
      ]
    }}
  ]
}}

"""),
    role=role.id)

