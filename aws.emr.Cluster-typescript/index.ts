import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const cluster = new aws.emr.Cluster("cluster", {
    additionalInfo: `{
  "instanceAwsClientConfiguration": {
    "proxyPort": 8099,
    "proxyHost": "myproxy.example.com"
  }
}
`,
    applications: ["Spark"],
    bootstrapActions: [{
        args: [
            "instance.isMaster=true",
            "echo running on master node",
        ],
        name: "runif",
        path: "s3://elasticmapreduce/bootstrap-actions/run-if",
    }],
    configurationsJson: `  [
    {
      "Classification": "hadoop-env",
      "Configurations": [
        {
          "Classification": "export",
          "Properties": {
            "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
          }
        }
      ],
      "Properties": {}
    },
    {
      "Classification": "spark-env",
      "Configurations": [
        {
          "Classification": "export",
          "Properties": {
            "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
          }
        }
      ],
      "Properties": {}
    }
  ]
`,
    coreInstanceGroup: {
        autoscalingPolicy: `{
"Constraints": {
  "MinCapacity": 1,
  "MaxCapacity": 2
},
"Rules": [
  {
    "Name": "ScaleOutMemoryPercentage",
    "Description": "Scale out if YARNMemoryAvailablePercentage is less than 15",
    "Action": {
      "SimpleScalingPolicyConfiguration": {
        "AdjustmentType": "CHANGE_IN_CAPACITY",
        "ScalingAdjustment": 1,
        "CoolDown": 300
      }
    },
    "Trigger": {
      "CloudWatchAlarmDefinition": {
        "ComparisonOperator": "LESS_THAN",
        "EvaluationPeriods": 1,
        "MetricName": "YARNMemoryAvailablePercentage",
        "Namespace": "AWS/ElasticMapReduce",
        "Period": 300,
        "Statistic": "AVERAGE",
        "Threshold": 15.0,
        "Unit": "PERCENT"
      }
    }
  }
]
}
`,
        bidPrice: "0.30",
        ebsConfigs: [{
            size: 40,
            type: "gp2",
            volumesPerInstance: 1,
        }],
        instanceCount: 1,
        instanceType: "c4.large",
    },
    ebsRootVolumeSize: 100,
    ec2Attributes: {
        emrManagedMasterSecurityGroup: aws_security_group_sg.id,
        emrManagedSlaveSecurityGroup: aws_security_group_sg.id,
        instanceProfile: aws_iam_instance_profile_emr_profile.arn,
        subnetId: aws_subnet_main.id,
    },
    keepJobFlowAliveWhenNoSteps: true,
    masterInstanceGroup: {
        instanceType: "m4.large",
    },
    releaseLabel: "emr-4.6.0",
    serviceRole: aws_iam_role_iam_emr_service_role.arn,
    tags: {
        env: "env",
        role: "rolename",
    },
    terminationProtection: false,
});

