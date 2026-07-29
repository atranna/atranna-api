package store

import "atranna-api/src/internal/models"

var nextInterfaceID int = 1

var Interfaces = []models.Interface{}

func AddInterface(interf models.Interface, device_id int) models.Interface {
	interf.Device_ID = device_id
	interf.ID = nextInterfaceID
	nextInterfaceID++
	Interfaces = append(Interfaces, interf)
	return interf
}

func GetInterfaceByID(id int) (models.Interface, bool) {
	for _, interf := range Interfaces {
		if interf.ID == id {
			return interf, true
		}
	}
	return models.Interface{}, false
}

func GetInterfacesByDeviceID(device_id int) []models.Interface {
	result := []models.Interface{}
	for _, interf := range Interfaces {
		if interf.Device_ID == device_id {
			result = append(result, interf)
		}
	}
	return result
}

func DeleteInterface(id int) (models.Interface, bool) {
	for i, interf := range Interfaces {
		if interf.ID == id {
			Interfaces = append(Interfaces[:i], Interfaces[i+1:]...)
			DeleteInterfacesByDeviceID(interf.Device_ID)
			return interf, true
		}
	}
	return models.Interface{}, false
}


func DeleteInterfacesByDeviceID(device_id int) {
	filtered := make([]models.Interface, 0, len(Interfaces))
	for _, interf := range Interfaces {
		if interf.Device_ID != device_id {
			filtered = append(filtered, interf)
		}
	}
	Interfaces = filtered
}