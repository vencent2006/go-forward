package orm

// 衍生类型
type op string

const (
	opEq  op = "="
	opNot op = "NOT"
	opAnd op = "AND"
	opOr  op = "OR"
)

func (o op) String() string {
	return string(o)
}

// 这种叫别名
//type op = string

// 谓词(Predicate) 中文领域一般喜欢叫条件(condition)
type Predicate struct {
	left  Expression
	op    op
	right Expression
}

type Column struct {
	name string
}

func C(name string) Column {
	return Column{name: name}
}

func (c Column) expr() {
}

// usage: sub.C("id").Eq(12), 保持链式调用的优雅
func (c Column) Eq(arg any) Predicate {
	return Predicate{
		left:  c,
		op:    opEq,
		right: value{val: arg},
	}
}

// Not(C("name").Eq("Tom"))
func Not(p Predicate) Predicate {
	return Predicate{
		op:    opNot,
		right: p,
	}
}

// C("id").Eq(12).And(C("name").Eq("Tom"))
func (left Predicate) And(right Predicate) Predicate {
	return Predicate{
		left:  left,
		op:    opAnd,
		right: right,
	}
}

// C("id").Eq(12).Or(C("name").Eq("Tom"))
func (left Predicate) Or(right Predicate) Predicate {
	return Predicate{
		left:  left,
		op:    opOr,
		right: right,
	}
}

func (left Predicate) expr() {
}

type value struct {
	val any
}

func (v value) expr() {
}

// Expression 是一个标记接口，代表表达式
type Expression interface {
	expr()
}
