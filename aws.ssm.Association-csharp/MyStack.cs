using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ssm.Association("example", new Aws.Ssm.AssociationArgs
        {
            Targets = 
            {
                new Aws.Ssm.Inputs.AssociationTargetArgs
                {
                    Key = "InstanceIds",
                    Values = 
                    {
                        aws_instance.Example.Id,
                    },
                },
            },
        });
    }

}

