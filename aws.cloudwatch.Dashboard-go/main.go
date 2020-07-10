package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewDashboard(ctx, "main", &cloudwatch.DashboardArgs{
			DashboardBody: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", " {\n", "   \"widgets\": [\n", "       {\n", "          \"type\":\"metric\",\n", "          \"x\":0,\n", "          \"y\":0,\n", "          \"width\":12,\n", "          \"height\":6,\n", "          \"properties\":{\n", "             \"metrics\":[\n", "                [\n", "                   \"AWS/EC2\",\n", "                   \"CPUUtilization\",\n", "                   \"InstanceId\",\n", "                   \"i-012345\"\n", "                ]\n", "             ],\n", "             \"period\":300,\n", "             \"stat\":\"Average\",\n", "             \"region\":\"us-east-1\",\n", "             \"title\":\"EC2 Instance CPU\"\n", "          }\n", "       },\n", "       {\n", "          \"type\":\"text\",\n", "          \"x\":0,\n", "          \"y\":7,\n", "          \"width\":3,\n", "          \"height\":3,\n", "          \"properties\":{\n", "             \"markdown\":\"Hello world\"\n", "          }\n", "       }\n", "   ]\n", " }\n", " \n")),
			DashboardName: pulumi.String("my-dashboard"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

