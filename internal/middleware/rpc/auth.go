package middleware_rpc

import (
	"context"
	"crypto/rsa"
	"errors"
	"strings"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/roy-hc310/fullmetal-product/internal/shared"
	"github.com/roy-hc310/fullmetal-product/pkg/constant"
	"github.com/roy-hc310/fullmetal-product/pkg/utils"
)

type claimsKey struct{}

func ClaimsFromCtx(ctx context.Context) *shared.Claims {
	if v := ctx.Value(claimsKey{}); v != nil {
		if claims, ok := v.(*shared.Claims); ok {
			return claims
		}
	}
	return nil
}

func JWTAuthMiddleware(pubKey *rsa.PublicKey) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req, resp interface{}) error {

			info := rpcinfo.GetRPCInfo(ctx)
			if info == nil || info.Invocation() == nil {
				return errors.New("missing RPC invocation metadata")
			}

			raw := info.Invocation().Extra("Authorization")

			tokenString, ok := raw.(string)
			if !ok || tokenString == "" {
				return errors.New("missing Authorization header")
			}

			parts := strings.SplitN(tokenString, " ", 2)
			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
				return errors.New("invalid token")
			}
			jwtStr := parts[1]

			claims, err := shared.VerifyJWT(jwtStr, pubKey)
			if err != nil {
				return err
			}
			userID, err := utils.StringToUUID(claims.UserID)
			if err != nil {
				return err
			}
			ctx = context.WithValue(ctx, constant.UserID, userID)
			ctx = context.WithValue(ctx, constant.Role, claims.Role)
			return next(ctx, req, resp)
		}
	}
}
