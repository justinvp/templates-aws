package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/athena"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		hogeBucket, err := s3.NewBucket(ctx, "hogeBucket", nil)
		if err != nil {
			return err
		}
		_, err = athena.NewDatabase(ctx, "hogeDatabase", &athena.DatabaseArgs{
			Bucket: hogeBucket.Bucket,
			Name:   pulumi.String("database_name"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

