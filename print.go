package bencode

import "fmt"

// Print B
func (b *B) Print() {
	b.print(0, 0)
}

func (b *B) print(pos, ident int) {
	switch (*(b.b))[pos].typ {
	case INT:
		fmt.Print((*(b.b))[pos].n, " ")
	case STRING:
		// if len((*(b.b))[pos].keys) != 0 && len((*(b.b))[pos].keys[0]) < 100 {
		// 	fmt.Print((*(b.b))[pos].keys[0], " ")
		// }
		if len((*(b.b))[pos].str) < 100 {
			fmt.Print((*(b.b))[pos].str, " ")
		}
	case LIST:
		for k := range (*(b.b))[pos].list {
			for i := 0; i < ident; i++ {
				fmt.Print(" ")
			}
			b.print((*(b.b))[pos].list[k], ident+2)
			fmt.Print("\n")
		}
	case DICT:
		for k := range (*(b.b))[pos].list {
			for i := 0; i < ident; i++ {
				fmt.Print(" ")
			}
			fmt.Print((*(b.b))[pos].keys[k], ": ")
			b.print((*(b.b))[pos].list[k], ident+2)
			fmt.Print("\n")
		}
	}
}

// func (bs *bstruct) print(ident int) {
// 	switch bs.typ {
// 	case INT:
// 		fmt.Print(bs.n, " ")
// 	case STRING:
// 		fmt.Print(bs.str, " ")
// 	case LIST:
// 		for k := range bs.list {
// 			for i := 0; i < ident; i++ {
// 				fmt.Print(" ")
// 			}
// 			bs.list[k].value.print(ident + 2)
// 			fmt.Print("\n")
// 		}
// 	case DICT:
// 		for k := range bs.list {
// 			for i := 0; i < ident; i++ {
// 				fmt.Print(" ")
// 			}
// 			fmt.Print(*bs.list[k].key, ":")
// 			bs.list[k].value.print(ident + 2)
// 			fmt.Print("\n")
// 		}
// 	}
// }
