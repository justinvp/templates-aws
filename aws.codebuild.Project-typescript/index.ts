import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucket = new aws.s3.Bucket("example", {
    acl: "private",
});
const exampleRole = new aws.iam.Role("example", {
    assumeRolePolicy: `{
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
`,
});
const exampleRolePolicy = new aws.iam.RolePolicy("example", {
    policy: pulumi.interpolate`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Resource": [
        "*"
      ],
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ]
    },
    {
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
    },
    {
      "Effect": "Allow",
      "Action": [
        "ec2:CreateNetworkInterfacePermission"
      ],
      "Resource": [
        "arn:aws:ec2:us-east-1:123456789012:network-interface/*"
      ],
      "Condition": {
        "StringEquals": {
          "ec2:Subnet": [
            "${aws_subnet_example1.arn}",
            "${aws_subnet_example2.arn}"
          ],
          "ec2:AuthorizedService": "codebuild.amazonaws.com"
        }
      }
    },
    {
      "Effect": "Allow",
      "Action": [
        "s3:*"
      ],
      "Resource": [
        "${exampleBucket.arn}",
        "${exampleBucket.arn}/*"
      ]
    }
  ]
}
`,
    role: exampleRole.name,
});
const exampleProject = new aws.codebuild.Project("example", {
    artifacts: {
        type: "NO_ARTIFACTS",
    },
    buildTimeout: 5,
    cache: {
        location: exampleBucket.bucket,
        type: "S3",
    },
    description: "test_codebuild_project",
    environment: {
        computeType: "BUILD_GENERAL1_SMALL",
        environmentVariables: [
            {
                name: "SOME_KEY1",
                value: "SOME_VALUE1",
            },
            {
                name: "SOME_KEY2",
                type: "PARAMETER_STORE",
                value: "SOME_VALUE2",
            },
        ],
        image: "aws/codebuild/standard:1.0",
        imagePullCredentialsType: "CODEBUILD",
        type: "LINUX_CONTAINER",
    },
    logsConfig: {
        cloudwatchLogs: {
            groupName: "log-group",
            streamName: "log-stream",
        },
        s3Logs: {
            location: pulumi.interpolate`${exampleBucket.id}/build-log`,
            status: "ENABLED",
        },
    },
    serviceRole: exampleRole.arn,
    source: {
        gitCloneDepth: 1,
        gitSubmodulesConfig: {
            fetchSubmodules: true,
        },
        location: "https://github.com/mitchellh/packer.git",
        type: "GITHUB",
    },
    sourceVersion: "master",
    tags: {
        Environment: "Test",
    },
    vpcConfig: {
        securityGroupIds: [
            aws_security_group_example1.id,
            aws_security_group_example2.id,
        ],
        subnets: [
            aws_subnet_example1.id,
            aws_subnet_example2.id,
        ],
        vpcId: aws_vpc_example.id,
    },
});
const project_with_cache = new aws.codebuild.Project("project-with-cache", {
    artifacts: {
        type: "NO_ARTIFACTS",
    },
    buildTimeout: 5,
    cache: {
        modes: [
            "LOCAL_DOCKER_LAYER_CACHE",
            "LOCAL_SOURCE_CACHE",
        ],
        type: "LOCAL",
    },
    description: "test_codebuild_project_cache",
    environment: {
        computeType: "BUILD_GENERAL1_SMALL",
        environmentVariables: [{
            name: "SOME_KEY1",
            value: "SOME_VALUE1",
        }],
        image: "aws/codebuild/standard:1.0",
        imagePullCredentialsType: "CODEBUILD",
        type: "LINUX_CONTAINER",
    },
    queuedTimeout: 5,
    serviceRole: exampleRole.arn,
    source: {
        gitCloneDepth: 1,
        location: "https://github.com/mitchellh/packer.git",
        type: "GITHUB",
    },
    tags: {
        Environment: "Test",
    },
});

