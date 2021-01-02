package file

import (
	"fmt"
)

func test() {

	fs := Default()
	roots := fs.Roots()
	for idx := range roots {
		root := roots[idx]
		fmt.Println(root.Path())
	}

}
