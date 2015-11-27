// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterSeqStmt struct {
	Sequence  *RangeVar `json:"sequence"` /* the sequence to alter */
	Options   []Node    `json:"options"`
	MissingOk bool      `json:"missing_ok"` /* skip error if a role is missing? */
}

func (node AlterSeqStmt) MarshalJSON() ([]byte, error) {
	type AlterSeqStmtMarshalAlias AlterSeqStmt
	return json.Marshal(map[string]interface{}{
		"ALTERSEQSTMT": (*AlterSeqStmtMarshalAlias)(&node),
	})
}

func (node *AlterSeqStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["sequence"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["sequence"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(RangeVar)
			node.Sequence = &val
		}
	}

	if fields["options"] != nil {
		node.Options, err = UnmarshalNodeArrayJSON(fields["options"])
		if err != nil {
			return
		}
	}

	if fields["missing_ok"] != nil {
		err = json.Unmarshal(fields["missing_ok"], &node.MissingOk)
		if err != nil {
			return
		}
	}

	return
}