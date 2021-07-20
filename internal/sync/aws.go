package sync

import (
	"context"
	"encoding/json"
	"github.com/crashdump/birdview/internal/model"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	eksTypes "github.com/aws/aws-sdk-go-v2/service/eks/types"
	elasticacheTypes "github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	elasticsearchTypes "github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	rdsTypes "github.com/aws/aws-sdk-go-v2/service/rds/types"
)

// Create an IAM role and attach the policy "AWSConfigRoleForOrganizations"

// Create an AWS config aggregator with the following parameters in the AWS root account.
// Name: "birdview"
// Source accounts: "Add to my organisation"
// Choose IAM role: "Chose a role from my account" / "birdview-config-role"
// Choose regions: "Select all regions" / "Include future regions"

//var eol map[string]map[string]string{
//	"AWS::EKS::Cluster": {
//		"1.14": "",
//		"1.15": "",
//		"1.16": "",
//	},
//	"AWS::Elasticsearch::Domain": {
//		"1.14": "",
//		"1.15": "",
//		"1.16": "",
//	},
//}

type clients struct {
	ConfigClient configservice.Client
}

func Init() clients {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("eu-west-1"),
		config.WithSharedConfigProfile("birdview"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return clients{
		 ConfigClient: *configservice.NewFromConfig(cfg),
	}
}

func (c *clients) GetResources() model.Resources {
	resourceIds, err := c.getDiscoveredResources()
	if err != nil {
		log.Fatalf(err.Error())
	}

	resourceConfigs, err := c.getResourcesConfig(resourceIds)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var results model.Resources

	for _, r := range resourceConfigs {
		if r.ResourceType == "AWS::EKS::Cluster" {
			var configuration eksTypes.Cluster
			_ = json.Unmarshal([]byte(*r.Configuration), &configuration)

			results = append(results, model.Resource{
				Account: *r.AccountId,
				Name:    *r.ResourceName,
				Type:    "AWS::EKS::Cluster",
				Region:  *r.AwsRegion,
				Version: *configuration.Version,
			})
		}

		if r.ResourceType == "AWS::RDS::Cluster" {
			var configuration rdsTypes.DBInstance
			_ = json.Unmarshal([]byte(*r.Configuration), &configuration)

			results = append(results, model.Resource{
				Account: *r.AccountId,
				Name:    *r.ResourceName,
				Type:    "AWS::RDS::Cluster",
				Region:  *r.AwsRegion,
				Version: *configuration.EngineVersion,
			})
		}

		if r.ResourceType == "AWS::Elasticsearch::Domain" {
			var configuration elasticsearchTypes.ElasticsearchDomainStatus
			_ = json.Unmarshal([]byte(*r.Configuration), &configuration)

			results = append(results, model.Resource{
				Account: *r.AccountId,
				Name:    *r.ResourceName,
				Type:    "AWS::Elasticsearch::Domain",
				Region:  *r.AwsRegion,
				Version: *configuration.ElasticsearchVersion,
			})
		}

		if r.ResourceType == "AWS::ElastiCache::CacheCluster" {
			var configuration elasticacheTypes.CacheCluster
			_ = json.Unmarshal([]byte(*r.Configuration), &configuration)

			results = append(results, model.Resource{
				Account: *r.AccountId,
				Name:    *r.ResourceName,
				Type:    "AWS::ElastiCache::CacheCluster",
				Region:  *r.AwsRegion,
				Version: *configuration.EngineVersion,
			})
		}
	}

	return results
}

