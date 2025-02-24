package openrouter

import (
	"context"
	"net/http"
)

// ListModels — API call to List all models.
func (c *Client) ListModels(ctx context.Context) (response *ListModelsResponse, err error) {
	urlSuffix := "/models"
	req, err := c.requestBuilder.Build(ctx, http.MethodGet, c.fullURL(urlSuffix), nil)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

// Free returns all models that are completely free to use.
func (r *ListModelsResponse) Free() []*Model {
	var models []*Model
	for _, model := range r.Models {
		if model.Pricing.Completion == "0" && model.Pricing.Prompt == "0" && model.Pricing.Request == "0" && model.Pricing.Image == "0" {
			models = append(models, model)
		}
	}
	return models
}
