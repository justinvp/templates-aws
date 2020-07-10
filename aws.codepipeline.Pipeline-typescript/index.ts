import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const codepipelineBucket = new aws.s3.Bucket("codepipeline_bucket", {
    acl: "private",
});
const codepipelineRole = new aws.iam.Role("codepipeline_role", {
    assumeRolePolicy: `{
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
`,
});
const codepipelinePolicy = new aws.iam.RolePolicy("codepipeline_policy", {
    policy: pulumi.interpolate`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect":"Allow",
      "Action": [
        "s3:GetObject",
        "s3:GetObjectVersion",
        "s3:GetBucketVersioning",
        "s3:PutObject"
      ],
      "Resource": [
        "${codepipelineBucket.arn}",
        "${codepipelineBucket.arn}/*"
      ]
    },
    {
      "Effect": "Allow",
      "Action": [
        "codebuild:BatchGetBuilds",
        "codebuild:StartBuild"
      ],
      "Resource": "*"
    }
  ]
}
`,
    role: codepipelineRole.id,
});
const s3kmskey = pulumi.output(aws.kms.getAlias({
    name: "alias/myKmsKey",
}, { async: true }));
const codepipeline = new aws.codepipeline.Pipeline("codepipeline", {
    artifactStores: {
        encryptionKey: {
            id: s3kmskey.arn,
            type: "KMS",
        },
        location: codepipelineBucket.bucket,
        type: "S3",
    },
    roleArn: codepipelineRole.arn,
    stages: [
        {
            actions: [{
                category: "Source",
                configuration: {
                    Branch: "master",
                    Owner: "my-organization",
                    Repo: "test",
                },
                name: "Source",
                outputArtifacts: ["source_output"],
                owner: "ThirdParty",
                provider: "GitHub",
                version: "1",
            }],
            name: "Source",
        },
        {
            actions: [{
                category: "Build",
                configuration: {
                    ProjectName: "test",
                },
                inputArtifacts: ["source_output"],
                name: "Build",
                outputArtifacts: ["build_output"],
                owner: "AWS",
                provider: "CodeBuild",
                version: "1",
            }],
            name: "Build",
        },
        {
            actions: [{
                category: "Deploy",
                configuration: {
                    ActionMode: "REPLACE_ON_FAILURE",
                    Capabilities: "CAPABILITY_AUTO_EXPAND,CAPABILITY_IAM",
                    OutputFileName: "CreateStackOutput.json",
                    StackName: "MyStack",
                    TemplatePath: "build_output::sam-templated.yaml",
                },
                inputArtifacts: ["build_output"],
                name: "Deploy",
                owner: "AWS",
                provider: "CloudFormation",
                version: "1",
            }],
            name: "Deploy",
        },
    ],
});

