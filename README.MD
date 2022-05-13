## README
Trabalho para a materia de serverless da Fiap com o professor Rafael,
o trabalho consiste em criar uma lambda via serverless e de bonus 
uma persistencia.

Foi realizado a criação de uma lambda get_books_handler cujo realiza uma busca
básica dentro do dynamoDB na tabela de books pelo titulo, a criação da tabela
foi feita via serverless, o serverless foi separado em 3 arquivos o principal 
e outros 2 onde um possui os recursos(dynamo) e o outro de funções(lambda)

Para a execução do trabalho é necessario configurar o aws cli na maquina e o serverless framework

Os comandos de configuração e build/deploy foram feitos com o MakeFile para facilitar o trabalho

A execução final do trabalho não pode ser realizada, pois o nosso ambiente dentro da academy
tem recursos limitados impossibilitando a execução (por causa da role que possuimos).

## Setup:

```sh
make setup
```

## Build and deploy:

```sh
make build-and-deploy
```

