// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// V1AdminCategoriesCategoryIDDelete invokes DELETE /v1/admin/categories/{categoryID} operation.
	//
	// Удаляет категорию по ее уникальному идентификатору.
	//
	// DELETE /v1/admin/categories/{categoryID}
	V1AdminCategoriesCategoryIDDelete(ctx context.Context, params V1AdminCategoriesCategoryIDDeleteParams) (V1AdminCategoriesCategoryIDDeleteRes, error)
	// V1AdminCategoriesCategoryIDPut invokes PUT /v1/admin/categories/{categoryID} operation.
	//
	// Обновляет категорию.
	//
	// PUT /v1/admin/categories/{categoryID}
	V1AdminCategoriesCategoryIDPut(ctx context.Context, request *CategoryPut, params V1AdminCategoriesCategoryIDPutParams) (V1AdminCategoriesCategoryIDPutRes, error)
	// V1AdminCategoriesGet invokes GET /v1/admin/categories operation.
	//
	// Возвращает список всех категорий с возможностью
	// пагинации.
	//
	// GET /v1/admin/categories
	V1AdminCategoriesGet(ctx context.Context, params V1AdminCategoriesGetParams) (V1AdminCategoriesGetRes, error)
	// V1AdminCategoriesPost invokes POST /v1/admin/categories operation.
	//
	// Создает новую категорию.
	//
	// POST /v1/admin/categories
	V1AdminCategoriesPost(ctx context.Context, request *CategoryPost) (V1AdminCategoriesPostRes, error)
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	sec       SecuritySource
	baseClient
}
type errorHandler interface {
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

var _ Handler = struct {
	errorHandler
	*Client
}{}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, sec SecuritySource, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		sec:        sec,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// V1AdminCategoriesCategoryIDDelete invokes DELETE /v1/admin/categories/{categoryID} operation.
//
// Удаляет категорию по ее уникальному идентификатору.
//
// DELETE /v1/admin/categories/{categoryID}
func (c *Client) V1AdminCategoriesCategoryIDDelete(ctx context.Context, params V1AdminCategoriesCategoryIDDeleteParams) (V1AdminCategoriesCategoryIDDeleteRes, error) {
	res, err := c.sendV1AdminCategoriesCategoryIDDelete(ctx, params)
	return res, err
}

func (c *Client) sendV1AdminCategoriesCategoryIDDelete(ctx context.Context, params V1AdminCategoriesCategoryIDDeleteParams) (res V1AdminCategoriesCategoryIDDeleteRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/v1/admin/categories/{categoryID}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "V1AdminCategoriesCategoryIDDelete",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v1/admin/categories/"
	{
		// Encode "categoryID" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "categoryID",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.CategoryID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BearerAuth"
			switch err := c.securityBearerAuth(ctx, "V1AdminCategoriesCategoryIDDelete", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BearerAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeV1AdminCategoriesCategoryIDDeleteResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// V1AdminCategoriesCategoryIDPut invokes PUT /v1/admin/categories/{categoryID} operation.
//
// Обновляет категорию.
//
// PUT /v1/admin/categories/{categoryID}
func (c *Client) V1AdminCategoriesCategoryIDPut(ctx context.Context, request *CategoryPut, params V1AdminCategoriesCategoryIDPutParams) (V1AdminCategoriesCategoryIDPutRes, error) {
	res, err := c.sendV1AdminCategoriesCategoryIDPut(ctx, request, params)
	return res, err
}

func (c *Client) sendV1AdminCategoriesCategoryIDPut(ctx context.Context, request *CategoryPut, params V1AdminCategoriesCategoryIDPutParams) (res V1AdminCategoriesCategoryIDPutRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/v1/admin/categories/{categoryID}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "V1AdminCategoriesCategoryIDPut",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v1/admin/categories/"
	{
		// Encode "categoryID" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "categoryID",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.CategoryID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeV1AdminCategoriesCategoryIDPutRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BearerAuth"
			switch err := c.securityBearerAuth(ctx, "V1AdminCategoriesCategoryIDPut", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BearerAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeV1AdminCategoriesCategoryIDPutResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// V1AdminCategoriesGet invokes GET /v1/admin/categories operation.
//
// Возвращает список всех категорий с возможностью
// пагинации.
//
// GET /v1/admin/categories
func (c *Client) V1AdminCategoriesGet(ctx context.Context, params V1AdminCategoriesGetParams) (V1AdminCategoriesGetRes, error) {
	res, err := c.sendV1AdminCategoriesGet(ctx, params)
	return res, err
}

func (c *Client) sendV1AdminCategoriesGet(ctx context.Context, params V1AdminCategoriesGetParams) (res V1AdminCategoriesGetRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/v1/admin/categories"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "V1AdminCategoriesGet",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/admin/categories"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "page" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Page.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "per_page" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "per_page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.PerPage.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BearerAuth"
			switch err := c.securityBearerAuth(ctx, "V1AdminCategoriesGet", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BearerAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeV1AdminCategoriesGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// V1AdminCategoriesPost invokes POST /v1/admin/categories operation.
//
// Создает новую категорию.
//
// POST /v1/admin/categories
func (c *Client) V1AdminCategoriesPost(ctx context.Context, request *CategoryPost) (V1AdminCategoriesPostRes, error) {
	res, err := c.sendV1AdminCategoriesPost(ctx, request)
	return res, err
}

func (c *Client) sendV1AdminCategoriesPost(ctx context.Context, request *CategoryPost) (res V1AdminCategoriesPostRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/v1/admin/categories"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "V1AdminCategoriesPost",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/admin/categories"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeV1AdminCategoriesPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BearerAuth"
			switch err := c.securityBearerAuth(ctx, "V1AdminCategoriesPost", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BearerAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeV1AdminCategoriesPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
