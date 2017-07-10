package txvm

//go:generate sh gen.sh

var OpNames = [...]string{
	Fail:          "fail",
	PC:            "pc",
	JumpIf:        "jumpif",
	Roll:          "roll",
	Bury:          "bury",
	Reverse:       "reverse",
	Depth:         "depth",
	ID:            "id",
	Len:           "len",
	Drop:          "drop",
	Dup:           "dup",
	ToAlt:         "toalt",
	FromAlt:       "fromalt",
	Inspect:       "inspect",
	Equal:         "equal",
	Not:           "not",
	And:           "and",
	Or:            "or",
	Add:           "add",
	Sub:           "sub",
	Mul:           "mul",
	Div:           "div",
	Mod:           "mod",
	Lshift:        "lshift",
	Rshift:        "rshift",
	GT:            "gt",
	Cat:           "cat",
	Slice:         "slice",
	BitNot:        "bitnot",
	BitAnd:        "bitand",
	BitOr:         "bitor",
	BitXor:        "bitxor",
	SHA256:        "sha256",
	SHA3:          "sha3",
	CheckSig:      "checksig",
	CheckMultiSig: "checkmultisig",
	PointAdd:      "pointadd",
	PointSub:      "pointsub",
	PointMul:      "pointmul",
	Encode:        "encode",
	Varint:        "varint",
	MakeTuple:     "maketuple",
	Untuple:       "untuple",
	Field:         "field",
	Type:          "type",
	Annotate:      "annotate",
	Defer:         "defer",
	Satisfy:       "satisfy",
	Unlock:        "unlock",
	UnlockOutput:  "unlockoutput",
	Merge:         "merge",
	Split:         "split",
	Lock:          "lock",
	Retire:        "retire",
	Nonce:         "nonce",
	Reanchor:      "reanchor",
	Issue:         "issue",
	IssueCA:       "issueca",
	Before:        "before",
	After:         "after",
	Summarize:     "summarize",
	Migrate:       "migrate",
	ProveRange:    "proverange",
	ProveValue:    "provevalue",
	ProveAsset:    "proveasset",
	Blind:         "blind",
	Nop0:          "nop0",
	Nop1:          "nop1",
	Nop2:          "nop2",
	Nop3:          "nop3",
	Nop4:          "nop4",
	Nop5:          "nop5",
	Nop6:          "nop6",
	Nop7:          "nop7",
	Nop8:          "nop8",
	Private:       "private",
}

var OpCodes = map[string]byte{
	"fail":          Fail,
	"pc":            PC,
	"jumpif":        JumpIf,
	"roll":          Roll,
	"bury":          Bury,
	"reverse":       Reverse,
	"depth":         Depth,
	"id":            ID,
	"len":           Len,
	"drop":          Drop,
	"dup":           Dup,
	"toalt":         ToAlt,
	"fromalt":       FromAlt,
	"inspect":       Inspect,
	"equal":         Equal,
	"not":           Not,
	"and":           And,
	"or":            Or,
	"add":           Add,
	"sub":           Sub,
	"mul":           Mul,
	"div":           Div,
	"mod":           Mod,
	"lshift":        Lshift,
	"rshift":        Rshift,
	"gt":            GT,
	"cat":           Cat,
	"slice":         Slice,
	"bitnot":        BitNot,
	"bitand":        BitAnd,
	"bitor":         BitOr,
	"bitxor":        BitXor,
	"sha256":        SHA256,
	"sha3":          SHA3,
	"checksig":      CheckSig,
	"checkmultisig": CheckMultiSig,
	"pointadd":      PointAdd,
	"pointsub":      PointSub,
	"pointmul":      PointMul,
	"encode":        Encode,
	"varint":        Varint,
	"maketuple":     MakeTuple,
	"untuple":       Untuple,
	"field":         Field,
	"type":          Type,
	"annotate":      Annotate,
	"defer":         Defer,
	"satisfy":       Satisfy,
	"unlock":        Unlock,
	"unlockoutput":  UnlockOutput,
	"merge":         Merge,
	"split":         Split,
	"lock":          Lock,
	"retire":        Retire,
	"nonce":         Nonce,
	"reanchor":      Reanchor,
	"issue":         Issue,
	"issueca":       IssueCA,
	"before":        Before,
	"after":         After,
	"summarize":     Summarize,
	"migrate":       Migrate,
	"proverange":    ProveRange,
	"provevalue":    ProveValue,
	"proveasset":    ProveAsset,
	"blind":         Blind,
	"nop0":          Nop0,
	"nop1":          Nop1,
	"nop2":          Nop2,
	"nop3":          Nop3,
	"nop4":          Nop4,
	"nop5":          Nop5,
	"nop6":          Nop6,
	"nop7":          Nop7,
	"nop8":          Nop8,
	"private":       Private,
}
