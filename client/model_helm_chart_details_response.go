/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: 0.3.0
 * Contact: info@banzaicloud.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

type HelmChartDetailsResponse struct {
	Chart HelmChartDetailsResponseChart `json:"chart,omitempty"`
	Values string `json:"values,omitempty"`
	Readme string `json:"readme,omitempty"`
}