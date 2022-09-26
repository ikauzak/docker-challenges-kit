package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/stretchr/testify/assert"
)

const image = "helloworld"
const tag = "v1"
const fullImageName = image + ":" + tag

// Função para validar se a imagem existe dentro do host
func TestIfImageExists(t *testing.T) {
	t.Logf("Validando se a imagem %s existe...", fullImageName)
	output := docker.DoesImageExist(t, fullImageName, logger.Discard)
	if assert.True(t, output) != true {
		t.Fatalf("Imagem %s não encontrada!", fullImageName)
	}
	t.Logf("Imagem %s encontrada com sucesso!", fullImageName)
}

// Função para validar a saída do arquivo que está dentro da imagem
func TestDockerHelloWorldExample(t *testing.T) {
	t.Logf("Validando o arquivo /test.txt existe na imagem %s...", fullImageName)
	opts := &docker.RunOptions{Command: []string{"cat", "/test.txt"}, Remove: true}
	output := docker.Run(t, fullImageName, opts)
	assert.Equal(t, "Hello, World!", output)
}