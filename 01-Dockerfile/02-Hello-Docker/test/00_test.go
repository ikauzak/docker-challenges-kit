package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/stretchr/testify/assert"
)

const image = "helloworld"
const tag = "v2"
const fullImageName = image + ":" + tag

// Função para validar se a imagem existe dentro do host
func TestIfImageExists(t *testing.T) {
	output := docker.DoesImageExist(t, fullImageName, logger.TestingT)
	assert.True(t, output)
}

// Função para validar a saída do arquivo que está dentro da imagem
func TestDockerHelloWorldExample(t *testing.T) {
	opts := &docker.RunOptions{Command: []string{"cat", "/test.txt"}, Remove: true}
	output := docker.Run(t, fullImageName, opts)
	assert.Equal(t, "Hello, Docker!", output)
}
