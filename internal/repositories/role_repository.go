package repositories

import (
	"context"
	"service-store/config"
	"service-store/internal/models"

	"github.com/uptrace/bun"
)

func GetRoleByName(name string) (*models.Role, error) {
	role := new(models.Role)
	err := config.DB.NewSelect().Model(role).Where("name = ?", name).Limit(1).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return role, nil
}

func CreateRole(role *models.Role) error {
	_, err := config.DB.NewInsert().Model(role).Exec(context.Background())
	return err
}

func GetRoleByID(id int64) (*models.Role, error) {
	role := new(models.Role)
	err := config.DB.NewSelect().Model(role).Where("id = ?", id).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return role, nil
}

func GetPermissionsByIDs(ids []int64) ([]models.Permission, error) {
	var perms []models.Permission
	err := config.DB.NewSelect().Model(&perms).Where("id IN (?)", bun.In(ids)).Scan(context.Background())
	return perms, err
}

func GetRolePermissions(roleID int64, permissionIDs []int64) ([]models.RolePermission, error) {
	var rps []models.RolePermission
	err := config.DB.NewSelect().
		Model(&rps).
		Where("role_id = ? AND permission_id IN (?)", roleID, bun.In(permissionIDs)).
		Scan(context.Background())
	return rps, err
}

func CreateRolePermissions(rps []models.RolePermission) error {
	_, err := config.DB.NewInsert().Model(&rps).Exec(context.Background())
	return err
}

func GetPaginatedRoles(page, limit int) ([]models.Role, int, error) {
	var roles []models.Role
	offset := (page - 1) * limit

	count, err := config.DB.NewSelect().Model(&roles).Count(context.Background())
	if err != nil {
		return nil, 0, err
	}

	err = config.DB.NewSelect().Model(&roles).Limit(limit).Offset(offset).Scan(context.Background())
	if err != nil {
		return nil, 0, err
	}

	return roles, count, nil
}

func GetRoleWithPermissions(roleID int64) (*models.RolePermission, error) {
	role := new(models.RolePermission)

	err := config.DB.NewSelect().
		Model(role).
		Relation("RolePermissions.Permission").
		Where("role.id = ?", roleID).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}

	return role, nil
}
