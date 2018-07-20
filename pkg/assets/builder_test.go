/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package assets

import (
	"testing"

	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/apis/kops/util"
)

func buildAssetBuilder(t *testing.T) *AssetBuilder {

	builder := &AssetBuilder{
		AssetsLocation:  &kops.Assets{},
		ContainerAssets: []*ContainerAsset{},
	}
	return builder
}

func TestValidate_RemapImage_ContainerProxy_AppliesToDockerHub(t *testing.T) {
	builder := buildAssetBuilder(t)

	proxyURL := "proxy.example.com/"
	image := "weaveworks/weave-kube"
	expected := "proxy.example.com/weaveworks/weave-kube"

	builder.AssetsLocation.ContainerProxy = &proxyURL

	remapped, err := builder.RemapImage(image)
	if err != nil {
		t.Error("Error remapping image", err)
	}

	if remapped != expected {
		t.Errorf("Error remapping image (Expecting: %s, got %s)", expected, remapped)
	}
}

func TestValidate_RemapImage_ContainerProxy_AppliesToSimplifiedDockerHub(t *testing.T) {
	builder := buildAssetBuilder(t)

	proxyURL := "proxy.example.com/"
	image := "debian"
	expected := "proxy.example.com/debian"

	builder.AssetsLocation.ContainerProxy = &proxyURL

	remapped, err := builder.RemapImage(image)
	if err != nil {
		t.Error("Error remapping image", err)
	}

	if remapped != expected {
		t.Errorf("Error remapping image (Expecting: %s, got %s)", expected, remapped)
	}
}

func TestValidate_RemapImage_ContainerProxy_AppliesToSimplifiedKubernetesURL(t *testing.T) {
	builder := buildAssetBuilder(t)

	proxyURL := "proxy.example.com/"
	image := "k8s.gcr.io/kube-apiserver"
	expected := "proxy.example.com/kube-apiserver"
	version, _ := util.ParseKubernetesVersion("1.10")

	builder.KubernetesVersion = *version
	builder.AssetsLocation.ContainerProxy = &proxyURL

	remapped, err := builder.RemapImage(image)
	if err != nil {
		t.Error("Error remapping image", err)
	}

	if remapped != expected {
		t.Errorf("Error remapping image (Expecting: %s, got %s)", expected, remapped)
	}
}

func TestValidate_RemapImage_ContainerProxy_AppliesToLegacyKubernetesURL(t *testing.T) {
	builder := buildAssetBuilder(t)

	proxyURL := "proxy.example.com/"
	image := "gcr.io/google_containers/kube-apiserver"
	expected := "proxy.example.com/google_containers/kube-apiserver"

	builder.AssetsLocation.ContainerProxy = &proxyURL

	remapped, err := builder.RemapImage(image)
	if err != nil {
		t.Error("Error remapping image", err)
	}

	if remapped != expected {
		t.Errorf("Error remapping image (Expecting: %s, got %s)", expected, remapped)
	}
}

func TestValidate_RemapImage_ContainerProxy_AppliesToImagesWithTags(t *testing.T) {
	builder := buildAssetBuilder(t)

	proxyURL := "proxy.example.com/"
	image := "k8s.gcr.io/kube-apiserver:1.2.3"
	expected := "proxy.example.com/kube-apiserver:1.2.3"
	version, _ := util.ParseKubernetesVersion("1.10")

	builder.KubernetesVersion = *version
	builder.AssetsLocation.ContainerProxy = &proxyURL

	remapped, err := builder.RemapImage(image)
	if err != nil {
		t.Error("Error remapping image", err)
	}

	if remapped != expected {
		t.Errorf("Error remapping image (Expecting: %s, got %s)", expected, remapped)
	}
}

func TestValidate_RemapImage_ContainerProxy_DoesNotApplyToSimplifiedDockerHubIfRegistryMirrorIsSet(t *testing.T) {
	builder := buildAssetBuilder(t)

	proxyURL := "proxy.example.com/"
	image := "debian"
	expected := image

	builder.AssetsLocation.ContainerProxy = &proxyURL
	builder.UseDockerRegistryMirror = true

	remapped, err := builder.RemapImage(image)
	if err != nil {
		t.Error("Error remapping image", err)
	}

	if remapped != expected {
		t.Errorf("Error remapping image (Expecting: %s, got %s)", expected, remapped)
	}
}

func TestValidate_RemapImage_ContainerProxy_AppliesToSimplifiedKubernetesURLIfRegistryMirrorIsSet(t *testing.T) {
	builder := buildAssetBuilder(t)

	proxyURL := "proxy.example.com/"
	image := "k8s.gcr.io/kube-apiserver"
	expected := "proxy.example.com/kube-apiserver"
	version, _ := util.ParseKubernetesVersion("1.10")

	builder.AssetsLocation.ContainerProxy = &proxyURL
	builder.UseDockerRegistryMirror = true
	builder.KubernetesVersion = *version

	remapped, err := builder.RemapImage(image)
	if err != nil {
		t.Error("Error remapping image", err)
	}

	if remapped != expected {
		t.Errorf("Error remapping image (Expecting: %s, got %s)", expected, remapped)
	}
}
