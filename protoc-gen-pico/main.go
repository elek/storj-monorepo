// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"flag"
	"fmt"
	"runtime/debug"
	"sort"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/pluginpb"
)

//go:generate protoc -I.. --go_out=paths=source_relative:. pico.proto

const (
	picobufPackage     = protogen.GoImportPath("storj.io/picobuf")
	timestamppbPackage = protogen.GoImportPath("google.golang.org/protobuf/types/known/timestamppb")
)

type config struct {
	suffix string
}

func main() {
	var flags flag.FlagSet
	var conf config
	flags.StringVar(&conf.suffix, "suffix", "", "Suffix to apply to all package level identifiers")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(plugin *protogen.Plugin) error {
		for _, f := range plugin.Files {
			if f.Generate {
				genFile(plugin, f, conf)
			}
		}
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		return nil
	})
}

type generator struct {
	*protogen.GeneratedFile
	*protogen.File
	suffix string
}

func genFile(plugin *protogen.Plugin, file *protogen.File, conf config) {
	gf := &generator{
		GeneratedFile: plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+".pico.go", file.GoImportPath),
		File:          file,
		suffix:        conf.suffix,
	}

	gf.P("// Code generated by protoc-gen-pico. DO NOT EDIT.")
	gf.P("// source: ", file.Desc.Path())
	gf.P("//")
	gf.P("// versions:")
	{
		picoVersion := "(unknown)"
		if bi, ok := debug.ReadBuildInfo(); ok {
			picoVersion = bi.Main.Version
		}
		gf.P("//     protoc-gen-pico: ", picoVersion)
	}
	{
		protocVersion := "(unknown)"
		if v := plugin.Request.GetCompilerVersion(); v != nil {
			protocVersion = fmt.Sprintf("v%v.%v.%v", v.GetMajor(), v.GetMinor(), v.GetPatch())
			if s := v.GetSuffix(); s != "" {
				protocVersion += "-" + s
			}
		}
		gf.P("//     protoc:          ", protocVersion)
	}
	gf.P()
	gf.P("package ", file.GoPackageName)
	gf.P()

	genWalk(gf, file, file.Enums, file.Messages)
}

func genWalk(gf *generator, file *protogen.File, enums []*protogen.Enum, msgs []*protogen.Message) {
	for _, e := range enums {
		genEnum(gf, e)
	}
	for _, m := range msgs {
		genMessage(gf, m)
	}
	for _, m := range msgs {
		genWalk(gf, file, m.Enums, m.Messages)
	}
}

func genEnum(gf *generator, e *protogen.Enum) {
	gf.P("type ", e.GoIdent, gf.suffix, " int32")

	gf.P("const (")
	for _, v := range e.Values {
		gf.P(v.GoIdent, gf.suffix, " ", e.GoIdent, gf.suffix, " = ", v.Desc.Number())
	}
	gf.P(")")

	gf.P()
}

func genMessage(gf *generator, m *protogen.Message) {
	if m.Desc.IsMapEntry() {
		return
	}

	gf.P("type ", m.GoIdent, gf.suffix, " struct {")
	for _, field := range m.Fields {
		genMessageField(gf, m, field)
	}
	gf.P("}")

	gf.P()

	genMessageMethods(gf, m)
	genMessageOneofWrapperTypes(gf, m)
}

func getMessageOpts(m *protogen.Message) *MessageOptions {
	opts, _ := proto.GetExtension(m.Desc.Options(), E_Message).(*MessageOptions)
	if opts == nil {
		return &MessageOptions{}
	}
	return opts
}

func genMessageField(gf *generator, m *protogen.Message, field *protogen.Field) {
	if field.Desc.IsWeak() {
		panic("unhandled: weak field")
	}
	if field.Desc.HasDefault() {
		panic("unsupported: default values")
	}

	if oneof := field.Oneof; oneof != nil && !oneof.Desc.IsSynthetic() {
		if oneof.Fields[0] != field {
			return // only generate for first appearance
		}
		gf.P(oneof.GoName, " ", oneofInterfaceName(gf, oneof))
		return
	}

	gf.P(field.GoName, " ", fieldGoType(gf, field))
}

func oneofInterfaceName(gf *generator, oneof *protogen.Oneof) string {
	return "is" + oneof.GoIdent.GoName + gf.suffix
}

func genMessageMethods(gf *generator, m *protogen.Message) {
	genMessageEncode(gf, m)
	genMessageDecode(gf, m)
}

func genMessageEncode(gf *generator, m *protogen.Message) {
	gf.P("func (m *", m.GoIdent, gf.suffix, ") Encode(c *", picobufPackage.Ident("Encoder"), ") bool {")
	gf.P("if m == nil { return false }")

	fields := append([]*protogen.Field(nil), m.Fields...)
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Desc.Number() < fields[j].Desc.Number()
	})

	for _, field := range fields {
		method := codecMethodName(gf, field)
		kind := field.Desc.Kind()
		cardinality := field.Desc.Cardinality()

		switch {
		case method == "Message":
			gf.P("c.Message(", field.Desc.Number(), ", m.", field.GoName, ".Encode)")
		case method == "PresentMessage":
			gf.P("c.PresentMessage(", field.Desc.Number(), ", m.", field.GoName, ".Encode)")
		case method == "RepeatedMessage":
			gf.P("c.RepeatedMessage(", field.Desc.Number(), ", func(c *", picobufPackage.Ident("Encoder"), ", index int) bool {")
			gf.P("  if index >= len(m.", field.GoName, ") { return false }")
			gf.P("  m.", field.GoName, "[index].Encode(c)")
			gf.P("  return true")
			gf.P("})")
		case kind == protoreflect.EnumKind && cardinality == protoreflect.Repeated:
			gf.P("c.RepeatedEnum(", field.Desc.Number(), ", func(index int) *int32 {")
			gf.P("  if index >= len(m.", field.GoName, ") { return nil }")
			gf.P("  return (*int32)(&m.", field.GoName, "[index])")
			gf.P("})")
		case kind == protoreflect.EnumKind:
			gf.P("c.", method, "(", field.Desc.Number(), ", (*int32)(&m.", field.GoName, "))")
		default:
			gf.P("c.", method, "(", field.Desc.Number(), ", &m.", field.GoName, ")")
		}
	}
	gf.P("return true")
	gf.P("}")
	gf.P()
}

func genMessageDecode(gf *generator, m *protogen.Message) {
	gf.P("func (m *", m.GoIdent, gf.suffix, ") Decode(c *", picobufPackage.Ident("Decoder"), ") {")
	gf.P("if m == nil { return }")

	fields := append([]*protogen.Field(nil), m.Fields...)
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Desc.Number() < fields[j].Desc.Number()
	})

	for _, field := range fields {
		method := codecMethodName(gf, field)
		kind := field.Desc.Kind()
		cardinality := field.Desc.Cardinality()

		switch {
		case method == "Message":
			gf.P("c.Message(", field.Desc.Number(), ", func(c *", picobufPackage.Ident("Decoder"), ") {")
			gf.P("  if m.", field.GoName, " == nil {")
			gf.P("    m.", field.GoName, " = new(", fieldGoType(gf, field)[1:], ")")
			gf.P("  }")
			gf.P("  m.", field.GoName, ".Decode(c)")
			gf.P("})")
		case method == "PresentMessage":
			gf.P("c.PresentMessage(", field.Desc.Number(), ", m.", field.GoName, ".Decode)")
		case method == "RepeatedMessage":
			gf.P("c.RepeatedMessage(", field.Desc.Number(), ", func(c *", picobufPackage.Ident("Decoder"), ") {")
			gf.P("  mm := new(", fieldGoType(gf, field)[3:], ")")
			gf.P("  c.Loop(mm.Decode)")
			gf.P("  m.", field.GoName, " = append(m.", field.GoName, ", mm)")
			gf.P("})")
		case kind == protoreflect.EnumKind && cardinality == protoreflect.Repeated:
			gf.P("c.RepeatedEnum(", field.Desc.Number(), ", func(x int32) {")
			gf.P("  m.", field.GoName, " = append(m.", field.GoName, ", (", fieldGoType(gf, field)[2:], ")(x))")
			gf.P("})")
		case kind == protoreflect.EnumKind:
			gf.P("c.", method, "(", field.Desc.Number(), ", (*int32)(&m.", field.GoName, "))")
		default:
			gf.P("c.", method, "(", field.Desc.Number(), ", &m.", field.GoName, ")")
		}
	}
	gf.P("}")
	gf.P()
}

type kind2 [2]protoreflect.Kind

func codecMethodName(gf *generator, field *protogen.Field) string {
	if field.Desc.IsMap() {
		switch (kind2{field.Desc.MapKey().Kind(), field.Desc.MapValue().Kind()}) {
		case kind2{protoreflect.StringKind, protoreflect.StringKind}:
			return "MapStringString"
		default:
			panic(fmt.Sprintf("unhandled: invalid map field kind: <%v,%v>", field.Desc.MapKey().Kind(), field.Desc.MapValue().Kind()))
		}
	}
	if method, ok := methodNames[field.Desc.Kind()]; ok {
		if field.Desc.IsList() {
			method = "Repeated" + method
		}
		if method == "Message" && !getFieldPresence(field) {
			method = "PresentMessage"
		}
		return method
	}
	panic(fmt.Sprintf("unhandled: invalid field kind: %v", field.Desc.Kind()))
}

var methodNames = map[protoreflect.Kind]string{
	protoreflect.BoolKind:     "Bool",
	protoreflect.Int32Kind:    "Int32",
	protoreflect.Int64Kind:    "Int64",
	protoreflect.Uint32Kind:   "Uint32",
	protoreflect.Uint64Kind:   "Uint64",
	protoreflect.Sint32Kind:   "Sint32",
	protoreflect.Sint64Kind:   "Sint64",
	protoreflect.Fixed32Kind:  "Fixed32",
	protoreflect.Fixed64Kind:  "Fixed64",
	protoreflect.Sfixed32Kind: "Sfixed32",
	protoreflect.Sfixed64Kind: "Sfixed64",
	protoreflect.FloatKind:    "Float",
	protoreflect.DoubleKind:   "Double",
	protoreflect.StringKind:   "String",
	protoreflect.BytesKind:    "Bytes",
	protoreflect.EnumKind:     "Int32",
	protoreflect.MessageKind:  "Message",
}

func fieldGoType(gf *generator, field *protogen.Field) (goType string) {
	if field.Desc.IsWeak() {
		panic("unhandled: weak field")
	}

	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		goType = "bool"
	case protoreflect.EnumKind:
		goType = gf.QualifiedGoIdent(field.Enum.GoIdent)
		if field.Enum.GoIdent.GoImportPath == gf.GoImportPath {
			goType += gf.suffix
		}
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		goType = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		goType = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		goType = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		goType = "uint64"
	case protoreflect.FloatKind:
		goType = "float32"
	case protoreflect.DoubleKind:
		goType = "float64"
	case protoreflect.StringKind:
		goType = "string"
	case protoreflect.BytesKind:
		goType = "[]byte"
	case protoreflect.MessageKind:
		// intercept the well-known timestamp type to be our own type
		switch field.Message.GoIdent {
		case timestamppbPackage.Ident("Timestamp"):
			goType = "picobuf.Timestamp"
		default:
			goType = gf.QualifiedGoIdent(field.Message.GoIdent)
			if field.Message.GoIdent.GoImportPath == gf.GoImportPath {
				goType += gf.suffix
			}
		}
	default:
		panic(fmt.Sprintf("unhandled: invalid field kind: %v", field.Desc.Kind()))
	}

	switch {
	case field.Desc.IsList() && field.Desc.Kind() == protoreflect.MessageKind:
		return "[]*" + goType
	case field.Desc.IsList():
		return "[]" + goType
	case getFieldPresence(field):
		return "*" + goType
	case field.Desc.IsMap():
		switch (kind2{field.Desc.MapKey().Kind(), field.Desc.MapValue().Kind()}) {
		case kind2{protoreflect.StringKind, protoreflect.StringKind}:
			return "map[string]string"
		default:
			panic(fmt.Sprintf("unhandled: invalid map field kind: <%v,%v>", field.Desc.MapKey().Kind(), field.Desc.MapValue().Kind()))
		}
	default:
		return goType
	}
}

func getFieldPresence(f *protogen.Field) bool {
	if f.Desc.Kind() == protoreflect.MessageKind && getMessageOpts(f.Message).AlwaysPresent {
		return false
	}
	return f.Desc.HasPresence() && !getFieldOpts(f).AlwaysPresent
}

func getFieldOpts(f *protogen.Field) *FieldOptions {
	opts, _ := proto.GetExtension(f.Desc.Options(), E_Field).(*FieldOptions)
	if opts == nil {
		return &FieldOptions{}
	}
	return opts
}

func genMessageOneofWrapperTypes(gf *generator, m *protogen.Message) {
	for _, oneof := range m.Oneofs {
		if oneof.Desc.IsSynthetic() {
			continue
		}
		ifName := oneofInterfaceName(gf, oneof)

		gf.P("type ", ifName, " interface {", ifName, "() }")
		gf.P()

		for _, field := range oneof.Fields {
			gf.P("type ", field.GoIdent, gf.suffix, " struct {")
			gf.P(field.GoName, " ", fieldGoType(gf, field))
			gf.P("}")
			gf.P()
		}
		gf.P()

		for _, field := range oneof.Fields {
			gf.P("func (*", field.GoIdent, gf.suffix, ") ", ifName, "() {}")
		}
		gf.P()
	}
}
