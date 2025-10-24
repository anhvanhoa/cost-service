package main

import (
	"context"
	"cost_service/bootstrap"
	"cost_service/infrastructure/grpc_client"
	"cost_service/infrastructure/grpc_service"
	cost_tracking_service "cost_service/infrastructure/grpc_service/cost_tracking"

	gc "github.com/anhvanhoa/service-core/domain/grpc_client"
)

func main() {
	StartGRPCServer()
}

func StartGRPCServer() {
	app := bootstrap.App()
	env := app.Env
	log := app.Log

	clientFactory := gc.NewClientFactory(env.GrpcClients...)
	permissionClient := grpc_client.NewPermissionClient(clientFactory.GetClient(env.PermissionServiceAddr))

	costTrackingService := cost_tracking_service.NewCostTrackingService(app.Repo)
	grpcSrv := grpc_service.NewGRPCServer(
		env, log, app.Cache, costTrackingService,
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	permissions := app.Helper.ConvertResourcesToPermissions(grpcSrv.GetResources())
	if _, err := permissionClient.PermissionServiceClient.RegisterPermission(ctx, permissions); err != nil {
		log.Fatal("Failed to register permission: " + err.Error())
	}
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}
