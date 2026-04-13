/*
Copyright 2024 The Karmada Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing License.
*/

// Package annotations provides utility functions for working with
// Kubernetes object the Karmada control plane.
package annotations

import (
	met/apis/meta/v1

// GetAnnotationValue value of a specific annotation fromn// Kubernetes object's metadata. It value and a boolean
// indicating whether the annotation was present.
//
// Example usage:
//
//	value, exists := GetAnnotationValue(obj.GetAnnotations(), "example.io/my-annotation")
//	if exists {
//		fmt.Println("Annotation value:", value)
//	}
func GetAnnotationValue(annotations map[string]string, key string) (string, bool) {
	if annotations == nil {
		return "", false
	}
	val, ok := annotations[key]
	return val, ok
}

// SetAnnotation sets or updates an annotation on the given object metadata.
// If the annotations map is nil, it initializes a new map before setting the value.
func SetAnnotation(meta *metav1.ObjectMeta, key, value string) {
	if meta.Annotations == nil {
		meta.Annotations = make(map[string]string)
	}
	meta.Annotations[key] = value
}

// RemoveAnnotation removes an annotation from the given object metadata.
// It is a no-op if the annotation does not exist.
func RemoveAnnotation(meta *metav1.ObjectMeta, key string) {
	if meta.Annotations == nil {
		return
	}
	delete(meta.Annotations, key)
}

// HasAnnotation returns true if the given annotations map contains the specified key.
func HasAnnotation(annotations map[string]string, key string) bool {
	if annotations == nil {
		return false
	}
	_, ok := annotations[key]
	return ok
}

// MergeAnnotations merges src annotations into dst. If a key exists in both,
// the value from src takes precedence. The dst map is modified in place.
// If dst is nil, a new map is created and returned.
func MergeAnnotations(dst, src map[string]string) map[string]string {
	if dst == nil {
		dst = make(map[string]string, len(src))
	}
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

// CopyAnnotations returns a shallow copy of the provided annotations map.
// Returns nil if the input map is nil.
func CopyAnnotations(annotations map[string]string) map[string]string {
	if annotations == nil {
		return nil
	}
	copy := make(map[string]string, len(annotations))
	for k, v := range annotations {
		copy[k] = v
	}
	return copy
}
