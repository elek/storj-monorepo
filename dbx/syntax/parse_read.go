// Copyright (C) 2019 Storj Labs, Inc.
// Copyright (C) 2017 Space Monkey, Inc.
// See LICENSE for copying information.

package syntax

import "storj.io/dbx/ast"

func parseRead(node *tupleNode) (*ast.Read, error) {
	read := new(ast.Read)
	read.Pos = node.getPos()

	view, err := parseView(node)
	if err != nil {
		return nil, err
	}
	read.View = view

	list_token, err := node.consumeList()
	if err != nil {
		return nil, err
	}

	err = list_token.consumeAnyTuples(tupleCases{
		"select": func(node *tupleNode) error {
			if read.Select != nil {
				return previouslyDefined(node.getPos(), "read", "select",
					read.Select.Pos)
			}

			refs, err := parseFieldRefs(node, false)
			if err != nil {
				return err
			}
			read.Select = refs

			return nil
		},
		"where": func(node *tupleNode) error {
			where, err := parseWhere(node)
			if err != nil {
				return err
			}
			read.Where = append(read.Where, where)

			return nil
		},
		"join": func(node *tupleNode) error {
			join, err := parseJoin(node)
			if err != nil {
				return err
			}
			read.Joins = append(read.Joins, join)

			return nil
		},
		"orderby": func(node *tupleNode) error {
			if read.OrderBy != nil {
				return previouslyDefined(node.getPos(), "read", "orderby",
					read.OrderBy.Pos)
			}

			order_by, err := parseOrderBy(node)
			if err != nil {
				return err
			}
			read.OrderBy = order_by

			return nil
		},
		"groupby": func(node *tupleNode) error {
			if read.GroupBy != nil {
				return previouslyDefined(node.getPos(), "read", "groupby",
					read.GroupBy.Pos)
			}

			group_by, err := parseGroupBy(node)
			if err != nil {
				return err
			}
			read.GroupBy = group_by

			return nil
		},
		"suffix": func(node *tupleNode) error {
			if read.Suffix != nil {
				return previouslyDefined(node.getPos(), "read", "suffix",
					read.Suffix.Pos)
			}

			suffix, err := parseSuffix(node)
			if err != nil {
				return err
			}
			read.Suffix = suffix

			return nil
		},
	})
	if err != nil {
		return nil, err
	}

	return read, nil
}
