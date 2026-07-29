package store

import (
	"atranna-api/src/internal/models"
)

var nextNetworkID int = 1

var Networks = []models.Network{}

func AddNetwork(network models.Network) models.Network {
	network.ID = nextNetworkID
	nextNetworkID++
	Networks = append(Networks, network)
	return network
}

func GetNetworkByID(id int) (models.Network, bool) {
	for _, network := range Networks {
		if network.ID == id {
			return network, true
		}
	}
	return models.Network{}, false
}

func DeleteNetwork(id int) (models.Network, bool) {
	for i, network := range Networks {
		if network.ID == id {
			Networks = append(Networks[:i], Networks[i+1:]...)
			return network, true
		}
	}

	return models.Network{}, false
}
