/*
Copyright 2023 KylinSoft  Co., Ltd.

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

package cert

import (
	"nestos-kubernetes-deployer/pkg/configmanager"
	"nestos-kubernetes-deployer/pkg/utils"

	"github.com/sirupsen/logrus"
)

// GenerateAllCertificates 生成所有证书和密钥
func GenerateAllFiles(clusterID string) ([]utils.StorageContent, error) {
	var certs []utils.StorageContent
	clusterconfig, _ := configmanager.GetClusterConfig(clusterID)

	// 生成root CA 证书和密钥
	rootCACert, err := GenerateAllCA(clusterconfig.CertAsset.RootCaCertPath,
		clusterconfig.CertAsset.RootCaKeyPath, "kubernetes", []string{"kubernetes"})
	if err != nil {
		logrus.Errorf("Error generating root CA:%v", err)
		return nil, err
	}
	//保存root CA证书和密钥到宿主机
	err = SaveCertificateToFile("./", rootCACert.CertRaw)
	if err != nil {
		return nil, err
	}

	err = SavePrivateKeyToFile("./", rootCACert.CertRaw)
	if err != nil {
		return nil, err
	}
	rootCACertContent := utils.StorageContent{
		Path:    utils.CaCrt,
		Mode:    int(utils.CertFileMode),
		Content: rootCACert.CertRaw,
	}
	rootCAKeyContent := utils.StorageContent{
		Path:    utils.CaKey,
		Mode:    int(utils.CertFileMode),
		Content: rootCACert.KeyRaw,
	}

	certs = append(certs, rootCACertContent, rootCAKeyContent)

	// 生成 etcd CA 证书
	etcdCACert, err := GenerateAllCA(clusterconfig.CertAsset.EtcdCaCertPath,
		clusterconfig.CertAsset.EtcdCaKeyPath, "etcd-ca", []string{"etcd-ca"})
	if err != nil {
		return nil, err
	}
	//保存etcd-ca和密钥到宿主机
	err = SaveCertificateToFile("./", etcdCACert.CertRaw)
	if err != nil {
		return nil, err
	}

	err = SavePrivateKeyToFile("./", etcdCACert.CertRaw)
	if err != nil {
		return nil, err
	}

	etcdCACertContent := utils.StorageContent{
		Path:    utils.EtcdCaCrt,
		Mode:    int(utils.CertFileMode),
		Content: etcdCACert.CertRaw,
	}
	etcdCAKeyContent := utils.StorageContent{
		Path:    utils.EtcdCaKey,
		Mode:    int(utils.CertFileMode),
		Content: etcdCACert.KeyRaw,
	}
	certs = append(certs, etcdCACertContent, etcdCAKeyContent)

	// 生成 front-proxy CA 证书
	frontProxyCACert, err := GenerateAllCA(clusterconfig.CertAsset.EtcdCaCertPath,
		clusterconfig.CertAsset.EtcdCaKeyPath, "front-proxy-ca", []string{"front-proxy-ca"})
	if err != nil {
		return nil, err
	}

	//保存front-proxy-ca和密钥到宿主机
	err = SaveCertificateToFile("./", frontProxyCACert.CertRaw)
	if err != nil {
		return nil, err
	}

	err = SavePrivateKeyToFile("./", frontProxyCACert.CertRaw)
	if err != nil {
		return nil, err
	}

	frontProxyCACertContent := utils.StorageContent{
		Path:    utils.FrontProxyCaCrt,
		Mode:    int(utils.CertFileMode),
		Content: frontProxyCACert.CertRaw,
	}
	frontProxyCAKeyContent := utils.StorageContent{
		Path:    utils.FrontProxyCaKey,
		Mode:    int(utils.CertFileMode),
		Content: frontProxyCACert.KeyRaw,
	}
	certs = append(certs, frontProxyCACertContent, frontProxyCAKeyContent)

	return certs, nil
}