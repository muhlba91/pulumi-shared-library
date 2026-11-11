package mocks

import (
	"sync"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Recorder is a mock that records the inputs (including secret markers) of every registered resource.
type Recorder struct {
	mu     sync.Mutex
	inputs map[string]resource.PropertyMap
}

// NewRecorder initializes and returns a new Recorder instance.
func NewRecorder() *Recorder {
	return &Recorder{inputs: make(map[string]resource.PropertyMap)}
}

// NewResource records the inputs of the resource and echoes them as outputs.
// args: MockResourceArgs containing information about the resource being mocked.
func (r *Recorder) NewResource(args pulumi.MockResourceArgs) (string, resource.PropertyMap, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.inputs[args.Name] = args.Inputs.Copy()

	return args.Name + "_id", args.Inputs, nil
}

// Call returns the call arguments unchanged.
// args: MockCallArgs containing information about the call being mocked.
func (r *Recorder) Call(args pulumi.MockCallArgs) (resource.PropertyMap, error) {
	return args.Args, nil
}

// Inputs returns the recorded inputs of the named resource.
// name: the resource name.
func (r *Recorder) Inputs(name string) (resource.PropertyMap, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	in, ok := r.inputs[name]

	return in, ok
}
