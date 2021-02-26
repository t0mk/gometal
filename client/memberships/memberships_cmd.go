// Code generated by go-swagger; DO NOT EDIT.

package memberships

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Command(ctx context.Context, client ClientService) *cobra.Command {
	command := cobra.Command{
		Use:   "memberships",
		Short: "Client for memberships",
	}

	command.AddCommand(deleteMembership(ctx, client))
	command.AddCommand(findMembershipById(ctx, client))
	command.AddCommand(updateMembership(ctx, client))

	return &command
}

func deleteMembership(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &DeleteMembershipParams{}

	command := &cobra.Command{
		Use: "DeleteMembership",
		RunE: func(cmd *cobra.Command, args []string) error {

			_, err := client.DeleteMembership(parameters, nil)
			return err

		},
	}

	parameters.Context = ctx

	// strfmt.UUID

	command.MarkPersistentFlagRequired("id")

	return command
}

func findMembershipById(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &FindMembershipByIDParams{}

	command := &cobra.Command{
		Use: "FindMembershipByID",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.FindMembershipByID(parameters, nil)
			if err != nil {
				return err
			}
			return printPayload(result.Payload)

		},
	}

	parameters.Context = ctx

	// strfmt.UUID

	command.MarkPersistentFlagRequired("id")

	// string

	parameters.Include = command.PersistentFlags().String("include", "", "related attributes to include")

	return command
}

func updateMembership(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &UpdateMembershipParams{}

	command := &cobra.Command{
		Use: "UpdateMembership",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.UpdateMembership(parameters, nil)
			if err != nil {
				return err
			}
			return printPayload(result.Payload)

		},
	}

	parameters.Context = ctx

	// strfmt.UUID

	command.MarkPersistentFlagRequired("id")

	// types.MembershipInput

	return command
}

func printPayload(value interface{}) error {
	bytes, err := json.MarshalIndent(value, "", "  ")

	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(bytes)
	return err
}

type dateTimeValue struct {
	v *strfmt.DateTime
}

func (dtv *dateTimeValue) String() string {
	return time.Time(*dtv.v).Format(strfmt.ISO8601LocalTime)
}

func (dtv *dateTimeValue) Set(flag string) error {
	return dtv.v.UnmarshalText([]byte(flag))
}

func (dtv *dateTimeValue) Type() string {
	return "dateTime"
}

func newDateTimeValue(v *strfmt.DateTime) pflag.Value {
	return &dateTimeValue{v}
}
