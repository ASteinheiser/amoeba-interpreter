package evaluator

import (
	"fmt"

	"github.com/ASteinheiser/amoeba-interpreter/color"
	"github.com/ASteinheiser/amoeba-interpreter/object"
)

var builtins = map[string]*object.Builtin{
	"print": {
		Fn: func(args ...object.Object) object.Object {
			fmt.Println()
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments passed to `len`: got %d, want 1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported: %s", args[0].Type())
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments passed to `first`: got %d, want 1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments passed to `last`: got %d, want 1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments passed to `rest`: got %d, want 1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments passed to `push`: got %d, want 2", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to `push` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"amoeba": {
		Fn: func(args ...object.Object) object.Object {
			color.Foreground(color.Green, false)
			fmt.Print("\n")
			fmt.Print("             ,,,,g,\n")
			fmt.Print("           #\"`    `@\n")
			fmt.Print("          @        \\b\n")
			fmt.Print("          jb    ##m @      ,smWWm\n")
			fmt.Print("           7m  ]#### '`7^\"\"      @\n")
			fmt.Print("             %n 7##b      #j@     b\n")
			fmt.Print("              @            ,,,,,,M`\n")
			fmt.Print("              @    ,w    ,M|'\n")
			fmt.Print("            ,#`   7m#`  ]b\n")
			fmt.Print("            @b         {^\n")
			fmt.Print("             %m     a#/\n")
			fmt.Print("               ^\"\"`^\n")
			fmt.Print("\n")
			color.ResetColor()

			return NULL
		},
	},
}
