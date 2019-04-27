package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	flag.Parse()

	args := flag.Args()

	for _, arg := range args {
		files, err := filepath.Glob(arg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
		for _, f := range files {
			Generate(f)
		}
	}
}

// Generate generates docs for aliases
func Generate(fileName string) {
	fset := token.NewFileSet()
	// Parse the file given in arguments
	node, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		fmt.Printf("Error parsing file %v", err)
		return
	}

	// need this because Inspect will skip comments in the room node (eg: the generate comment)
	origignalComments := make(map[token.Pos]*ast.CommentGroup, len(node.Comments))
	for _, c := range node.Comments {
		origignalComments[c.Pos()] = c
	}

	comments := make([]*ast.CommentGroup, 0)

	ast.Inspect(node, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.CommentGroup:
			delete(origignalComments, t.Pos())
			comments = append(comments, t)

		case *ast.ValueSpec:
			if len(t.Values) != len(t.Names) {
				return true
			}

			cg := getCommentGroup(t.Names, t.Values, t.Pos())
			if cg != nil {
				if t.Doc != nil {
					for _, l := range t.Doc.List {
						delete(origignalComments, l.Slash)
					}
				}
				t.Doc = cg
			}

		case *ast.TypeSpec:
			cg := getCommentGroup([]*ast.Ident{t.Name}, []ast.Expr{t.Type}, t.Pos())
			if cg != nil {
				if t.Doc != nil {
					for _, l := range t.Doc.List {
						delete(origignalComments, l.Slash)
					}
				}
				t.Doc = cg
			}
		default:
		}

		return true
	})
	for _, c := range origignalComments {
		comments = append(comments, c)
		for _, l := range c.List {
			fmt.Printf("%+#v %v\n", l.Text, l.Slash)
		}
		sort.Slice(comments, func(i, j int) bool {
			return comments[i].Pos() < comments[j].Pos()
		})
	}

	node.Comments = comments

	// overwrite the file with modified version of ast.
	write, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error opening file %v", err)
		return
	}
	defer func() {
		err = write.Close()
		if err != nil {
			fmt.Print(err)
		}
	}()
	w := bufio.NewWriter(write)
	err = format.Node(w, fset, node)
	if err != nil {
		fmt.Printf("Error formating file %s", err)
		return
	}
	err = w.Flush()
	if err != nil {
		fmt.Printf("Error writing file %s", err)
		return
	}
}

func getCommentGroup(names []*ast.Ident, values []ast.Expr, pos token.Pos) *ast.CommentGroup {
	if len(names) != len(values) {
		return nil
	}

	cg := &ast.CommentGroup{
		List: make([]*ast.Comment, 0),
	}
	for i, name := range names {
		v, ok := values[i].(*ast.SelectorExpr)
		if !ok {
			continue
		}

		comment := &ast.Comment{
			Slash: pos - 1,
			// TODO: collect docs from aliased package and use them here
			Text: fmt.Sprintf("// %s is an alias for %s.%s", name, v.X, v.Sel.Name),
		}
		cg.List = append(cg.List, comment)
	}

	if len(cg.List) == 0 {
		return nil
	}

	return cg
}
