// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type ArrayCoerceExpr struct {
	Xpr          Expr         `json:"xpr"`
	Arg          *Expr        `json:"arg"`          /* input expression (yields an array) */
	Elemfuncid   Oid          `json:"elemfuncid"`   /* OID of element coercion function, or 0 */
	Resulttype   Oid          `json:"resulttype"`   /* output type of coercion (an array type) */
	Resulttypmod int32        `json:"resulttypmod"` /* output typmod (also element typmod) */
	Resultcollid Oid          `json:"resultcollid"` /* OID of collation, or InvalidOid if none */
	IsExplicit   bool         `json:"isExplicit"`   /* conversion semantics flag to pass to func */
	Coerceformat CoercionForm `json:"coerceformat"` /* how to display this node */
	Location     int          `json:"location"`     /* token location, or -1 if unknown */
}

func (node ArrayCoerceExpr) MarshalJSON() ([]byte, error) {
	type ArrayCoerceExprMarshalAlias ArrayCoerceExpr
	return json.Marshal(map[string]interface{}{
		"ARRAYCOERCEEXPR": (*ArrayCoerceExprMarshalAlias)(&node),
	})
}

func (node *ArrayCoerceExpr) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["xpr"] != nil {
		err = json.Unmarshal(fields["xpr"], &node.Xpr)
		if err != nil {
			return
		}
	}

	if fields["arg"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["arg"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Expr)
			node.Arg = &val
		}
	}

	if fields["elemfuncid"] != nil {
		err = json.Unmarshal(fields["elemfuncid"], &node.Elemfuncid)
		if err != nil {
			return
		}
	}

	if fields["resulttype"] != nil {
		err = json.Unmarshal(fields["resulttype"], &node.Resulttype)
		if err != nil {
			return
		}
	}

	if fields["resulttypmod"] != nil {
		err = json.Unmarshal(fields["resulttypmod"], &node.Resulttypmod)
		if err != nil {
			return
		}
	}

	if fields["resultcollid"] != nil {
		err = json.Unmarshal(fields["resultcollid"], &node.Resultcollid)
		if err != nil {
			return
		}
	}

	if fields["isExplicit"] != nil {
		err = json.Unmarshal(fields["isExplicit"], &node.IsExplicit)
		if err != nil {
			return
		}
	}

	if fields["coerceformat"] != nil {
		err = json.Unmarshal(fields["coerceformat"], &node.Coerceformat)
		if err != nil {
			return
		}
	}

	if fields["location"] != nil {
		err = json.Unmarshal(fields["location"], &node.Location)
		if err != nil {
			return
		}
	}

	return
}