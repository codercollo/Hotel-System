package roles

import "context"

// Repository defines persistence for roles and permissions.
type Repository interface {
	ListRoles(ctx context.Context) ([]*Role, error)
	FindRoleByName(ctx context.Context, name string) (*Role, error)
	FindPermissionsForRole(ctx context.Context, roleID string) ([]Permission, error)
	HasPermission(ctx context.Context, roleName, permission string) (bool, error)
	SeedDefaults(ctx context.Context) error
}
