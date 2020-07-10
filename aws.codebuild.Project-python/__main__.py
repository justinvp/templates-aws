import pulumi
import pulumi_aws as aws

example_bucket = aws.s3.Bucket("exampleBucket", acl="private")
example_role = aws.iam.Role("exampleRole", assume_role_policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "codebuild.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}

""")
example_role_policy = aws.iam.RolePolicy("exampleRolePolicy",
    policy=pulumi.Output.all(example_bucket.arn, example_bucket.arn).apply(lambda exampleBucketArn, exampleBucketArn1: f"""{{
  "Version": "2012-10-17",
  "Statement": [
    {{
      "Effect": "Allow",
      "Resource": [
        "*"
      ],
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ]
    }},
    {{
      "Effect": "Allow",
      "Action": [
        "ec2:CreateNetworkInterface",
        "ec2:DescribeDhcpOptions",
        "ec2:DescribeNetworkInterfaces",
        "ec2:DeleteNetworkInterface",
        "ec2:DescribeSubnets",
        "ec2:DescribeSecurityGroups",
        "ec2:DescribeVpcs"
      ],
      "Resource": "*"
    }},
    {{
      "Effect": "Allow",
      "Action": [
        "ec2:CreateNetworkInterfacePermission"
      ],
      "Resource": [
        "arn:aws:ec2:us-east-1:123456789012:network-interface/*"
      ],
      "Condition": {{
        "StringEquals": {{
          "ec2:Subnet": [
            "{aws_subnet["example1"]["arn"]}",
            "{aws_subnet["example2"]["arn"]}"
          ],
          "ec2:AuthorizedService": "codebuild.amazonaws.com"
        }}
      }}
    }},
    {{
      "Effect": "Allow",
      "Action": [
        "s3:*"
      ],
      "Resource": [
        "{example_bucket_arn}",
        "{example_bucket_arn1}/*"
      ]
    }}
  ]
}}

"""),
    role=example_role.name)
example_project = aws.codebuild.Project("exampleProject",
    artifacts={
        "type": "NO_ARTIFACTS",
    },
    build_timeout="5",
    cache={
        "location": example_bucket.bucket,
        "type": "S3",
    },
    description="test_codebuild_project",
    environment={
        "computeType": "BUILD_GENERAL1_SMALL",
        "environmentVariable": [
            {
                "name": "SOME_KEY1",
                "value": "SOME_VALUE1",
            },
            {
                "name": "SOME_KEY2",
                "type": "PARAMETER_STORE",
                "value": "SOME_VALUE2",
            },
        ],
        "image": "aws/codebuild/standard:1.0",
        "imagePullCredentialsType": "CODEBUILD",
        "type": "LINUX_CONTAINER",
    },
    logs_config={
        "cloudwatchLogs": {
            "group_name": "log-group",
            "streamName": "log-stream",
        },
        "s3Logs": {
            "location": example_bucket.id.apply(lambda id: f"{id}/build-log"),
            "status": "ENABLED",
        },
    },
    service_role=example_role.arn,
    source={
        "gitCloneDepth": 1,
        "gitSubmodulesConfig": {
            "fetchSubmodules": True,
        },
        "location": "https://github.com/mitchellh/packer.git",
        "type": "GITHUB",
    },
    source_version="master",
    tags={
        "Environment": "Test",
    },
    vpc_config={
        "security_group_ids": [
            aws_security_group["example1"]["id"],
            aws_security_group["example2"]["id"],
        ],
        "subnets": [
            aws_subnet["example1"]["id"],
            aws_subnet["example2"]["id"],
        ],
        "vpc_id": aws_vpc["example"]["id"],
    })
project_with_cache = aws.codebuild.Project("project-with-cache",
    artifacts={
        "type": "NO_ARTIFACTS",
    },
    build_timeout="5",
    cache={
        "modes": [
            "LOCAL_DOCKER_LAYER_CACHE",
            "LOCAL_SOURCE_CACHE",
        ],
        "type": "LOCAL",
    },
    description="test_codebuild_project_cache",
    environment={
        "computeType": "BUILD_GENERAL1_SMALL",
        "environmentVariable": [{
            "name": "SOME_KEY1",
            "value": "SOME_VALUE1",
        }],
        "image": "aws/codebuild/standard:1.0",
        "imagePullCredentialsType": "CODEBUILD",
        "type": "LINUX_CONTAINER",
    },
    queued_timeout="5",
    service_role=example_role.arn,
    source={
        "gitCloneDepth": 1,
        "location": "https://github.com/mitchellh/packer.git",
        "type": "GITHUB",
    },
    tags={
        "Environment": "Test",
    })

