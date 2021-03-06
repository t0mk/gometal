/*
 * Metal API
 *
 * This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>. 
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"time"
)

// DeviceCreateInput struct for DeviceCreateInput
type DeviceCreateInput struct {
	// The datacenter where the device should be created.  The API will accept either a single facility `{ \"facility\": \"f1\" }`, or it can be instructed to create the device in the best available datacenter `{ \"facility\": \"any\" }`.  Additionally it is possible to set a prioritized location selection. For example `{ \"facility\": [\"f3\", \"f2\", \"any\"] }` can be used to prioritize `f3` and then `f2` before accepting `any` facility. If none of the facilities provided have availability for the requested device the request will fail.
	Facility string `json:"facility"`
	// The slug of the device plan to provision.
	Plan string `json:"plan"`
	// The hostname to use within the operating system. The same hostname may be used on multiple devices within a project.
	Hostname *string `json:"hostname,omitempty"`
	// Any description of the device or how it will be used. This may be used to inform other API consumers with project access.
	Description *string `json:"description,omitempty"`
	// The billing cycle of the device.
	BillingCycle *string `json:"billing_cycle,omitempty"`
	// The slug of the operating system to provision. Check the Equinix Metal operating system documentation for rules that may be imposed per operating system, including restrictions on IP address options and device plans.
	OperatingSystem string `json:"operating_system"`
	// When true, devices with a `custom_ipxe` OS will always boot to iPXE. The default setting of false ensures that iPXE will be used on only the first boot.
	AlwaysPxe *bool `json:"always_pxe,omitempty"`
	// When set, the device will chainload an iPXE Script at boot fetched from the supplied URL.        See [Custom iPXE](https://metal.equinix.com/developers/docs/operating-systems/custom-ipxe/) for more details.
	IpxeScriptUrl *string `json:"ipxe_script_url,omitempty"`
	// The userdata presented in the metadata service for this device.  Userdata is fetched and interpretted by the operating system installed on the device. Acceptable formats are determined by the operating system, with the exception of a special iPXE enabling syntax which is handled before the operating system starts.        See [Server User Data](https://metal.equinix.com/developers/docs/servers/user-data/) and [Provisioning with Custom iPXE](https://metal.equinix.com/developers/docs/operating-systems/custom-ipxe/#provisioning-with-custom-ipxe) for more details.
	Userdata *string `json:"userdata,omitempty"`
	// Whether the device should be locked, preventing accidental deletion.
	Locked *bool `json:"locked,omitempty"`
	// Customdata is an arbitrary JSON value that can be accessed via the metadata service.
	Customdata *map[string]interface{} `json:"customdata,omitempty"`
	// Metro code or ID of where the instance should be provisioned in.
	Metro *string `json:"metro,omitempty"`
	// The Hardware Reservation UUID to provision. Alternatively, `next-available` can be specified to select from any of the available hardware reservations. An error will be returned if the requested reservation option is not available.        See [Reserved Hardware](https://metal.equinix.com/developers/docs/deploy/reserved/) for more details.
	HardwareReservationId *string `json:"hardware_reservation_id,omitempty"`
	SpotInstance *bool `json:"spot_instance,omitempty"`
	SpotPriceMax *float32 `json:"spot_price_max,omitempty"`
	TerminationTime *time.Time `json:"termination_time,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	// A list of UUIDs identifying the device parent project that should be authorized to access this device (typically via /root/.ssh/authorized_keys). These keys will also appear in the device metadata.  If no SSH keys are specified (`user_ssh_keys`, `project_ssh_keys`, and `ssh_keys` are all empty lists or omitted), all parent project keys, parent project members keys and organization members keys will be included. This behaviour can be changed with 'no_ssh_keys' option to omit any SSH key being added.  
	ProjectSshKeys []string `json:"project_ssh_keys,omitempty"`
	// A list of UUIDs identifying the users that should be authorized to access this device (typically via /root/.ssh/authorized_keys).  These keys will also appear in the device metadata.  The users must be members of the project or organization.  If no SSH keys are specified (`user_ssh_keys`, `project_ssh_keys`, and `ssh_keys` are all empty lists or omitted), all parent project keys, parent project members keys and organization members keys will be included. This behaviour can be changed with 'no_ssh_keys' option to omit any SSH key being added. 
	UserSshKeys []string `json:"user_ssh_keys,omitempty"`
	// A list of new or existing project ssh_keys that should be authorized to access this device (typically via /root/.ssh/authorized_keys). These keys will also appear in the device metadata.  These keys are added in addition to any keys defined by   `project_ssh_keys` and `user_ssh_keys`. 
	SshKeys []SSHKeyInput `json:"ssh_keys,omitempty"`
	// Overrides default behaviour of attaching all of the organization members ssh keys and project ssh keys to device if no specific keys specified
	NoSshKeys NullableBool `json:"no_ssh_keys,omitempty"`
	// The features attribute allows you to optionally specify what features your server should have.  In the API shorthand syntax, all features listed are `required`:  ``` { \"features\": [\"tpm\"] } ```  Alternatively, if you do not require a certain feature, but would prefer to be assigned a server with that feature if there are any available, you may specify that feature with a `preferred` value. The request will not fail if we have no servers with that feature in our inventory. The API offers an alternative syntax for mixing preferred and required features:  ``` { \"features\": { \"tpm\": \"required\", \"raid\": \"preferred\" } } ```  The request will only fail if there are no available servers matching the required `tpm` criteria.
	Features *[]string `json:"features,omitempty"`
	// Deprecated. Use ip_addresses. Subnet range for addresses allocated to this device. Your project must have addresses available for a non-default request.
	PublicIpv4SubnetSize *float32 `json:"public_ipv4_subnet_size,omitempty"`
	// Deprecated. Use ip_addresses. Subnet range for addresses allocated to this device.
	PrivateIpv4SubnetSize *float32 `json:"private_ipv4_subnet_size,omitempty"`
	// The `ip_addresses attribute will allow you to specify the addresses you want created with your device.  The default value configures public IPv4, public IPv6, and private IPv4.  Private IPv4 address is required. When specifying `ip_addresses`, one of the array items must enable private IPv4.  Some operating systems require public IPv4 address. In those cases you will receive an error message if public IPv4 is not enabled.  For example, to only configure your server with a private IPv4 address, you can send `{ \"ip_addresses\": [{ \"address_family\": 4, \"public\": false }] }`.  It is possible to request a subnet size larger than a `/30` by assigning addresses using the UUID(s) of ip_reservations in your project.  For example, `{ \"ip_addresses\": [..., {\"address_family\": 4, \"public\": true, \"ip_reservations\": [\"uuid1\", \"uuid2\"]}] }`  To access a server without public IPs, you can use our Out-of-Band console access (SOS) or proxy through another server in the project with public IPs enabled.
	IpAddresses *[]DeviceCreateInputIpAddresses `json:"ip_addresses,omitempty"`
}

// NewDeviceCreateInput instantiates a new DeviceCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceCreateInput(facility string, plan string, operatingSystem string) *DeviceCreateInput {
	this := DeviceCreateInput{}
	this.Facility = facility
	this.Plan = plan
	this.OperatingSystem = operatingSystem
	var alwaysPxe bool = false
	this.AlwaysPxe = &alwaysPxe
	var locked bool = false
	this.Locked = &locked
	var hardwareReservationId string = ""
	this.HardwareReservationId = &hardwareReservationId
	var noSshKeys bool = false
	this.NoSshKeys = *NewNullableBool(&noSshKeys)
	return &this
}

// NewDeviceCreateInputWithDefaults instantiates a new DeviceCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceCreateInputWithDefaults() *DeviceCreateInput {
	this := DeviceCreateInput{}
	var alwaysPxe bool = false
	this.AlwaysPxe = &alwaysPxe
	var locked bool = false
	this.Locked = &locked
	var hardwareReservationId string = ""
	this.HardwareReservationId = &hardwareReservationId
	var noSshKeys bool = false
	this.NoSshKeys = *NewNullableBool(&noSshKeys)
	return &this
}

// GetFacility returns the Facility field value
func (o *DeviceCreateInput) GetFacility() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetFacilityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Facility, true
}

// SetFacility sets field value
func (o *DeviceCreateInput) SetFacility(v string) {
	o.Facility = v
}

// GetPlan returns the Plan field value
func (o *DeviceCreateInput) GetPlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetPlanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *DeviceCreateInput) SetPlan(v string) {
	o.Plan = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *DeviceCreateInput) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *DeviceCreateInput) SetHostname(v string) {
	o.Hostname = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeviceCreateInput) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeviceCreateInput) SetDescription(v string) {
	o.Description = &v
}

// GetBillingCycle returns the BillingCycle field value if set, zero value otherwise.
func (o *DeviceCreateInput) GetBillingCycle() string {
	if o == nil || o.BillingCycle == nil {
		var ret string
		return ret
	}
	return *o.BillingCycle
}

// GetBillingCycleOk returns a tuple with the BillingCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetBillingCycleOk() (*string, bool) {
	if o == nil || o.BillingCycle == nil {
		return nil, false
	}
	return o.BillingCycle, true
}

// HasBillingCycle returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasBillingCycle() bool {
	if o != nil && o.BillingCycle != nil {
		return true
	}

	return false
}

// SetBillingCycle gets a reference to the given string and assigns it to the BillingCycle field.
func (o *DeviceCreateInput) SetBillingCycle(v string) {
	o.BillingCycle = &v
}

// GetOperatingSystem returns the OperatingSystem field value
func (o *DeviceCreateInput) GetOperatingSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetOperatingSystemOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OperatingSystem, true
}

// SetOperatingSystem sets field value
func (o *DeviceCreateInput) SetOperatingSystem(v string) {
	o.OperatingSystem = v
}

// GetAlwaysPxe returns the AlwaysPxe field value if set, zero value otherwise.
func (o *DeviceCreateInput) GetAlwaysPxe() bool {
	if o == nil || o.AlwaysPxe == nil {
		var ret bool
		return ret
	}
	return *o.AlwaysPxe
}

// GetAlwaysPxeOk returns a tuple with the AlwaysPxe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetAlwaysPxeOk() (*bool, bool) {
	if o == nil || o.AlwaysPxe == nil {
		return nil, false
	}
	return o.AlwaysPxe, true
}

// HasAlwaysPxe returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasAlwaysPxe() bool {
	if o != nil && o.AlwaysPxe != nil {
		return true
	}

	return false
}

// SetAlwaysPxe gets a reference to the given bool and assigns it to the AlwaysPxe field.
func (o *DeviceCreateInput) SetAlwaysPxe(v bool) {
	o.AlwaysPxe = &v
}

// GetIpxeScriptUrl returns the IpxeScriptUrl field value if set, zero value otherwise.
func (o *DeviceCreateInput) GetIpxeScriptUrl() string {
	if o == nil || o.IpxeScriptUrl == nil {
		var ret string
		return ret
	}
	return *o.IpxeScriptUrl
}

// GetIpxeScriptUrlOk returns a tuple with the IpxeScriptUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetIpxeScriptUrlOk() (*string, bool) {
	if o == nil || o.IpxeScriptUrl == nil {
		return nil, false
	}
	return o.IpxeScriptUrl, true
}

// HasIpxeScriptUrl returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasIpxeScriptUrl() bool {
	if o != nil && o.IpxeScriptUrl != nil {
		return true
	}

	return false
}

// SetIpxeScriptUrl gets a reference to the given string and assigns it to the IpxeScriptUrl field.
func (o *DeviceCreateInput) SetIpxeScriptUrl(v string) {
	o.IpxeScriptUrl = &v
}

// GetUserdata returns the Userdata field value if set, zero value otherwise.
func (o *DeviceCreateInput) GetUserdata() string {
	if o == nil || o.Userdata == nil {
		var ret string
		return ret
	}
	return *o.Userdata
}

// GetUserdataOk returns a tuple with the Userdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetUserdataOk() (*string, bool) {
	if o == nil || o.Userdata == nil {
		return nil, false
	}
	return o.Userdata, true
}

// HasUserdata returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasUserdata() bool {
	if o != nil && o.Userdata != nil {
		return true
	}

	return false
}

// SetUserdata gets a reference to the given string and assigns it to the Userdata field.
func (o *DeviceCreateInput) SetUserdata(v string) {
	o.Userdata = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *DeviceCreateInput) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *DeviceCreateInput) SetLocked(v bool) {
	o.Locked = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *DeviceCreateInput) GetCustomdata() map[string]interface{} {
	if o == nil || o.Customdata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetCustomdataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Customdata == nil {
		return nil, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasCustomdata() bool {
	if o != nil && o.Customdata != nil {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *DeviceCreateInput) SetCustomdata(v map[string]interface{}) {
	o.Customdata = &v
}

// GetMetro returns the Metro field value if set, zero value otherwise.
func (o *DeviceCreateInput) GetMetro() string {
	if o == nil || o.Metro == nil {
		var ret string
		return ret
	}
	return *o.Metro
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetMetroOk() (*string, bool) {
	if o == nil || o.Metro == nil {
		return nil, false
	}
	return o.Metro, true
}

// HasMetro returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasMetro() bool {
	if o != nil && o.Metro != nil {
		return true
	}

	return false
}

// SetMetro gets a reference to the given string and assigns it to the Metro field.
func (o *DeviceCreateInput) SetMetro(v string) {
	o.Metro = &v
}

// GetHardwareReservationId returns the HardwareReservationId field value if set, zero value otherwise.
func (o *DeviceCreateInput) GetHardwareReservationId() string {
	if o == nil || o.HardwareReservationId == nil {
		var ret string
		return ret
	}
	return *o.HardwareReservationId
}

// GetHardwareReservationIdOk returns a tuple with the HardwareReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetHardwareReservationIdOk() (*string, bool) {
	if o == nil || o.HardwareReservationId == nil {
		return nil, false
	}
	return o.HardwareReservationId, true
}

// HasHardwareReservationId returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasHardwareReservationId() bool {
	if o != nil && o.HardwareReservationId != nil {
		return true
	}

	return false
}

// SetHardwareReservationId gets a reference to the given string and assigns it to the HardwareReservationId field.
func (o *DeviceCreateInput) SetHardwareReservationId(v string) {
	o.HardwareReservationId = &v
}

// GetSpotInstance returns the SpotInstance field value if set, zero value otherwise.
func (o *DeviceCreateInput) GetSpotInstance() bool {
	if o == nil || o.SpotInstance == nil {
		var ret bool
		return ret
	}
	return *o.SpotInstance
}

// GetSpotInstanceOk returns a tuple with the SpotInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetSpotInstanceOk() (*bool, bool) {
	if o == nil || o.SpotInstance == nil {
		return nil, false
	}
	return o.SpotInstance, true
}

// HasSpotInstance returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasSpotInstance() bool {
	if o != nil && o.SpotInstance != nil {
		return true
	}

	return false
}

// SetSpotInstance gets a reference to the given bool and assigns it to the SpotInstance field.
func (o *DeviceCreateInput) SetSpotInstance(v bool) {
	o.SpotInstance = &v
}

// GetSpotPriceMax returns the SpotPriceMax field value if set, zero value otherwise.
func (o *DeviceCreateInput) GetSpotPriceMax() float32 {
	if o == nil || o.SpotPriceMax == nil {
		var ret float32
		return ret
	}
	return *o.SpotPriceMax
}

// GetSpotPriceMaxOk returns a tuple with the SpotPriceMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetSpotPriceMaxOk() (*float32, bool) {
	if o == nil || o.SpotPriceMax == nil {
		return nil, false
	}
	return o.SpotPriceMax, true
}

// HasSpotPriceMax returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasSpotPriceMax() bool {
	if o != nil && o.SpotPriceMax != nil {
		return true
	}

	return false
}

// SetSpotPriceMax gets a reference to the given float32 and assigns it to the SpotPriceMax field.
func (o *DeviceCreateInput) SetSpotPriceMax(v float32) {
	o.SpotPriceMax = &v
}

// GetTerminationTime returns the TerminationTime field value if set, zero value otherwise.
func (o *DeviceCreateInput) GetTerminationTime() time.Time {
	if o == nil || o.TerminationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.TerminationTime
}

// GetTerminationTimeOk returns a tuple with the TerminationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetTerminationTimeOk() (*time.Time, bool) {
	if o == nil || o.TerminationTime == nil {
		return nil, false
	}
	return o.TerminationTime, true
}

// HasTerminationTime returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasTerminationTime() bool {
	if o != nil && o.TerminationTime != nil {
		return true
	}

	return false
}

// SetTerminationTime gets a reference to the given time.Time and assigns it to the TerminationTime field.
func (o *DeviceCreateInput) SetTerminationTime(v time.Time) {
	o.TerminationTime = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DeviceCreateInput) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DeviceCreateInput) SetTags(v []string) {
	o.Tags = &v
}

// GetProjectSshKeys returns the ProjectSshKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceCreateInput) GetProjectSshKeys() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.ProjectSshKeys
}

// GetProjectSshKeysOk returns a tuple with the ProjectSshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceCreateInput) GetProjectSshKeysOk() (*[]string, bool) {
	if o == nil || o.ProjectSshKeys == nil {
		return nil, false
	}
	return &o.ProjectSshKeys, true
}

// HasProjectSshKeys returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasProjectSshKeys() bool {
	if o != nil && o.ProjectSshKeys != nil {
		return true
	}

	return false
}

// SetProjectSshKeys gets a reference to the given []string and assigns it to the ProjectSshKeys field.
func (o *DeviceCreateInput) SetProjectSshKeys(v []string) {
	o.ProjectSshKeys = v
}

// GetUserSshKeys returns the UserSshKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceCreateInput) GetUserSshKeys() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.UserSshKeys
}

// GetUserSshKeysOk returns a tuple with the UserSshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceCreateInput) GetUserSshKeysOk() (*[]string, bool) {
	if o == nil || o.UserSshKeys == nil {
		return nil, false
	}
	return &o.UserSshKeys, true
}

// HasUserSshKeys returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasUserSshKeys() bool {
	if o != nil && o.UserSshKeys != nil {
		return true
	}

	return false
}

// SetUserSshKeys gets a reference to the given []string and assigns it to the UserSshKeys field.
func (o *DeviceCreateInput) SetUserSshKeys(v []string) {
	o.UserSshKeys = v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceCreateInput) GetSshKeys() []SSHKeyInput {
	if o == nil  {
		var ret []SSHKeyInput
		return ret
	}
	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceCreateInput) GetSshKeysOk() (*[]SSHKeyInput, bool) {
	if o == nil || o.SshKeys == nil {
		return nil, false
	}
	return &o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasSshKeys() bool {
	if o != nil && o.SshKeys != nil {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []SSHKeyInput and assigns it to the SshKeys field.
func (o *DeviceCreateInput) SetSshKeys(v []SSHKeyInput) {
	o.SshKeys = v
}

// GetNoSshKeys returns the NoSshKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceCreateInput) GetNoSshKeys() bool {
	if o == nil || o.NoSshKeys.Get() == nil {
		var ret bool
		return ret
	}
	return *o.NoSshKeys.Get()
}

// GetNoSshKeysOk returns a tuple with the NoSshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceCreateInput) GetNoSshKeysOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NoSshKeys.Get(), o.NoSshKeys.IsSet()
}

// HasNoSshKeys returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasNoSshKeys() bool {
	if o != nil && o.NoSshKeys.IsSet() {
		return true
	}

	return false
}

// SetNoSshKeys gets a reference to the given NullableBool and assigns it to the NoSshKeys field.
func (o *DeviceCreateInput) SetNoSshKeys(v bool) {
	o.NoSshKeys.Set(&v)
}
// SetNoSshKeysNil sets the value for NoSshKeys to be an explicit nil
func (o *DeviceCreateInput) SetNoSshKeysNil() {
	o.NoSshKeys.Set(nil)
}

// UnsetNoSshKeys ensures that no value is present for NoSshKeys, not even an explicit nil
func (o *DeviceCreateInput) UnsetNoSshKeys() {
	o.NoSshKeys.Unset()
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *DeviceCreateInput) GetFeatures() []string {
	if o == nil || o.Features == nil {
		var ret []string
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetFeaturesOk() (*[]string, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *DeviceCreateInput) SetFeatures(v []string) {
	o.Features = &v
}

// GetPublicIpv4SubnetSize returns the PublicIpv4SubnetSize field value if set, zero value otherwise.
func (o *DeviceCreateInput) GetPublicIpv4SubnetSize() float32 {
	if o == nil || o.PublicIpv4SubnetSize == nil {
		var ret float32
		return ret
	}
	return *o.PublicIpv4SubnetSize
}

// GetPublicIpv4SubnetSizeOk returns a tuple with the PublicIpv4SubnetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetPublicIpv4SubnetSizeOk() (*float32, bool) {
	if o == nil || o.PublicIpv4SubnetSize == nil {
		return nil, false
	}
	return o.PublicIpv4SubnetSize, true
}

// HasPublicIpv4SubnetSize returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasPublicIpv4SubnetSize() bool {
	if o != nil && o.PublicIpv4SubnetSize != nil {
		return true
	}

	return false
}

// SetPublicIpv4SubnetSize gets a reference to the given float32 and assigns it to the PublicIpv4SubnetSize field.
func (o *DeviceCreateInput) SetPublicIpv4SubnetSize(v float32) {
	o.PublicIpv4SubnetSize = &v
}

// GetPrivateIpv4SubnetSize returns the PrivateIpv4SubnetSize field value if set, zero value otherwise.
func (o *DeviceCreateInput) GetPrivateIpv4SubnetSize() float32 {
	if o == nil || o.PrivateIpv4SubnetSize == nil {
		var ret float32
		return ret
	}
	return *o.PrivateIpv4SubnetSize
}

// GetPrivateIpv4SubnetSizeOk returns a tuple with the PrivateIpv4SubnetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetPrivateIpv4SubnetSizeOk() (*float32, bool) {
	if o == nil || o.PrivateIpv4SubnetSize == nil {
		return nil, false
	}
	return o.PrivateIpv4SubnetSize, true
}

// HasPrivateIpv4SubnetSize returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasPrivateIpv4SubnetSize() bool {
	if o != nil && o.PrivateIpv4SubnetSize != nil {
		return true
	}

	return false
}

// SetPrivateIpv4SubnetSize gets a reference to the given float32 and assigns it to the PrivateIpv4SubnetSize field.
func (o *DeviceCreateInput) SetPrivateIpv4SubnetSize(v float32) {
	o.PrivateIpv4SubnetSize = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *DeviceCreateInput) GetIpAddresses() []DeviceCreateInputIpAddresses {
	if o == nil || o.IpAddresses == nil {
		var ret []DeviceCreateInputIpAddresses
		return ret
	}
	return *o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInput) GetIpAddressesOk() (*[]DeviceCreateInputIpAddresses, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *DeviceCreateInput) HasIpAddresses() bool {
	if o != nil && o.IpAddresses != nil {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []DeviceCreateInputIpAddresses and assigns it to the IpAddresses field.
func (o *DeviceCreateInput) SetIpAddresses(v []DeviceCreateInputIpAddresses) {
	o.IpAddresses = &v
}

func (o DeviceCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["facility"] = o.Facility
	}
	if true {
		toSerialize["plan"] = o.Plan
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.BillingCycle != nil {
		toSerialize["billing_cycle"] = o.BillingCycle
	}
	if true {
		toSerialize["operating_system"] = o.OperatingSystem
	}
	if o.AlwaysPxe != nil {
		toSerialize["always_pxe"] = o.AlwaysPxe
	}
	if o.IpxeScriptUrl != nil {
		toSerialize["ipxe_script_url"] = o.IpxeScriptUrl
	}
	if o.Userdata != nil {
		toSerialize["userdata"] = o.Userdata
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.Customdata != nil {
		toSerialize["customdata"] = o.Customdata
	}
	if o.Metro != nil {
		toSerialize["metro"] = o.Metro
	}
	if o.HardwareReservationId != nil {
		toSerialize["hardware_reservation_id"] = o.HardwareReservationId
	}
	if o.SpotInstance != nil {
		toSerialize["spot_instance"] = o.SpotInstance
	}
	if o.SpotPriceMax != nil {
		toSerialize["spot_price_max"] = o.SpotPriceMax
	}
	if o.TerminationTime != nil {
		toSerialize["termination_time"] = o.TerminationTime
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.ProjectSshKeys != nil {
		toSerialize["project_ssh_keys"] = o.ProjectSshKeys
	}
	if o.UserSshKeys != nil {
		toSerialize["user_ssh_keys"] = o.UserSshKeys
	}
	if o.SshKeys != nil {
		toSerialize["ssh_keys"] = o.SshKeys
	}
	if o.NoSshKeys.IsSet() {
		toSerialize["no_ssh_keys"] = o.NoSshKeys.Get()
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.PublicIpv4SubnetSize != nil {
		toSerialize["public_ipv4_subnet_size"] = o.PublicIpv4SubnetSize
	}
	if o.PrivateIpv4SubnetSize != nil {
		toSerialize["private_ipv4_subnet_size"] = o.PrivateIpv4SubnetSize
	}
	if o.IpAddresses != nil {
		toSerialize["ip_addresses"] = o.IpAddresses
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceCreateInput struct {
	value *DeviceCreateInput
	isSet bool
}

func (v NullableDeviceCreateInput) Get() *DeviceCreateInput {
	return v.value
}

func (v *NullableDeviceCreateInput) Set(val *DeviceCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceCreateInput(val *DeviceCreateInput) *NullableDeviceCreateInput {
	return &NullableDeviceCreateInput{value: val, isSet: true}
}

func (v NullableDeviceCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


