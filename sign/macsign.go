//|------------------------------------------------------------------
//| Ali LMQ MQTT
//| Author:Tommy.Jin Dtime:2018-3-20
//|-------------------------------------------------------------------

package sign

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"io"
	"strings"
)

/**
* @param text      要签名的文本
* @param secretKey 阿里云 MQ SecretKey
* @return 加密后的字符串
 */
func MacSignature(text string, secretKey string) string {
	h := sha1.New()
	io.WriteString(h, text)
	key := []byte(secretKey)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(text))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

/**
* 发送方签名方法
*
* @param clientId  MQTT ClientID
* @param secretKey 阿里云 MQ SecretKey
* @return 加密后的字符串
 */
func PublishSignature(clientId string, secretKey string) string {
	return MacSignature(clientId, secretKey)
}

/**
* 接收方签名方法
*
* @param clientId  MQTT ClientID
* @param secretKey 阿里云 MQ SecretKey
* @return 加密后的字符串
 */
func SublishSignature(clientId string, secretKey string) string {
	return MacSignature(clientId, secretKey)
}

/**
 * 订阅方签名方法
 *
 * @param topics    要订阅的 Topic 集合
 * @param clientId  MQTT ClientID
 * @param secretKey 阿里云 MQ SecretKey
 * @return 加密后的字符串
 */
func SubSignatureArr(topics []string, clientId string, secretKey string) string {
	topicText := ""
	for _, value := range topics {
		topicText = topicText + value + "\n"
	}
	text := topicText + clientId
	return MacSignature(text, secretKey)
}

/**
 * 订阅方签名方法
 *
 * @param topics    要订阅的 Topic 集合
 * @param clientId  MQTT ClientID
 * @param secretKey 阿里云 MQ SecretKey
 * @return 加密后的字符串
 */
func SubSignature(topics string, clientId string, secretKey string) string {
	topicsArr := strings.Split(topics, "/")
	topicText := ""
	for _, value := range topicsArr {
		topicText = topicText + value + "\n"
	}
	text := topicText + clientId
	return MacSignature(text, secretKey)
}
