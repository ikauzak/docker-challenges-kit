package test

import (
	"io"
	"os"
	"testing"
	"time"

	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/stretchr/testify/assert"
)

// Teste para validar se o registry está em execução local
func TestRegistryIsUpAndRunning(t *testing.T) {
	hostPort := 5000

	registry, containerId := findLocalRegistry()
	if assert.True(t, registry) != true {
		t.Fatalf("%s", containerId)
	}
	t.Logf("Container registry com ID %s encontrado com sucesso!", containerId[:12])

	// Localizar a porta do host que está vinculada a porta 5000 do container
	containerPorts := docker.Inspect(t, containerId).GetExposedHostPort(5000)

	if assert.Equal(t, int(containerPorts), hostPort) != true {
		t.Fatalf("Registry deveria estar vinculada à porta %v do host", hostPort)
	}
}

// Teste para validar envio de imagens ao registry local
func TestPushToLocalRegistry(t *testing.T) {
	imageToPush := pullImage("busybox:1.35.0")
	docker.Push(t, logger.Discard, imageToPush)
	t.Logf("Enviado imagem %s de teste com sucesso!", imageToPush)

	// Remover imagem temporária armazenada localmente
	docker.DeleteImage(t, imageToPush, logger.Discard)
	t.Logf("Removido imagem %s localmente com sucesso!", imageToPush)
}

func pullImage(testImage string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	reader, err := cli.ImagePull(ctx, testImage, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer reader.Close()
	io.Copy(os.Stdout, reader)

	t := time.Now()
	newTag := "localhost:5000/gotest:" + t.Format("200610102150405")
	cli.ImageTag(ctx, testImage, newTag)

	return newTag
}

func tagImage(newTag string) string {
	t := time.Now()
	newTag := "localhost:5000/gotest:" + t.Format("200610102150405")
	cli.ImageTag(ctx, testImage, newTag)

	return newTag
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
