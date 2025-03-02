// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The specified layer upload doesn't contain any layer parts.
type EmptyUploadException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *EmptyUploadException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EmptyUploadException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EmptyUploadException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EmptyUploadException"
	}
	return *e.ErrorCodeOverride
}
func (e *EmptyUploadException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified image has already been pushed, and there were no changes to the
// manifest or image tag after the last push.
type ImageAlreadyExistsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ImageAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ImageAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ImageAlreadyExistsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ImageAlreadyExistsException"
	}
	return *e.ErrorCodeOverride
}
func (e *ImageAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified image digest doesn't match the digest that Amazon ECR calculated
// for the image.
type ImageDigestDoesNotMatchException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ImageDigestDoesNotMatchException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ImageDigestDoesNotMatchException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ImageDigestDoesNotMatchException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ImageDigestDoesNotMatchException"
	}
	return *e.ErrorCodeOverride
}
func (e *ImageDigestDoesNotMatchException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The image requested doesn't exist in the specified repository.
type ImageNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ImageNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ImageNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ImageNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ImageNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ImageNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified image is tagged with a tag that already exists. The repository is
// configured for tag immutability.
type ImageTagAlreadyExistsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ImageTagAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ImageTagAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ImageTagAlreadyExistsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ImageTagAlreadyExistsException"
	}
	return *e.ErrorCodeOverride
}
func (e *ImageTagAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The layer digest calculation performed by Amazon ECR when the image layer
// doesn't match the digest specified.
type InvalidLayerException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidLayerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidLayerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidLayerException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidLayerException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidLayerException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The layer part size isn't valid, or the first byte specified isn't consecutive
// to the last byte of a previous layer part upload.
type InvalidLayerPartException struct {
	Message *string

	ErrorCodeOverride *string

	RegistryId            *string
	RepositoryName        *string
	UploadId              *string
	LastValidByteReceived *int64

	noSmithyDocumentSerde
}

func (e *InvalidLayerPartException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidLayerPartException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidLayerPartException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidLayerPartException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidLayerPartException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified parameter is invalid. Review the available parameters for the API
// request.
type InvalidParameterException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidParameterException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An invalid parameter has been specified. Tag keys can have a maximum character
// length of 128 characters, and tag values can have a maximum length of 256
// characters.
type InvalidTagParameterException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidTagParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTagParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTagParameterException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidTagParameterException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidTagParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The image layer already exists in the associated repository.
type LayerAlreadyExistsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LayerAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LayerAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LayerAlreadyExistsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LayerAlreadyExistsException"
	}
	return *e.ErrorCodeOverride
}
func (e *LayerAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Layer parts must be at least 5 MiB in size.
type LayerPartTooSmallException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LayerPartTooSmallException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LayerPartTooSmallException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LayerPartTooSmallException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LayerPartTooSmallException"
	}
	return *e.ErrorCodeOverride
}
func (e *LayerPartTooSmallException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified layers can't be found, or the specified layer isn't valid for
// this repository.
type LayersNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LayersNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LayersNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LayersNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LayersNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *LayersNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation didn't succeed because it would have exceeded a service limit for
// your account. For more information, see [Amazon ECR Service Quotas]in the Amazon Elastic Container
// Registry User Guide.
//
// [Amazon ECR Service Quotas]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/service-quotas.html
type LimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The manifest list is referencing an image that doesn't exist.
type ReferencedImagesNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ReferencedImagesNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReferencedImagesNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReferencedImagesNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ReferencedImagesNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ReferencedImagesNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The registry doesn't exist.
type RegistryNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *RegistryNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RegistryNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RegistryNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RegistryNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *RegistryNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified repository already exists in the specified registry.
type RepositoryAlreadyExistsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *RepositoryAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RepositoryAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RepositoryAlreadyExistsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RepositoryAlreadyExistsException"
	}
	return *e.ErrorCodeOverride
}
func (e *RepositoryAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The repository catalog data doesn't exist.
type RepositoryCatalogDataNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *RepositoryCatalogDataNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RepositoryCatalogDataNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RepositoryCatalogDataNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RepositoryCatalogDataNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *RepositoryCatalogDataNotFoundException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified repository contains images. To delete a repository that contains
// images, you must force the deletion with the force parameter.
type RepositoryNotEmptyException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *RepositoryNotEmptyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RepositoryNotEmptyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RepositoryNotEmptyException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RepositoryNotEmptyException"
	}
	return *e.ErrorCodeOverride
}
func (e *RepositoryNotEmptyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified repository can't be found. Check the spelling of the specified
// repository and ensure that you're performing operations on the correct registry.
type RepositoryNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *RepositoryNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RepositoryNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RepositoryNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RepositoryNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *RepositoryNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified repository and registry combination doesn't have an associated
// repository policy.
type RepositoryPolicyNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *RepositoryPolicyNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RepositoryPolicyNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RepositoryPolicyNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RepositoryPolicyNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *RepositoryPolicyNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// These errors are usually caused by a server-side issue.
type ServerException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServerException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServerException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The list of tags on the repository is over the limit. The maximum number of
// tags that can be applied to a repository is 50.
type TooManyTagsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TooManyTagsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTagsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTagsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyTagsException"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyTagsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The action isn't supported in this Region.
type UnsupportedCommandException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UnsupportedCommandException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedCommandException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedCommandException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UnsupportedCommandException"
	}
	return *e.ErrorCodeOverride
}
func (e *UnsupportedCommandException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The upload can't be found, or the specified upload ID isn't valid for this
// repository.
type UploadNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UploadNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UploadNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UploadNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UploadNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *UploadNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
