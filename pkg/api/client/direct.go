package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/kubeshop/kubetest/pkg/api/kubetest"
	"github.com/kubeshop/kubetest/pkg/problem"
)

const (
	ClientHTTPTimeout = 20 * time.Second
)

type Config struct {
	URI string `default:"http://localhost:8080"`
}

var config Config

func init() {
	envconfig.Process("KUBETEST_API", &config)
}
func NewDirectScriptsAPI(uri string) DirectScriptsAPI {
	return DirectScriptsAPI{
		URI: uri,
		client: &http.Client{
			Timeout: ClientHTTPTimeout,
		},
	}
}

func NewDefaultDirectScriptsAPI() DirectScriptsAPI {
	return NewDirectScriptsAPI(config.URI)
}

type DirectScriptsAPI struct {
	URI    string
	client HTTPClient
}

func (c DirectScriptsAPI) GetScript(id string) (script kubetest.Script, err error) {
	uri := fmt.Sprintf(c.URI+"/v1/scripts/%s", id)
	resp, err := c.client.Get(uri)
	if err != nil {
		return script, err
	}

	if err := c.responseError(resp); err != nil {
		return script, fmt.Errorf("api/get-script returned error: %w", err)
	}

	return c.getScriptFromResponse(resp)
}

func (c DirectScriptsAPI) GetExecution(scriptID, executionID string) (execution kubetest.ScriptExecution, err error) {
	uri := fmt.Sprintf(c.URI+"/v1/scripts/%s/executions/%s", scriptID, executionID)
	resp, err := c.client.Get(uri)
	if err != nil {
		return execution, err
	}

	if err := c.responseError(resp); err != nil {
		return execution, fmt.Errorf("api/get-execution returned error: %w", err)
	}

	return c.getExecutionFromResponse(resp)
}

// ListExecutions list all executions for given script name
func (c DirectScriptsAPI) ListExecutions(scriptID string) (executions kubetest.ScriptExecutions, err error) {
	uri := fmt.Sprintf(c.URI+"/v1/scripts/%s/executions", scriptID)
	resp, err := c.client.Get(uri)
	if err != nil {
		return executions, err
	}

	if err := c.responseError(resp); err != nil {
		return executions, fmt.Errorf("api/get-executions returned error: %w", err)
	}

	return c.getExecutionsFromResponse(resp)
}

// CreateScript creates new Script Custom Resource
func (c DirectScriptsAPI) CreateScript(name, scriptType, content, namespace string) (script kubetest.Script, err error) {
	uri := fmt.Sprintf(c.URI + "/v1/scripts")

	request := kubetest.ScriptCreateRequest{
		Name:      name,
		Content:   content,
		Type_:     scriptType,
		Namespace: namespace,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return script, err
	}

	resp, err := c.client.Post(uri, "application/json", bytes.NewReader(body))
	if err != nil {
		return script, err
	}

	if err := c.responseError(resp); err != nil {
		return script, fmt.Errorf("api/create-script returned error: %w", err)
	}

	return c.getScriptFromResponse(resp)
}

// ExecuteScript starts new external script execution, reads data and returns ID
// Execution is started asynchronously client can check later for results
func (c DirectScriptsAPI) ExecuteScript(id, namespace, executionName string, executionParams map[string]string) (execution kubetest.ScriptExecution, err error) {
	// TODO call executor API - need to get parameters (what executor?) taken from CRD?
	uri := fmt.Sprintf(c.URI+"/v1/scripts/%s/executions", id)

	request := kubetest.ScriptExecutionRequest{
		Name:      executionName,
		Namespace: namespace,
		Params:    executionParams,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return execution, err
	}

	resp, err := c.client.Post(uri, "application/json", bytes.NewReader(body))
	if err != nil {
		return execution, err
	}

	if err := c.responseError(resp); err != nil {
		return execution, fmt.Errorf("api/execute-script returned error: %w", err)
	}

	return c.getExecutionFromResponse(resp)
}

// GetExecutions list all executions in given script
func (c DirectScriptsAPI) ListScripts(namespace string) (scripts kubetest.Scripts, err error) {
	uri := fmt.Sprintf(c.URI+"/v1/scripts?namespace=%s", namespace)
	resp, err := c.client.Get(uri)
	if err != nil {
		return scripts, fmt.Errorf("GET client error: %w", err)
	}
	defer resp.Body.Close()

	if err := c.responseError(resp); err != nil {
		return scripts, fmt.Errorf("api/list-scripts returned error: %w", err)
	}

	err = json.NewDecoder(resp.Body).Decode(&scripts)
	return
}

func (c DirectScriptsAPI) getExecutionFromResponse(resp *http.Response) (execution kubetest.ScriptExecution, err error) {
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&execution)
	return
}

func (c DirectScriptsAPI) getExecutionsFromResponse(resp *http.Response) (executions kubetest.ScriptExecutions, err error) {
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&executions)
	return
}

func (c DirectScriptsAPI) getScriptFromResponse(resp *http.Response) (script kubetest.Script, err error) {
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&script)
	return
}
func (c DirectScriptsAPI) getProblemFromResponse(resp *http.Response) (p problem.Problem, err error) {
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&p)
	return
}

func (c DirectScriptsAPI) responseError(resp *http.Response) error {
	if resp.StatusCode >= 400 {
		pr, err := c.getProblemFromResponse(resp)
		if err != nil {
			return fmt.Errorf("can't get problem from api response: %w", err)
		}

		return fmt.Errorf("problem: %+v", pr.Detail)
	}

	return nil
}
