package services

import (
	"service-store/internal/models"
	"service-store/internal/repositories"
	"service-store/utils"
)

type CreateRoleInput struct {
	Name string `json:"name" validate:"required"`
	Desc string `json:"desc"`
}

type AssignPermissionInput struct {
	RoleID        int64   `json:"roleId" validate:"required"`
	PermissionIDs []int64 `json:"permissionIds" validate:"required,min=1"`
}

func CreateRole(input CreateRoleInput) (*models.Role, error) {

	role := &models.Role{
		Name: input.Name,
		Desc: input.Desc,
	}
	if err := repositories.CreateRole(role); err != nil {
		return nil, err
	}
	return role, nil
}

// IsRoleNameTaken checks if the name already exists
func IsRoleNameTaken(name string) (bool, error) {
	role, err := repositories.GetRoleByName(name)

	if err != nil {
		if !utils.IsNotFoundError(err) {
			return false, err
		}
		return false, nil
	}

	return role != nil, nil
}

// func AssignPermissions(input AssignPermissionInput) error {
// 	role, _ := repositories.GetRoleByID(input.RoleID)
// 	if role == nil {
// 		return utils.BadRequestError("roleId", "This role not found")
// 	}

// 	validPerms, err := repositories.GetPermissionsByIDs(input.PermissionIDs)
// 	if err != nil {
// 		return err
// 	}

// 	validPermIDs := make(map[int64]bool)
// 	for _, p := range validPerms {
// 		validPermIDs[p.ID] = true
// 	}

// 	var invalidIDs []int64
// 	for _, id := range input.PermissionIDs {
// 		if !validPermIDs[id] {
// 			invalidIDs = append(invalidIDs, id)
// 		}
// 	}

// 	if len(invalidIDs) > 0 {
// 		return utils.BadRequestError("permissionIds", fmt.Sprintf("Invalid permission IDs: %v", invalidIDs))
// 	}

// 	existingPerms, _ := repositories.GetRolePermissions(input.RoleID, input.PermissionIDs)
// 	existingIDs := make(map[int64]bool)
// 	for _, rp := range existingPerms {
// 		existingIDs[rp.PermissionID] = true
// 	}

// 	var toInsert []models.RolePermission
// 	for _, id := range input.PermissionIDs {
// 		if !existingIDs[id] {
// 			toInsert = append(toInsert, models.RolePermission{
// 				RoleID:       input.RoleID,
// 				PermissionID: id,
// 			})
// 		}
// 	}

// 	if len(toInsert) == 0 {
// 		return nil
// 	}

// 	return repositories.CreateRolePermissions(toInsert)
// }

// func GetAllRoles(page, limit int) ([]models.Role, int, error) {
// 	return repositories.GetPaginatedRoles(page, limit)
// }

// func GetRoleByID(roleID int64) (*models.RoleWithPermissions, error) {
// 	return repositories.GetRoleWithPermissions(roleID)
// }
