# Primeiro estágio
FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build-env
WORKDIR /App

# Copiar o conteúdo do diretório local para a imagem
COPY . ./
RUN dotnet restore
# Compilar pacotes
RUN dotnet publish -c Release -o out

# Segundo estágio
FROM mcr.microsoft.com/dotnet/aspnet:6.0
WORKDIR /App
# Copiando artefatos criados no primeiro estágio
COPY --from=build-env /App/out .
ENTRYPOINT ["dotnet", "DotNet.Docker.dll"]
