package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	todo "github/avner.oliveira/todoApp"
	"io"
	"os"
	"strings"
)

const (
	todoFile = ".todo.json"
)

func main() {
	new := flag.Bool("new", false, "Adicionar nova tarefa")
	complete := flag.Int("complete", 0, "Completar tarefa")
	delete := flag.Int("delete", 0, "Excluir tarefa")
	list := flag.Bool("list", false, "Listar todas as tarefas")

	flag.Parse()
	todo := &todo.Todos{}

	if err := todo.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *new:
		task, err := getInput(os.Stdin, flag.Args()...)
		checkError(err)

		todo.Add(task)
		err = todo.Store(todoFile)
		checkError(err)

	case *complete != 0:
		err := todo.Complete(*complete)
		checkError(err)

		err = todo.Store(todoFile)
		checkError(err)

	case *delete != 0:
		err := todo.Delete(*delete)
		checkError(err)
		err = todo.Store(todoFile)
		checkError(err)
	case *list:
		todo.Print()

	default:
		fmt.Fprintln(os.Stdout, "Comando invalido")
		os.Exit(0)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	descricao := scanner.Text()
	if len(descricao) == 0 {
		return "", errors.New("Tarefa sem nome não aceita")
	}
	return descricao, nil
}
