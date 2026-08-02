package postgres

import (
	"database/sql"

	"atranna-api/src/internal/models"
	"atranna-api/src/internal/repository"
)

type InterfaceRepository struct {
	db *sql.DB
	devices repository.DeviceRepository
}

func NewInterfaceRepository(db *sql.DB, devices repository.DeviceRepository) *InterfaceRepository {
	return &InterfaceRepository{db: db, devices: devices}
}

func (r *InterfaceRepository) Add(interf models.Interface) (models.Interface, error) {
	err := r.db.QueryRow(
		`INSERT INTO interfaces (name, device_id, ip_address, mac_address, state, speed) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		interf.Name, interf.DeviceID, interf.IPAddress, interf.MACAddress, interf.State, interf.Speed,
	).Scan(&interf.ID)
	if err != nil {
		return models.Interface{}, err
	}
	return interf, nil
}

func (r *InterfaceRepository) GetByID(id int) (models.Interface, bool) {
	row := r.db.QueryRow(`SELECT id, name, device_id, ip_address, mac_address, state, speed FROM interfaces WHERE id = $1`, id)
	var interf models.Interface
	err := row.Scan(&interf.ID, &interf.Name, &interf.DeviceID, &interf.IPAddress, &interf.MACAddress, &interf.State, &interf.Speed)
	if err != nil {
		return models.Interface{}, false
	}
	return interf, true
}

func (r *InterfaceRepository) GetAll() ([]models.Interface) {
	rows, err := r.db.Query(`SELECT id, name, device_id, ip_address, mac_address, state, speed FROM interfaces`)
	if err != nil {
		return []models.Interface{}
	}
	defer rows.Close()

	var interfaces []models.Interface
	for rows.Next() {
		var interf models.Interface
		err := rows.Scan(&interf.ID, &interf.Name, &interf.DeviceID, &interf.IPAddress, &interf.MACAddress, &interf.State, &interf.Speed)
		if err != nil {
			continue
		}
		interfaces = append(interfaces, interf)
	}
	if err := rows.Err(); err != nil {
		return []models.Interface{}
	}
	return interfaces
}

func (r *InterfaceRepository) GetByDeviceID(deviceID int) []models.Interface {
	rows, err := r.db.Query(`SELECT id, name, device_id, ip_address, mac_address, state, speed FROM interfaces WHERE device_id = $1`, deviceID)
	if err != nil {
		return []models.Interface{}
	}
	defer rows.Close()

	var interfaces []models.Interface
	for rows.Next() {
		var interf models.Interface
		err := rows.Scan(&interf.ID, &interf.Name, &interf.DeviceID, &interf.IPAddress, &interf.MACAddress, &interf.State, &interf.Speed)
		if err != nil {
			continue
		}
		interfaces = append(interfaces, interf)
	}
	if err := rows.Err(); err != nil {
		return []models.Interface{}
	}
	return interfaces
}

func (r *InterfaceRepository) Delete(id int) (models.Interface, bool) {
	interf, found := r.GetByID(id)
	if !found {
		return models.Interface{}, false
	}
	_, err := r.db.Exec(`DELETE FROM interfaces WHERE id = $1`, id)
	if err != nil {
		return models.Interface{}, false
	}
	return interf, true
}

func (r *InterfaceRepository) DeleteByDeviceID(deviceID int) error {
	_, err := r.db.Exec(`DELETE FROM interfaces WHERE device_id = $1`, deviceID)
	return err
}