/*
 * Minio Client (C) 2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import "fmt"

/// Collection of standard errors

// APINotImplemented - api not implemented
type APINotImplemented struct {
	API     string
	APIType string
}

func (e APINotImplemented) Error() string {
	return "‘" + e.API + "’ feature " + "is not implemented for ‘" + e.APIType + "’."
}

// GenericBucketError - generic bucket operations error
type GenericBucketError struct {
	Bucket string
}

// BucketDoesNotExist - bucket does not exist.
type BucketDoesNotExist GenericBucketError

func (e BucketDoesNotExist) Error() string {
	return "Bucket ‘" + e.Bucket + "’ does not exist."
}

// BucketExists - bucket exists.
type BucketExists GenericBucketError

func (e BucketExists) Error() string {
	return "Bucket ‘" + e.Bucket + "’ exists."
}

// InvalidBucketName - bucket name invalid (http://goo.gl/wJlzDz)
type InvalidBucketName GenericBucketError

func (e InvalidBucketName) Error() string {
	return "Invalid bucketname [‘" + e.Bucket + "’], please read http://goo.gl/wJlzDz."
}

// BucketNameEmpty - bucket name empty (http://goo.gl/wJlzDz)
type BucketNameEmpty struct{}

func (e BucketNameEmpty) Error() string {
	return "Bucket name cannot be empty."
}

// ObjectAlreadyExists - typed return for MethodNotAllowed
type ObjectAlreadyExists struct {
	Object string
}

func (e ObjectAlreadyExists) Error() string {
	return "Object ‘" + e.Object + "’ already exists."
}

// BucketNameTopLevel - generic error
type BucketNameTopLevel struct{}

func (e BucketNameTopLevel) Error() string {
	return "Buckets can only be created at the top level."
}

// GenericFileError - generic file error.
type GenericFileError struct {
	Path string
}

// PathNotFound (ENOENT) - file not found.
type PathNotFound GenericFileError

func (e PathNotFound) Error() string {
	return "Requested file ‘" + e.Path + "’ not found"
}

// PathIsNotRegular (ENOTREG) - file is not a regular file.
type PathIsNotRegular GenericFileError

func (e PathIsNotRegular) Error() string {
	return "Requested file ‘" + e.Path + "’ is not a regular file."
}

// PathInsufficientPermission (EPERM) - permission denied.
type PathInsufficientPermission GenericFileError

func (e PathInsufficientPermission) Error() string {
	return "Insufficient permissions to access this file ‘" + e.Path + "’"
}

// BrokenSymlink (ENOTENT) - file has broken symlink.
type BrokenSymlink GenericFileError

func (e BrokenSymlink) Error() string {
	return "Requested file ‘" + e.Path + "’ has broken symlink"
}

// TooManyLevelsSymlink (ELOOP) - file has too many levels of symlinks.
type TooManyLevelsSymlink GenericFileError

func (e TooManyLevelsSymlink) Error() string {
	return "Requested file ‘" + e.Path + "’ has too many levels of symlinks"
}

// EmptyPath (EINVAL) - invalid argument.
type EmptyPath struct{}

func (e EmptyPath) Error() string {
	return "Invalid path, path cannot be empty"
}

// ObjectMissing (EINVAL) - object key missing.
type ObjectMissing struct{}

func (e ObjectMissing) Error() string {
	return "Object key is missing, object key cannot be empty"
}

// UnexpectedShortWrite - write wrote less bytes than expected.
type UnexpectedShortWrite struct {
	InputSize int
	WriteSize int
}

func (e UnexpectedShortWrite) Error() string {
	msg := fmt.Sprintf("Wrote less data than requested. Expected ‘%d’ bytes, but only wrote ‘%d’ bytes.", e.InputSize, e.WriteSize)
	return msg
}

// UnexpectedEOF (EPIPE) - reader closed prematurely.
type UnexpectedEOF struct {
	TotalSize    int64
	TotalWritten int64
}

func (e UnexpectedEOF) Error() string {
	msg := fmt.Sprintf("Input reader closed pre-maturely. Expected ‘%d’ bytes, but only received ‘%d’ bytes.", e.TotalSize, e.TotalWritten)
	return msg
}

// UnexpectedExcessRead - reader wrote more data than requested.
type UnexpectedExcessRead UnexpectedEOF

func (e UnexpectedExcessRead) Error() string {
	msg := fmt.Sprintf("Received excess data on input reader. Expected only ‘%d’ bytes, but received ‘%d’ bytes.", e.TotalSize, e.TotalWritten)
	return msg
}
