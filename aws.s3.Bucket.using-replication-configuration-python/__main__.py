import pulumi
import pulumi_aws as aws
import pulumi_pulumi as pulumi

central = pulumi.providers.Aws("central", region="eu-central-1")
replication_role = aws.iam.Role("replicationRole", assume_role_policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "s3.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}

""")
destination = aws.s3.Bucket("destination",
    region="eu-west-1",
    versioning={
        "enabled": True,
    })
bucket = aws.s3.Bucket("bucket",
    acl="private",
    region="eu-central-1",
    replication_configuration={
        "role": replication_role.arn,
        "rules": [{
            "destination": {
                "bucket": destination.arn,
                "storage_class": "STANDARD",
            },
            "id": "foobar",
            "prefix": "foo",
            "status": "Enabled",
        }],
    },
    versioning={
        "enabled": True,
    },
    opts=ResourceOptions(provider="aws.central"))
replication_policy = aws.iam.Policy("replicationPolicy", policy=pulumi.Output.all(bucket.arn, bucket.arn, destination.arn).apply(lambda bucketArn, bucketArn1, destinationArn: f"""{{
  "Version": "2012-10-17",
  "Statement": [
    {{
      "Action": [
        "s3:GetReplicationConfiguration",
        "s3:ListBucket"
      ],
      "Effect": "Allow",
      "Resource": [
        "{bucket_arn}"
      ]
    }},
    {{
      "Action": [
        "s3:GetObjectVersion",
        "s3:GetObjectVersionAcl"
      ],
      "Effect": "Allow",
      "Resource": [
        "{bucket_arn1}/*"
      ]
    }},
    {{
      "Action": [
        "s3:ReplicateObject",
        "s3:ReplicateDelete"
      ],
      "Effect": "Allow",
      "Resource": "{destination_arn}/*"
    }}
  ]
}}

"""))
replication_role_policy_attachment = aws.iam.RolePolicyAttachment("replicationRolePolicyAttachment",
    policy_arn=replication_policy.arn,
    role=replication_role.name)

