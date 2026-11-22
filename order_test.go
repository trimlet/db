package main

import "testing"

// TestCannotOrder: text is PostgreSQL's usual string type, SQL Server's old long text type.
func TestCannotOrder(t *testing.T) {
	cases := []struct {
		e    engine
		base string
		want bool
	}{
		{postgres{}, "text", false},
		{postgres{}, "varchar", false},
		{postgres{}, "json", true},
		{sqlServer{}, "text", true},
		{sqlServer{}, "varchar", false},
		{sqlServer{}, "nvarchar", false},
	}
	for _, c := range cases {
		if got := cannotOrder(c.e, c.base); got != c.want {
			t.Errorf("%s %s: cannotOrder = %v, want %v", c.e.name(), c.base, got, c.want)
		}
	}
}
