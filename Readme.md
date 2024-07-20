# Olist challenge

O Sistema de Gerenciamento de Biblioteca foi desenvolvido para gerenciar um conjunto abrangente de funcionalidades relacionadas a livros e autores. Este sistema oferece uma API RESTful HTTP que permite aos usuários interagir com os dados da biblioteca, incluindo a importação de informações de autores, gerenciamento de registros de livros e acesso a listas de autores paginadas.

## Inicío

1. Clone o reposotório

```
  git clone https://github.com/LucasBiazon/Olist_challange.git
```

2. Entrar no arquivo

```
  cd Olist_challange
```

3. Build cmd

```
  go build -o ./api/cmd/CSVCmd.go
```

4. Mover arquivo build

```
  mv gocsv /usr/local/bin
```

5. Rodar projeto

```
  go run main.go
```

## Create authors

```
  gocsv -csv "./data.csv"
```

## Routes:

https://documenter.getpostman.com/view/33427106/2sA3kUG2d5
