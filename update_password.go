package pwrotator

import (
  "context"
  "fmt"
  "log"
  "golang.org/x/oauth2/google"
  "google.golang.org/api/sqladmin/v1beta4"
  secretmanager "cloud.google.com/go/secretmanager/apiv1"
  secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// PubSubMessage is the payload of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type PubSubMessage struct {
  Data []byte `json:"data"`
}

// SecretsUpdater consumes a Pub/Sub message and updates a secret.
func SecretsUpdater(ctx context.Context, m PubSubMessage) error {
  parent := "projects/<PROJECT_NUMBER>/secrets/my-test-secret-1"
  password := []byte("insert your generated password here")
  updateCloudSQL(password)
  addSecretVersion(parent, password)
  return nil
}

// addSecretVersion adds a new secret version to the given secret with the
// provided payload.
func addSecretVersion(parent string, payload []byte) error {
  // Create the client.
  ctx := context.Background()
  client, err := secretmanager.NewClient(ctx)
  if err != nil {
    return fmt.Errorf("failed to create secretmanager client: %v", err)
  }

  // Build the request.
  req := &secretmanagerpb.AddSecretVersionRequest{
    Parent: parent,
    Payload: &secretmanagerpb.SecretPayload{
      Data: payload,
    },
  }

  // Call the API.
  result, err := client.AddSecretVersion(ctx, req)
  if err != nil {
    return fmt.Errorf("failed to add secret version: %v", err)
  }
  log.Println("Updated the secret version")
  log.Println(result.Name)
  return nil
}

func updateCloudSQL(password []byte) {
  ctx := context.Background()
  c, err := google.DefaultClient(ctx, sqladmin.CloudPlatformScope)
  if err != nil {
    log.Fatal(err)
  }

  sqladminService, err := sqladmin.New(c)
  if err != nil {
    log.Fatal(err)
  }

  // Project ID of the project that contains the instance.
  project := "<PROJECT_NUMBER>"
  // Database instance ID. This does not include the project ID.
  instance := "my-db-instance"
  rb := &sqladmin.User{
    Host:     "%",
    Name:     "root",
    Password: string(password[:]),
  }

  resp, err := sqladminService.Users.Update(project, instance, rb).Context(ctx).Do()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("%#v\n", resp)
}
