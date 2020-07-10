using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = new Aws.Batch.JobDefinition("test", new Aws.Batch.JobDefinitionArgs
        {
            ContainerProperties = @"{
	""command"": [""ls"", ""-la""],
	""image"": ""busybox"",
	""memory"": 1024,
	""vcpus"": 1,
	""volumes"": [
      {
        ""host"": {
          ""sourcePath"": ""/tmp""
        },
        ""name"": ""tmp""
      }
    ],
	""environment"": [
		{""name"": ""VARNAME"", ""value"": ""VARVAL""}
	],
	""mountPoints"": [
		{
          ""sourceVolume"": ""tmp"",
          ""containerPath"": ""/tmp"",
          ""readOnly"": false
        }
	],
    ""ulimits"": [
      {
        ""hardLimit"": 1024,
        ""name"": ""nofile"",
        ""softLimit"": 1024
      }
    ]
}

",
            Type = "container",
        });
    }

}

