// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package search

import original "github.com/Azure/azure-sdk-for-go/services/search/mgmt/2015-08-19/search"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AdminKeyKind = original.AdminKeyKind

const (
	Primary   AdminKeyKind = original.Primary
	Secondary AdminKeyKind = original.Secondary
)

type HostingMode = original.HostingMode

const (
	Default     HostingMode = original.Default
	HighDensity HostingMode = original.HighDensity
)

type IdentityType = original.IdentityType

const (
	None           IdentityType = original.None
	SystemAssigned IdentityType = original.SystemAssigned
)

type ProvisioningState = original.ProvisioningState

const (
	Failed       ProvisioningState = original.Failed
	Provisioning ProvisioningState = original.Provisioning
	Succeeded    ProvisioningState = original.Succeeded
)

type ServiceStatus = original.ServiceStatus

const (
	ServiceStatusDegraded     ServiceStatus = original.ServiceStatusDegraded
	ServiceStatusDeleting     ServiceStatus = original.ServiceStatusDeleting
	ServiceStatusDisabled     ServiceStatus = original.ServiceStatusDisabled
	ServiceStatusError        ServiceStatus = original.ServiceStatusError
	ServiceStatusProvisioning ServiceStatus = original.ServiceStatusProvisioning
	ServiceStatusRunning      ServiceStatus = original.ServiceStatusRunning
)

type SkuName = original.SkuName

const (
	Basic              SkuName = original.Basic
	Free               SkuName = original.Free
	Standard           SkuName = original.Standard
	Standard2          SkuName = original.Standard2
	Standard3          SkuName = original.Standard3
	StorageOptimizedL1 SkuName = original.StorageOptimizedL1
	StorageOptimizedL2 SkuName = original.StorageOptimizedL2
)

type UnavailableNameReason = original.UnavailableNameReason

const (
	AlreadyExists UnavailableNameReason = original.AlreadyExists
	Invalid       UnavailableNameReason = original.Invalid
)

type AdminKeyResult = original.AdminKeyResult
type AdminKeysClient = original.AdminKeysClient
type BaseClient = original.BaseClient
type CheckNameAvailabilityInput = original.CheckNameAvailabilityInput
type CheckNameAvailabilityOutput = original.CheckNameAvailabilityOutput
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type Identity = original.Identity
type ListQueryKeysResult = original.ListQueryKeysResult
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type QueryKey = original.QueryKey
type QueryKeysClient = original.QueryKeysClient
type Resource = original.Resource
type Service = original.Service
type ServiceListResult = original.ServiceListResult
type ServiceProperties = original.ServiceProperties
type ServicesClient = original.ServicesClient
type ServicesCreateOrUpdateFuture = original.ServicesCreateOrUpdateFuture
type Sku = original.Sku

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAdminKeysClient(subscriptionID string) AdminKeysClient {
	return original.NewAdminKeysClient(subscriptionID)
}
func NewAdminKeysClientWithBaseURI(baseURI string, subscriptionID string) AdminKeysClient {
	return original.NewAdminKeysClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewQueryKeysClient(subscriptionID string) QueryKeysClient {
	return original.NewQueryKeysClient(subscriptionID)
}
func NewQueryKeysClientWithBaseURI(baseURI string, subscriptionID string) QueryKeysClient {
	return original.NewQueryKeysClientWithBaseURI(baseURI, subscriptionID)
}
func NewServicesClient(subscriptionID string) ServicesClient {
	return original.NewServicesClient(subscriptionID)
}
func NewServicesClientWithBaseURI(baseURI string, subscriptionID string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAdminKeyKindValues() []AdminKeyKind {
	return original.PossibleAdminKeyKindValues()
}
func PossibleHostingModeValues() []HostingMode {
	return original.PossibleHostingModeValues()
}
func PossibleIdentityTypeValues() []IdentityType {
	return original.PossibleIdentityTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleServiceStatusValues() []ServiceStatus {
	return original.PossibleServiceStatusValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleUnavailableNameReasonValues() []UnavailableNameReason {
	return original.PossibleUnavailableNameReasonValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
