package main

import (
	"../pbforms"
	"../forms"
	"os"
)

func main() {
	f := forms.MyForm{}
	fw := pbforms.NewFormWriter(f)
	fw.Write(os.Stdout)
}
