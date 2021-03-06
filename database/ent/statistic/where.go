// Code generated by entc, DO NOT EDIT.

package statistic

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/godcong/telebot/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirstName), v))
	})
}

// LatName applies equality check predicate on the "lat_name" field. It's identical to LatNameEQ.
func LatName(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLatName), v))
	})
}

// UserName applies equality check predicate on the "user_name" field. It's identical to UserNameEQ.
func UserName(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserName), v))
	})
}

// FromUser applies equality check predicate on the "from_user" field. It's identical to FromUserEQ.
func FromUser(v int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFromUser), v))
	})
}

// ChannelID applies equality check predicate on the "channel_id" field. It's identical to ChannelIDEQ.
func ChannelID(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChannelID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// JoinTime applies equality check predicate on the "join_time" field. It's identical to JoinTimeEQ.
func JoinTime(v time.Time) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldJoinTime), v))
	})
}

// Invited applies equality check predicate on the "invited" field. It's identical to InvitedEQ.
func Invited(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInvited), v))
	})
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// LastMessage applies equality check predicate on the "last_message" field. It's identical to LastMessageEQ.
func LastMessage(v time.Time) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastMessage), v))
	})
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirstName), v))
	})
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFirstName), v))
	})
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFirstName), v...))
	})
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFirstName), v...))
	})
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFirstName), v))
	})
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFirstName), v))
	})
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFirstName), v))
	})
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFirstName), v))
	})
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFirstName), v))
	})
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFirstName), v))
	})
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFirstName), v))
	})
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFirstName), v))
	})
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFirstName), v))
	})
}

// LatNameEQ applies the EQ predicate on the "lat_name" field.
func LatNameEQ(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLatName), v))
	})
}

// LatNameNEQ applies the NEQ predicate on the "lat_name" field.
func LatNameNEQ(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLatName), v))
	})
}

// LatNameIn applies the In predicate on the "lat_name" field.
func LatNameIn(vs ...string) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLatName), v...))
	})
}

// LatNameNotIn applies the NotIn predicate on the "lat_name" field.
func LatNameNotIn(vs ...string) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLatName), v...))
	})
}

// LatNameGT applies the GT predicate on the "lat_name" field.
func LatNameGT(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLatName), v))
	})
}

// LatNameGTE applies the GTE predicate on the "lat_name" field.
func LatNameGTE(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLatName), v))
	})
}

// LatNameLT applies the LT predicate on the "lat_name" field.
func LatNameLT(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLatName), v))
	})
}

// LatNameLTE applies the LTE predicate on the "lat_name" field.
func LatNameLTE(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLatName), v))
	})
}

// LatNameContains applies the Contains predicate on the "lat_name" field.
func LatNameContains(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLatName), v))
	})
}

// LatNameHasPrefix applies the HasPrefix predicate on the "lat_name" field.
func LatNameHasPrefix(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLatName), v))
	})
}

// LatNameHasSuffix applies the HasSuffix predicate on the "lat_name" field.
func LatNameHasSuffix(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLatName), v))
	})
}

// LatNameEqualFold applies the EqualFold predicate on the "lat_name" field.
func LatNameEqualFold(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLatName), v))
	})
}

// LatNameContainsFold applies the ContainsFold predicate on the "lat_name" field.
func LatNameContainsFold(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLatName), v))
	})
}

// UserNameEQ applies the EQ predicate on the "user_name" field.
func UserNameEQ(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserName), v))
	})
}

// UserNameNEQ applies the NEQ predicate on the "user_name" field.
func UserNameNEQ(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserName), v))
	})
}

// UserNameIn applies the In predicate on the "user_name" field.
func UserNameIn(vs ...string) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserName), v...))
	})
}

// UserNameNotIn applies the NotIn predicate on the "user_name" field.
func UserNameNotIn(vs ...string) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserName), v...))
	})
}

// UserNameGT applies the GT predicate on the "user_name" field.
func UserNameGT(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserName), v))
	})
}

// UserNameGTE applies the GTE predicate on the "user_name" field.
func UserNameGTE(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserName), v))
	})
}

// UserNameLT applies the LT predicate on the "user_name" field.
func UserNameLT(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserName), v))
	})
}

// UserNameLTE applies the LTE predicate on the "user_name" field.
func UserNameLTE(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserName), v))
	})
}

// UserNameContains applies the Contains predicate on the "user_name" field.
func UserNameContains(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserName), v))
	})
}

// UserNameHasPrefix applies the HasPrefix predicate on the "user_name" field.
func UserNameHasPrefix(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserName), v))
	})
}

// UserNameHasSuffix applies the HasSuffix predicate on the "user_name" field.
func UserNameHasSuffix(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserName), v))
	})
}

// UserNameEqualFold applies the EqualFold predicate on the "user_name" field.
func UserNameEqualFold(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserName), v))
	})
}

// UserNameContainsFold applies the ContainsFold predicate on the "user_name" field.
func UserNameContainsFold(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserName), v))
	})
}

// FromUserEQ applies the EQ predicate on the "from_user" field.
func FromUserEQ(v int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFromUser), v))
	})
}

// FromUserNEQ applies the NEQ predicate on the "from_user" field.
func FromUserNEQ(v int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFromUser), v))
	})
}

// FromUserIn applies the In predicate on the "from_user" field.
func FromUserIn(vs ...int) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFromUser), v...))
	})
}

// FromUserNotIn applies the NotIn predicate on the "from_user" field.
func FromUserNotIn(vs ...int) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFromUser), v...))
	})
}

// FromUserGT applies the GT predicate on the "from_user" field.
func FromUserGT(v int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFromUser), v))
	})
}

// FromUserGTE applies the GTE predicate on the "from_user" field.
func FromUserGTE(v int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFromUser), v))
	})
}

// FromUserLT applies the LT predicate on the "from_user" field.
func FromUserLT(v int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFromUser), v))
	})
}

// FromUserLTE applies the LTE predicate on the "from_user" field.
func FromUserLTE(v int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFromUser), v))
	})
}

// ChannelIDEQ applies the EQ predicate on the "channel_id" field.
func ChannelIDEQ(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChannelID), v))
	})
}

// ChannelIDNEQ applies the NEQ predicate on the "channel_id" field.
func ChannelIDNEQ(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChannelID), v))
	})
}

// ChannelIDIn applies the In predicate on the "channel_id" field.
func ChannelIDIn(vs ...int64) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldChannelID), v...))
	})
}

// ChannelIDNotIn applies the NotIn predicate on the "channel_id" field.
func ChannelIDNotIn(vs ...int64) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldChannelID), v...))
	})
}

// ChannelIDGT applies the GT predicate on the "channel_id" field.
func ChannelIDGT(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChannelID), v))
	})
}

// ChannelIDGTE applies the GTE predicate on the "channel_id" field.
func ChannelIDGTE(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChannelID), v))
	})
}

// ChannelIDLT applies the LT predicate on the "channel_id" field.
func ChannelIDLT(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChannelID), v))
	})
}

// ChannelIDLTE applies the LTE predicate on the "channel_id" field.
func ChannelIDLTE(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChannelID), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// JoinTimeEQ applies the EQ predicate on the "join_time" field.
func JoinTimeEQ(v time.Time) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldJoinTime), v))
	})
}

// JoinTimeNEQ applies the NEQ predicate on the "join_time" field.
func JoinTimeNEQ(v time.Time) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldJoinTime), v))
	})
}

// JoinTimeIn applies the In predicate on the "join_time" field.
func JoinTimeIn(vs ...time.Time) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldJoinTime), v...))
	})
}

// JoinTimeNotIn applies the NotIn predicate on the "join_time" field.
func JoinTimeNotIn(vs ...time.Time) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldJoinTime), v...))
	})
}

// JoinTimeGT applies the GT predicate on the "join_time" field.
func JoinTimeGT(v time.Time) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldJoinTime), v))
	})
}

// JoinTimeGTE applies the GTE predicate on the "join_time" field.
func JoinTimeGTE(v time.Time) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldJoinTime), v))
	})
}

// JoinTimeLT applies the LT predicate on the "join_time" field.
func JoinTimeLT(v time.Time) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldJoinTime), v))
	})
}

// JoinTimeLTE applies the LTE predicate on the "join_time" field.
func JoinTimeLTE(v time.Time) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldJoinTime), v))
	})
}

// InvitedEQ applies the EQ predicate on the "invited" field.
func InvitedEQ(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInvited), v))
	})
}

// InvitedNEQ applies the NEQ predicate on the "invited" field.
func InvitedNEQ(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInvited), v))
	})
}

// InvitedIn applies the In predicate on the "invited" field.
func InvitedIn(vs ...int64) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInvited), v...))
	})
}

// InvitedNotIn applies the NotIn predicate on the "invited" field.
func InvitedNotIn(vs ...int64) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInvited), v...))
	})
}

// InvitedGT applies the GT predicate on the "invited" field.
func InvitedGT(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInvited), v))
	})
}

// InvitedGTE applies the GTE predicate on the "invited" field.
func InvitedGTE(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInvited), v))
	})
}

// InvitedLT applies the LT predicate on the "invited" field.
func InvitedLT(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInvited), v))
	})
}

// InvitedLTE applies the LTE predicate on the "invited" field.
func InvitedLTE(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInvited), v))
	})
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMessage), v))
	})
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...int64) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMessage), v...))
	})
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...int64) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMessage), v...))
	})
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMessage), v))
	})
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMessage), v))
	})
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMessage), v))
	})
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v int64) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMessage), v))
	})
}

// LastMessageEQ applies the EQ predicate on the "last_message" field.
func LastMessageEQ(v time.Time) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastMessage), v))
	})
}

// LastMessageNEQ applies the NEQ predicate on the "last_message" field.
func LastMessageNEQ(v time.Time) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastMessage), v))
	})
}

// LastMessageIn applies the In predicate on the "last_message" field.
func LastMessageIn(vs ...time.Time) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastMessage), v...))
	})
}

// LastMessageNotIn applies the NotIn predicate on the "last_message" field.
func LastMessageNotIn(vs ...time.Time) predicate.Statistic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastMessage), v...))
	})
}

// LastMessageGT applies the GT predicate on the "last_message" field.
func LastMessageGT(v time.Time) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastMessage), v))
	})
}

// LastMessageGTE applies the GTE predicate on the "last_message" field.
func LastMessageGTE(v time.Time) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastMessage), v))
	})
}

// LastMessageLT applies the LT predicate on the "last_message" field.
func LastMessageLT(v time.Time) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastMessage), v))
	})
}

// LastMessageLTE applies the LTE predicate on the "last_message" field.
func LastMessageLTE(v time.Time) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastMessage), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Statistic) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Statistic) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Statistic) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		p(s.Not())
	})
}
