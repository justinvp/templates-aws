package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myapi, err := apigateway.NewRestApi(ctx, "myapi", nil)
		if err != nil {
			return err
		}
		dev, err := apigateway.NewDeployment(ctx, "dev", &apigateway.DeploymentArgs{
			RestApi:   myapi.ID(),
			StageName: pulumi.String("dev"),
		})
		if err != nil {
			return err
		}
		prod, err := apigateway.NewDeployment(ctx, "prod", &apigateway.DeploymentArgs{
			RestApi:   myapi.ID(),
			StageName: pulumi.String("prod"),
		})
		if err != nil {
			return err
		}
		_, err = apigateway.NewUsagePlan(ctx, "myUsagePlan", &apigateway.UsagePlanArgs{
			ApiStages: apigateway.UsagePlanApiStageArray{
				&apigateway.UsagePlanApiStageArgs{
					ApiId: myapi.ID(),
					Stage: dev.StageName,
				},
				&apigateway.UsagePlanApiStageArgs{
					ApiId: myapi.ID(),
					Stage: prod.StageName,
				},
			},
			Description: pulumi.String("my description"),
			ProductCode: pulumi.String("MYCODE"),
			QuotaSettings: &apigateway.UsagePlanQuotaSettingsArgs{
				Limit:  pulumi.Int(20),
				Offset: pulumi.Int(2),
				Period: pulumi.String("WEEK"),
			},
			ThrottleSettings: &apigateway.UsagePlanThrottleSettingsArgs{
				BurstLimit: pulumi.Int(5),
				RateLimit:  pulumi.Float64(10),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

