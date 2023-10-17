/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package akscluster

import "time"

const (
	ResourceName                               = "tanzu-mission-control_akscluster"
	RetryInterval                              = retryInterval("retry-interval")
	defaultTimeout                             = 30 * time.Minute
	defaultInterval                            = 10 * time.Second
	CredentialNameKey                          = "credential_name" //nolint:gosec
	SubscriptionIDKey                          = "subscription_id"
	ResourceGroupNameKey                       = "resource_group"
	NameKey                                    = "name"
	clusterSpecKey                             = "spec"
	nodepoolSpecKey                            = "spec"
	waitKey                                    = "ready_wait_timeout"
	waitForHealthyKey                          = "wait_for_healthy"
	clusterGroupKey                            = "cluster_group"
	clusterGroupDefaultValue                   = "default"
	proxyNameKey                               = "proxy"
	configKey                                  = "config"
	agentNameKey                               = "agent_name"
	resourceIDKey                              = "resource_id"
	nodepoolKey                                = "nodepool"
	kubernetesVersionKey                       = "kubernetes_version"
	locationKey                                = "location"
	nodeResourceGroupNameKey                   = "node_resource_group_name"
	diskEncryptionSetKey                       = "disk_encryption_set"
	tagsKey                                    = "tags"
	skuKey                                     = "sku"
	accessConfigKey                            = "access_config"
	enableRbacKey                              = "enable_rbac"
	disableLocalAccountsKey                    = "disable_local_accounts"
	aadConfigKey                               = "aad_config"
	managedKey                                 = "managed"
	tenantIDKey                                = "tenant_id"
	adminGroupIDsKey                           = "admin_group_ids"
	enableAzureRbacKey                         = "enable_azure_rbac"
	apiServerAccessConfigKey                   = "api_server_access_config"
	authorizedIPRangesKey                      = "authorized_ip_ranges"
	enablePrivateClusterKey                    = "enable_private_cluster"
	linuxConfigKey                             = "linux_config"
	adminUserNameKey                           = "admin_username"
	sshkeysKey                                 = "ssh_keys"
	networkConfigKey                           = "network_config"
	loadBalancerSkuKey                         = "load_balancer_sku"
	networkPluginKey                           = "network_plugin"
	networkPolicyKey                           = "network_policy"
	dnsServiceIPKey                            = "dns_service_ip"
	dockerBridgeCidrKey                        = "docker_bridge_cidr"
	podCidrKey                                 = "pod_cidr"
	serviceCidrKey                             = "service_cidr"
	dnsPrefixKey                               = "dns_prefix"
	enableHTTPApplicationRoutingKey            = "enable_http_application_routing"
	storageConfigKey                           = "storage_config"
	enableDiskCsiDriverKey                     = "enable_disk_csi_driver"
	enableFileCsiDriverKey                     = "enable_file_csi_driver"
	enableSnapshotControllerKey                = "enable_snapshot_controller"
	addonsConfigKey                            = "addon_config"
	azureKeyvaultSecretsProviderAddonConfigKey = "azure_keyvault_secrets_provider_addon_config" //nolint:gosec
	enableKey                                  = "enable"
	keyVaultSecretsProviderConfigKey           = "keyvault_secrets_provider_config"
	enableSecretsRotationKey                   = "enable_secret_rotation"
	rotationPollIntervalKey                    = "rotation_poll_interval"
	monitorAddonConfigKey                      = "monitor_addon_config"
	logAnalyticsWorkspaceIDKey                 = "log_analytics_workspace_id"
	azurePolicyAddonConfigKey                  = "azure_policy_addon_config"
	autoUpgradeConfigKey                       = "auto_upgrade_config"
	upgradeChannelKey                          = "upgrade_channel"
	skuNameKey                                 = "name"
	skuTierKey                                 = "tier"
	infoKey                                    = "info"
	modeKey                                    = "mode"
	nodeImageVersionKey                        = "node_image_version"
	typeKey                                    = "type"
	availabilityZonesKey                       = "availability_zones"
	countKey                                   = "count"
	vmSizeKey                                  = "vm_size"
	osTypeKey                                  = "os_type"
	osDiskTypeKey                              = "os_disk_type"
	osDiskSizeKey                              = "os_disk_size_gb"
	maxPodsKey                                 = "max_pods"
	enableNodePublicIPKey                      = "enable_node_public_ip"
	taintsKey                                  = "taints"
	effectKey                                  = "effect"
	keyKey                                     = "key"
	valueKey                                   = "value"
	vnetSubnetKey                              = "vnet_subnet_id"
	nodeLabelsKey                              = "node_labels"
	autoscalingConfigKey                       = "auto_scaling_config"
	minCountKey                                = "min_count"
	maxCountKey                                = "max_count"
	scaleSetPriorityKey                        = "scale_set_priority"
	scaleSetEvictionPolicyKey                  = "scale_set_eviction_policy"
	maxSpotPriceKey                            = "spot_max_price"
	upgradeConfigKey                           = "upgrade_config"
	maxSurgeKey                                = "max_surge"
	kubeconfigKey                              = "kubeconfig"
)
