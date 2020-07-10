using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.GetAmi.InvokeAsync(new Aws.GetAmiArgs
        {
            ExecutableUsers = 
            {
                "self",
            },
            Filters = 
            {
                new Aws.Inputs.GetAmiFilterArgs
                {
                    Name = "name",
                    Values = 
                    {
                        "myami-*",
                    },
                },
                new Aws.Inputs.GetAmiFilterArgs
                {
                    Name = "root-device-type",
                    Values = 
                    {
                        "ebs",
                    },
                },
                new Aws.Inputs.GetAmiFilterArgs
                {
                    Name = "virtualization-type",
                    Values = 
                    {
                        "hvm",
                    },
                },
            },
            MostRecent = true,
            NameRegex = "^myami-\\d{3}",
            Owners = 
            {
                "self",
            },
        }));
    }

}

