package postgres

import (
	"database/sql"

	"github.com/atranna/atranna-api/src/internal/models"
)

type DeviceRepository struct {
	db *sql.DB
}

func NewDeviceRepository(db *sql.DB) *DeviceRepository {
	return &DeviceRepository{db: db}
}

func (r *DeviceRepository) Add(device models.Device) (models.Device, error) {
	err := r.db.QueryRow(
		`INSERT INTO devices (hostname, ip, vendor, model, type) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		device.Hostname, device.IP, device.Vendor, device.Model, device.Type,
	).Scan(&device.ID)
	if err != nil {
		return models.Device{}, err
	}
	return device, nil
}

func (r *DeviceRepository) GetByID(id int) (models.Device, bool) {
	row := r.db.QueryRow(`SELECT id, hostname, ip, vendor, model, type FROM devices WHERE id = $1`, id)
	var device models.Device
	err := row.Scan(&device.ID, &device.Hostname, &device.IP, &device.Vendor, &device.Model, &device.Type)
	if err != nil {
		return models.Device{}, false
	}
	return device, true
}

func (r *DeviceRepository) GetAll() []models.Device {
	rows, err := r.db.Query(`SELECT id, hostname, ip, vendor, model, type FROM devices`)
	if err != nil {
		return []models.Device{}
	}
	defer rows.Close()

	var devices []models.Device
	for rows.Next() {
		var device models.Device
		err := rows.Scan(&device.ID, &device.Hostname, &device.IP, &device.Vendor, &device.Model, &device.Type)
		if err != nil {
			continue
		}
		devices = append(devices, device)
	}
	if err := rows.Err(); err != nil {
		return []models.Device{}
	}
	return devices
}

func (r *DeviceRepository) Delete(id int) (models.Device, bool) {
	device, found := r.GetByID(id)
	if !found {
		return models.Device{}, false
	}
	_, err := r.db.Exec(`DELETE FROM devices WHERE id = $1`, id)
	if err != nil {
		return models.Device{}, false
	}
	return device, true
}