using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var main = new Aws.CloudWatch.Dashboard("main", new Aws.CloudWatch.DashboardArgs
        {
            DashboardBody = @" {
   ""widgets"": [
       {
          ""type"":""metric"",
          ""x"":0,
          ""y"":0,
          ""width"":12,
          ""height"":6,
          ""properties"":{
             ""metrics"":[
                [
                   ""AWS/EC2"",
                   ""CPUUtilization"",
                   ""InstanceId"",
                   ""i-012345""
                ]
             ],
             ""period"":300,
             ""stat"":""Average"",
             ""region"":""us-east-1"",
             ""title"":""EC2 Instance CPU""
          }
       },
       {
          ""type"":""text"",
          ""x"":0,
          ""y"":7,
          ""width"":3,
          ""height"":3,
          ""properties"":{
             ""markdown"":""Hello world""
          }
       }
   ]
 }
 
",
            DashboardName = "my-dashboard",
        });
    }

}

