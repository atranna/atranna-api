package sqlite

import (
	"database/sql"

	"github.com/atranna/atranna-api/src/internal/models"
)

type OrganizationMemberRepository struct {
	db *sql.DB
}

func NewOrganizationMemberRepository(db *sql.DB) *OrganizationMemberRepository {
	return &OrganizationMemberRepository{db: db}
}

func (r *OrganizationMemberRepository) Add(member models.OrganizationMember) (models.OrganizationMember, error) {
	var createdMember models.OrganizationMember
	err := r.db.QueryRow(
		`INSERT INTO organization_members (organization_id, user_id, role) VALUES (?, ?, ?) RETURNING organization_id, user_id, role`,
		member.OrganizationID, member.UserID, member.Role,
	).Scan(&createdMember.OrganizationID, &createdMember.UserID, &createdMember.Role)
	if err != nil {
		return models.OrganizationMember{}, err
	}
	return createdMember, nil
}

func (r *OrganizationMemberRepository) GetByOrganizationID(organizationID int) ([]models.OrganizationMember, error) {
	rows, err := r.db.Query(`SELECT organization_id, user_id, role FROM organization_members WHERE organization_id = ?`, organizationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []models.OrganizationMember
	for rows.Next() {
		var member models.OrganizationMember
		err := rows.Scan(&member.OrganizationID, &member.UserID, &member.Role)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return members, nil
}

func (r *OrganizationMemberRepository) GetByUserID(userID int) []models.OrganizationMember {
	rows, err := r.db.Query(`SELECT organization_id, user_id, role FROM organization_members WHERE user_id = ?`, userID)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var members []models.OrganizationMember
	for rows.Next() {
		var member models.OrganizationMember
		err := rows.Scan(&member.OrganizationID, &member.UserID, &member.Role)
		if err != nil {
			return nil
		}
		members = append(members, member)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return members
}

func (r *OrganizationMemberRepository) Delete(organizationID int, userID int) (models.OrganizationMember, bool) {
	row := r.db.QueryRow(`SELECT organization_id, user_id, role FROM organization_members WHERE organization_id = ? AND user_id = ?`, organizationID, userID)
	var member models.OrganizationMember
	err := row.Scan(&member.OrganizationID, &member.UserID, &member.Role)
	if err != nil {
		return models.OrganizationMember{}, false
	}

	_, err = r.db.Exec(`DELETE FROM organization_members WHERE organization_id = ? AND user_id = ?`, organizationID, userID)
	if err != nil {
		return models.OrganizationMember{}, false
	}

	return member, true
}

func (r *OrganizationMemberRepository) GetRole(organizationID int, userID int) (string, bool) {
	row := r.db.QueryRow(`SELECT role FROM organization_members WHERE organization_id = ? AND user_id = ?`, organizationID, userID)
	var role string
	err := row.Scan(&role)
	if err != nil {
		return "", false
	}
	return role, true
}