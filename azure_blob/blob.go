// azureblob/blob.go
package azureblob

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
)

const containerName = "animezonecontatiner"

// UploadFile uploads a file to Azure Blob Storage and returns the file URL
func UploadFile(fileName string, content []byte, contentType string) (string, error) {
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
	blockBlobClient := containerClient.NewBlockBlobClient(fileName)

	// Setting blob properties
	var httpHeaders = blob.HTTPHeaders{
		BlobContentType: &contentType,
	}
	// Upload the file content
	_, err = blockBlobClient.UploadBuffer(context.TODO(), content, &azblob.UploadBufferOptions{
		HTTPHeaders: &httpHeaders,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	// Construct the URL for the uploaded file
	fileURL := blockBlobClient.URL()

	fmt.Println("File uploaded successfully! URL: ", fileURL)
	return fileURL, nil
}
