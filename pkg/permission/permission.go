package permission

import (
	"context"
	"encoding/json"

	"github.com/OpenSlides/openslides-permission-service/internal/definitions"

	"github.com/OpenSlides/openslides-permission-service/internal/core"
)

// New returns a new permission service.
func New(dataprovider ExternalDataProvider) Permission {
	return core.NewPermissionService(dataprovider)
}

// ExternalDataProvider is the connection to the datastore. It returns the data
// required by the permission service.
type ExternalDataProvider interface {
	// If a field does not exist, it is not returned.
	Get(ctx context.Context, fqfields ...string) ([]json.RawMessage, error)
}

// Permission can tell, if a user has the permission for some data.
//
// See https://github.com/FinnStutzenstein/OpenSlides/blob/permissionService/docs/interfaces/permission-service.txt
type Permission interface {
	IsAllowed(ctx context.Context, name string, userID definitions.Id, data definitions.FqfieldData) (bool, map[string]interface{}, error)
	RestrictFQIDs(ctx context.Context, userID definitions.Id, fqids []definitions.Fqid) (map[definitions.Fqid]bool, error)
	RestrictFQFields(ctx context.Context, userID definitions.Id, fqfields []definitions.Fqfield) (map[definitions.Fqfield]bool, error)
	AdditionalUpdate(ctx context.Context, updated definitions.FqfieldData) ([]definitions.Id, error)
}
