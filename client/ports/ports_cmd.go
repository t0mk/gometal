// Code generated by go-swagger; DO NOT EDIT.

package ports

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
		Use:   "ports",
		Short: "Client for ports",
	}

	command.AddCommand(assignNativeVlan(ctx, client))
	command.AddCommand(assignPort(ctx, client))
	command.AddCommand(bondPort(ctx, client))
	command.AddCommand(convertLayer2(ctx, client))
	command.AddCommand(convertLayer3(ctx, client))
	command.AddCommand(deleteNativeVlan(ctx, client))
	command.AddCommand(disbondPort(ctx, client))
	command.AddCommand(findPortById(ctx, client))
	command.AddCommand(unassignPort(ctx, client))

	return &command
}

func assignNativeVlan(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &AssignNativeVLANParams{}

	command := &cobra.Command{
		Use: "AssignNativeVLAN",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.AssignNativeVLAN(parameters, nil)
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

	command.PersistentFlags().StringVar(&parameters.Vnid, "vnid", "", "UUID or VNID of the virtual network to assign")

	command.MarkPersistentFlagRequired("vnid")

	return command
}

func assignPort(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &AssignPortParams{}

	command := &cobra.Command{
		Use: "AssignPort",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.AssignPort(parameters, nil)
			if err != nil {
				return err
			}
			return printPayload(result.Payload)

		},
	}

	parameters.Context = ctx

	// strfmt.UUID

	command.MarkPersistentFlagRequired("id")

	// types.PortAssignInput

	return command
}

func bondPort(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &BondPortParams{}

	command := &cobra.Command{
		Use: "BondPort",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.BondPort(parameters, nil)
			if err != nil {
				return err
			}
			return printPayload(result.Payload)

		},
	}

	parameters.Context = ctx

	// bool

	parameters.BulkEnable = command.PersistentFlags().Bool("bulkEnable", false, "enable both ports")

	// strfmt.UUID

	command.MarkPersistentFlagRequired("id")

	return command
}

func convertLayer2(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &ConvertLayer2Params{}

	command := &cobra.Command{
		Use: "ConvertLayer2",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.ConvertLayer2(parameters, nil)
			if err != nil {
				return err
			}
			return printPayload(result.Payload)

		},
	}

	parameters.Context = ctx

	// strfmt.UUID

	command.MarkPersistentFlagRequired("id")

	// types.PortAssignInput

	return command
}

func convertLayer3(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &ConvertLayer3Params{}

	command := &cobra.Command{
		Use: "ConvertLayer3",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.ConvertLayer3(parameters, nil)
			if err != nil {
				return err
			}
			return printPayload(result.Payload)

		},
	}

	parameters.Context = ctx

	// strfmt.UUID

	command.MarkPersistentFlagRequired("id")

	// types.PortConvertLayer3Input

	return command
}

func deleteNativeVlan(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &DeleteNativeVLANParams{}

	command := &cobra.Command{
		Use: "DeleteNativeVLAN",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.DeleteNativeVLAN(parameters, nil)
			if err != nil {
				return err
			}
			return printPayload(result.Payload)

		},
	}

	parameters.Context = ctx

	// strfmt.UUID

	command.MarkPersistentFlagRequired("id")

	return command
}

func disbondPort(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &DisbondPortParams{}

	command := &cobra.Command{
		Use: "DisbondPort",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.DisbondPort(parameters, nil)
			if err != nil {
				return err
			}
			return printPayload(result.Payload)

		},
	}

	parameters.Context = ctx

	// bool

	parameters.BulkDisable = command.PersistentFlags().Bool("bulkDisable", false, "disable both ports")

	// strfmt.UUID

	command.MarkPersistentFlagRequired("id")

	return command
}

func findPortById(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &FindPortByIDParams{}

	command := &cobra.Command{
		Use: "FindPortByID",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.FindPortByID(parameters, nil)
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

func unassignPort(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &UnassignPortParams{}

	command := &cobra.Command{
		Use: "UnassignPort",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.UnassignPort(parameters, nil)
			if err != nil {
				return err
			}
			return printPayload(result.Payload)

		},
	}

	parameters.Context = ctx

	// strfmt.UUID

	command.MarkPersistentFlagRequired("id")

	// types.PortAssignInput

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
