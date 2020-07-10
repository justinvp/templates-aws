using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = Output.Create(Aws.Kms.GetKey.InvokeAsync(new Aws.Kms.GetKeyArgs
        {
            KeyId = "arn:aws:kms:us-east-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab",
        }));
    }

}

