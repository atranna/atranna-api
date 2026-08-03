package postgres

import (
	"database/sql"

	"github.com/atranna/atranna-api/src/internal/models"
)

type OrganizationRepository struct {
	db *sql.DB
}

func NewOrganizationRepository(db *sql.DB) *OrganizationRepository {
	return &OrganizationRepository{db: db}
}

func (r *OrganizationRepository) Add(organization models.Organization) (models.Organization, error) {
	err := r.db.QueryRow(
		`INSERT INTO organizations (name, slug) VALUES ($1, $2) RETURNING id`,
		organization.Name, organization.Slug,
	).Scan(&organization.ID)
	if err != nil {
		return models.Organization{}, err
	}
	return organization, nil
}

func (r *OrganizationRepository) GetByID(id int) (models.Organization, bool) {
	row := r.db.QueryRow(`SELECT id, name, slug FROM organizations WHERE id = $1`, id)
	var organization models.Organization
	err := row.Scan(&organization.ID, &organization.Name, &organization.Slug)
	if err != nil {
		return models.Organization{}, false
	}
	return organization, true
}

func (r *OrganizationRepository) GetAll() []models.Organization {
	rows, err := r.db.Query(`SELECT id, name, slug FROM organizations`)
	if err != nil {
		return []models.Organization{}
	}
	defer rows.Close()

	var organizations []models.Organization
	for rows.Next() {
		var organization models.Organization
		err := rows.Scan(&organization.ID, &organization.Name, &organization.Slug)
		if err != nil {
			continue
		}
		organizations = append(organizations, organization)
	}
	if err := rows.Err(); err != nil {
		return []models.Organization{}
	}
	return organizations
}

func (r *OrganizationRepository) Delete(id int) (models.Organization, bool) {
	row := r.db.QueryRow(`SELECT id, name, slug FROM organizations WHERE id = $1`, id)
	var organization models.Organization
	err := row.Scan(&organization.ID, &organization.Name, &organization.Slug)
	if err != nil {
		return models.Organization{}, false
	}

	_, err = r.db.Exec(`DELETE FROM organizations WHERE id = $1`, id)
	if err != nil {
		return models.Organization{}, false
	}

	return organization, true
}