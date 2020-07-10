using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleBucket = new Aws.S3.Bucket("exampleBucket", new Aws.S3.BucketArgs
        {
            Acl = "private",
        });
        var exampleRole = new Aws.Iam.Role("exampleRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Effect"": ""Allow"",
      ""Principal"": {
        ""Service"": ""codebuild.amazonaws.com""
      },
      ""Action"": ""sts:AssumeRole""
    }
  ]
}

",
        });
        var exampleRolePolicy = new Aws.Iam.RolePolicy("exampleRolePolicy", new Aws.Iam.RolePolicyArgs
        {
            Policy = Output.Tuple(exampleBucket.Arn, exampleBucket.Arn).Apply(values =>
            {
                var exampleBucketArn = values.Item1;
                var exampleBucketArn1 = values.Item2;
                return @$"{{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {{
      ""Effect"": ""Allow"",
      ""Resource"": [
        ""*""
      ],
      ""Action"": [
        ""logs:CreateLogGroup"",
        ""logs:CreateLogStream"",
        ""logs:PutLogEvents""
      ]
    }},
    {{
      ""Effect"": ""Allow"",
      ""Action"": [
        ""ec2:CreateNetworkInterface"",
        ""ec2:DescribeDhcpOptions"",
        ""ec2:DescribeNetworkInterfaces"",
        ""ec2:DeleteNetworkInterface"",
        ""ec2:DescribeSubnets"",
        ""ec2:DescribeSecurityGroups"",
        ""ec2:DescribeVpcs""
      ],
      ""Resource"": ""*""
    }},
    {{
      ""Effect"": ""Allow"",
      ""Action"": [
        ""ec2:CreateNetworkInterfacePermission""
      ],
      ""Resource"": [
        ""arn:aws:ec2:us-east-1:123456789012:network-interface/*""
      ],
      ""Condition"": {{
        ""StringEquals"": {{
          ""ec2:Subnet"": [
            ""{aws_subnet.Example1.Arn}"",
            ""{aws_subnet.Example2.Arn}""
          ],
          ""ec2:AuthorizedService"": ""codebuild.amazonaws.com""
        }}
      }}
    }},
    {{
      ""Effect"": ""Allow"",
      ""Action"": [
        ""s3:*""
      ],
      ""Resource"": [
        ""{exampleBucketArn}"",
        ""{exampleBucketArn1}/*""
      ]
    }}
  ]
}}

";
            }),
            Role = exampleRole.Name,
        });
        var exampleProject = new Aws.CodeBuild.Project("exampleProject", new Aws.CodeBuild.ProjectArgs
        {
            Artifacts = new Aws.CodeBuild.Inputs.ProjectArtifactsArgs
            {
                Type = "NO_ARTIFACTS",
            },
            BuildTimeout = 5,
            Cache = new Aws.CodeBuild.Inputs.ProjectCacheArgs
            {
                Location = exampleBucket.BucketName,
                Type = "S3",
            },
            Description = "test_codebuild_project",
            Environment = new Aws.CodeBuild.Inputs.ProjectEnvironmentArgs
            {
                ComputeType = "BUILD_GENERAL1_SMALL",
                EnvironmentVariable = 
                {
                    
                    {
                        { "name", "SOME_KEY1" },
                        { "value", "SOME_VALUE1" },
                    },
                    
                    {
                        { "name", "SOME_KEY2" },
                        { "type", "PARAMETER_STORE" },
                        { "value", "SOME_VALUE2" },
                    },
                },
                Image = "aws/codebuild/standard:1.0",
                ImagePullCredentialsType = "CODEBUILD",
                Type = "LINUX_CONTAINER",
            },
            LogsConfig = new Aws.CodeBuild.Inputs.ProjectLogsConfigArgs
            {
                CloudwatchLogs = new Aws.CodeBuild.Inputs.ProjectLogsConfigCloudwatchLogsArgs
                {
                    GroupName = "log-group",
                    StreamName = "log-stream",
                },
                S3Logs = new Aws.CodeBuild.Inputs.ProjectLogsConfigS3LogsArgs
                {
                    Location = exampleBucket.Id.Apply(id => $"{id}/build-log"),
                    Status = "ENABLED",
                },
            },
            ServiceRole = exampleRole.Arn,
            Source = new Aws.CodeBuild.Inputs.ProjectSourceArgs
            {
                GitCloneDepth = 1,
                GitSubmodulesConfig = new Aws.CodeBuild.Inputs.ProjectSourceGitSubmodulesConfigArgs
                {
                    FetchSubmodules = true,
                },
                Location = "https://github.com/mitchellh/packer.git",
                Type = "GITHUB",
            },
            SourceVersion = "master",
            Tags = 
            {
                { "Environment", "Test" },
            },
            VpcConfig = new Aws.CodeBuild.Inputs.ProjectVpcConfigArgs
            {
                SecurityGroupIds = 
                {
                    aws_security_group.Example1.Id,
                    aws_security_group.Example2.Id,
                },
                Subnets = 
                {
                    aws_subnet.Example1.Id,
                    aws_subnet.Example2.Id,
                },
                VpcId = aws_vpc.Example.Id,
            },
        });
        var project_with_cache = new Aws.CodeBuild.Project("project-with-cache", new Aws.CodeBuild.ProjectArgs
        {
            Artifacts = new Aws.CodeBuild.Inputs.ProjectArtifactsArgs
            {
                Type = "NO_ARTIFACTS",
            },
            BuildTimeout = 5,
            Cache = new Aws.CodeBuild.Inputs.ProjectCacheArgs
            {
                Modes = 
                {
                    "LOCAL_DOCKER_LAYER_CACHE",
                    "LOCAL_SOURCE_CACHE",
                },
                Type = "LOCAL",
            },
            Description = "test_codebuild_project_cache",
            Environment = new Aws.CodeBuild.Inputs.ProjectEnvironmentArgs
            {
                ComputeType = "BUILD_GENERAL1_SMALL",
                EnvironmentVariable = 
                {
                    
                    {
                        { "name", "SOME_KEY1" },
                        { "value", "SOME_VALUE1" },
                    },
                },
                Image = "aws/codebuild/standard:1.0",
                ImagePullCredentialsType = "CODEBUILD",
                Type = "LINUX_CONTAINER",
            },
            QueuedTimeout = 5,
            ServiceRole = exampleRole.Arn,
            Source = new Aws.CodeBuild.Inputs.ProjectSourceArgs
            {
                GitCloneDepth = 1,
                Location = "https://github.com/mitchellh/packer.git",
                Type = "GITHUB",
            },
            Tags = 
            {
                { "Environment", "Test" },
            },
        });
    }

}

