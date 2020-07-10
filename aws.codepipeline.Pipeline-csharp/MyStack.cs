using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var codepipelineBucket = new Aws.S3.Bucket("codepipelineBucket", new Aws.S3.BucketArgs
        {
            Acl = "private",
        });
        var codepipelineRole = new Aws.Iam.Role("codepipelineRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Effect"": ""Allow"",
      ""Principal"": {
        ""Service"": ""codepipeline.amazonaws.com""
      },
      ""Action"": ""sts:AssumeRole""
    }
  ]
}

",
        });
        var codepipelinePolicy = new Aws.Iam.RolePolicy("codepipelinePolicy", new Aws.Iam.RolePolicyArgs
        {
            Policy = Output.Tuple(codepipelineBucket.Arn, codepipelineBucket.Arn).Apply(values =>
            {
                var codepipelineBucketArn = values.Item1;
                var codepipelineBucketArn1 = values.Item2;
                return @$"{{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {{
      ""Effect"":""Allow"",
      ""Action"": [
        ""s3:GetObject"",
        ""s3:GetObjectVersion"",
        ""s3:GetBucketVersioning"",
        ""s3:PutObject""
      ],
      ""Resource"": [
        ""{codepipelineBucketArn}"",
        ""{codepipelineBucketArn1}/*""
      ]
    }},
    {{
      ""Effect"": ""Allow"",
      ""Action"": [
        ""codebuild:BatchGetBuilds"",
        ""codebuild:StartBuild""
      ],
      ""Resource"": ""*""
    }}
  ]
}}

";
            }),
            Role = codepipelineRole.Id,
        });
        var s3kmskey = Output.Create(Aws.Kms.GetAlias.InvokeAsync(new Aws.Kms.GetAliasArgs
        {
            Name = "alias/myKmsKey",
        }));
        var codepipeline = new Aws.CodePipeline.Pipeline("codepipeline", new Aws.CodePipeline.PipelineArgs
        {
            ArtifactStore = new Aws.CodePipeline.Inputs.PipelineArtifactStoreArgs
            {
                EncryptionKey = new Aws.CodePipeline.Inputs.PipelineArtifactStoreEncryptionKeyArgs
                {
                    Id = s3kmskey.Apply(s3kmskey => s3kmskey.Arn),
                    Type = "KMS",
                },
                Location = codepipelineBucket.BucketName,
                Type = "S3",
            },
            RoleArn = codepipelineRole.Arn,
            Stages = 
            {
                new Aws.CodePipeline.Inputs.PipelineStageArgs
                {
                    Action = 
                    {
                        
                        {
                            { "category", "Source" },
                            { "configuration", 
                            {
                                { "Branch", "master" },
                                { "Owner", "my-organization" },
                                { "Repo", "test" },
                            } },
                            { "name", "Source" },
                            { "outputArtifacts", 
                            {
                                "source_output",
                            } },
                            { "owner", "ThirdParty" },
                            { "provider", "GitHub" },
                            { "version", "1" },
                        },
                    },
                    Name = "Source",
                },
                new Aws.CodePipeline.Inputs.PipelineStageArgs
                {
                    Action = 
                    {
                        
                        {
                            { "category", "Build" },
                            { "configuration", 
                            {
                                { "ProjectName", "test" },
                            } },
                            { "inputArtifacts", 
                            {
                                "source_output",
                            } },
                            { "name", "Build" },
                            { "outputArtifacts", 
                            {
                                "build_output",
                            } },
                            { "owner", "AWS" },
                            { "provider", "CodeBuild" },
                            { "version", "1" },
                        },
                    },
                    Name = "Build",
                },
                new Aws.CodePipeline.Inputs.PipelineStageArgs
                {
                    Action = 
                    {
                        
                        {
                            { "category", "Deploy" },
                            { "configuration", 
                            {
                                { "ActionMode", "REPLACE_ON_FAILURE" },
                                { "Capabilities", "CAPABILITY_AUTO_EXPAND,CAPABILITY_IAM" },
                                { "OutputFileName", "CreateStackOutput.json" },
                                { "StackName", "MyStack" },
                                { "TemplatePath", "build_output::sam-templated.yaml" },
                            } },
                            { "inputArtifacts", 
                            {
                                "build_output",
                            } },
                            { "name", "Deploy" },
                            { "owner", "AWS" },
                            { "provider", "CloudFormation" },
                            { "version", "1" },
                        },
                    },
                    Name = "Deploy",
                },
            },
        });
    }

}

