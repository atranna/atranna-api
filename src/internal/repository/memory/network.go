package memory

import (
	"github.com/atranna/atranna-api/src/internal/models"
)

type NetworkRepository struct {
	networks []models.Network
	nextID   int
}

func NewNetworkRepository() *NetworkRepository {
	return &NetworkRepository{}
}

func (r *NetworkRepository) Add(network models.Network) (models.Network, error) {
	r.nextID++
	network.ID = r.nextID
	r.networks = append(r.networks, network)
	return network, nil
}

func (r *NetworkRepository) GetByID(id int) (models.Network, bool) {
	for _, network := range r.networks {
		if network.ID == id {
			return network, true
		}
	}
	return models.Network{}, false
}

func (r *NetworkRepository) GetAll() []models.Network {
	return r.networks
}

func (r *NetworkRepository) Delete(id int) (models.Network, bool) {
	for i, network := range r.networks {
		if network.ID == id {
			r.networks = append(r.networks[:i], r.networks[i+1:]...)
			return network, true
		}
	}
	return models.Network{}, false
}
