package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	authorizationQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the msg authorization module",
		Long:                       "",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	authorizationQueryCmd.AddCommand(
		GetCmdQueryAuthorization(),
		GetCmdQueryAuthorizations(),
	)

	return authorizationQueryCmd
}

// GetCmdQueryAuthorizations implements the query authorizations command.
func GetCmdQueryAuthorizations() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "authorizations [granter-addr] [grantee-addr]",
		Args:  cobra.ExactArgs(2),
		Short: "query list of authorizations for a granter-grantee pair",
		Long:  "query list of authorizations for a granter-grantee pair",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			granterAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			granteeAddr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			res, err := queryClient.Authorizations(
				context.Background(),
				&types.QueryAuthorizationsRequest{
					GranterAddr: granterAddr.String(),
					GranteeAddr: granteeAddr.String(),
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryAuthorization implements the query authorization command.
func GetCmdQueryAuthorization() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "authorization [granter-addr] [grantee-addr] [msg-type]",
		Args:  cobra.ExactArgs(3),
		Short: "query authorization for a granter-grantee pair",
		Long:  "query authorization for a granter-grantee pair",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			granterAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			granteeAddr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msgAuthorized := args[2]

			res, err := queryClient.Authorization(
				context.Background(),
				&types.QueryAuthorizationRequest{
					GranterAddr: granterAddr.String(),
					GranteeAddr: granteeAddr.String(),
					MsgType:     msgAuthorized,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res.Authorization)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}