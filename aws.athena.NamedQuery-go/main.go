package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/athena"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		hogeBucket, err := s3.NewBucket(ctx, "hogeBucket", nil)
		if err != nil {
			return err
		}
		testKey, err := kms.NewKey(ctx, "testKey", &kms.KeyArgs{
			DeletionWindowInDays: pulumi.Int(7),
			Description:          pulumi.String("Athena KMS Key"),
		})
		if err != nil {
			return err
		}
		testWorkgroup, err := athena.NewWorkgroup(ctx, "testWorkgroup", &athena.WorkgroupArgs{
			Configuration: &athena.WorkgroupConfigurationArgs{
				ResultConfiguration: &athena.WorkgroupConfigurationResultConfigurationArgs{
					EncryptionConfiguration: &athena.WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs{
						EncryptionOption: pulumi.String("SSE_KMS"),
						KmsKeyArn:        testKey.Arn,
					},
				},
			},
		})
		if err != nil {
			return err
		}
		hogeDatabase, err := athena.NewDatabase(ctx, "hogeDatabase", &athena.DatabaseArgs{
			Bucket: hogeBucket.ID(),
			Name:   pulumi.String("users"),
		})
		if err != nil {
			return err
		}
		_, err = athena.NewNamedQuery(ctx, "foo", &athena.NamedQueryArgs{
			Database: hogeDatabase.Name,
			Query: hogeDatabase.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("%v%v%v", "SELECT * FROM ", name, " limit 10;"), nil
			}).(pulumi.StringOutput),
			Workgroup: testWorkgroup.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

