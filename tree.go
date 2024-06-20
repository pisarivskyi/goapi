package goapi

import (
	"fmt"
	"slices"
	"strings"
)

type Tree struct {
	root *Node
}

type Node struct {
	configs  map[string]*RouteConfig
	pattern  string
	children []*Node
}

func (tr *Tree) AddNode(path string, config *RouteConfig) {
	nd := tr.root
	pathSplit := splitPath(path)
	// keep track of tree depth
	depth := 0

	for index, pathPart := range pathSplit {
		// if nil consider it as root node and create new one
		if nd == nil {
			newNode := &Node{
				configs:  map[string]*RouteConfig{},
				pattern:  pathPart,
				children: []*Node{},
			}

			tr.root = newNode
			nd = newNode
		} else {
			// node exists, skip
			if nd.pattern == pathPart && depth == index {
				continue
			}

			if len(nd.children) > 0 {
				existingIndex := slices.IndexFunc(nd.children, func(child *Node) bool {
					return child.pattern == pathPart
				})

				if existingIndex != -1 {
					nd = nd.children[existingIndex]
					depth++
					continue
				}
			}

			// create new node for 'pathPart'
			newNode := &Node{
				configs:  map[string]*RouteConfig{},
				pattern:  pathPart,
				children: []*Node{},
			}

			nd.children = append(nd.children, newNode)

			nd = newNode
			depth++
		}

		if index == len(pathSplit)-1 {
			nd.configs[config.Method] = config
		}
	}
}

func (tr *Tree) FindByPath(path string) (*Node, error) {
	node := tr.root
	pathSplit := splitPath(path)
	length := len(pathSplit)
	index := 0

	return walk(node, pathSplit, index, length)
}

func walk(node *Node, pathSplit []string, index int, length int) (*Node, error) {
	if index == length-1 && node.pattern == pathSplit[index] {
		return node, nil
	}

	for _, child := range node.children {
		res, err := walk(child, pathSplit, index+1, length)

		if err == nil {
			return res, nil
		}
	}

	return nil, fmt.Errorf("not found")
}

func splitPath(path string) []string {
	if path == "/" {
		return []string{"/"}
	}

	sp := strings.Split(path, "/")

	sp = append([]string{"/"}, sp[1:]...)

	return sp
}
