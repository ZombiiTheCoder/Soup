package ast

import "strconv"

func Create_Program (body []Stmt) Stmt {
	return Program{ Body: body }
}

func Create_VarDec (cont bool, idnt string, valu any) Stmt {
	return VarDec{ Cont: cont, Idnt: idnt, Valu: valu }
}

func Create_Import_Stmt (file string, rel bool) Stmt {
	return ImpStmt{ File: file, Rel:rel }
}

func Create_AssignExpr (assigner, valu Expr) Expr {
	return AssignExpr{ Assigner: assigner, Valu: valu }
}

func Create_Property (key string, val any) Expr {
	return Property{ Key: key, Val: val }
}

func Create_ObjectLiteral (properties []Property) Expr {
	return ObjectLiteral{ Properties: properties }
}

func Create_MemberExpr (obj, property Expr, computed bool) Expr {
	return MemberExpr{ Obj: obj, Property: property, Computed: computed}
}

func Create_BinaryExpr (left, right Expr, op string) Expr {
	return BinaryExpr{ Left: left, Right: right, Op: op }
}

func Create_UnaryExpr (left Expr, op string, prefix bool) Expr {
	return UnaryExpr{ Left: left, Op: op, Prefix: prefix }
}

func Create_CallExpr (caller Expr, args []Expr) Expr {
	return CallExpr{ Caller: caller, Args:args }
}

func Create_Function (name string, params []string, body []Stmt) Stmt {
	return FuncDec{ Name:name, Params:params, Body:body }
}

func Create_Ret_Stmt (val Expr) Stmt {
	return RetStmt{ Valu:val }
}

func Create_Identifier (symb string) Expr {
	return Identifier{ Symb: symb }
}

func Create_StringLiteral (valu string) Expr {
	return StringLiteral{ Valu: valu }
}

func Create_NumericLiteral (valu string) Expr {
	valu2, _:=strconv.ParseInt(valu, 10, 64)
	return NumericLiteral{ Valu: valu2 }
}

func Create_FloatLiteral (valu string) Expr {
	valu2, _:=strconv.ParseFloat(valu, 64)
	return FloatLiteral{ Valu: valu2 }
}

func Create_NullLiteral () Expr {
	return NullLiteral{ Valu: nul }
}