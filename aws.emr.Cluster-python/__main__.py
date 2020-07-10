import pulumi
import pulumi_aws as aws

cluster = aws.emr.Cluster("cluster",
    additional_info="""{
  "instanceAwsClientConfiguration": {
    "proxyPort": 8099,
    "proxyHost": "myproxy.example.com"
  }
}

""",
    applications=["Spark"],
    bootstrap_actions=[{
        "args": [
            "instance.isMaster=true",
            "echo running on master node",
        ],
        "name": "runif",
        "path": "s3://elasticmapreduce/bootstrap-actions/run-if",
    }],
    configurations_json="""  [
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

""",
    core_instance_group={
        "autoscaling_policy": """{
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

""",
        "bid_price": "0.30",
        "ebsConfig": [{
            "size": "40",
            "type": "gp2",
            "volumesPerInstance": 1,
        }],
        "instance_count": 1,
        "instance_type": "c4.large",
    },
    ebs_root_volume_size=100,
    ec2_attributes={
        "emrManagedMasterSecurityGroup": aws_security_group["sg"]["id"],
        "emrManagedSlaveSecurityGroup": aws_security_group["sg"]["id"],
        "instanceProfile": aws_iam_instance_profile["emr_profile"]["arn"],
        "subnet_id": aws_subnet["main"]["id"],
    },
    keep_job_flow_alive_when_no_steps=True,
    master_instance_group={
        "instance_type": "m4.large",
    },
    release_label="emr-4.6.0",
    service_role=aws_iam_role["iam_emr_service_role"]["arn"],
    tags={
        "env": "env",
        "role": "rolename",
    },
    termination_protection=False)

