package broker

// Service represents a service (of which there may be many variants-- "plans")
// offered by a service broker
//
// https://github.com/openservicebrokerapi/servicebroker/blob/v2.12/spec.md#service-objects
type Service struct {
	Name            string        `json:"name"`
	ID              string        `json:"id"`
	Description     string        `json:"description"`
	Tags            []string      `json:"tags,omitempty"`
	Requires        []string      `json:"requires,omitempty"`
	Bindable        bool          `json:"bindable"`
	Metadata        interface{}   `json:"metadata,omitempty"`
	DashboardClient interface{}   `json:"dashboard_client"`
	PlanUpdateable  bool          `json:"plan_updateable,omitempty"`
	Plans           []ServicePlan `json:"plans"`
}

// ServicePlan is the Open Service API compatible struct for service plans.
// It comes with with JSON struct tags to match the API spec
type ServicePlan struct {
	Name        string      `json:"name"`
	ID          string      `json:"id"`
	Description string      `json:"description"`
	Metadata    interface{} `json:"metadata,omitempty"`
	Free        bool        `json:"free,omitempty"`
	Bindable    *bool       `json:"bindable,omitempty"`
	Schemas     *Schemas    `json:"schemas,omitempty"`
}

// ServiceInstance represents an instance of a service
type ServiceInstance struct {
	ID               string `json:"id"`
	DashboardURL     string `json:"dashboard_url"`
	InternalID       string `json:"internal_id,omitempty"`
	ServiceID        string `json:"service_id"`
	PlanID           string `json:"plan_id"`
	OrganizationGUID string `json:"organization_guid"`
	SpaceGUID        string `json:"space_guid"`

	LastOperation *LastOperationResponse `json:"last_operation,omitempty"`

	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// CreateServiceInstanceRequest represents a request to a broker to provision an
// instance of a service
type CreateServiceInstanceRequest struct {
	OrgID             string                 `json:"organization_guid,omitempty"`
	PlanID            string                 `json:"plan_id,omitempty"`
	ServiceID         string                 `json:"service_id,omitempty"`
	SpaceID           string                 `json:"space_guid,omitempty"`
	Parameters        map[string]interface{} `json:"parameters,omitempty"`
	AcceptsIncomplete bool                   `json:"accepts_incomplete,omitempty"`
	ContextProfile    ContextProfile         `json:"context,omitempty"`
}

// ContextProfilePlatformKubernetes is a constant to send when the
// client is representing a kubernetes style ecosystem.
const ContextProfilePlatformKubernetes string = "kubernetes"

// ContextProfile implements the optional OSB field
// https://github.com/duglin/servicebroker/blob/CFisms/context-profiles.md#kubernetes
type ContextProfile struct {
	// Platform is always `kubernetes`
	Platform string `json:"platform,omitempty"`
	// Namespace is the Kubernetes namespace in which the service instance will be visible.
	Namespace string `json:"namespace,omitempty"`
}

// CreateServiceInstanceResponse represents the response from a broker after a
// request to provision an instance of a service
type CreateServiceInstanceResponse struct {
	DashboardURL string `json:"dashboard_url,omitempty"`
	Operation    string `json:"operation,omitempty"`
	Code         int    `json:"-"`
}

// DeleteServiceInstanceRequest represents a request to a broker to deprovision an
// instance of a service
type DeleteServiceInstanceRequest struct {
	ServiceID         string `json:"service_id"`
	PlanID            string `json:"plan_id"`
	AcceptsIncomplete bool   `json:"accepts_incomplete,omitempty"`
}

// DeleteServiceInstanceResponse represents the response from a broker after a request
// to deprovision an instance of a service
type DeleteServiceInstanceResponse struct {
	Operation string `json:"operation,omitempty"`
}

// LastOperationRequest represents a request to a broker to give the state of the action
// it is completing asynchronously
type LastOperationRequest struct {
	ServiceID string `json:"service_id,omitempty"`
	PlanID    string `json:"plan_id,omitempty"`
	Operation string `json:"operation,omitempty"`
}

// LastOperationResponse represents the broker response with the state of a discrete action
// that the broker is completing asynchronously
type LastOperationResponse struct {
	State       string `json:"state"`
	Description string `json:"description,omitempty"`
}

type BrokerResponseError struct {
	Code        int    `json:"-"`
	Description string `json:"description,omitempty"`
}

// Defines the possible states of an asynchronous request to a broker
const (
	StateInProgress = "in progress"
	StateSucceeded  = "succeeded"
	StateFailed     = "failed"
)

// ServiceBinding represents a binding to a service instance
type ServiceBinding struct {
	ID                string                 `json:"id"`
	ServiceID         string                 `json:"service_id"`
	AppID             string                 `json:"app_id"`
	ServicePlanID     string                 `json:"service_plan_id"`
	PrivateKey        string                 `json:"private_key"`
	ServiceInstanceID string                 `json:"service_instance_id"`
	BindResource      map[string]interface{} `json:"bind_resource,omitempty"`
	Parameters        map[string]interface{} `json:"parameters,omitempty"`
}

// BindingRequest represents a request to bind to a service instance
type BindingRequest struct {
	AppGUID      string                 `json:"app_guid,omitempty"`
	PlanID       string                 `json:"plan_id,omitempty"`
	ServiceID    string                 `json:"service_id,omitempty"`
	BindResource map[string]interface{} `json:"bind_resource,omitempty"`
	Parameters   map[string]interface{} `json:"parameters,omitempty"`
}

// CreateServiceBindingResponse represents a response to a service binding
// request
type CreateServiceBindingResponse struct {
	Credentials Credential `json:"credentials"`
}

// Credential represents connection details, username, and password that are
// provisioned when a consumer binds to a service instance
type Credential map[string]interface{}

// Schemas represents a plan's schemas for service instance and binding create
// and update.
type Schemas struct {
	ServiceInstance *ServiceInstanceSchema `json:"service_instance,omitempty"`
	ServiceBinding  *ServiceBindingSchema  `json:"service_binding,omitempty"`
}

// ServiceInstanceSchema represents a plan's schemas for a create and update
// of a service instance.
type ServiceInstanceSchema struct {
	Create *InputParametersSchema `json:"create,omitempty"`
	Update *InputParametersSchema `json:"update,omitempty"`
}

// ServiceBindingSchema represents a plan's schemas for the parameters
// accepted for binding creation.
type ServiceBindingSchema struct {
	Create *RequestResponseSchema `json:"create,omitempty"`
}

// InputParametersSchema represents a schema for input parameters for
// creation or update of an API resource.
type InputParametersSchema struct {
	Parameters interface{} `json:"parameters,omitempty"`
}

// RequestResponseSchema represents a schema for both input parameters and
// the broker's response to the binding request
type RequestResponseSchema struct {
	InputParametersSchema
	Response interface{} `json:"response,omitempty"`
}

// Catalog is a JSON-compatible type to be used to decode the result from a /v2/catalog call
// to an open service broker compatible API
type Catalog struct {
	Services []*Service `json:"services"`
}
