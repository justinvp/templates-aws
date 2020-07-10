using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.DirectoryService.GetDirectory.InvokeAsync(new Aws.DirectoryService.GetDirectoryArgs
        {
            DirectoryId = aws_directory_service_directory.Main.Id,
        }));
    }

}

