using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @group = new Aws.Iam.Group("group", new Aws.Iam.GroupArgs
        {
        });
        var policy = new Aws.Iam.Policy("policy", new Aws.Iam.PolicyArgs
        {
            Description = "A test policy",
            Policy = "",
        });
        // insert policy here
        var test_attach = new Aws.Iam.GroupPolicyAttachment("test-attach", new Aws.Iam.GroupPolicyAttachmentArgs
        {
            Group = @group.Name,
            PolicyArn = policy.Arn,
        });
    }

}

