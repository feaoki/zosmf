package jobs

// Estrutura base para o serviço de jobs do z/OSMF

import (
	"encoding/json"
	"io/ioutil"
	"zosmf/transport"
)

// Service representa o serviço de jobs, responsável por interagir com a API de jobs do z/OSMF.
type Service struct {
	tr *transport.Transport
}

// New cria uma nova instância do serviço de jobs.
func New(tr *transport.Transport) *Service {
	return &Service{tr: tr}
}

// JobInfo contém informações básicas sobre um job no z/OSMF.
type JobInfo struct {
	JobName string `json:"jobname"` // Nome do job
	JobID   string `json:"jobid"`   // ID do job
}

// GetJobs retorna a lista de jobs disponíveis no z/OSMF.
func (s *Service) GetJobs() ([]JobInfo, error) {
	resp, err := s.tr.DoRequest("GET", "/restjobs/jobs", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var list []JobInfo
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &list)
	return list, nil
}
