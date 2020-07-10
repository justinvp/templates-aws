using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var ni = new Aws.Sagemaker.NotebookInstance("ni", new Aws.Sagemaker.NotebookInstanceArgs
        {
            InstanceType = "ml.t2.medium",
            RoleArn = aws_iam_role.Role.Arn,
            Tags = 
            {
                { "Name", "foo" },
            },
        });
    }

}

