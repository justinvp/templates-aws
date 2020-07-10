package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/emr"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := emr.NewCluster(ctx, "cluster", &emr.ClusterArgs{
			AdditionalInfo: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v", "{\n", "  \"instanceAwsClientConfiguration\": {\n", "    \"proxyPort\": 8099,\n", "    \"proxyHost\": \"myproxy.example.com\"\n", "  }\n", "}\n", "\n")),
			Applications: pulumi.StringArray{
				pulumi.String("Spark"),
			},
			BootstrapActions: emr.ClusterBootstrapActionArray{
				&emr.ClusterBootstrapActionArgs{
					Args: pulumi.StringArray{
						pulumi.String("instance.isMaster=true"),
						pulumi.String("echo running on master node"),
					},
					Name: pulumi.String("runif"),
					Path: pulumi.String("s3://elasticmapreduce/bootstrap-actions/run-if"),
				},
			},
			ConfigurationsJson: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "  [\n", "    {\n", "      \"Classification\": \"hadoop-env\",\n", "      \"Configurations\": [\n", "        {\n", "          \"Classification\": \"export\",\n", "          \"Properties\": {\n", "            \"JAVA_HOME\": \"/usr/lib/jvm/java-1.8.0\"\n", "          }\n", "        }\n", "      ],\n", "      \"Properties\": {}\n", "    },\n", "    {\n", "      \"Classification\": \"spark-env\",\n", "      \"Configurations\": [\n", "        {\n", "          \"Classification\": \"export\",\n", "          \"Properties\": {\n", "            \"JAVA_HOME\": \"/usr/lib/jvm/java-1.8.0\"\n", "          }\n", "        }\n", "      ],\n", "      \"Properties\": {}\n", "    }\n", "  ]\n", "\n")),
			CoreInstanceGroup: &emr.ClusterCoreInstanceGroupArgs{
				AutoscalingPolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "\"Constraints\": {\n", "  \"MinCapacity\": 1,\n", "  \"MaxCapacity\": 2\n", "},\n", "\"Rules\": [\n", "  {\n", "    \"Name\": \"ScaleOutMemoryPercentage\",\n", "    \"Description\": \"Scale out if YARNMemoryAvailablePercentage is less than 15\",\n", "    \"Action\": {\n", "      \"SimpleScalingPolicyConfiguration\": {\n", "        \"AdjustmentType\": \"CHANGE_IN_CAPACITY\",\n", "        \"ScalingAdjustment\": 1,\n", "        \"CoolDown\": 300\n", "      }\n", "    },\n", "    \"Trigger\": {\n", "      \"CloudWatchAlarmDefinition\": {\n", "        \"ComparisonOperator\": \"LESS_THAN\",\n", "        \"EvaluationPeriods\": 1,\n", "        \"MetricName\": \"YARNMemoryAvailablePercentage\",\n", "        \"Namespace\": \"AWS/ElasticMapReduce\",\n", "        \"Period\": 300,\n", "        \"Statistic\": \"AVERAGE\",\n", "        \"Threshold\": 15.0,\n", "        \"Unit\": \"PERCENT\"\n", "      }\n", "    }\n", "  }\n", "]\n", "}\n", "\n")),
				BidPrice:          pulumi.String("0.30"),
				EbsConfig: pulumi.MapArray{
					pulumi.Map{
						"size":               pulumi.String("40"),
						"type":               pulumi.String("gp2"),
						"volumesPerInstance": pulumi.Float64(1),
					},
				},
				InstanceCount: pulumi.Int(1),
				InstanceType:  pulumi.String("c4.large"),
			},
			EbsRootVolumeSize: pulumi.Int(100),
			Ec2Attributes: &emr.ClusterEc2AttributesArgs{
				EmrManagedMasterSecurityGroup: pulumi.Any(aws_security_group.Sg.Id),
				EmrManagedSlaveSecurityGroup:  pulumi.Any(aws_security_group.Sg.Id),
				InstanceProfile:               pulumi.Any(aws_iam_instance_profile.Emr_profile.Arn),
				SubnetId:                      pulumi.Any(aws_subnet.Main.Id),
			},
			KeepJobFlowAliveWhenNoSteps: pulumi.Bool(true),
			MasterInstanceGroup: &emr.ClusterMasterInstanceGroupArgs{
				InstanceType: pulumi.String("m4.large"),
			},
			ReleaseLabel: pulumi.String("emr-4.6.0"),
			ServiceRole:  pulumi.Any(aws_iam_role.Iam_emr_service_role.Arn),
			Tags: pulumi.StringMap{
				"env":  pulumi.String("env"),
				"role": pulumi.String("rolename"),
			},
			TerminationProtection: pulumi.Bool(false),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

