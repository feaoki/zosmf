// transport/transport.go
package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"zosmf/config"
)

// Transport representa o transporte HTTP para comunicação com o servidor z/OSMF
type Transport struct {
	client *http.Client   // Cliente HTTP utilizado para as requisições
	cfg    *config.Config // Configurações do servidor z/OSMF
}

// New cria uma nova instância de Transport com base nas configurações fornecidas
func New(cfg *config.Config) *Transport {
	return &Transport{
		client: cfg.HttpClient,
		cfg:    cfg,
	}
}

// DoRequest executa uma requisição HTTP com o método, caminho e corpo especificados
func (t *Transport) DoRequest(method, path string, body any) (*http.Response, error) {
	url := fmt.Sprintf("%s:%d%s%s", t.cfg.Host, t.cfg.Port, t.cfg.BasePath, path)

	var reader io.Reader
	if body != nil {
		b, _ := json.Marshal(body) // Serializa o corpo para JSON
		reader = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(t.cfg.Username, t.cfg.Password)   // Define autenticação básica
	req.Header.Set("Content-Type", "application/json") // Define o tipo de conteúdo como JSON

	return t.client.Do(req) // Executa a requisição HTTP
}
