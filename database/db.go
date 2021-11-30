package database

import (
	"context"
	"fmt"

	"github.com/motomototv/telebot/database/ent"
	"github.com/motomototv/telebot/database/ent/message"
	"github.com/motomototv/telebot/database/ent/schema"
	"github.com/motomototv/telebot/database/ent/statistic"
)

const sqlDSN = "file:%v?cache=shared&_journal=WAL&_fk=1"

type DB struct {
	*ent.Client
}

func (d DB) QueryMessages(ctx context.Context, t schema.MessageType) ([]*ent.Message, error) {
	return d.Message.Query().Where(message.TypeEQ(int(t))).All(ctx)
}

func (d DB) UpdateNewMemberStatistic(ctx context.Context, stc *ent.Statistic) error {
	_, err := d.Statistic.Query().Where(statistic.UserID(stc.UserID), statistic.ChannelID(stc.ChannelID)).First(ctx)
	if err != nil {
		_, err = d.Statistic.Create().
			SetUserID(stc.UserID).
			SetChannelID(stc.ChannelID).
			SetFromUser(stc.FromUser).
			SetFirstName(stc.FirstName).
			SetLatName(stc.LatName).
			SetUserName(stc.UserName).
			//SetLastMessage(stc.LastMessage).
			SetJoinTime(stc.JoinTime).Save(ctx)
	} else {
		//_, err = d.Statistic.UpdateOneID(s.ID).SetMessage(s.Message + 1).SetLastMessage(stc.LastMessage).Save(ctx)
		return nil
	}
	return err
	//s, err = d.Statistic.Query().Where(statistic.UserID(stc.FromUser)).First(ctx)
	//if err != nil {
	//	return err
	//}
	//_, err = d.Statistic.UpdateOneID(s.ID).SetInvited(s.Invited + 1).Save(ctx)
	//return err
}

func (d *DB) UpdateChatStatistic(ctx context.Context, stc *ent.Statistic) error {
	s, err := d.Statistic.Query().Where(statistic.UserID(stc.UserID), statistic.ChannelID(stc.ChannelID)).First(ctx)
	if err != nil {
		_, err = d.Statistic.Create().
			SetUserID(stc.UserID).
			SetChannelID(stc.ChannelID).
			SetFromUser(stc.FromUser).
			SetFirstName(stc.FirstName).
			SetLatName(stc.LatName).
			SetUserName(stc.UserName).
			SetMessage(1).
			SetLastMessage(stc.LastMessage).
			SetJoinTime(stc.JoinTime).Save(ctx)
	} else {
		_, err = d.Statistic.UpdateOneID(s.ID).SetMessage(s.Message + 1).SetLastMessage(stc.LastMessage).Save(ctx)
	}
	return err
}

func (d *DB) UpdateInviteStatistic(ctx context.Context, stc *ent.Statistic) error {
	s, err := d.Statistic.Query().Where(statistic.UserID(stc.UserID), statistic.ChannelID(stc.ChannelID)).First(ctx)
	if err != nil {
		_, err = d.Statistic.Create().
			SetUserID(stc.UserID).
			SetChannelID(stc.ChannelID).
			SetFromUser(stc.FromUser).
			SetFirstName(stc.FirstName).
			SetLatName(stc.LatName).
			SetUserName(stc.UserName).
			SetInvited(1).
			SetLastMessage(stc.LastMessage).
			SetJoinTime(stc.JoinTime).Save(ctx)
	} else {
		_, err = d.Statistic.UpdateOneID(s.ID).SetInvited(s.Invited + 1).Save(ctx)
	}
	return err
}

func (d *DB) QueryStatistics(ctx context.Context) ([]*ent.Statistic, error) {
	return d.Statistic.Query().All(ctx)
}

func (d *DB) UpdateMessage(ctx context.Context, e *ent.Message) (*ent.Message, error) {
	return d.Message.UpdateOneID(e.ID).
		SetMessage(e.Message).
		SetType(e.Type).
		SetAction(e.Action).
		SetAutoRemoveTime(e.AutoRemoveTime).
		SetAutoRemove(e.AutoRemove).
		Save(ctx)
}

func Open(ctx context.Context, file string, debug bool) (*DB, error) {
	var options []ent.Option

	if debug {
		options = append(options, ent.Debug())
	}

	client, err := ent.Open("sqlite3", dsnFile(file), options...)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to database: %v", err)
	}

	if err := client.Schema.Create(ctx); err != nil {
		return nil, err
	}

	return &DB{
		Client: client,
	}, nil
}

func dsnFile(file string) string {
	return fmt.Sprintf(sqlDSN, file)
}
