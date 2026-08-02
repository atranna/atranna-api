package sqlite

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

func (r *NetworkRepository) Add(network models.Network) (models.Network, error) {
	res, err := r.db.Exec(
		`INSERT INTO networks (name, cidr, gateway, vlan) VALUES (?, ?, ?, ?)`,
		network.Name, network.CIDR, network.Gateway, network.Vlan,
	)
	if err != nil {
		return models.Network{}, err
	}
	id, _ := res.LastInsertId()
	network.ID = int(id)
	return network, nil
}

func (r *NetworkRepository) GetByID(id int) (models.Network, bool) {
	row := r.db.QueryRow(`SELECT id, name, cidr, gateway, vlan FROM networks WHERE id = ?`, id)
	var network models.Network
	var vlan sql.NullInt64
	err := row.Scan(&network.ID, &network.Name, &network.CIDR, &network.Gateway, &vlan)
	if err != nil {
		return models.Network{}, false
	}
	network.Vlan = int(vlan.Int64)
	return network, true
}

func (r *NetworkRepository) GetAll() []models.Network {
	rows, err := r.db.Query(`SELECT id, name, cidr, gateway, vlan FROM networks`)
	if err != nil {
		return []models.Network{}
	}
	defer rows.Close()

	var networks []models.Network
	for rows.Next() {
		var network models.Network
		var vlan sql.NullInt64
		err := rows.Scan(&network.ID, &network.Name, &network.CIDR, &network.Gateway, &vlan)
		if err != nil {
			continue
		}
		network.Vlan = int(vlan.Int64)
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