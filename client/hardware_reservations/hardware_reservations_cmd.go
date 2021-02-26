// Code generated by go-swagger; DO NOT EDIT.

package hardware_reservations

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
		Use:   "hardware_reservations",
		Short: "Client for hardware_reservations",
	}

	command.AddCommand(postHardwareReservationsIdMove(ctx, client))
	command.AddCommand(findHardwareReservationById(ctx, client))

	return &command
}

func postHardwareReservationsIdMove(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &PostHardwareReservationsIDMoveParams{}

	command := &cobra.Command{
		Use: "PostHardwareReservationsIDMove",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.PostHardwareReservationsIDMove(parameters, nil)
			if err != nil {
				return err
			}
			return printPayload(result.Payload)

		},
	}

	parameters.Context = ctx

	// strfmt.UUID

	command.MarkPersistentFlagRequired("id")

	// strfmt.UUID

	command.MarkPersistentFlagRequired("projectId")

	return command
}

func findHardwareReservationById(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &FindHardwareReservationByIDParams{}

	command := &cobra.Command{
		Use: "FindHardwareReservationByID",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.FindHardwareReservationByID(parameters, nil)
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
