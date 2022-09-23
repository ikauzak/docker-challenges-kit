package test

import (
	"testing"

	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

// Teste para validar se o registry está em execução local
func TestRegistryIsUpAndRunning(t *testing.T) {
	registry, containerId := findLocalRegistry()
	if assert.True(t, registry) != true {
		t.Fatalf("%s", containerId)
	}
	t.Logf("Container %s encontrado com sucesso!", containerId[:12])

	// Localizar a porta do host que está vinculada a porta 5000 do container
	containerPorts := docker.Inspect(t, containerId).GetExposedHostPort(5000)

	assert.Equal(t, int(containerPorts), 5000, "sao igyuais")
}

// Função para localizar o ID do container com o nome de registry em execução
func findLocalRegistry() (bool, string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if container.Names[0] == "/registry" {
			return true, container.ID
		}
	}
	return false, "Container não localizado"
}
