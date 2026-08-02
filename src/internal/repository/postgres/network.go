package postgres

import (
	"database/sql"

	"atranna-api/src/internal/models"
)

type NetworkRepository struct {
	db *sql.DB
}

func NewNetworkRepository(db *sql.DB) *NetworkRepository {
	return &NetworkRepository{db: db}
}

func (r *NetworkRepository) Add(network models.Network) models.Network {
	res, err := r.db.Exec(
		`INSERT INTO networks (name, cidr, gateway) VALUES (?, ?, ?)`,
		network.Name, network.CIDR, network.Gateway,
	)
	if err != nil {
		return models.Network{}
	}
	id, _ := res.LastInsertId()
	network.ID = int(id)
	return network
}

func (r *NetworkRepository) GetByID(id int) (models.Network, bool) {
	row := r.db.QueryRow(`SELECT id, name, cidr, gateway FROM networks WHERE id = ?`, id)
	var network models.Network
	err := row.Scan(&network.ID, &network.Name, &network.CIDR, &network.Gateway)
	if err != nil {
		return models.Network{}, false
	}
	return network, true
}

func (r *NetworkRepository) GetAll() []models.Network {
	rows, err := r.db.Query(`SELECT id, name, cidr, gateway FROM networks`)
	if err != nil {
		return []models.Network{}
	}
	defer rows.Close()

	var networks []models.Network
	for rows.Next() {
		var network models.Network
		err := rows.Scan(&network.ID, &network.Name, &network.CIDR, &network.Gateway)
		if err != nil {
			continue
		}
		networks = append(networks, network)
	}
	if err := rows.Err(); err != nil {
		return []models.Network{}
	}
	return networks
}

func (r *NetworkRepository) Delete(id int) (models.Network, bool) {
	res, err := r.db.Exec(`DELETE FROM networks WHERE id = ?`, id)
	if err != nil {
		return models.Network{}, false
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return models.Network{}, false
	}
	return models.Network{ID: id}, true
}