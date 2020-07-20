package server

import (
	"github.com/go-redis/redis"
	"golang.org/x/net/context"
	"sbs-entrytask-template/agent/cache"
	grpc_cache "sbs-entrytask-template/apps/example/grpc/proto/cache_proto"
	"sbs-entrytask-template/libs/response/Errorcode"
)


var EXPIRATION_TIME  int32 = 86400
var WRONG_RETURN_VALUE int32 = -1


type CacheServer struct{
	grpc_cache.CacheServiceServer
}


func build_grpc_response(Value int32, Retcode int32, Message string) *grpc_cache.CacheResponse{
	if Retcode == 0 {
		return &grpc_cache.CacheResponse{
			Value:    Value,
			Retcode: nil,
		}
	}

	return &grpc_cache.CacheResponse{
		Value: WRONG_RETURN_VALUE,
		Retcode: &grpc_cache.RetCodeResponse{
			Errorcode:  Retcode,
			Message:  Message,
		},
	}
}


func (c *CacheServer) GetInfo(ctx context.Context, req *grpc_cache.CacheRequest) (*grpc_cache.CacheResponse, error){
	key := req.Key
	if key == "" {
		res := build_grpc_response( WRONG_RETURN_VALUE, Errorcode.PARAMS_WRONG, "Key is empty.")
		return res, nil
	}

	value, err := cache.CommonCache.Get(key)
	if err != nil {
		var errorcode int32
		var message string
		if err == redis.Nil {
			errorcode = Errorcode.CACHE_KEY_NOT_EXIST
			message = "Cannot get cookie info."
		}else{
			errorcode = Errorcode.CACHE_PROBLEM
			message = "Cache Wrong."
		}
		res := build_grpc_response(WRONG_RETURN_VALUE, errorcode , message)
		return res, nil
	}
	res := build_grpc_response(value.(int32), 0 , "")
	return res, nil
}


func (c *CacheServer) SetInfo(ctx context.Context, req *grpc_cache.CacheRequest) (*grpc_cache.CacheResponse, error){
	key := req.Key
	value := req.Value

	if key == "" {
		res := build_grpc_response(WRONG_RETURN_VALUE, Errorcode.PARAMS_WRONG, "Key is empty.")
		return res, nil
	}

	err := cache.CommonCache.Set(key, value, EXPIRATION_TIME)
	if err != nil {
		res := build_grpc_response(WRONG_RETURN_VALUE, Errorcode.CACHE_PROBLEM , "Cache Wrong.")
		return res, nil
	}

	res := build_grpc_response(value, 0, "")
	return res, nil
}