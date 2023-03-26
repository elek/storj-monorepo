// Copyright (C) 2019 Storj Labs, Inc.
// Copyright (C) 2017 Space Monkey, Inc.
// See LICENSE for copying information.

// DBX implements code generation for database schemas and accessors.
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	cli "github.com/jawher/mow.cli"
	"github.com/spacemonkeygo/errors"

	"storj.io/dbx/ast"
	"storj.io/dbx/code/golang"
	"storj.io/dbx/errutil"
	"storj.io/dbx/ir"
	"storj.io/dbx/ir/xform"
	"storj.io/dbx/sql"
	"storj.io/dbx/sqlgen"
	"storj.io/dbx/syntax"
	"storj.io/dbx/templates"
	"storj.io/dbx/tmplutil"
)

// dbx golang (-p package) (-d dialect) (-i DBXFILE) DBXFILE OUTDIR
// dbx schema (-d dialect) (-i DBXFILE) DBXFILE OUTDIR

// Global contains data for running commands.
type Global struct {
	loadedData map[string][]byte
}

func newGlobal() *Global {
	return &Global{
		loadedData: make(map[string][]byte),
	}
}

func main() {
	global := newGlobal()

	die := func(err error) {
		if err != nil {
			// if the error came from errutil, don't bother with the dbx prefix
			err_string := strings.TrimPrefix(errors.GetMessage(err), "dbx: ")
			fmt.Fprintln(os.Stderr, err_string)
			if context := errutil.GetContext(global.loadedData, err); context != "" {
				fmt.Fprintln(os.Stderr)
				fmt.Fprintln(os.Stderr, "context:")
				fmt.Fprintln(os.Stderr, context)
			}
			cli.Exit(1)
		}
	}

	app := cli.App("dbx", "generate SQL schema and matching code")

	app.Command("golang", "generate Go code", func(cmd *cli.Cmd) {
		package_opt := cmd.StringOpt("p package", "",
			"package name for generated code")
		dialects_opt := cmd.StringsOpt("d dialect", nil,
			"SQL dialects (defaults to postgres)")
		templatedir_opt := cmd.StringOpt("t templates", "",
			"override the template directory")
		rx_opt := cmd.BoolOpt("rx", true,
			"generate Rx support")
		userdata_opt := cmd.BoolOpt("userdata", false,
			"generate userdata interface and mutex on models")
		dbxfile_arg := cmd.StringArg("DBXFILE", "",
			"path to dbx file")
		dbxincludes_opt := cmd.StringsOpt("i", []string{}, "include additional dbx files")
		outdir_arg := cmd.StringArg("OUTDIR", "",
			"output directory")
		cmd.Action = func() {
			dbxfiles := append([]string{*dbxfile_arg}, *dbxincludes_opt...)
			die(global.golangCmd(*package_opt, *dialects_opt, *templatedir_opt,
				*rx_opt, *userdata_opt, dbxfiles, *outdir_arg))
		}
	})

	app.Command("schema", "generate table schema", func(cmd *cli.Cmd) {
		dialects_opt := cmd.StringsOpt("d dialect", nil,
			"SQL dialects (default is postgres)")
		dbxincludes_opt := cmd.StringsOpt("i", []string{}, "include additional dbx files")
		dbxfile_arg := cmd.StringArg("DBXFILE", "",
			"path to dbx file")
		outdir_arg := cmd.StringArg("OUTDIR", "",
			"output directory")
		cmd.Action = func() {
			dbxfiles := append([]string{*dbxfile_arg}, *dbxincludes_opt...)
			die(global.schemaCmd(*dialects_opt, dbxfiles, *outdir_arg))
		}
	})

	app.Command("format", "format dbx file on stdin", func(cmd *cli.Cmd) {
		cmd.Action = func() { die(global.formatCmd()) }
	})

	die(app.Run(os.Args))
}

func (global *Global) golangCmd(pkg string, dialects_opt []string, template_dir string,
	rx bool, userdata bool, dbxfiles []string, outdir string) (err error) {

	if pkg == "" {
		base := filepath.Base(dbxfiles[0])
		pkg = base[:len(base)-len(filepath.Ext(base))]
	}

	fw := newFileWriter(outdir, dbxfiles[0])

	root, err := global.parseDBX(dbxfiles...)
	if err != nil {
		return err
	}

	dialects, err := global.createDialects(dialects_opt)
	if err != nil {
		return err
	}

	loader := global.getLoader(template_dir)

	renderer, err := golang.New(loader, &golang.Options{
		Package:         pkg,
		SupportRx:       rx,
		SupportUserdata: userdata,
	})
	if err != nil {
		return err
	}

	rendered, err := renderer.RenderCode(root, dialects)
	if err != nil {
		return err
	}

	if err := fw.writeFile("go", rendered); err != nil {
		return err
	}

	return nil
}

func (global *Global) schemaCmd(dialects_opt []string, dbxfiles []string, outdir string) (err error) {
	fw := newFileWriter(outdir, dbxfiles[0])

	root, err := global.parseDBX(dbxfiles...)
	if err != nil {
		return err
	}

	dialects, err := global.createDialects(dialects_opt)
	if err != nil {
		return err
	}

	for _, dialect := range dialects {
		err = fw.writeFile(dialect.Name()+".sql", global.renderSchema(dialect, root))
		if err != nil {
			return err
		}
	}

	return nil
}

func (global *Global) formatCmd() (err error) {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	global.loadedData[""] = data

	formatted, err := syntax.Format("", data)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(formatted)
	return err
}

func (global *Global) renderSchema(dialect sql.Dialect, root *ir.Root) []byte {
	const schema_hdr = `-- AUTOGENERATED BY storj.io/dbx
-- DO NOT EDIT`

	rendered := sqlgen.Render(dialect,
		sql.SchemaSQL(root, dialect),
		sqlgen.NoTerminate, sqlgen.NoFlatten)

	return []byte(schema_hdr + "\n" + rendered + "\n")
}

func (global *Global) parseDBX(in ...string) (*ir.Root, error) {
	var root ast.Root
	for _, in := range in {
		data, err := os.ReadFile(in)
		if err != nil {
			return nil, err
		}
		global.loadedData[in] = data

		ast, err := syntax.Parse(in, data)
		if err != nil {
			return nil, err
		}

		root.Add(ast)
	}

	return xform.Transform(&root)
}

func (global *Global) getLoader(dir string) tmplutil.Loader {
	loader := tmplutil.FSLoader(templates.Asset)
	if dir != "" {
		return tmplutil.DirLoader(dir, loader)
	}
	return loader
}

func (global *Global) createDialects(which []string) (out []sql.Dialect, err error) {
	if len(which) == 0 {
		which = append(which, "postgres")
	}
	for _, name := range which {
		var d sql.Dialect
		switch name {
		case "postgres":
			d = sql.Postgres()
		case "sqlite3":
			d = sql.SQLite3()
		case "cockroach":
			d = sql.Cockroach()
		case "pgx":
			d = sql.PGX()
		case "pgxcockroach":
			d = sql.PGXCockroach()
		default:
			return nil, fmt.Errorf("unknown dialect %q", name)
		}
		out = append(out, d)
	}
	return out, nil
}

type fileWriter struct {
	dir    string
	prefix string
}

func newFileWriter(outdir, dbxfile string) *fileWriter {
	return &fileWriter{
		dir:    outdir,
		prefix: filepath.Base(dbxfile),
	}
}

func (fw *fileWriter) writeFile(suffix string, data []byte) (err error) {
	file_path := filepath.Join(fw.dir, fw.prefix+"."+suffix)
	tmp_path := file_path + ".tmp"

	if err := os.WriteFile(tmp_path, data, 0644); err != nil {
		return fmt.Errorf("unable to write %s: %v", tmp_path, err)
	}

	if err := os.Rename(tmp_path, file_path); err != nil {
		os.Remove(tmp_path)
		return fmt.Errorf("unable to rename %s over %s: %v", tmp_path,
			file_path, err)
	}

	return nil
}