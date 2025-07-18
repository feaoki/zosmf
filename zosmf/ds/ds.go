package ds

import (
	"encoding/json"
	"fmt"
	"io"
	"zosmf/transport"
)

// Serviço para manipulação de datasets via z/OSMF.
type Service struct {
	tr *transport.Transport
}

// Cria uma nova instância do serviço de datasets.
func New(tr *transport.Transport) *Service {
	return &Service{tr: tr}
}

// Estrutura que representa informações de um dataset.
type DatasetInfo struct {
	DSName string `json:"dsname"`
	Type   string `json:"type"`
}

// List recupera uma lista de datasets do z/OSMF.
func (s *Service) List() ([]DatasetInfo, error) {
	resp, err := s.tr.DoRequest("GET", "/restfiles/ds", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var list []DatasetInfo
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		return nil, err
	}
	return list, nil
}

// Get recupera informações sobre um dataset específico.
func (s *Service) Get(dsname string) (*DatasetInfo, error) {
	path := fmt.Sprintf("/restfiles/ds/%s", dsname)
	resp, err := s.tr.DoRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var info DatasetInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}
	return &info, nil
}

// Create cria um novo dataset.
func (s *Service) Create(info DatasetInfo) error {
	resp, err := s.tr.DoRequest("POST", "/restfiles/ds", info)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 && resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("falha ao criar dataset: %s", string(body))
	}
	return nil
}

// Delete remove um dataset.
func (s *Service) Delete(dsname string) error {
	path := fmt.Sprintf("/restfiles/ds/%s", dsname)
	resp, err := s.tr.DoRequest("DELETE", path, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 204 && resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("falha ao remover dataset: %s", string(body))
	}
	return nil
}

/*
 Serviço para manipulação de datasets via z/OSMF.

 Métodos disponíveis:
 - List: Lista todos os datasets.
 - Get: Obtém informações de um dataset específico.
 - Create: Cria um novo dataset.
 - Delete: Remove um dataset existente.
*/
