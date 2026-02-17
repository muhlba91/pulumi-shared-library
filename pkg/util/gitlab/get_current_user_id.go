package gitlab

import (
	"strconv"

	gl "github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GetCurrentUserID is a helper function that retrieves the current user's ID using the GetCurrentUser data source. It returns a pointer to the user ID string, or nil if there was an error retrieving the user information.
// ctx: The Pulumi context used for invoking the GetCurrentUser data source.
func GetCurrentUserID(ctx *pulumi.Context) *int {
	user, uErr := gl.GetCurrentUser(ctx)
	if uErr != nil {
		return nil
	}

	id, err := strconv.Atoi(user.Id)
	if err != nil {
		return nil
	}
	return &id
}
