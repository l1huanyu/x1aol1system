// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/l1huanyu/x1aol1system/infrastructure/db/ent/mission"
)

// Mission is the model entity for the Mission schema.
type Mission struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Mission) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case mission.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Mission", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Mission fields.
func (m *Mission) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Mission.
// Note that you need to call Mission.Unwrap() before calling this method if this Mission
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Mission) Update() *MissionUpdateOne {
	return (&MissionClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Mission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Mission) Unwrap() *Mission {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Mission is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Mission) String() string {
	var builder strings.Builder
	builder.WriteString("Mission(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Missions is a parsable slice of Mission.
type Missions []*Mission

func (m Missions) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}