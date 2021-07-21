package sync

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	configTypes "github.com/aws/aws-sdk-go-v2/service/configservice/types"
	eksTypes "github.com/aws/aws-sdk-go-v2/service/eks/types"
	elasticacheTypes "github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	elasticsearchTypes "github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	rdsTypes "github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/crashdump/birdview/internal/model"
	"log"
)


func (c *clients) GetResources() model.Resources {
	resourceIds, err := c.configDiscoverResources()
	if err != nil {
		log.Fatalf(err.Error())
	}

	resourceConfigs, err := c.configGetResourceDetails(resourceIds)
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

		if r.ResourceType == "AWS::RDS::DBCluster" {
			var configuration rdsTypes.DBInstance
			_ = json.Unmarshal([]byte(*r.Configuration), &configuration)

			results = append(results, model.Resource{
				Account: *r.AccountId,
				Name:    *r.ResourceName,
				Type:    "AWS::RDS::DBCluster",
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

func (c *clients) configGetResourceDetails(resourceIdentifiers []configTypes.AggregateResourceIdentifier) ([]configTypes.BaseConfigurationItem, error) {
	aggregatorName := "callsign-sub-accounts"
	var results []configTypes.BaseConfigurationItem

	batch := 100
	for i := 0; i < len(resourceIdentifiers); i += 100 {
		j := i + batch
		if j > len(resourceIdentifiers) {
			j = len(resourceIdentifiers)
		}

		input := &configservice.BatchGetAggregateResourceConfigInput{
			ConfigurationAggregatorName: &aggregatorName,
			ResourceIdentifiers:         resourceIdentifiers[i:j],
		}

		r, err := c.ConfigClient.BatchGetAggregateResourceConfig(context.TODO(), input)
		if err != nil {
			log.Fatalf("Error BatchGetAggregateResourceConfig: %v\n", err)
			return nil, err
		}
		results = append(results, r.BaseConfigurationItems...)
	}

	return results, nil
}


func (c *clients) configDiscoverResources() ([]configTypes.AggregateResourceIdentifier, error) {
	aggregatorName := "callsign-sub-accounts"
	var results []configTypes.AggregateResourceIdentifier

	// List of resource types pulled from
	// github.com/sync/sync-sdk-go/models/apis/config/2014-11-12/api-2.json
	var resourceTypes = [...]configTypes.ResourceType{
		"AWS::RDS::DBCluster",
		"AWS::EKS::Cluster",
		"AWS::ElastiCache::CacheCluster",
		"AWS::Elasticsearch::Domain",
		"AWS::MSK::Cluster",
	}

	for _, t := range &resourceTypes {
		input := &configservice.ListAggregateDiscoveredResourcesInput{
			ResourceType: t,
			ConfigurationAggregatorName: &aggregatorName,
		}

		result, err := c.ConfigClient.ListAggregateDiscoveredResources(context.TODO(), input)
		if err != nil {
			log.Fatalf("Error ListAggregateDiscoveredResources (ResourceType: %s): %v\n", t, err)
			return nil, err
		}
		results = append(results, result.ResourceIdentifiers...)
	}


	return results, nil
}