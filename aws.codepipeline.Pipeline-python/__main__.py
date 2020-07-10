import pulumi
import pulumi_aws as aws

codepipeline_bucket = aws.s3.Bucket("codepipelineBucket", acl="private")
codepipeline_role = aws.iam.Role("codepipelineRole", assume_role_policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "codepipeline.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}

""")
codepipeline_policy = aws.iam.RolePolicy("codepipelinePolicy",
    policy=pulumi.Output.all(codepipeline_bucket.arn, codepipeline_bucket.arn).apply(lambda codepipelineBucketArn, codepipelineBucketArn1: f"""{{
  "Version": "2012-10-17",
  "Statement": [
    {{
      "Effect":"Allow",
      "Action": [
        "s3:GetObject",
        "s3:GetObjectVersion",
        "s3:GetBucketVersioning",
        "s3:PutObject"
      ],
      "Resource": [
        "{codepipeline_bucket_arn}",
        "{codepipeline_bucket_arn1}/*"
      ]
    }},
    {{
      "Effect": "Allow",
      "Action": [
        "codebuild:BatchGetBuilds",
        "codebuild:StartBuild"
      ],
      "Resource": "*"
    }}
  ]
}}

"""),
    role=codepipeline_role.id)
s3kmskey = aws.kms.get_alias(name="alias/myKmsKey")
codepipeline = aws.codepipeline.Pipeline("codepipeline",
    artifact_store={
        "encryption_key": {
            "id": s3kmskey.arn,
            "type": "KMS",
        },
        "location": codepipeline_bucket.bucket,
        "type": "S3",
    },
    role_arn=codepipeline_role.arn,
    stages=[
        {
            "action": [{
                "category": "Source",
                "configuration": {
                    "Branch": "master",
                    "Owner": "my-organization",
                    "Repo": "test",
                },
                "name": "Source",
                "outputArtifacts": ["source_output"],
                "owner": "ThirdParty",
                "provider": "GitHub",
                "version": "1",
            }],
            "name": "Source",
        },
        {
            "action": [{
                "category": "Build",
                "configuration": {
                    "ProjectName": "test",
                },
                "inputArtifacts": ["source_output"],
                "name": "Build",
                "outputArtifacts": ["build_output"],
                "owner": "AWS",
                "provider": "CodeBuild",
                "version": "1",
            }],
            "name": "Build",
        },
        {
            "action": [{
                "category": "Deploy",
                "configuration": {
                    "ActionMode": "REPLACE_ON_FAILURE",
                    "Capabilities": "CAPABILITY_AUTO_EXPAND,CAPABILITY_IAM",
                    "OutputFileName": "CreateStackOutput.json",
                    "StackName": "MyStack",
                    "TemplatePath": "build_output::sam-templated.yaml",
                },
                "inputArtifacts": ["build_output"],
                "name": "Deploy",
                "owner": "AWS",
                "provider": "CloudFormation",
                "version": "1",
            }],
            "name": "Deploy",
        },
    ])

