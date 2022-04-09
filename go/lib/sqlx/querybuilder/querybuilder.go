package querybuilder

import (
	"fmt"
	"strings"
)

// Query is main struct for building query
type Query struct {
	sb        strings.Builder
	whereFmt  []string
	whereArgs []interface{}
	extraStr  []string
}

// New will create new query
func New(template string) *Query {
	q := &Query{
		sb: strings.Builder{},
	}

	q.sb.WriteString(strings.ToLower(template))
	return q
}

// AddCondition will add where string format & arguments
func (q *Query) AddCondition(format string, data ...interface{}) {
	q.whereFmt = append(q.whereFmt, strings.ToLower(format))
	q.whereArgs = append(q.whereArgs, data...)
}

// AddString will append string at the end of query, after where condition finished
func (q *Query) AddString(str string) {
	q.sb.WriteString(" " + str)
}

func (q *Query) buildWhere() {
	var (
		whereSb  strings.Builder
		whereFmt string
		counter  int
	)

	if len(q.whereFmt) == 0 {
		return
	}

	if !strings.Contains(q.sb.String(), "where") {
		whereSb.WriteString(" where ")
	}

	// append all where conditions
	for _, condition := range q.whereFmt {
		whereSb.WriteString(" " + condition)
	}
	whereFmt = whereSb.String()

	// replace ? with $counter
	for strings.Contains(whereFmt, "?") {
		counter++
		whereFmt = strings.Replace(whereFmt, "?", fmt.Sprintf("$%d", counter), 1)
	}
	q.sb.WriteString(whereFmt)
}

func (q *Query) buildExtraStr() {
	// add extra string
	for _, str := range q.extraStr {
		q.sb.WriteString(" " + str)
	}

	if !strings.HasSuffix(q.sb.String(), ";") {
		q.sb.WriteString(";")
	}
}

func (q *Query) buildQuery() {
	q.buildWhere()
	q.buildExtraStr()
}

// String will return formatted query string
func (q *Query) String() string {
	q.buildQuery()
	return q.sb.String()
}

// Params will return where arguments
func (q *Query) Params() []interface{} {
	return q.whereArgs
}
