package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const usage string = `Usage: groupbycolumn path-to-file.csv`

func main() {

	args := os.Args

	if len(args) != 2 {
		fmt.Println(usage)
		return
	}

	pth, err := filepath.Abs(args[1])
	if err != nil {
		fmt.Println(err)
		fmt.Println(usage)
		return
	}

	f, err := os.Open(pth)
	if err != nil {
		fmt.Println("error opening csv file")
		fmt.Println(err)
		fmt.Println(usage)
		return
	}

	rd := csv.NewReader(f)

	rec, err := rd.ReadAll()
	if err != nil {
		fmt.Println("error parsing csv file")
		fmt.Println(err)
		fmt.Println(usage)
		return
	}

	if len(rec) == 0 {
		fmt.Println("no records")
		return
	}

	ncols := len(rec[0])
	unique := make([]map[string]bool, ncols)

	for i := 0; i < ncols; i++ {
		unique[i] = map[string]bool{}
	}

	baseNode := NewTreeNode()
	var nd *TreeNode

	for _, rw := range rec {
		nd = baseNode

		for _, cell := range rw {
			nd = nd.NavTo(cell)
		}
	}

	baseNode.PrintGroups(0)
}

type TreeNode struct {
	Value    string
	SubNodes map[string]*TreeNode
}

func NewTreeNode() *TreeNode {
	return &TreeNode{
		SubNodes: map[string]*TreeNode{},
	}
}

func (t *TreeNode) NavTo(value string) *TreeNode {
	if v, ok := t.SubNodes[value]; ok {
		return v
	}

	nv := NewTreeNode()
	nv.Value = value
	t.SubNodes[value] = nv
	return nv
}

func (t *TreeNode) PrintGroups(indentLevel int) {

	fmt.Println(strings.Repeat("\t", indentLevel) + t.Value)
	for _, v := range t.SubNodes {
		v.PrintGroups(indentLevel + 1)
	}
}
