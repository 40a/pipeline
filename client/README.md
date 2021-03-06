# Go API client for client

Pipeline v0.3.0 swagger

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.3.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:
```
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:
```golang
import "./client"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:9090*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApplicationsApi* | [**CreateApplication**](docs/ApplicationsApi.md#createapplication) | **Post** /api/v1/orgs/{orgId}/applications | Create new application based on catalog
*ApplicationsApi* | [**GetApplication**](docs/ApplicationsApi.md#getapplication) | **Get** /api/v1/orgs/{orgId}/applications/{appId} | Get application details
*ApplicationsApi* | [**ListApplications**](docs/ApplicationsApi.md#listapplications) | **Get** /api/v1/orgs/{orgId}/applications | List application catalogs
*AuthApi* | [**CreateToken**](docs/AuthApi.md#createtoken) | **Post** /api/v1/tokens | Create token
*AuthApi* | [**DeleteToken**](docs/AuthApi.md#deletetoken) | **Delete** /api/v1/tokens/{tokenId} | Delete an API token
*AuthApi* | [**ListTokens**](docs/AuthApi.md#listtokens) | **Get** /api/v1/tokens | List all API tokens
*CatalogsApi* | [**GetCatalogDetail**](docs/CatalogsApi.md#getcatalogdetail) | **Get** /api/v1/orgs/{orgId}/catalogs/{name} | Get catalog details
*CatalogsApi* | [**ListCatalogs**](docs/CatalogsApi.md#listcatalogs) | **Get** /api/v1/orgs/{orgId}/catalogs | List catalogs
*CatalogsApi* | [**UpdateCatalogs**](docs/CatalogsApi.md#updatecatalogs) | **Put** /api/v1/orgs/{orgId}/catalogs/update | Update repository for catalog
*ClustersApi* | [**ClusterPostHooks**](docs/ClustersApi.md#clusterposthooks) | **Put** /api/v1/orgs/{orgId}/clusters/{id}/posthooks | Run posthook functions
*ClustersApi* | [**CreateCluster**](docs/ClustersApi.md#createcluster) | **Post** /api/v1/orgs/{orgId}/clusters | Create cluster
*ClustersApi* | [**DeleteCluster**](docs/ClustersApi.md#deletecluster) | **Delete** /api/v1/orgs/{orgId}/clusters/{id} | Delete cluster
*ClustersApi* | [**GetAPIEndpoint**](docs/ClustersApi.md#getapiendpoint) | **Get** /api/v1/orgs/{orgId}/clusters/{id}/apiendpoint | Get API endpoint
*ClustersApi* | [**GetCluster**](docs/ClustersApi.md#getcluster) | **Get** /api/v1/orgs/{orgId}/clusters/{id} | Get cluster status
*ClustersApi* | [**GetClusterApplication**](docs/ClustersApi.md#getclusterapplication) | **Get** /api/v1/orgs/{orgId}/clusters/{id}/application | Get cluster&#39;s applications
*ClustersApi* | [**GetClusterConfig**](docs/ClustersApi.md#getclusterconfig) | **Get** /api/v1/orgs/{orgId}/clusters/{id}/config | Get a cluster config
*ClustersApi* | [**GetClusterDetails**](docs/ClustersApi.md#getclusterdetails) | **Get** /api/v1/orgs/{orgId}/clusters/{id}/details | Get cluster details
*ClustersApi* | [**GetClusterStatus**](docs/ClustersApi.md#getclusterstatus) | **Head** /api/v1/orgs/{orgId}/clusters/{id} | Get cluster status
*ClustersApi* | [**HelmInit**](docs/ClustersApi.md#helminit) | **Post** /api/v1/orgs/{orgId}/clusters/{id}/helminit | Initialize Helm
*ClustersApi* | [**InstallSecrets**](docs/ClustersApi.md#installsecrets) | **Post** /api/v1/orgs/{orgId}/clusters/{id}/secrets | Install secrets into cluster
*ClustersApi* | [**ListClusters**](docs/ClustersApi.md#listclusters) | **Get** /api/v1/orgs/{orgId}/clusters | List clusters
*ClustersApi* | [**ListEndpoints**](docs/ClustersApi.md#listendpoints) | **Get** /api/v1/orgs/{orgId}/clusters/{id}/endpoints | List service public endpoints
*ClustersApi* | [**ListNodes**](docs/ClustersApi.md#listnodes) | **Get** /api/v1/orgs/{orgId}/clusters/{id}/nodes | List cluser nodes
*ClustersApi* | [**UpdateCluster**](docs/ClustersApi.md#updatecluster) | **Put** /api/v1/orgs/{orgId}/clusters/{id} | Update cluster
*ClustersApi* | [**UpdateMonitoring**](docs/ClustersApi.md#updatemonitoring) | **Post** /api/v1/orgs/{orgId}/clusters/{id}/monitoring | Update monitoring
*CommonApi* | [**ListEndpoints**](docs/CommonApi.md#listendpoints) | **Get** /api | List Pipeline API endpoints
*DeploymentApi* | [**DeleteDeployment**](docs/DeploymentApi.md#deletedeployment) | **Delete** /api/v1/orgs/{orgId}/clusters/{id}/deployments/{name} | Delete deployment
*DeploymentApi* | [**GetDeployment**](docs/DeploymentApi.md#getdeployment) | **Get** /api/v1/orgs/{orgId}/clusters/{id}/deployments/{name} | Get deployment details
*DeploymentApi* | [**HelmDeploymentStatus**](docs/DeploymentApi.md#helmdeploymentstatus) | **Head** /api/v1/orgs/{orgId}/clusters/{id}/deployments/{name} | Check deployment status
*DeploymentApi* | [**UpdateDeployment**](docs/DeploymentApi.md#updatedeployment) | **Put** /api/v1/orgs/{orgId}/clusters/{id}/deployments/{name} | Update deployment
*DeploymentsApi* | [**CreateDeployment**](docs/DeploymentsApi.md#createdeployment) | **Post** /api/v1/orgs/{orgId}/clusters/{id}/deployments | Create a Helm deployment
*DeploymentsApi* | [**GetTillerStatus**](docs/DeploymentsApi.md#gettillerstatus) | **Head** /api/v1/orgs/{orgId}/clusters/{id}/deployments | Get tiller status
*DeploymentsApi* | [**ListDeployments**](docs/DeploymentsApi.md#listdeployments) | **Get** /api/v1/orgs/{orgId}/clusters/{id}/deployments | List deployments
*HelmApi* | [**HelmChartDetails**](docs/HelmApi.md#helmchartdetails) | **Get** /api/v1/orgs/{orgId}/helm/chart/{repoName}/{chartName} | Chart details
*HelmApi* | [**HelmChartList**](docs/HelmApi.md#helmchartlist) | **Get** /api/v1/orgs/{orgId}/helm/charts/ | Chart List
*HelmApi* | [**HelmInit**](docs/HelmApi.md#helminit) | **Get** /api/v1/orgs/{orgId}/helm/repos | List repositories
*HelmApi* | [**HelmReposAdd**](docs/HelmApi.md#helmreposadd) | **Post** /api/v1/orgs/{orgId}/helm/repos | Add Repo
*HelmApi* | [**HelmReposDelete**](docs/HelmApi.md#helmreposdelete) | **Delete** /api/v1/orgs/{orgId}/helm/repos/{repoName} | Delete Repo
*HelmApi* | [**HelmReposModify**](docs/HelmApi.md#helmreposmodify) | **Put** /api/v1/orgs/{orgId}/helm/repos/{repoName} | Modify Repo
*HelmApi* | [**HelmReposUpdate**](docs/HelmApi.md#helmreposupdate) | **Put** /api/v1/orgs/{orgId}/helm/repos/{repoName}/update | Update Repo
*InfoApi* | [**GetAmazonConfig**](docs/InfoApi.md#getamazonconfig) | **Get** /api/v1/orgs/{orgId}/cloudinfo/amazon | Get all amazon config
*InfoApi* | [**GetAzureConfig**](docs/InfoApi.md#getazureconfig) | **Get** /api/v1/orgs/{orgId}/cloudinfo/azure | Get all azure config
*InfoApi* | [**GetGoogleConfig**](docs/InfoApi.md#getgoogleconfig) | **Get** /api/v1/orgs/{orgId}/cloudinfo/google | Get all google config
*InfoApi* | [**GetSupportedClouds**](docs/InfoApi.md#getsupportedclouds) | **Get** /api/v1/orgs/{orgId}/cloudinfo | Get supported cloud types
*OrganizationsApi* | [**CreateOrg**](docs/OrganizationsApi.md#createorg) | **Post** /api/v1/orgs | Create organization
*OrganizationsApi* | [**GetOrg**](docs/OrganizationsApi.md#getorg) | **Get** /api/v1/orgs/{orgId} | Get organization
*OrganizationsApi* | [**ListOrgs**](docs/OrganizationsApi.md#listorgs) | **Get** /api/v1/orgs | List organizations
*ProfilesApi* | [**AddProfiles**](docs/ProfilesApi.md#addprofiles) | **Post** /api/v1/orgs/{orgId}/profiles/cluster | Add cluster profiles
*ProfilesApi* | [**DeleteProfiles**](docs/ProfilesApi.md#deleteprofiles) | **Delete** /api/v1/orgs/{orgId}/profiles/cluster/{type}/{name} | Delete cluster profiles
*ProfilesApi* | [**ListProfiles**](docs/ProfilesApi.md#listprofiles) | **Get** /api/v1/orgs/{orgId}/profiles/cluster/{type} | List cluster profiles
*ProfilesApi* | [**UpdateProfiles**](docs/ProfilesApi.md#updateprofiles) | **Put** /api/v1/orgs/{orgId}/profiles/cluster | Update cluster profiles
*SecretsApi* | [**AddSecrets**](docs/SecretsApi.md#addsecrets) | **Post** /api/v1/orgs/{orgId}/secrets | Add secrets
*SecretsApi* | [**AllowedSecretsTypes**](docs/SecretsApi.md#allowedsecretstypes) | **Get** /api/v1/allowed/secrets | List allowed secret types
*SecretsApi* | [**AllowedSecretsTypesKeys**](docs/SecretsApi.md#allowedsecretstypeskeys) | **Get** /api/v1/allowed/secrets/{type} | List required keys
*SecretsApi* | [**DeleteSecrets**](docs/SecretsApi.md#deletesecrets) | **Delete** /api/v1/orgs/{orgId}/secrets/{secretId} | Delete secrets
*SecretsApi* | [**GetSecret**](docs/SecretsApi.md#getsecret) | **Get** /api/v1/orgs/{orgId}/secrets/{secretId} | Get secret
*SecretsApi* | [**GetSecrets**](docs/SecretsApi.md#getsecrets) | **Get** /api/v1/orgs/{orgId}/secrets | List secrets
*SecretsApi* | [**UpdateSecrets**](docs/SecretsApi.md#updatesecrets) | **Put** /api/v1/orgs/{orgId}/secrets/{secretId} | Update secrets
*SecretsApi* | [**ValidateSecret**](docs/SecretsApi.md#validatesecret) | **Get** /api/v1/orgs/{orgId}/secrets/{secretId}/validate | Validate secret
*StorageApi* | [**CreateObjectStoreBucket**](docs/StorageApi.md#createobjectstorebucket) | **Post** /api/v1/orgs/{orgId}/buckets | Creates a new object store bucket with the given params
*StorageApi* | [**DeleteObjectStoreBucket**](docs/StorageApi.md#deleteobjectstorebucket) | **Delete** /api/v1/orgs/{orgId}/buckets/{name} | Deletes the object store bucket with the given name
*StorageApi* | [**GetObjectStoreBucketStatus**](docs/StorageApi.md#getobjectstorebucketstatus) | **Head** /api/v1/orgs/{orgId}/buckets/{name} | Get object store bucket status
*StorageApi* | [**ListObjectStoreBuckets**](docs/StorageApi.md#listobjectstorebuckets) | **Get** /api/v1/orgs/{orgId}/buckets | List object storage buckets
*UsersApi* | [**GetUsers**](docs/UsersApi.md#getusers) | **Get** /api/v1/orgs/{orgId}/users/{userId} | Get user
*UsersApi* | [**ListUsers**](docs/UsersApi.md#listusers) | **Get** /api/v1/orgs/{orgId}/users | List users


## Documentation For Models

 - [AddClusterProfileAmazon](docs/AddClusterProfileAmazon.md)
 - [AddClusterProfileAmazonAmazon](docs/AddClusterProfileAmazonAmazon.md)
 - [AddClusterProfileAmazonAmazonNodePools](docs/AddClusterProfileAmazonAmazonNodePools.md)
 - [AddClusterProfileAmazonAmazonNodePoolsPool1](docs/AddClusterProfileAmazonAmazonNodePoolsPool1.md)
 - [AddClusterProfileAzure](docs/AddClusterProfileAzure.md)
 - [AddClusterProfileAzureAzure](docs/AddClusterProfileAzureAzure.md)
 - [AddClusterProfileAzureAzureNodePools](docs/AddClusterProfileAzureAzureNodePools.md)
 - [AddClusterProfileAzureAzureNodePoolsPool1](docs/AddClusterProfileAzureAzureNodePoolsPool1.md)
 - [AddClusterProfileGoogle](docs/AddClusterProfileGoogle.md)
 - [AddClusterProfileGoogleGoogle](docs/AddClusterProfileGoogleGoogle.md)
 - [AddClusterProfileGoogleGoogleNodePools](docs/AddClusterProfileGoogleGoogleNodePools.md)
 - [AddClusterProfileGoogleGoogleNodePoolsPool1](docs/AddClusterProfileGoogleGoogleNodePoolsPool1.md)
 - [AddClusterProfileRequest](docs/AddClusterProfileRequest.md)
 - [AllowedSecretTypeResponse](docs/AllowedSecretTypeResponse.md)
 - [AllowedSecretTypeResponseFields](docs/AllowedSecretTypeResponseFields.md)
 - [AllowedSecretTypesResponse](docs/AllowedSecretTypesResponse.md)
 - [AmazonConfigResponse](docs/AmazonConfigResponse.md)
 - [AmazonConfigResponseImage](docs/AmazonConfigResponseImage.md)
 - [AmazonConfigResponseInstanceType](docs/AmazonConfigResponseInstanceType.md)
 - [ApplicationDeploymentItem](docs/ApplicationDeploymentItem.md)
 - [ApplicationDetailsResponse](docs/ApplicationDetailsResponse.md)
 - [ApplicationListItem](docs/ApplicationListItem.md)
 - [ApplicationListResponse](docs/ApplicationListResponse.md)
 - [AzureBlobStorageProps](docs/AzureBlobStorageProps.md)
 - [AzureConfigResponse](docs/AzureConfigResponse.md)
 - [AzureConfigResponseInstanceType](docs/AzureConfigResponseInstanceType.md)
 - [BaseError400](docs/BaseError400.md)
 - [BaseError500](docs/BaseError500.md)
 - [BasePostHook](docs/BasePostHook.md)
 - [Body](docs/Body.md)
 - [BucketInfo](docs/BucketInfo.md)
 - [CatalogChartInfo](docs/CatalogChartInfo.md)
 - [CatalogChartInfoAnnotations](docs/CatalogChartInfoAnnotations.md)
 - [CatalogDetailsResponse](docs/CatalogDetailsResponse.md)
 - [CatalogDetailsResponseSpotguide](docs/CatalogDetailsResponseSpotguide.md)
 - [CatalogNotFound](docs/CatalogNotFound.md)
 - [CatalogOptionsMonitor](docs/CatalogOptionsMonitor.md)
 - [CatalogOptionsMysqlDatabaseName](docs/CatalogOptionsMysqlDatabaseName.md)
 - [CatalogOptionsMysqlDatabaseSize](docs/CatalogOptionsMysqlDatabaseSize.md)
 - [CatalogOptionsMysqlVersion](docs/CatalogOptionsMysqlVersion.md)
 - [ChartNotFound](docs/ChartNotFound.md)
 - [ClusterApplicationList](docs/ClusterApplicationList.md)
 - [ClusterConfig](docs/ClusterConfig.md)
 - [ClusterDelete200](docs/ClusterDelete200.md)
 - [ClusterDetailsResponse](docs/ClusterDetailsResponse.md)
 - [ClusterDetailsResponseNodePools](docs/ClusterDetailsResponseNodePools.md)
 - [ClusterDetailsResponseNodePoolsPool1](docs/ClusterDetailsResponseNodePoolsPool1.md)
 - [ClusterDetailsResponseTotalSummary](docs/ClusterDetailsResponseTotalSummary.md)
 - [ClusterListResponse](docs/ClusterListResponse.md)
 - [ClusterNotFound](docs/ClusterNotFound.md)
 - [ClusterProfileAmazon](docs/ClusterProfileAmazon.md)
 - [ClusterProfileAmazonAmazon](docs/ClusterProfileAmazonAmazon.md)
 - [ClusterProfileAmazonAmazonNodePools](docs/ClusterProfileAmazonAmazonNodePools.md)
 - [ClusterProfileAmazonAmazonNodePoolsPool1](docs/ClusterProfileAmazonAmazonNodePoolsPool1.md)
 - [ClusterProfileAzure](docs/ClusterProfileAzure.md)
 - [ClusterProfileAzureAzure](docs/ClusterProfileAzureAzure.md)
 - [ClusterProfileAzureAzureNodePools](docs/ClusterProfileAzureAzureNodePools.md)
 - [ClusterProfileAzureAzureNodePoolsPool1](docs/ClusterProfileAzureAzureNodePoolsPool1.md)
 - [ClusterProfileGoogle](docs/ClusterProfileGoogle.md)
 - [ClusterProfileGoogleGoogle](docs/ClusterProfileGoogleGoogle.md)
 - [ClusterProfileGoogleGoogleNodePools](docs/ClusterProfileGoogleGoogleNodePools.md)
 - [ClusterProfileGoogleGoogleNodePoolsPool1](docs/ClusterProfileGoogleGoogleNodePoolsPool1.md)
 - [ClusterProfileNotFound](docs/ClusterProfileNotFound.md)
 - [Conflict](docs/Conflict.md)
 - [CreateAmazonObjectStoreBucketProperties](docs/CreateAmazonObjectStoreBucketProperties.md)
 - [CreateAmazonProperties](docs/CreateAmazonProperties.md)
 - [CreateAmazonPropertiesAmazon](docs/CreateAmazonPropertiesAmazon.md)
 - [CreateAmazonPropertiesAmazonMaster](docs/CreateAmazonPropertiesAmazonMaster.md)
 - [CreateApplicationRequest](docs/CreateApplicationRequest.md)
 - [CreateAzureObjectStoreBucketProperties](docs/CreateAzureObjectStoreBucketProperties.md)
 - [CreateAzureProperties](docs/CreateAzureProperties.md)
 - [CreateAzurePropertiesAzure](docs/CreateAzurePropertiesAzure.md)
 - [CreateClusterRequest](docs/CreateClusterRequest.md)
 - [CreateClusterResponse202](docs/CreateClusterResponse202.md)
 - [CreateClusterResponse400](docs/CreateClusterResponse400.md)
 - [CreateGoogleObjectStoreBucketProperties](docs/CreateGoogleObjectStoreBucketProperties.md)
 - [CreateGoogleProperties](docs/CreateGoogleProperties.md)
 - [CreateGooglePropertiesGoogle](docs/CreateGooglePropertiesGoogle.md)
 - [CreateGooglePropertiesGoogleMaster](docs/CreateGooglePropertiesGoogleMaster.md)
 - [CreateObjectStoreBucketRequest](docs/CreateObjectStoreBucketRequest.md)
 - [CreateObjectStoreBucketResponse](docs/CreateObjectStoreBucketResponse.md)
 - [CreateSecretRequest](docs/CreateSecretRequest.md)
 - [CreateSecretResponse](docs/CreateSecretResponse.md)
 - [CreateUpdateDeploymentRequest](docs/CreateUpdateDeploymentRequest.md)
 - [CreateUpdateDeploymentResponse](docs/CreateUpdateDeploymentResponse.md)
 - [DeleteDeploymentResponse](docs/DeleteDeploymentResponse.md)
 - [EndpointItem](docs/EndpointItem.md)
 - [GetClusterStatusResponse](docs/GetClusterStatusResponse.md)
 - [GetClusterStatusResponseNodePools](docs/GetClusterStatusResponseNodePools.md)
 - [GetDeploymentResponse](docs/GetDeploymentResponse.md)
 - [GoogleConfigResponse](docs/GoogleConfigResponse.md)
 - [GoogleConfigResponseInstanceType](docs/GoogleConfigResponseInstanceType.md)
 - [GoogleConfigResponseKubernetesVersions](docs/GoogleConfigResponseKubernetesVersions.md)
 - [HelmChartDetailsResponse](docs/HelmChartDetailsResponse.md)
 - [HelmChartDetailsResponseChart](docs/HelmChartDetailsResponseChart.md)
 - [HelmChartDetailsResponseChartMaintainers](docs/HelmChartDetailsResponseChartMaintainers.md)
 - [HelmChartDetailsResponseVersions](docs/HelmChartDetailsResponseVersions.md)
 - [HelmChartsListResponse](docs/HelmChartsListResponse.md)
 - [HelmChartsListResponseInner](docs/HelmChartsListResponseInner.md)
 - [HelmInitRequest](docs/HelmInitRequest.md)
 - [HelmInitResponse](docs/HelmInitResponse.md)
 - [HelmRepoListItem](docs/HelmRepoListItem.md)
 - [HelmReposAddRequest](docs/HelmReposAddRequest.md)
 - [HelmReposAddResponse](docs/HelmReposAddResponse.md)
 - [HelmReposDeleteResponse](docs/HelmReposDeleteResponse.md)
 - [HelmReposListResponse](docs/HelmReposListResponse.md)
 - [HelmReposModifyRequest](docs/HelmReposModifyRequest.md)
 - [HelmReposUpdateResponse](docs/HelmReposUpdateResponse.md)
 - [InstallSecretsRequest](docs/InstallSecretsRequest.md)
 - [InstallSecretsRequestQuery](docs/InstallSecretsRequestQuery.md)
 - [InstallSecretsResponse](docs/InstallSecretsResponse.md)
 - [InstallSecretsResponseItem](docs/InstallSecretsResponseItem.md)
 - [ListCatalogItem1](docs/ListCatalogItem1.md)
 - [ListCatalogItem2](docs/ListCatalogItem2.md)
 - [ListCatalogResponse](docs/ListCatalogResponse.md)
 - [ListDeploymentsResponse](docs/ListDeploymentsResponse.md)
 - [ListDeploymentsResponseInner](docs/ListDeploymentsResponseInner.md)
 - [ListEndpointsResponse](docs/ListEndpointsResponse.md)
 - [ListNodesResponse](docs/ListNodesResponse.md)
 - [ListNodesResponseMetadata](docs/ListNodesResponseMetadata.md)
 - [ListStorageBucketsResponse](docs/ListStorageBucketsResponse.md)
 - [ListUserResponse](docs/ListUserResponse.md)
 - [LoggingPostHook](docs/LoggingPostHook.md)
 - [LoggingPostHookInstallLogging](docs/LoggingPostHookInstallLogging.md)
 - [NodeItem](docs/NodeItem.md)
 - [NodeItemMetadata](docs/NodeItemMetadata.md)
 - [NodeItemMetadataAnnotations](docs/NodeItemMetadataAnnotations.md)
 - [NodeItemMetadataLabels](docs/NodeItemMetadataLabels.md)
 - [NodeItemSpec](docs/NodeItemSpec.md)
 - [NodeItemStatus](docs/NodeItemStatus.md)
 - [NodeItemStatusAddresses](docs/NodeItemStatusAddresses.md)
 - [NodeItemStatusAllocatable](docs/NodeItemStatusAllocatable.md)
 - [NodeItemStatusCapacity](docs/NodeItemStatusCapacity.md)
 - [NodeItemStatusConditions](docs/NodeItemStatusConditions.md)
 - [NodeItemStatusDaemonEndpoints](docs/NodeItemStatusDaemonEndpoints.md)
 - [NodeItemStatusImages](docs/NodeItemStatusImages.md)
 - [NodeItemStatusNodeInfo](docs/NodeItemStatusNodeInfo.md)
 - [NodePoolStatusAmazon](docs/NodePoolStatusAmazon.md)
 - [NodePoolStatusAzure](docs/NodePoolStatusAzure.md)
 - [NodePoolStatusGoogle](docs/NodePoolStatusGoogle.md)
 - [NodePoolsAmazon](docs/NodePoolsAmazon.md)
 - [NodePoolsAzure](docs/NodePoolsAzure.md)
 - [NodePoolsGoogle](docs/NodePoolsGoogle.md)
 - [OrganizationCreateResponse](docs/OrganizationCreateResponse.md)
 - [OrganizationListItemResponse](docs/OrganizationListItemResponse.md)
 - [OrganizationListResponse](docs/OrganizationListResponse.md)
 - [OrganizationNotFound](docs/OrganizationNotFound.md)
 - [ProfileListResponse](docs/ProfileListResponse.md)
 - [ReRunPostHook](docs/ReRunPostHook.md)
 - [RepoNotFound](docs/RepoNotFound.md)
 - [RequestedResources](docs/RequestedResources.md)
 - [ResourceItem](docs/ResourceItem.md)
 - [ResourceSummaryItem](docs/ResourceSummaryItem.md)
 - [ResourceSummaryItemIp100100180Euwest1ComputeInternal](docs/ResourceSummaryItemIp100100180Euwest1ComputeInternal.md)
 - [RunPostHook](docs/RunPostHook.md)
 - [SecretItem](docs/SecretItem.md)
 - [SecretKeyValueAmazon](docs/SecretKeyValueAmazon.md)
 - [SecretKeyValueAzure](docs/SecretKeyValueAzure.md)
 - [SecretKeyValueGeneric](docs/SecretKeyValueGeneric.md)
 - [SecretKeyValueGoogle](docs/SecretKeyValueGoogle.md)
 - [SecretKeyValueKubernetes](docs/SecretKeyValueKubernetes.md)
 - [SecretKeyValueTls](docs/SecretKeyValueTls.md)
 - [SecretsListResponse](docs/SecretsListResponse.md)
 - [SecretsNotFound](docs/SecretsNotFound.md)
 - [SupportedCloudItem](docs/SupportedCloudItem.md)
 - [SupportedCloudsResponse](docs/SupportedCloudsResponse.md)
 - [TokenCreateRequest](docs/TokenCreateRequest.md)
 - [TokenCreateResponse](docs/TokenCreateResponse.md)
 - [TokenListResponse](docs/TokenListResponse.md)
 - [TokenListResponseItem](docs/TokenListResponseItem.md)
 - [Unauthorized](docs/Unauthorized.md)
 - [UpdateAmazonProperties](docs/UpdateAmazonProperties.md)
 - [UpdateAmazonPropertiesAmazon](docs/UpdateAmazonPropertiesAmazon.md)
 - [UpdateAmazonPropertiesAmazonNodePools](docs/UpdateAmazonPropertiesAmazonNodePools.md)
 - [UpdateAzureProperties](docs/UpdateAzureProperties.md)
 - [UpdateAzurePropertiesAzure](docs/UpdateAzurePropertiesAzure.md)
 - [UpdateAzurePropertiesAzureNodePools](docs/UpdateAzurePropertiesAzureNodePools.md)
 - [UpdateAzurePropertiesAzureNodePoolsPool1](docs/UpdateAzurePropertiesAzureNodePoolsPool1.md)
 - [UpdateClusterRequest](docs/UpdateClusterRequest.md)
 - [UpdateGoogleProperties](docs/UpdateGoogleProperties.md)
 - [UpdateGooglePropertiesMaster](docs/UpdateGooglePropertiesMaster.md)
 - [UpdateGooglePropertiesNodePools](docs/UpdateGooglePropertiesNodePools.md)
 - [UpdateGooglePropertiesNodePoolsPool1](docs/UpdateGooglePropertiesNodePoolsPool1.md)
 - [UpdateNodePoolsAmazon](docs/UpdateNodePoolsAmazon.md)
 - [UrlItem](docs/UrlItem.md)
 - [User](docs/User.md)


## Documentation For Authorization

## bearerAuth
- **Type**: HTTP basic authentication

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

## Author

info@banzaicloud.com

