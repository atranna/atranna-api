package sqlite

import (
	"database/sql"

	"github.com/atranna/atranna-api/src/internal/models"
	"github.com/atranna/atranna-api/src/internal/repository"
)

type InterfaceRepository struct {
	db *sql.DB
	devices repository.DeviceRepository
}

func NewInterfaceRepository(db *sql.DB, devices repository.DeviceRepository) *InterfaceRepository {
	return &InterfaceRepository{db: db, devices: devices}
}

func (r *InterfaceRepository) Add(interf models.Interface) (models.Interface, error) {
	res, err := r.db.Exec(
		`INSERT INTO interfaces (name, device_id, ip_address, mac_address, state, speed) VALUES (?, ?, ?, ?, ?, ?)`,
		interf.Name, interf.DeviceID, interf.IPAddress, interf.MACAddress, interf.State, interf.Speed,
	)
	if err != nil {
		return models.Interface{}, err
	}
	id, _ := res.LastInsertId()
	interf.ID = int(id)
	return interf, nil
}

func (r *InterfaceRepository) GetByID(id int) (models.Interface, bool) {
	row := r.db.QueryRow(`SELECT id, name, device_id, ip_address, mac_address, state, speed FROM interfaces WHERE id = ?`, id)
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
	rows, err := r.db.Query(`SELECT id, name, device_id, ip_address, mac_address, state, speed FROM interfaces WHERE device_id = ?`, deviceID)
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
	_, err := r.db.Exec(`DELETE FROM interfaces WHERE id = ?`, id)
	if err != nil {
		return models.Interface{}, false
	}
	return interf, true
}

func (r *InterfaceRepository) DeleteByDeviceID(deviceID int) error {
	_, err := r.db.Exec(`DELETE FROM interfaces WHERE device_id = ?`, deviceID)
	return err
}