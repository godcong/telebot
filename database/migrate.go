package database

import (
	"context"

	"github.com/godcong/telebot/database/ent"
	"github.com/godcong/telebot/database/ent/message"
	"github.com/godcong/telebot/database/ent/schema"
)

func (d *DB) Migrate(ctx context.Context) error {
	if err := initWelcomeHook(ctx, d.Message.Create()); err != nil {
		return err
	}
	if err := initNewMemberStatisticHook(ctx, d.Message.Create()); err != nil {
		return err
	}
	if err := initStatisticHook(ctx, d.Message.Create()); err != nil {
		return err
	}
	return nil
}

func initWelcomeHook(ctx context.Context, create *ent.MessageCreate) error {
	_, err := create.SetMessage("Welcome %v join to channel").
		SetAction(message.ActionWelcome).
		SetAutoRemove(true).
		SetAutoRemoveTime(30).
		SetType(int(schema.MessageTypeChatMember)).
		Save(ctx)
	return err
}

func initNewMemberStatisticHook(ctx context.Context, create *ent.MessageCreate) error {
	_, err := create.SetMessage("").
		SetAction(message.ActionStatistic).
		SetAutoRemove(false).
		SetAutoRemoveTime(0).
		SetType(int(schema.MessageTypeChatMember)).
		Save(ctx)
	return err
}

func initStatisticHook(ctx context.Context, create *ent.MessageCreate) error {
	_, err := create.SetMessage("").
		SetAction(message.ActionStatistic).
		SetAutoRemove(false).
		SetAutoRemoveTime(0).
		SetType(int(schema.MessageTypeMessage)).
		Save(ctx)
	return err
}
