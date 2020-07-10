using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var cluster = new Aws.Emr.Cluster("cluster", new Aws.Emr.ClusterArgs
        {
            AdditionalInfo = @"{
  ""instanceAwsClientConfiguration"": {
    ""proxyPort"": 8099,
    ""proxyHost"": ""myproxy.example.com""
  }
}

",
            Applications = 
            {
                "Spark",
            },
            BootstrapActions = 
            {
                new Aws.Emr.Inputs.ClusterBootstrapActionArgs
                {
                    Args = 
                    {
                        "instance.isMaster=true",
                        "echo running on master node",
                    },
                    Name = "runif",
                    Path = "s3://elasticmapreduce/bootstrap-actions/run-if",
                },
            },
            ConfigurationsJson = @"  [
    {
      ""Classification"": ""hadoop-env"",
      ""Configurations"": [
        {
          ""Classification"": ""export"",
          ""Properties"": {
            ""JAVA_HOME"": ""/usr/lib/jvm/java-1.8.0""
          }
        }
      ],
      ""Properties"": {}
    },
    {
      ""Classification"": ""spark-env"",
      ""Configurations"": [
        {
          ""Classification"": ""export"",
          ""Properties"": {
            ""JAVA_HOME"": ""/usr/lib/jvm/java-1.8.0""
          }
        }
      ],
      ""Properties"": {}
    }
  ]

",
            CoreInstanceGroup = new Aws.Emr.Inputs.ClusterCoreInstanceGroupArgs
            {
                AutoscalingPolicy = @"{
""Constraints"": {
  ""MinCapacity"": 1,
  ""MaxCapacity"": 2
},
""Rules"": [
  {
    ""Name"": ""ScaleOutMemoryPercentage"",
    ""Description"": ""Scale out if YARNMemoryAvailablePercentage is less than 15"",
    ""Action"": {
      ""SimpleScalingPolicyConfiguration"": {
        ""AdjustmentType"": ""CHANGE_IN_CAPACITY"",
        ""ScalingAdjustment"": 1,
        ""CoolDown"": 300
      }
    },
    ""Trigger"": {
      ""CloudWatchAlarmDefinition"": {
        ""ComparisonOperator"": ""LESS_THAN"",
        ""EvaluationPeriods"": 1,
        ""MetricName"": ""YARNMemoryAvailablePercentage"",
        ""Namespace"": ""AWS/ElasticMapReduce"",
        ""Period"": 300,
        ""Statistic"": ""AVERAGE"",
        ""Threshold"": 15.0,
        ""Unit"": ""PERCENT""
      }
    }
  }
]
}

",
                BidPrice = "0.30",
                EbsConfig = 
                {
                    
                    {
                        { "size", "40" },
                        { "type", "gp2" },
                        { "volumesPerInstance", 1 },
                    },
                },
                InstanceCount = 1,
                InstanceType = "c4.large",
            },
            EbsRootVolumeSize = 100,
            Ec2Attributes = new Aws.Emr.Inputs.ClusterEc2AttributesArgs
            {
                EmrManagedMasterSecurityGroup = aws_security_group.Sg.Id,
                EmrManagedSlaveSecurityGroup = aws_security_group.Sg.Id,
                InstanceProfile = aws_iam_instance_profile.Emr_profile.Arn,
                SubnetId = aws_subnet.Main.Id,
            },
            KeepJobFlowAliveWhenNoSteps = true,
            MasterInstanceGroup = new Aws.Emr.Inputs.ClusterMasterInstanceGroupArgs
            {
                InstanceType = "m4.large",
            },
            ReleaseLabel = "emr-4.6.0",
            ServiceRole = aws_iam_role.Iam_emr_service_role.Arn,
            Tags = 
            {
                { "env", "env" },
                { "role", "rolename" },
            },
            TerminationProtection = false,
        });
    }

}

