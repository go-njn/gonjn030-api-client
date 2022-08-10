package internal

import (
	"context"
	"encoding/json"
	"github.com/go-njn/gonjn030-api-client/pkg/domain"
	"github.com/go-njn/gonjn030-api-client/shared"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

func NewUserApiClient(config shared.Config) shared.UserApiClient {
	var logger = logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	return &apiClient{
		config: config,
		logger: logger,
		client: WithLogger(&http.Client{
			Transport: http.DefaultTransport,
			Timeout:   time.Duration(config.TimeoutSeconds) * time.Second,
		}, logger),
	}
}

type apiClient struct {
	config shared.Config
	logger *logrus.Logger
	client *http.Client
}

func (c apiClient) GetAll(ctx context.Context) ([]domain.User, error) {
	var result []domain.User

	err := c.actionForRef(ctx, http.MethodGet, c.config.BaseUserApiUrl, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c apiClient) GetById(ctx context.Context, id domain.ItemId) (domain.User, error) {
	apiUrl := joinPath(c.config.BaseUserApiUrl, id)

	var result domain.User

	err := c.actionForRef(ctx, http.MethodGet, apiUrl, nil, &result)
	if err != nil {
		return domain.User{}, err
	}

	return result, nil
}

func (c apiClient) Create(ctx context.Context, user domain.User) (domain.ItemId, error) {
	body := asJSON(user)

	var result domain.ItemId

	err := c.actionForRef(ctx, http.MethodPost, c.config.BaseUserApiUrl, body, &result)
	if err != nil {
		return result, err
	}

	c.logger.Tracef("%+v", result)

	return result, err
}

func (c apiClient) Update(ctx context.Context, id domain.ItemId, item domain.User) error {
	apiUrl := joinPath(c.config.BaseUserApiUrl, id)
	body := asJSON(item)

	var trace domain.User

	err := c.actionForRef(ctx, http.MethodPut, apiUrl, body, &trace)
	if err != nil {
		return err
	}

	c.logger.Tracef("%+v", trace)

	return nil
}

func (c apiClient) UpdateStatus(ctx context.Context, id domain.ItemId, status domain.UserStatus) error {
	apiUrl := joinPath(c.config.BaseUserApiUrl, id)
	userData := domain.User{Status: status}
	body := asJSON(userData)

	var trace domain.User

	err := c.actionForRef(ctx, http.MethodPut, apiUrl, body, &trace)
	if err != nil {
		return err
	}

	//TODO: looks like REST violation - partial PUT must remove all non-mentioned fields, but it do a PATCH instead

	c.logger.Tracef("%+v", trace)

	return nil
}

func (c apiClient) UpdateGender(ctx context.Context, id domain.ItemId, gender domain.UserGender) error {
	apiUrl := joinPath(c.config.BaseUserApiUrl, id)
	userData := domain.User{Gender: gender}
	body := asJSON(userData)

	var trace domain.User

	err := c.actionForRef(ctx, http.MethodPatch, apiUrl, body, &trace)
	if err != nil {
		return err
	}

	//TODO: looks like REST violation - partial PUT must remove all non-mentioned fields, but it do a PATCH instead

	c.logger.Tracef("%+v", trace)

	return nil
}

func (c apiClient) Delete(ctx context.Context, id domain.ItemId) error {
	apiUrl := joinPath(c.config.BaseUserApiUrl, id)

	err := c.actionForRef(ctx, http.MethodDelete, apiUrl, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c apiClient) actionForRef(ctx context.Context, httpMethod string, apiUrl string, reqBodyJSON io.Reader, resultRef any) error {
	c.logger.Tracef("apiClient::do %T %s %s", resultRef, httpMethod, apiUrl)

	request, err := http.NewRequestWithContext(ctx, httpMethod, apiUrl, reqBodyJSON)
	if err != nil {
		return err
	}

	request.Header.Set(ContentType, ApplicationJson)
	request.Header.Set(Authorization, c.config.BearerToken)

	response, err := c.client.Do(request)
	if err != nil {
		return err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if resultRef == nil {
		return nil
	}

	if response.ContentLength == 0 {
		return nil
	}

	if err := json.NewDecoder(response.Body).Decode(&resultRef); err != nil {
		return err
	}

	return nil
}
