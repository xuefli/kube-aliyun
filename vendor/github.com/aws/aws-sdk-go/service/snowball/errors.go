// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package snowball

const (

	// ErrCodeClusterLimitExceededException for service response error code
	// "ClusterLimitExceededException".
	//
	// Job creation failed. Currently, clusters support five nodes. If you have
	// less than five nodes for your cluster and you have more nodes to create for
	// this cluster, try again and create jobs until your cluster has exactly five
	// notes.
	ErrCodeClusterLimitExceededException = "ClusterLimitExceededException"

	// ErrCodeInvalidAddressException for service response error code
	// "InvalidAddressException".
	//
	// The address provided was invalid. Check the address with your region's carrier,
	// and try again.
	ErrCodeInvalidAddressException = "InvalidAddressException"

	// ErrCodeInvalidInputCombinationException for service response error code
	// "InvalidInputCombinationException".
	//
	// Job or cluster creation failed. One ore more inputs were invalid. Confirm
	// that the CreateClusterRequest$SnowballType value supports your CreateJobRequest$JobType,
	// and try again.
	ErrCodeInvalidInputCombinationException = "InvalidInputCombinationException"

	// ErrCodeInvalidJobStateException for service response error code
	// "InvalidJobStateException".
	//
	// The action can't be performed because the job's current state doesn't allow
	// that action to be performed.
	ErrCodeInvalidJobStateException = "InvalidJobStateException"

	// ErrCodeInvalidResourceException for service response error code
	// "InvalidResourceException".
	//
	// The specified resource can't be found. Check the information you provided
	// in your last request, and try again.
	ErrCodeInvalidResourceException = "InvalidResourceException"

	// ErrCodeKMSRequestFailedException for service response error code
	// "KMSRequestFailedException".
	//
	// The provided AWS Key Management Service key lacks the permissions to perform
	// the specified CreateJob or UpdateJob action.
	ErrCodeKMSRequestFailedException = "KMSRequestFailedException"

	// ErrCodeUnsupportedAddressException for service response error code
	// "UnsupportedAddressException".
	//
	// The address is either outside the serviceable area for your region, or an
	// error occurred. Check the address with your region's carrier and try again.
	// If the issue persists, contact AWS Support.
	ErrCodeUnsupportedAddressException = "UnsupportedAddressException"
)