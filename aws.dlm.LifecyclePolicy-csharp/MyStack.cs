using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var dlmLifecycleRole = new Aws.Iam.Role("dlmLifecycleRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Action"": ""sts:AssumeRole"",
      ""Principal"": {
        ""Service"": ""dlm.amazonaws.com""
      },
      ""Effect"": ""Allow"",
      ""Sid"": """"
    }
  ]
}

",
        });
        var dlmLifecycle = new Aws.Iam.RolePolicy("dlmLifecycle", new Aws.Iam.RolePolicyArgs
        {
            Policy = @"{
   ""Version"": ""2012-10-17"",
   ""Statement"": [
      {
         ""Effect"": ""Allow"",
         ""Action"": [
            ""ec2:CreateSnapshot"",
            ""ec2:DeleteSnapshot"",
            ""ec2:DescribeVolumes"",
            ""ec2:DescribeSnapshots""
         ],
         ""Resource"": ""*""
      },
      {
         ""Effect"": ""Allow"",
         ""Action"": [
            ""ec2:CreateTags""
         ],
         ""Resource"": ""arn:aws:ec2:*::snapshot/*""
      }
   ]
}

",
            Role = dlmLifecycleRole.Id,
        });
        var example = new Aws.Dlm.LifecyclePolicy("example", new Aws.Dlm.LifecyclePolicyArgs
        {
            Description = "example DLM lifecycle policy",
            ExecutionRoleArn = dlmLifecycleRole.Arn,
            PolicyDetails = new Aws.Dlm.Inputs.LifecyclePolicyPolicyDetailsArgs
            {
                ResourceTypes = 
                {
                    "VOLUME",
                },
                Schedule = 
                {
                    
                    {
                        { "copyTags", false },
                        { "createRule", 
                        {
                            { "interval", 24 },
                            { "intervalUnit", "HOURS" },
                            { "times", "23:45" },
                        } },
                        { "name", "2 weeks of daily snapshots" },
                        { "retainRule", 
                        {
                            { "count", 14 },
                        } },
                        { "tagsToAdd", 
                        {
                            { "SnapshotCreator", "DLM" },
                        } },
                    },
                },
                TargetTags = 
                {
                    { "Snapshot", "true" },
                },
            },
            State = "ENABLED",
        });
    }

}

