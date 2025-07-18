package config

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// Config armazena a configuração para conexão com os serviços z/OSMF.
type Config struct {
	Host       string       // Nome do host ou IP do servidor z/OSMF
	Port       int          // Número da porta do servidor z/OSMF
	Username   string       // Nome de usuário para autenticação
	Password   string       // Senha para autenticação
	BasePath   string       // Caminho base para os endpoints da API (ex: "/zosmf")
	Insecure   bool         // Permitir conexões SSL inseguras
	HttpClient *http.Client // Cliente HTTP customizado (opcional)
}

// LoadConfigFromEnv carrega a configuração das variáveis de ambiente.
func LoadConfigFromEnv() (*Config, error) {
	host := os.Getenv("ZOSMF_HOST")
	portStr := os.Getenv("ZOSMF_PORT")
	username := os.Getenv("ZOSMF_USERNAME")
	password := os.Getenv("ZOSMF_PASSWORD")
	basePath := os.Getenv("ZOSMF_BASEPATH")
	insecureStr := os.Getenv("ZOSMF_INSECURE")

	if host == "" || portStr == "" || username == "" || password == "" {
		return nil, errors.New("variáveis de ambiente obrigatórias ausentes")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	insecure := false
	if insecureStr == "true" || insecureStr == "1" {
		insecure = true
	}

	return &Config{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		BasePath: basePath,
		Insecure: insecure,
	}, nil
}

// LoadConfigFromFile carrega a configuração de um arquivo JSON.
func LoadConfigFromFile(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// loadConfigFromDefaultPaths tenta carregar a configuração dos caminhos padrão de arquivos de configuração.
func loadConfigFromDefaultPaths() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	paths := []string{
		filepath.Join(home, ".config", "zosmf", "config.json"),
		filepath.Join(home, ".zosmf_config.json"),
	}
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return LoadConfigFromFile(path)
		}
	}
	return nil, errors.New("nenhum arquivo de configuração encontrado nos locais padrão")
}

// LoadDefaultConfig carrega a configuração das variáveis de ambiente, depois dos arquivos de configuração, ou retorna um erro.
func LoadDefaultConfig() (*Config, error) {
	cfg, err := LoadConfigFromEnv()
	if err == nil {
		return cfg, nil
	}
	cfg, err = loadConfigFromDefaultPaths()
	if err == nil {
		return cfg, nil
	}
	return nil, errors.New("não foi possível carregar a configuração das variáveis de ambiente ou dos arquivos de configuração")
}
