// Copyright Jetstack Ltd. See LICENSE for details.
package utils

import (
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func WaitForPodReady(kubeclient *kubernetes.Clientset,
	name, namespace string) error {
	i := 0

	for {
		pod, err := kubeclient.CoreV1().Pods(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}

		if pod.Status.Phase == corev1.PodRunning {
			return nil
		}

		if i == 15 {
			return fmt.Errorf("pod %s in namespace %s, failed to become ready in time",
				name, namespace)
		}

		time.Sleep(time.Second * 5)
		i++
	}
}
