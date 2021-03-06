// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/godcong/telebot/database/ent/statistic"
)

// Statistic is the model entity for the Statistic schema.
type Statistic struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LatName holds the value of the "lat_name" field.
	LatName string `json:"lat_name,omitempty"`
	// UserName holds the value of the "user_name" field.
	UserName string `json:"user_name,omitempty"`
	// FromUser holds the value of the "from_user" field.
	FromUser int `json:"from_user,omitempty"`
	// ChannelID holds the value of the "channel_id" field.
	ChannelID int64 `json:"channel_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// JoinTime holds the value of the "join_time" field.
	JoinTime time.Time `json:"join_time,omitempty"`
	// Invited holds the value of the "invited" field.
	Invited int64 `json:"invited,omitempty"`
	// Message holds the value of the "message" field.
	Message int64 `json:"message,omitempty"`
	// LastMessage holds the value of the "last_message" field.
	LastMessage time.Time `json:"last_message,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Statistic) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case statistic.FieldID, statistic.FieldFromUser, statistic.FieldChannelID, statistic.FieldUserID, statistic.FieldInvited, statistic.FieldMessage:
			values[i] = new(sql.NullInt64)
		case statistic.FieldFirstName, statistic.FieldLatName, statistic.FieldUserName:
			values[i] = new(sql.NullString)
		case statistic.FieldJoinTime, statistic.FieldLastMessage:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Statistic", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Statistic fields.
func (s *Statistic) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case statistic.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case statistic.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				s.FirstName = value.String
			}
		case statistic.FieldLatName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lat_name", values[i])
			} else if value.Valid {
				s.LatName = value.String
			}
		case statistic.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_name", values[i])
			} else if value.Valid {
				s.UserName = value.String
			}
		case statistic.FieldFromUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field from_user", values[i])
			} else if value.Valid {
				s.FromUser = int(value.Int64)
			}
		case statistic.FieldChannelID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field channel_id", values[i])
			} else if value.Valid {
				s.ChannelID = value.Int64
			}
		case statistic.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				s.UserID = int(value.Int64)
			}
		case statistic.FieldJoinTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field join_time", values[i])
			} else if value.Valid {
				s.JoinTime = value.Time
			}
		case statistic.FieldInvited:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field invited", values[i])
			} else if value.Valid {
				s.Invited = value.Int64
			}
		case statistic.FieldMessage:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				s.Message = value.Int64
			}
		case statistic.FieldLastMessage:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_message", values[i])
			} else if value.Valid {
				s.LastMessage = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Statistic.
// Note that you need to call Statistic.Unwrap() before calling this method if this Statistic
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Statistic) Update() *StatisticUpdateOne {
	return (&StatisticClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Statistic entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Statistic) Unwrap() *Statistic {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Statistic is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Statistic) String() string {
	var builder strings.Builder
	builder.WriteString("Statistic(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", first_name=")
	builder.WriteString(s.FirstName)
	builder.WriteString(", lat_name=")
	builder.WriteString(s.LatName)
	builder.WriteString(", user_name=")
	builder.WriteString(s.UserName)
	builder.WriteString(", from_user=")
	builder.WriteString(fmt.Sprintf("%v", s.FromUser))
	builder.WriteString(", channel_id=")
	builder.WriteString(fmt.Sprintf("%v", s.ChannelID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", s.UserID))
	builder.WriteString(", join_time=")
	builder.WriteString(s.JoinTime.Format(time.ANSIC))
	builder.WriteString(", invited=")
	builder.WriteString(fmt.Sprintf("%v", s.Invited))
	builder.WriteString(", message=")
	builder.WriteString(fmt.Sprintf("%v", s.Message))
	builder.WriteString(", last_message=")
	builder.WriteString(s.LastMessage.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Statistics is a parsable slice of Statistic.
type Statistics []*Statistic

func (s Statistics) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
