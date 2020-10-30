# golang-managed-identity
 
Considerations:

- This application must run in Azure environment, since it's working with managed identity.
- The application container needs the following environment variables: 
    - KEYVAULT_NAME: The name of the key vault.
    - AZURE_CLIENT_ID: The managed identity client id.

Ensure you follow the configuration mentioned here: https://github.com/Azure-Samples/azure-sdk-for-go-samples/tree/master/keyvault/examples

API Invokation:

Method: Post

Url: https://localhost/api/v1/order/secret

Body:
{
    "SecretName" : "THE NAME OF THE KEY VAULT SECRET"
}