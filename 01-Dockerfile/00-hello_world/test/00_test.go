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
func TestIfImageExist(t *testing.T) {
	output := docker.DoesImageExist(t, fullImageName, logger.TestingT)
	assert.True(t, output)
}

// Função para validar output do arquivo dentro da imagem
func TestDockerHelloWorldExample(t *testing.T) {
	/*
		tag := "gruntwork/docker-hello-world-example"
		buildOptions := &docker.BuildOptions{
			Tags: []string{tag},
		}

		docker.Build(t, "../", buildOptions)
	*/

	opts := &docker.RunOptions{Command: []string{"cat", "/test.txt"}, Remove: true}
	output := docker.Run(t, fullImageName, opts)
	assert.Equal(t, "Hello, World!", output)
}
