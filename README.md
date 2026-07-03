## Instalação github

git init
git add .
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/marcosbras/...
git push -u origin main

## Branchs
O git branch serve para criar, listar ou excluir ramificações (branches). Ele cria o ponteiro, mas você continua na mesma branch em que estava.

O git checkout serve para navegar entre branches ou restaurar arquivos. Ele altera os arquivos do seu diretório para a versão da branch escolhida.

Para entender melhor:

git branch nome-da-branch: Apenas cria a nova branch.

git checkout nome-da-branch: Muda você de uma branch para outra já existente.

git checkout -b nome-da-branch: Cria a branch e já pula (faz o checkout) para ela no mesmo comando.


## Instalação golang
https://go.dev/dl/

## Tutoriais
golang: 
https://www.youtube.com/watch?v=E-VNDPIVhs4

## Comandos go
iniciar projeto: go mod init <nome-do-projeto>
Rodar programa: go run main.go
Build programa: go build 


Gerar binário em outro sistema operacional:
GOOS=windows go build

## Particularidade do go
Qualquer variável que eu declare eu sou obrigado a usa-la
Funções podem retornar mais que um valor



