package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		test, err := apigateway.NewRestApi(ctx, "test", nil)
		if err != nil {
			return err
		}
		myusageplan, err := apigateway.NewUsagePlan(ctx, "myusageplan", &apigateway.UsagePlanArgs{
			ApiStages: apigateway.UsagePlanApiStageArray{
				&apigateway.UsagePlanApiStageArgs{
					ApiId: test.ID(),
					Stage: pulumi.Any(aws_api_gateway_deployment.Foo.Stage_name),
				},
			},
		})
		if err != nil {
			return err
		}
		mykey, err := apigateway.NewApiKey(ctx, "mykey", nil)
		if err != nil {
			return err
		}
		_, err = apigateway.NewUsagePlanKey(ctx, "main", &apigateway.UsagePlanKeyArgs{
			KeyId:       mykey.ID(),
			KeyType:     pulumi.String("API_KEY"),
			UsagePlanId: myusageplan.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

