using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = Output.Create(Aws.Ec2.GetLaunchTemplate.InvokeAsync(new Aws.Ec2.GetLaunchTemplateArgs
        {
            Filters = 
            {
                new Aws.Ec2.Inputs.GetLaunchTemplateFilterArgs
                {
                    Name = "launch-template-name",
                    Values = 
                    {
                        "some-template",
                    },
                },
            },
        }));
    }

}

