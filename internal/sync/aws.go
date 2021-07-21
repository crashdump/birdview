package sync

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/health"
)

// Create an IAM role and attach the policy "AWSConfigRoleForOrganizations"

// Create an AWS config aggregator with the following parameters in the AWS root account.
// Name: "birdview"
// Source accounts: "Add to my organisation"
// Choose IAM role: "Chose a role from my account" / "birdview-config-role"
// Choose regions: "Select all regions" / "Include future regions"

type clients struct {
	ConfigClient configservice.Client
	HealthClient health.Client
}

func Init() clients {
	awsConfigCfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("eu-west-1"),
		config.WithSharedConfigProfile("birdview"),
	)
	if err != nil {
		log.Fatalf("unable to load AWS SDK config, %v", err)
	}

	awsHealthCfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithSharedConfigProfile("birdview"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return clients{
		 ConfigClient: *configservice.NewFromConfig(awsConfigCfg),
		 HealthClient: *health.NewFromConfig(awsHealthCfg),
	}
}
