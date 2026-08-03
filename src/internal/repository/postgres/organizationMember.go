package postgres

import (
	"atranna-api/src/internal/models"
	"database/sql"
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
		`INSERT INTO organization_members (organization_id, user_id) VALUES ($1, $2) RETURNING organization_id, user_id`,
		member.OrganizationID, member.UserID,
	).Scan(&createdMember.OrganizationID, &createdMember.UserID)
	if err != nil {
		return models.OrganizationMember{}, err
	}
	return createdMember, nil
}

func (r *OrganizationMemberRepository) GetByOrganizationID(organizationID int) ([]models.OrganizationMember, error) {
	rows, err := r.db.Query(`SELECT organization_id, user_id FROM organization_members WHERE organization_id = $1`, organizationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []models.OrganizationMember
	for rows.Next() {
		var member models.OrganizationMember
		err := rows.Scan(&member.OrganizationID, &member.UserID)
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
	rows, err := r.db.Query(`SELECT organization_id, user_id FROM organization_members WHERE user_id = $1`, userID)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var members []models.OrganizationMember
	for rows.Next() {
		var member models.OrganizationMember
		err := rows.Scan(&member.OrganizationID, &member.UserID)
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
	row := r.db.QueryRow(`DELETE FROM organization_members WHERE organization_id = $1 AND user_id = $2 RETURNING organization_id, user_id`, organizationID, userID)
	var member models.OrganizationMember
	err := row.Scan(&member.OrganizationID, &member.UserID)
	if err != nil {
		return models.OrganizationMember{}, false
	}
	return member, true
}