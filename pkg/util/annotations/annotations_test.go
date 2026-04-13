/*
Copyright 2024 The Karmada Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

UnWITHOUT WARRANTIES OR CONDITIONS OF ANY KIND permissions and
limitations under the License.
*/

 (
	"testing\tav1 "k8s.io/apimachinery/pkg/apis/meta/v1	 *testing.T) {		name        jectMeta  metav1.ObjectMeta
		annotation  string
		expected    string
	}{
		{
			name: "annotation exists",
			objectMeta: metav1.ObjectMeta{
				Annotations: map[string]string{
					"key": "value",
				},
			},
			annotation: "key",
			expected:   "value",
		},
		{
			name:       "annotation does not exist",
			objectMeta: metav1.ObjectMeta{},
			annotation: "missing-key",
			expected:   "",
		},
		{
			name: "nil annotations map",
			objectMeta: metav1.ObjectMeta{
				Annotations: nil,
			},
			annotation: "key",
			expected:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAnnotationValue(tt.objectMeta, tt.annotation)
			if got != tt.expected {
				t.Errorf("GetAnnotationValue() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSetAnnotation(t *testing.T) {
	tests := []struct {
		name       string
		objectMeta *metav1.ObjectMeta
		key        string
		value      string
	}{
		{
			name:       "set annotation on empty object",
			objectMeta: &metav1.ObjectMeta{},
			key:        "new-key",
			value:      "new-value",
		},
		{
			name: "overwrite existing annotation",
			objectMeta: &metav1.ObjectMeta{
				Annotations: map[string]string{
					"existing-key": "old-value",
				},
			},
			key:   "existing-key",
			value: "new-value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetAnnotation(tt.objectMeta, tt.key, tt.value)
			if got := tt.objectMeta.Annotations[tt.key]; got != tt.value {
				t.Errorf("SetAnnotation() annotation value = %v, want %v", got, tt.value)
			}
		})
	}
}

func TestHasAnnotation(t *testing.T) {
	tests := []struct {
		name       string
		objectMeta metav1.ObjectMeta
		key        string
		expected   bool
	}{
		{
			name: "annotation exists",
			objectMeta: metav1.ObjectMeta{
				Annotations: map[string]string{"key": "value"},
			},
			key:      "key",
			expected: true,
		},
		{
			name:       "annotation does not exist",
			objectMeta: metav1.ObjectMeta{},
			key:        "missing",
			expected:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasAnnotation(tt.objectMeta, tt.key)
			if got != tt.expected {
				t.Errorf("HasAnnotation() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestRemoveAnnotation(t *testing.T) {
	objectMeta := &metav1.ObjectMeta{
		Annotations: map[string]string{
			"key-to-remove": "value",
			"key-to-keep":   "value",
		},
	}

	RemoveAnnotation(objectMeta, "key-to-remove")

	if _, exists := objectMeta.Annotations["key-to-remove"]; exists {
		t.Errorf("RemoveAnnotation() failed: annotation 'key-to-remove' still exists")
	}
	if _, exists := objectMeta.Annotations["key-to-keep"]; !exists {
		t.Errorf("RemoveAnnotation() incorrectly removed 'key-to-keep'")
	}
}
