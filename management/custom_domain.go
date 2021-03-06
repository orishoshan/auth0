package management

type CustomDomain struct {

	// The id of the custom domain
	ID *string `json:"custom_domain_id,omitempty"`

	// The custom domain.
	Domain *string `json:"domain,omitempty"`

	// The custom domain provisioning type. Can be either "auth0_managed_certs"
	// or "self_managed_certs"
	Type *string `json:"type,omitempty"`

	// Primary is true if the domain was marked as "primary", false otherwise.
	Primary *bool `json:"primary,omitempty"`

	// The custom domain configuration status. Can be any of the following:
	//
	// "disabled", "pending", "pending_verification" or "ready"
	Status *string `json:"status,omitempty"`

	// The custom domain verification method. The only allowed value is "txt".
	VerificationMethod *string `json:"verification_method,omitempty"`

	Verification *CustomDomainVerification `json:"verification,omitempty"`
}

func (c *CustomDomain) String() string {
	return Stringify(c)
}

type CustomDomainVerification struct {

	// The custom domain verification methods.
	Methods []map[string]interface{} `json:"methods,omitempty"`
}

type CustomDomainManager struct {
	m *Management
}

func NewCustomDomainManager(m *Management) *CustomDomainManager {
	return &CustomDomainManager{m}
}

// Create a new custom domain.
//
// Note: The custom domain will need to be verified before it starts accepting
// requests.
//
// See: https://auth0.com/docs/api/management/v2#!/Custom_Domains/post_custom_domains
func (cm *CustomDomainManager) Create(c *CustomDomain) (err error) {
	return cm.m.post(cm.m.uri("custom-domains"), c)
}

// Retrieve a custom domain configuration and status.
//
// See: https://auth0.com/docs/api/management/v2#!/Custom_Domains/get_custom_domains_by_id
func (cm *CustomDomainManager) Read(id string, opts ...ReqOption) (*CustomDomain, error) {
	c := new(CustomDomain)
	err := cm.m.get(cm.m.uri("custom-domains", id)+cm.m.q(opts), c)
	return c, err
}

// Run the verification process on a custom domain.
//
// See: https://auth0.com/docs/api/management/v2#!/Custom_Domains/post_verify
func (cm *CustomDomainManager) Verify(id string) (*CustomDomain, error) {
	c := new(CustomDomain)
	err := cm.m.post(cm.m.uri("custom-domains", id, "verify"), c)
	return c, err
}

// Delete a custom domain and stop serving requests for it.
//
// See: https://auth0.com/docs/api/management/v2#!/Custom_Domains/delete_custom_domains_by_id
func (cm *CustomDomainManager) Delete(id string) (err error) {
	return cm.m.delete(cm.m.uri("custom-domains", id))
}

// Retrieve a list of custom domains.
//
// See: https://auth0.com/docs/api/management/v2#!/Custom_Domains/get_custom_domains
func (cm *CustomDomainManager) List(opts ...ReqOption) ([]*CustomDomain, error) {
	var c []*CustomDomain
	err := cm.m.get(cm.m.uri("custom-domains")+cm.m.q(opts), &c)
	return c, err
}
