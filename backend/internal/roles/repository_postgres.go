package roles

import (
	"context"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/codercollo/hotel-system/backend/pkg/idgen"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepository struct {
	db *pgxpool.Pool
}

// NewPostgresRepository creates a Postgres-backed roles repository.
func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) ListRoles(ctx context.Context) ([]*Role, error) {
	rows, err := r.db.Query(ctx, `SELECT id, name, description, created_at, updated_at FROM roles ORDER BY name`)
	if err != nil {
		return nil, apierror.Internal(err)
	}
	defer rows.Close()

	var roles []*Role
	for rows.Next() {
		role := &Role{}
		if err := rows.Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt); err != nil {
			return nil, apierror.Internal(err)
		}
		role.Permissions, _ = r.FindPermissionsForRole(ctx, role.ID)
		roles = append(roles, role)
	}
	return roles, nil
}

func (r *postgresRepository) FindRoleByName(ctx context.Context, name string) (*Role, error) {
	role := &Role{}
	err := r.db.QueryRow(ctx,
		`SELECT id, name, description, created_at, updated_at FROM roles WHERE name=$1`, name,
	).Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		return nil, apierror.NotFound(apierror.CodeRoleNotFound, "role not found")
	}
	role.Permissions, _ = r.FindPermissionsForRole(ctx, role.ID)
	return role, nil
}

func (r *postgresRepository) FindPermissionsForRole(ctx context.Context, roleID string) ([]Permission, error) {
	rows, err := r.db.Query(ctx,
		`SELECT p.id, p.name, p.description, p.created_at
		 FROM permissions p
		 JOIN role_permissions rp ON p.id = rp.permission_id
		 WHERE rp.role_id = $1`, roleID,
	)
	if err != nil {
		return nil, apierror.Internal(err)
	}
	defer rows.Close()

	var perms []Permission
	for rows.Next() {
		p := Permission{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt); err != nil {
			return nil, apierror.Internal(err)
		}
		perms = append(perms, p)
	}
	return perms, nil
}

func (r *postgresRepository) HasPermission(ctx context.Context, roleName, permission string) (bool, error) {
	var count int
	err := r.db.QueryRow(ctx,
		`SELECT COUNT(*) FROM permissions p
		 JOIN role_permissions rp ON p.id = rp.permission_id
		 JOIN roles ro ON ro.id = rp.role_id
		 WHERE ro.name=$1 AND p.name=$2`, roleName, permission,
	).Scan(&count)
	return count > 0, err
}

func (r *postgresRepository) SeedDefaults(ctx context.Context) error {
	// Seed roles
	defaultRoles := []struct{ name, desc string }{
		{"admin", "Full platform access"},
		{"user", "Standard guest access"},
	}
	for _, role := range defaultRoles {
		roleID := idgen.NewULID()
		_, err := r.db.Exec(ctx,
			`INSERT INTO roles (id, name, description) VALUES ($1,$2,$3)
			 ON CONFLICT (name) DO NOTHING`,
			roleID, role.name, role.desc,
		)
		if err != nil {
			return apierror.Internal(err)
		}
	}

	// Seed permissions
	allPerms := []string{
		"items:create", "items:read", "items:update", "items:delete",
		"orders:create", "orders:read", "orders:update", "orders:cancel",
		"payments:read", "payments:refund",
		"users:read", "users:update", "users:delete",
		"roles:read", "roles:update",
		"notifications:read", "uploads:create",
	}
	for _, perm := range allPerms {
		permID := idgen.NewULID()
		_, err := r.db.Exec(ctx,
			`INSERT INTO permissions (id, name, description) VALUES ($1,$2,$3)
			 ON CONFLICT (name) DO NOTHING`,
			permID, perm, perm,
		)
		if err != nil {
			return apierror.Internal(err)
		}
	}

	return nil
}
