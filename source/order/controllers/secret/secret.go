package secret

import (
	"context"
	"fmt"
	"order/structs"
	"os"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/keyvault/keyvault"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/jongio/azidext/go/azidext"
	"github.com/kataras/iris/v12"
)

func init() {
	fmt.Println("package: order.controllers.secret - initialized")
}

// PostAuthHandler function
func PostAuthHandler(ctx iris.Context) {

	// verify body content
	var body structs.SecretBody
	readJSONErr := ctx.ReadJSON(&body)

	if readJSONErr != nil {
		ctx.Writef(fmt.Sprintf("Error verifying body: %v", readJSONErr))
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}

	// set variables
	vaultName := os.Getenv("KEYVAULT_NAME")

	if vaultName == "" {
		ctx.Writef(fmt.Sprintf("Error verifying vault name"))
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	cliendID := os.Getenv("AZURE_CLIENT_ID")

	if cliendID == "" {
		ctx.Writef(fmt.Sprintf("Error verifying client id"))
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	// instantiate a new ClientSecretCredential as specified in the documentation
	cred, err := azidentity.NewManagedIdentityCredential(cliendID, nil)

	if err != nil {
		ctx.Writef(fmt.Sprintf("Failed to get credential: %v", err))
		ctx.StatusCode(iris.StatusInternalServerError)
	}

	// call azidext.NewAzureIdentityCredentialAdapter with the azidentity credential and necessary scope
	// NOTE: Scopes define the set of resources and permissions that the credential will have assigned to it.
	// 		 To read more about scopes, see: https://docs.microsoft.com/en-us/azure/active-directory/develop/v2-permissions-and-consent
	authorizer := azidext.NewAzureIdentityCredentialAdapter(
		cred,
		azcore.AuthenticationPolicyOptions{
			Options: azcore.TokenRequestOptions{
				Scopes: []string{"https://vault.azure.net"}}})

	if err != nil {
		ctx.Writef(fmt.Sprintf("Failed to get credential: %v", err))
		ctx.StatusCode(iris.StatusInternalServerError)
	}

	basicClient := keyvault.New()
	basicClient.Authorizer = authorizer

	secretResp, getSecretErr := basicClient.GetSecret(context.Background(), "https://"+vaultName+".vault.azure.net", body.SecretName, "")

	if getSecretErr != nil {
		ctx.Writef(fmt.Sprintf("Unable to get the value for secret: %v", getSecretErr))
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}

	ctx.Writef(fmt.Sprintf("Secret value: %v", *secretResp.Value))
	ctx.StatusCode(iris.StatusOK)
}
