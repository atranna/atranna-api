package memory

import "github.com/atranna/atranna-api/src/internal/models"

type OrganizationMemberRepository struct {
	organizationMembers []models.OrganizationMember
}

func NewOrganizationMemberRepository() *OrganizationMemberRepository {
	return &OrganizationMemberRepository{}
}

func (r *OrganizationMemberRepository) Add(member models.OrganizationMember) (models.OrganizationMember, error) {
	r.organizationMembers = append(r.organizationMembers, member)
	return member, nil
}

func (r *OrganizationMemberRepository) GetByOrganizationID(organizationID int) ([]models.OrganizationMember, error) {
	var members []models.OrganizationMember
	for _, member := range r.organizationMembers {
		if member.OrganizationID == organizationID {
			members = append(members, member)
		}
	}
	return members, nil
}

func (r *OrganizationMemberRepository) GetByUserID(userID int) []models.OrganizationMember {
	var members []models.OrganizationMember
	for _, member := range r.organizationMembers {
		if member.UserID == userID {
			members = append(members, member)
		}
	}
	return members
}

func (r *OrganizationMemberRepository) Delete(organizationID int, userID int) (models.OrganizationMember, bool) {
	for i, member := range r.organizationMembers {
		if member.OrganizationID == organizationID && member.UserID == userID {
			r.organizationMembers = append(r.organizationMembers[:i], r.organizationMembers[i+1:]...)
			return member, true
		}
	}
	return models.OrganizationMember{}, false
}