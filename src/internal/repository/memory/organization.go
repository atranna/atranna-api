package memory

import "atranna-api/src/internal/models"

type OrganizationRepository struct {
	organizations []models.Organization
	nextID int
}

func NewOrganizationRepository() *OrganizationRepository {
	return &OrganizationRepository{}
}

func (r *OrganizationRepository) Add(organization models.Organization) (models.Organization, error) {
	r.nextID++
	organization.ID = r.nextID
	r.organizations = append(r.organizations, organization)
	return organization, nil
}

func (r *OrganizationRepository) GetByID(id int) (models.Organization, bool) {
	for _, organization := range r.organizations {
		if organization.ID == id {
			return organization, true
		}
	}
	return models.Organization{}, false
}

func (r *OrganizationRepository) GetAll() []models.Organization {
	return r.organizations
}

func (r *OrganizationRepository) Delete(id int) (models.Organization, bool) {
	for i, organization := range r.organizations {
		if organization.ID == id {
			r.organizations = append(r.organizations[:i], r.organizations[i+1:]...)
			return organization, true
		}
	}
	return models.Organization{}, false
}