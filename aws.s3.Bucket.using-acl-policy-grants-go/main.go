package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		currentUser, err := aws.GetCanonicalUserId(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
			Grants: s3.BucketGrantArray{
				&s3.BucketGrantArgs{
					Id: pulumi.String(currentUser.Id),
					Permissions: pulumi.StringArray{
						pulumi.String("FULL_CONTROL"),
					},
					Type: pulumi.String("CanonicalUser"),
				},
				&s3.BucketGrantArgs{
					Permissions: pulumi.StringArray{
						pulumi.String("READ"),
						pulumi.String("WRITE"),
					},
					Type: pulumi.String("Group"),
					Uri:  pulumi.String("http://acs.amazonaws.com/groups/s3/LogDelivery"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

