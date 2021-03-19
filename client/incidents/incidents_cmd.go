// Code generated by go-swagger; DO NOT EDIT.

package incidents

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
		Use:   "incidents",
		Short: "Client for incidents",
	}

	command.AddCommand(getIncidents(ctx, client))

	return &command
}

func getIncidents(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &GetIncidentsParams{}

	command := &cobra.Command{
		Use: "GetIncidents",
		RunE: func(cmd *cobra.Command, args []string) error {

			_, err := client.GetIncidents(parameters, nil)
			return err

		},
	}

	parameters.Context = ctx

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