package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elasticbeanstalk"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tftest, err := elasticbeanstalk.NewApplication(ctx, "tftest", &elasticbeanstalk.ApplicationArgs{
			Description: pulumi.String("tf-test-desc"),
		})
		if err != nil {
			return err
		}
		_, err = elasticbeanstalk.NewConfigurationTemplate(ctx, "tfTemplate", &elasticbeanstalk.ConfigurationTemplateArgs{
			Application:       tftest.Name,
			SolutionStackName: pulumi.String("64bit Amazon Linux 2015.09 v2.0.8 running Go 1.4"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

