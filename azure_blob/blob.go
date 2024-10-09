// azureblob/blob.go
package azureblob

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

const containerName = "animezonecontatiner"

// UploadFile uploads a file to Azure Blob Storage and returns the file URL
func UploadFile(fileName string, content []byte) (string, error) {
	// Get the Blob service client
	serviceClient, err := GetBlobServiceClient()
	if err != nil {
		return "", fmt.Errorf("error getting Blob service client: %w", err)
	}

	// Create the container client
	containerClient := serviceClient.ServiceClient().NewContainerClient(containerName)

	// Create the container if it doesn't exist
	_, err = containerClient.Create(context.TODO(), nil)
	if err != nil {
		fmt.Println("Container may already exist, continuing...")
	}

	// Create the blob client and upload the file
	blobClient := containerClient.NewBlockBlobClient(fileName)

	// Upload the file content
	_, err = blobClient.UploadBuffer(context.TODO(), content, &azblob.UploadBufferOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	// Construct the URL for the uploaded file
	fileURL := blobClient.URL()

	fmt.Println("File uploaded successfully! URL: ", fileURL)
	return fileURL, nil
}
