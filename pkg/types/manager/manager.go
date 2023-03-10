// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package manager

import (
	"context"

	"github.com/anydotcloud/grm/pkg/compare"
	"github.com/anydotcloud/grm/pkg/types/resource"
)

// Manager is responsible for providing a consistent way to perform
// CRUD operations in a backend cloud service API for Resources.
//
// Use a Factory to create a Manager for a
// particular Kind and Credentials.
type Manager interface {
	// Get returns the currently-observed state of the supplied Resource in
	// the backend cloud service API.
	//
	// Implementers should return (nil, pkg/errors.NotFound) when the backend
	// cloud service API cannot find the resource.
	Get(
		context.Context,
		resource.Identifiers,
	) (resource.Resource, error)
	// Create attempts to create the supplied Resource in the backend cloud
	// service API, returning a Resource representing the newly-created
	// resource
	Create(
		context.Context,
		resource.Resource, // desired
	) (resource.Resource, error)
	// Update attempts to mutate the supplied desired Resource in the backend
	// cloud service API, returning a Resource representing the newly-mutated
	// resource.
	//
	// A `compare.Delta` is provided to help implementing structs understand
	// which fields between the desired and latest resources have changed.
	Update(
		context.Context,
		resource.Resource, // desired
		resource.Resource, // latest
		*compare.Delta,
	) (resource.Resource, error)
	// Delete attempts to destroy the supplied Resource in the backend cloud
	// service API, returning a Resource representing the resource being
	// deleted (if delete is asynchronous and takes time)
	Delete(
		context.Context,
		resource.Resource, // latest
	) (resource.Resource, error)
}
