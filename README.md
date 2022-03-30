# Hands On Hyperledger Fabric e-Workshop

Este e-Workshop tem como objetivo construir um Chaincode Fabric em Go, simples e abstrato, capaz de suportar a rastreabilidade de **Lotes** e **Atividades** sobre esses mesmos lotes no protocolo do Hyperledger Fabric.

Teremos a possibilidade de criar Lotes provenientes de outros Lotes através das Atividades, registos esses que serão rastráveis no ledger do Fabric. 

![Diagrama de rastreabilidade](https://github.com/lcvalves/fabric_e-workshop/blob/master/diagrams/traceability-diagram.png?raw=true)

Para suportar estas funcionalidades, iremos desenvolver 1 único smart contract capaz de representar os **Lotes** e **Atividades**, com restrições a nível de validação e verificação dos dados definidos, nomeadamente as quantidades dos lotes.

![Modelo Structs Go](https://github.com/lcvalves/fabric_e-workshop/blob/master/diagrams/go-struct-model.png?raw=true)

---

## Ambiente de desenvolvimento & Software
O software a instalar deve ser instalado no SO do ambiente de desenvolvimento:

 - [ ] Sistema operativo baseado em **Unix**:
	 - [ ] **Linux** / **macOS**, etc...
  	> ⚠️ Utilizadores Windows podem utilizar (preferencialmente) o **[WSL2](https://docs.microsoft.com/en-us/windows/wsl/install)** ou máquinas virtuais como **VirtualBox**, **VMware** ou **Hyper-V** em conjunto com uma distribuição **Linux (ex: [Ubuntu LTS](https://ubuntu.com/wsl))**

 - [ ] **[Fabric Development Environment Setup](https://hyperledger-fabric.readthedocs.io/en/release-2.2/dev-setup/devenv.html#prerequisites)**
  
 - [ ] **[Fabric Prerequisites](https://hyperledger-fabric.readthedocs.io/en/release-2.2/prereqs.html#prerequisites)**
  
 - [ ] **[Docker Desktop](https://docs.docker.com/desktop/#download-and-install)**
  
 - [ ] **[Visual Studio Code](https://code.visualstudio.com/Download) + Extensões ⬇️**:
  
	 - [ ] **[*Docker*](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker)**
  
	 - [ ] **[*Remote Development*](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack)**
  
	 - [ ] **[*IBM Blockchain Platform*](https://marketplace.visualstudio.com/items?itemName=IBMBlockchain.ibm-blockchain-platform)**
  
	 - [ ] **[*Go*](https://marketplace.visualstudio.com/items?itemName=golang.go)**
  
 - [ ] Aceder ao **[ambiente Postman](https://app.getpostman.com/join-team?invite_code=4223fbd84753d939313742a6aeb2f9b3&target_code=2f99ae3d099475ee13339fd7a4448f47)** do e-Workshop

### TODO: VERSÕES DO FABRIC, GO, NODE, etc...

---

## Desenvolvimento do Chaincode

### Microfab Setup

Para iniciar o desenvolvimento do chaincode, precisamos dum ambiente Fabric onde o possamos desenvolver. Para isso, vamos utilizar o **[Microfab](https://github.com/IBM-Blockchain/microfab)**, um runtime de Fabric "containerizado" em Docker para ambientes de desenvolvimento/teste.

Inicializamos o Microfab com o seguinte comando:

`docker run -p 8080:8080 ibmcom/ibp-microfab`

> ✔️ É esperada uma resposta no terminal de execução do tipo `Microfab started in ...ms` a confirmar que o Mircofab está a correr corretamente.

De seguida, vamos adicionar este Microfab à extensão ***IBM Blockchain Platform*** do VS Code para podermos gerar um sample Smart Contract.

O processo para adicionar o Microfab na extensão é o seguinte:

1. Na janela *Fabric Environments*, adicionar um ambiente novo;
    
2. Clicar *Add a Microfab network*;
   
3. Inserir URL pré-definido: **`http://console.127-0-01.nip.io:8080`**;
   
4. Inserir nome do ambiente (ex: **`microfab`**);
   
5. Clicar no ambiente acabado de criar;
   > ✔️ É esperada a resposta `'Connected to environment: xxxxxx'`;
   
6. Na janela *Fabric Gateways*, clicar em *Org1 Gateway* e em *Org1 Admin* para interagirmos como administrador da Org1.
   > ✔️ É esperada a resposta `'Connected via gateway: xxxxxx`';

Após estes passos, estaremos prontos a utilizar o Microfab como runtime de desenvolvimento dentro da extensão ***IBM Blockchain Platform***.

### Sample Smart Contract

Iremos proceder agora à criação da estrutura do Smart Contract base na extensão.

> ⚠️ Idealmente, seria criado um Smart Contract para cada um dos ativos (**Lotes** & **Atividades**), mas como a gestão dos Lotes depende da gestão das Atividades (ex: a criação de um lote irá depender da criação de uma atividade), iremos juntar os 2 ativos num só Smart Contract chamado **Workshop**.

Sendo assim, eis os passos para a criação dos Smart Contract:

1. 1º temos de criar uma diretoria onde vamos desenvolver o contrato. Segundo a documentação do **Go** & **Fabric**, para otimizar o proceso de desenvolvimento, devemos colocar as diretorias com aplicações Go no caminho `~/go/src/github.com/<github_username>/<nome_da_pasta>`.
   Assim sendo, vamos criar a pasta `workshop` nessa mesma diretoria;
   
2. Na janela *Smart Contracts* da extensão, clicamos nas reticências (`...`) e em *Create New Project* de seguida;
   
3. Selecionamos *Default Contract* e a linguagem *Go*;
   
4. Inserimos o nome do ativo do contrato (neste caso é **`Workshop`**);
   
5. Selecionamos a diretoria da pasta acabada de criar e a opção *Open in current window*.

Como é possível ver na diretoria do projeto, a extensão ***IBM Blockchain Platform*** criou vários ficheiros do chaincode:

- `go.mod` & `go.sum` - dependências e versionamento do projeto (em alguns casos será necessário executar o comando `go get` na diretoria do projeto para descarregar packages em falta);
  
- `main.go` - ficheiro de entrada do chaincode. Contém atribuições e definições de metadados dos contratos do chaincode;
  
- `workshop.go` - definição das estruturas (structs) de dados;
  
- `workshop-contract.go` - definição dos métodos das transações + BLL;
  
- `workshop-contract_test.go` - definição de testes unitários dos métodos definidos em `workshop-code` utilizando o package **Testify**;
  
- `transaction_data` - pasta que contém definição de argumentos/parâmetros pré-definidos das transações de `workshop-contract.go` para serem usados na extensão.

Neste sample chaincode, o modelo de dados definido em `workshop.go` apenas contém a seguinte struct `Workshop` que armazena uma string `Value`. Este atributo tem uma tag `json:"value"` para ser indexado ao campo `value` na notação JSON dos dados:

```go
// Workshop stores a value
type Workshop struct {
	Value string `json:"value"`
}
```

>⚠️ O atributo `ID` pode ser omitido.

Agora iremos empacotar o smart contract para poder ser instalado nos peers do ambiente **Microfab** que temos a correr na extensão. Eis os passos para o fazer:

1. Na janela *Smart Contracts* da extensão, clicamos nas reticências (`...`) e em *Package Open Project* de seguida;
   
2. Selecionamos o formato do package *tar.gz* pois é o formato compatível com a versão do Fabric usada nesta demo (**2.2.4**);
   
3. Introduzimos o nome do package (neste caso, `workshop`).
   >⚠️ Nomes de packages de chaincodes preferencialmente só com letras minúsculas, `'-'` e `'_'`;

4. Introduzir versão do package (a extensão incrementa **+1** à versão automaticamente, por isso colocamos `1`).

Terminando estes passos, na janela *Smart Contracts* da extensão já conseguimos ver o package acabado de criar `workshop@1(tar.gz)`.

Vamos agora instalar o package `workshop@1` no **Microfab**:

1. Na janela *Fabric Environments* da extensão, clicamos em *Deploy smart contract*;
   
2. No Step 1, selecionamos o package acabado de criar `workshop@1` e clicamos *Next*.
   A extensão permite fazer o empacotamento diretamente aqui e indica se os projetos estão ainda por empacotar;
   
3. Step 2 podemos ver a definição do contrato com o nome e versão préviamente definidos.
	Clicamos *Next*;
   
4. Por fim, no Step 3, clicamos *Deploy*.
   > ✔️ Após alguns segundos, a janela de output da extensão deve retornar uma mensagem `[SUCCESS] Successfully deployed smart contract`.

> ⚠️ É possível que hajam erros de listagem das transações nas janelas da extensão, nesse caso é recomendado atualizar as janela *Fabric Gateways* com o ícone 🔄. Caso contrário, já devemos ter listadas as transações do ficheiro `workshop-contract.go` em *Channels* > *channel1* > *workshop@1*.

Clicando no método *CreateWorkshop*, conseguimos ver 2 parâmetros no campo *Transaction arguments*;
- `param0` representa o ID omitido da struct `Workshop`;
- `param1` representa o 1º e único atributo da mesma struct - `Value`.

Sabendo isto, vamos proceder ao teste das transações na seguinte ordem:

1. No método de transação *CreateWorkshop*, introduzimos os valores `"test-id"` e `"test-value"` nos `param0` & `param1` respetivamente;
   De seguida clicamos no botão azul **Submit transaction**.
   > ✔️ O resultado esperado é `No value returned from CreateWorkshop` no campo *Transaction output* pois o método não retorna nada em caso de sucesso. No entanto, na janela de output da extensão conseguimos ver o log abaixo, o que nos indica que a transação foi executada com sucesso; 
   ```log
   [INFO] submitting transaction CreateWorkshop with args test-id,test-value on channel channel1 to peers org1peer-api.127-0-0-1.nip.io:8080
   [SUCCESS] No value returned from CreateWorkshop
	```

2. No menu dropdown *Transaction name*  da mesma aba de *Transaction View*, selecionamos o método *ReadWorkshop* e clicamos no botão cinza **Evaluate transaction**;
   > ✔️ O resultado esperado é `{"value":"test-value"}`. Caso seja, o método de leitura por `ID` está funcional com sucesso.

   > ⚠️ Para fazer leituras ao estado atual do ledger - a state database - fazemos **Evaluate transaction**, em caso de escrita fazemos **Submit transaction**.

3. Selecionamos o método *UpdateWorkshop* e alteramos o valor de `param1` para `new-value` e clicamos em **Submit transaction**;
   > ✔️ O resultado esperado é `No value returned from CreateWorkshop`.

4. Selecionamos o método *ReadWorkshop* de novo e clicamos em **Evaluate transaction** para ler o workshop;
   > ✔️ O resultado esperado agora é `{"value":"new-value"}`.

5. Selecionamos o método *DeleteWorkshop* e clicamos em **Submit transaction** para apagar o workshop da base de dados de estado atual;
   > ✔️ O resultado esperado é `No value returned from DeleteWorkshop`.

6. Por fim, para verificar que o workshop foi apagado, eelecionamos o método *WorkshopExists* e clicamos em **Evaluate transaction**.
   > ✔️ O resultado esperado é `false`.

Após estes passos, tendo em conta os resultados esperados, testamos as nossas transações definidas no contrato. Agora podemos prosseguir para o desenvolvimento do contrato com os **Lotes** & **Atividades**.


### Traceability Smart Contract



#### `workshop.go`

```go
// Lot armazena informação relativa aos lotes da cadeia de valor
type Lot struct {
	DocType string  `json:"docType"` // docType ("lot") tem de ser usado para se distinguir de outros documentos da state database
	ID      string  `json:"ID"`
	Product string  `json:"product"`
	Amount  float32 `json:"amount"`
	Unit    string  `json:"unit"`
	Owner   string  `json:"owner"`
}

// Activity armazena informação sobre as atividades da cadeia de valor
type Activity struct {
	DocType   string             `json:"docType"` // docType ("act") tem de ser usado para se distinguir de outros documentos da state database
	ID        string             `json:"ID"`
	InputLots map[string]float32 `json:"inputLots,omitempty" metadata:",optional"` // inputLots é opcional porque podemos ter atividades que apenas registam lotes que vêm de fora da cadeira de valor
	OutputLot Lot                `json:"outputLot"`
	Date      string             `json:"date"`
	Issuer    string             `json:"issuer"`
}
```

#### `workshop-contract.go`

##### CreateLot

```go
// Lot armazena informação relativa aos lotes da cadeia de valor
type Lot struct {
	DocType string  `json:"docType"` // docType ("lot") tem de ser usado para se distinguir de outros documentos da state database
	ID      string  `json:"ID"`
	Product string  `json:"product"`
	Amount  float32 `json:"amount"`
	Unit    string  `json:"unit"`
	Owner   string  `json:"owner"`
}

// Activity armazena informação sobre as atividades da cadeia de valor
type Activity struct {
	DocType   string             `json:"docType"` // docType ("act") tem de ser usado para se distinguir de outros documentos da state database
	ID        string             `json:"ID"`
	InputLots map[string]float32 `json:"inputLots,omitempty" metadata:",optional"` // inputLots é opcional porque podemos ter atividades que apenas registam lotes que vêm de fora da cadeira de valor
	OutputLot Lot                `json:"outputLot"`
	Date      string             `json:"date"`
	Issuer    string             `json:"issuer"`
}
```


## Deploy on Fablo

```bash
curl -Lf https://github.com/hyperledger-labs/fablo/releases/download/1.0.0/fablo.sh -o ./fablo && chmod +x ./fablo
```


## Test Explorer

...