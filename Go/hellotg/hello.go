package main

import (
	"fmt"

	tg "github.com/galeone/tfgo"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func main() {
	root := tg.NewRoot()
	hello := tg.Const(root, "Hello, TensorFlow!")
	helloT := tg.NewTensor(root, hello)

	results := tg.Exec(root, []tf.Output{helloT.Output}, nil, &tf.SessionOptions{})
	fmt.Println(results[0].Value())
}
