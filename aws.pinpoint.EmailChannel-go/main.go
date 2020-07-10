package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/pinpoint"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		app, err := pinpoint.NewApp(ctx, "app", nil)
		if err != nil {
			return err
		}
		identity, err := ses.NewDomainIdentity(ctx, "identity", &ses.DomainIdentityArgs{
			Domain: pulumi.String("example.com"),
		})
		if err != nil {
			return err
		}
		role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"sts:AssumeRole\",\n", "      \"Principal\": {\n", "        \"Service\": \"pinpoint.amazonaws.com\"\n", "      },\n", "      \"Effect\": \"Allow\",\n", "      \"Sid\": \"\"\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = pinpoint.NewEmailChannel(ctx, "email", &pinpoint.EmailChannelArgs{
			ApplicationId: app.ApplicationId,
			FromAddress:   pulumi.String("user@example.com"),
			Identity:      identity.Arn,
			RoleArn:       role.Arn,
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicy(ctx, "rolePolicy", &iam.RolePolicyArgs{
			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": {\n", "    \"Action\": [\n", "      \"mobileanalytics:PutEvents\",\n", "      \"mobileanalytics:PutItems\"\n", "    ],\n", "    \"Effect\": \"Allow\",\n", "    \"Resource\": [\n", "      \"*\"\n", "    ]\n", "  }\n", "}\n", "\n")),
			Role:   role.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

