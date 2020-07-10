using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Glue.Job("example", new Aws.Glue.JobArgs
        {
            Command = new Aws.Glue.Inputs.JobCommandArgs
            {
                ScriptLocation = $"s3://{aws_s3_bucket.Example.Bucket}/example.py",
            },
            RoleArn = aws_iam_role.Example.Arn,
        });
    }

}

