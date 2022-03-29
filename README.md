# Hands On Hyperledeger Fabric e-Workshop

Este e-Workshop tem como objetivo construir um Chaincode Fabric em Go, simples e abstrato, capaz de suportar a rastreabilidade de **Lotes** e **Atividades** sobre esses mesmos lotes no protocolo do Hyperledger Fabric.

Teremos a possibilidade de criar Lotes provenientes de outros Lotes através das Atividades, registos esses que serão rastráveis no ledger do Fabric. 

![Diagrama de rastreabilidade](https://github.com/lcvalves/fabric_e-workshop/blob/master/diagrams/traceability-diagram.png?raw=true)

Para suportar estas funcionalidades, iremos desenvolver 1 único smart contract capaz de representar os **Lotes** e **Atividades**, com restrições a nível de validação e verificação dos dados definidos, nomeadamente as quantidades dos lotes.

![Modelo Structs Go](https://github.com/lcvalves/fabric_e-workshop/blob/master/diagrams/go-struct-model.png?raw=true)

# Ambiente de desenvolvimento & Software
O software a instalar deve ser instalado no SO do ambiente de desenvolvimento:

 - [ ] Sistema operativo baseado em **Unix**:
	 - [ ] **Linux**, **macOS**, etc... Utilizadores Windows podem utilizar (preferencialmente) o **[WSL2](https://docs.microsoft.com/en-us/windows/wsl/install)** ou máquinas virtuais como **VirtualBox**, **VMware** ou **Hyper-V** em conjunto com uma distribuição **Linux (ex: [Ubuntu LTS](https://ubuntu.com/wsl))**)
 - [ ] **[Fabric prerequisites](https://hyperledger-fabric.readthedocs.io/en/release-2.2/dev-setup/devenv.html#prerequisites)**
 - [ ] **[Docker Desktop](https://docs.docker.com/desktop/#download-and-install)**
 - [ ] **[Visual Studio Code](https://code.visualstudio.com/Download) + Extensões ⬇️**:
	 - [ ] [Docker](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker)
	 - [ ] [Remote Development](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack)
	 - [ ] [IBM Blockchain Platform](https://marketplace.visualstudio.com/items?itemName=IBMBlockchain.ibm-blockchain-platform)
	 - [ ] [Go](https://marketplace.visualstudio.com/items?itemName=golang.go)
 - [ ] Aceder ao **[ambiente Postman](https://app.getpostman.com/join-team?invite_code=4223fbd84753d939313742a6aeb2f9b3&target_code=2f99ae3d099475ee13339fd7a4448f47)** do e-Workshop


## Desenvolvimento do Chaincode

...

## Deploy on Fablo

...

## Test via Postman

...

## Test Explorer

...