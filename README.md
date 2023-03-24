# api desenvolvida para a disciplina de web avançada

Api-Rest de uma aplicação CRUD em GO estruturada em camadas MVC.
Persistência de dados realizada no MongoDB.
Aplicação Dockerizada


Para estruturar o projeto, deve-se iniciar o go.mod e go.sum.

- go.mod: 
    cria o modulo para exportar o projeto como package, apontando para a versão do go.

- go.sum:
    lista os packages dependencies do projeto

No arquivo .env se encontra algumas variáveis constantes para gerenciamento do projeto, adicionando
ele no .gitignore, não será enviada suas configurações privadas do projeto.


- Na pasta configuration/handling_err se encontra as configurações de padronização de erros


- Utilizado o gin-gonic para handle das rotas da aplicação

- Na pasta routes foi criada alguns endpoints e realizado um bind ao respectivo controller

- No arquivo main é necessário instanciar um router gin e iniciar nossas rotas, assim ao inicializar o projeto, já é possível visualizar os endpoints

- É necessário bindar uma porta livre para o gin, caso contrário, a aplicação não será inicializada

