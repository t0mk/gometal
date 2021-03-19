// Code generated by go-swagger; DO NOT EDIT.

package licenses

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
		Use:   "licenses",
		Short: "Client for licenses",
	}

	command.AddCommand(deleteLicense(ctx, client))
	command.AddCommand(findLicenseById(ctx, client))
	command.AddCommand(updateLicense(ctx, client))

	return &command
}

func deleteLicense(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &DeleteLicenseParams{}

	command := &cobra.Command{
		Use: "DeleteLicense",
		RunE: func(cmd *cobra.Command, args []string) error {

			_, err := client.DeleteLicense(parameters, nil)
			return err

		},
	}

	parameters.Context = ctx

	// strfmt.UUID

	command.MarkPersistentFlagRequired("id")

	return command
}

func findLicenseById(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &FindLicenseByIDParams{}

	command := &cobra.Command{
		Use: "FindLicenseByID",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.FindLicenseByID(parameters, nil)
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

func updateLicense(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &UpdateLicenseParams{}

	command := &cobra.Command{
		Use: "UpdateLicense",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.UpdateLicense(parameters, nil)
			if err != nil {
				return err
			}
			return printPayload(result.Payload)

		},
	}

	parameters.Context = ctx

	// strfmt.UUID

	command.MarkPersistentFlagRequired("id")

	// types.LicenseUpdateInput

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
