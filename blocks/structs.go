package blocks

import (
	"errors"
	"fmt"
	"log"
)


type CodeInterface interface {
	Code(scopes ...ScopeInterface) string
	GetName() string
	GetType() string
}

func Operate(operation string, Vars ...*Variable) *String{
	vals := make([]interface{}, len(Vars))

	for i, v := range Vars {
		vals[i] = v.Value.Code()
	}

	return NewString(fmt.Sprintf(operation, vals...)).NoQuotes()
}

type ScopeInterface interface {
	CodeInterface
	getBlocks() []CodeInterface
}


type Parameter struct {
	Type  string
	Name  string
}

func NewParameter() *Parameter{
	return &Parameter{}
}

func (P *Parameter) SetType(t string) *Parameter {
	var err error

	if t == "" {
		err = errors.New("SetType(): Empty type argument")
		log.Println(err)
	} else {
		P.Type = t
	}

	return P
}

func (P *Parameter) GetType() string {
	return P.Type
}

func (P *Parameter) SetName(n string) *Parameter {
	var err error

	if n == "" {
		err = errors.New("SetName(): Empty name argument")
		log.Println(err)
	} else {
		P.Name = n
	}

	return P
}

func (P *Parameter) GetName() string {
	return P.Name
}


type StructParameter struct {
	Parameter
	Json string
}

func (SP *StructParameter) Code(scopes ...ScopeInterface) string {
	jsonField := ""

	if SP.Json != "" {
		jsonField = "`json: \"" + SP.Json + "\"`"
	}
	return SP.Name + " " + SP.Type + " " + jsonField
}



type CodeBlock struct {
	Parameter
	Value CodeInterface
	Comment string

}


func NewCodeBlock() *CodeBlock{
	return &CodeBlock{}
}


func (C *CodeBlock) SetType(t string) *CodeBlock {
	C.Type = t
	return C
}


func (C *CodeBlock) SetName(n string) *CodeBlock {
	C.Name = n
	return C
}



func (C *CodeBlock) GetValue() CodeInterface{
	return C.Value
}


func (C *CodeBlock) SetValue(val CodeInterface) *CodeBlock {
	C.Value = val
	return C
}

func (C *CodeBlock) Code(scopes ...ScopeInterface) string{
	return C.Value.Code(scopes...)
}


type Variable struct {
	CodeBlock
}


func NewVariable() *Variable{
	return &Variable{}
}

var funcVariable = "%s := %s"

var globalVariable = "var %s %s"

var existsVariable = "%s = %s"


func (V *Variable) SetType(t string) *Variable {
	V.Type = t
	return V
}


func (V *Variable) SetName(n string) *Variable {
	V.Name = n
	return V
}

func (V *Variable) SetValue(val CodeInterface) *Variable {
	V.Value = val
	return V
}


func (V *Variable) Code(scopes ...ScopeInterface) string {

	funcScope := false

	alreadyExists := false

	code := ""

	if V.GetType() == "" && V.GetValue() == nil {
		return ""
	}

	scope := getScope(scopes, -1)

	if scope != nil {
		switch scope.(type) {

		case *Function:

			funcScope = true
		}
	}

	if funcScope {
		v := findInScope(scope.getBlocks(), "name", V.GetName())

		if v != nil && V != v {
			alreadyExists = true
		}

		if alreadyExists {
			code = fmt.Sprintf(existsVariable, V.GetName(), V.GetValue().Code())
		} else {
			code = fmt.Sprintf(funcVariable, V.GetName(), V.GetValue().Code())
		}


	} else {

		val := ""

		if V.GetValue() != nil {
			val = "= " + V.GetValue().Code()
		} else {
			val = V.GetType()
		}

		if val == "" || V.GetName() == "" {return val}

		code = fmt.Sprintf(globalVariable, V.GetName(), val)
	}


	return code
}


func getScope(scopes []ScopeInterface, index int) ScopeInterface{
	l := len(scopes)

	if l == 0 || index > l - 1 {
		return nil
	}

	if index == -1 {
		index = l - 1
	}

	return scopes[index]
}

func findInScope(codeBlocks []CodeInterface, field, value string) CodeInterface {

	if field == "" { field = "name" }
	if value == "" {return nil}

	for _, block := range codeBlocks {
		if field == "name" && block.GetName() == value {return block}
		if field == "type" && block.GetType() == value {return block}
	}

    return nil
}


type Struct struct {
	CodeBlock
}

func NewStruct() *Struct{
	return &Struct{}
}

type Interface struct {
	CodeBlock
}

func NewInterface() *Interface{
	return &Interface{}
}

type Function struct {
	CodeBlock
	Skeleton
	Input       []*Variable
	Output      []*Variable
	Goroutine   bool
	Execute     bool
	MethodOf    *Struct
}

func NewFunction() *Function{
	return &Function{}
}


func (F *Function) GetName() string {
	return F.Name
}

func (F *Function) SetType(t string) *Function {
	F.Type = t
	return F
}


func (F *Function) SetName(n string) *Function {
	F.Name = n
	return F
}

func (F *Function) SetValue(val CodeInterface) *Function {
	F.Value = val
	return F
}

func (F *Function) SetInput(params []*Variable) *Function {
	F.Input = params
	return F
}

func (F *Function) AddInput(params ...*Variable) *Function {
	F.Input = append(F.Input, params...)
	return F
}

func (F *Function) SetOutput(output []*Variable) *Function {
	F.Output = output
	return F
}

func (F *Function) SetBlock(blocks []CodeInterface) *Function {
	F.Blocks = blocks
	return F
}

func (F *Function) AppendBlock(block CodeInterface) *Function {
	F.Blocks = append(F.Blocks, block)
	return F
}

func (F *Function) Executing() *Function{
	FE := *F
	FE.Execute = true
    return &FE
}

func (F *Function) isMethodOf() *Struct {
	return F.MethodOf
}

func (F *Function) isExecuting() bool {
	return F.Execute
}

func (F *Function) isGoroutine() bool {
	return F.Goroutine
}


var functionHeader = "%sfunc %s%s%s %s {"
var functionFooter = "}"
var functionExecute = "%s%s"


func (F *Function) Code(scopes ...ScopeInterface) string {

	Parent := F.isMethodOf()

	pName := ""

	goroutine := ""

	scopes = append(scopes, F)


	if Parent != nil {
		pName = "(" + string(Parent.GetName()[0]) + Parent.GetName() + ") "
	}

	if F.isGoroutine() {
		goroutine = "go "
	}

	output := paramsToCode(F.Output, false, true)


	if output != "" {
		output = "(" + output + ")"
	}

	if F.isExecuting() {

		return fmt.Sprintf(functionExecute, F.GetName(), "("+paramsToCode(F.Input, true, false)+")")

	} else {
		header := fmt.Sprintf(functionHeader, goroutine, pName, F.GetName(), "("+paramsToCode(F.Input, true, true)+")",
			output)

		returning := paramsToCode(F.Output, true, false)

		if returning != "" {returning = "return " + returning + "\n"}

		return header + "\n" + F.Skeleton.Code(scopes...) + returning + functionFooter
	}

}

func paramsToCode(Params []*Variable, withNames bool, withTypes bool) string{

	code := ""

	for i, Param := range Params {

		if i > 0 {
			code += ", "
		}

		if withNames {
			if Param.GetName() == "" {
				code += Param.GetValue().Code() + " "
			} else {
				code += Param.GetName() + " "
			}
		}
		if withTypes {
			code += Param.GetType()
		}
	}

	return code
}


type Skeleton struct {
	Blocks    []CodeInterface
}

func NewSkeleton() *Skeleton{
	return &Skeleton{}
}

func (S Skeleton) GetName() string{
	return "Skeleton"
}

func (S Skeleton) GetType() string{
	return "Skeleton"
}

func (S *Skeleton) getBlocks() []CodeInterface{
	return S.Blocks
}

func (S *Skeleton) AppendBlocks(blocks ...CodeInterface) *Skeleton {
	S.Blocks = append(S.Blocks, blocks...)
	return S
}

func (S *Skeleton) RemoveBlocks() *Skeleton {
	S.Blocks = []CodeInterface{}
	return S
}


func (S *Skeleton) Code(scopes ...ScopeInterface) string {

	code := ""

	for _, block := range S.getBlocks() {
		code += block.Code(scopes...) + "\n"
	}

	return code
}


type Import struct {
	Alias  string
	Source string
}

func NewImport() *Import {
	return &Import{}
}

func (I *Import) SetSource(source string) *Import{
	I.Source = source
	return I
}

func (I *Import) SetAlias(alias string) *Import{
	I.Alias = alias
	return I
}

func (I *Import) Code(scopes ...ScopeInterface) string {
	return I.Alias + " \"" + I.Source + "\""
}


type Heading struct {
	Package  string
	Imports  []*Import
}


func NewHeading() *Heading {
	return &Heading{}
}


func (H *Heading) SetPackageName(name string) *Heading{
	H.Package = name
	return H
}


func (H *Heading) AppendImports(imports ...*Import) *Heading{
	H.Imports = append(H.Imports, imports...)
	return H
}

func (H *Heading) Code(scopes ...ScopeInterface) string {

	p := "package " + H.Package

	i := "import (\n"

	for _, imp := range H.Imports {
		i += imp.Code() + "\n"
	}

	i += ")"

    return p + "\n\n" + i
}


type File struct {
	*Heading
	*Skeleton
	Name string
	Path string
}


func NewFile() *File {
	return &File{}
}

func (F *File) SetFileName(name string) *File{
	F.Name = name
	return F
}

func (F *File) GetFileName() string{
	return F.Name
}

func (F *File) SetFilePath(path string) *File{
	F.Path = path
	return F
}

func (F *File) GetFilePath() string{
	return F.Path
}


func (F *File) SetHeading(h *Heading) *File{
	F.Heading = h
	return F
}

func (F *File) SetSkeleton(s *Skeleton) *File{
	F.Skeleton = s
	return F
}

func (F *File) Code(scopes ...ScopeInterface) string {

	code := F.Heading.Code() + "\n"

	for _, block := range F.getBlocks() {
		code += block.Code(scopes...) + "\n"
	}

	return code
}
