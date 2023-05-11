// Copyright 2023 Woodpecker Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grpc

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/woodpecker-ci/woodpecker/pipeline/rpc/proto"
	"github.com/woodpecker-ci/woodpecker/server"
	"github.com/woodpecker-ci/woodpecker/server/model"
	"github.com/woodpecker-ci/woodpecker/server/store"
)

type WoodpeckerAuthServer struct {
	proto.UnimplementedWoodpeckerAuthServer
	jwtManager       *JWTManager
	agentMasterToken string
	store            store.Store
}

func NewWoodpeckerAuthServer(jwtManager *JWTManager, agentMasterToken string, store store.Store) *WoodpeckerAuthServer {
	return &WoodpeckerAuthServer{jwtManager: jwtManager, agentMasterToken: agentMasterToken, store: store}
}

func (s *WoodpeckerAuthServer) Auth(_ context.Context, req *proto.AuthRequest) (*proto.AuthResponse, error) {
	agent, err := s.getAgent(req.AgentId, req.AgentToken)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.jwtManager.Generate(agent.ID)
	if err != nil {
		return nil, err
	}

	return &proto.AuthResponse{
		Status:      "ok",
		AgentId:     agent.ID,
		AccessToken: accessToken,
	}, nil
}

func (s *WoodpeckerAuthServer) getAgent(agentID int64, agentToken string) (*model.Agent, error) {
	if agentToken == s.agentMasterToken && agentID == -1 {
		agent := new(model.Agent)
		agent.Name = ""
		agent.OwnerID = -1 // system agent
		agent.Token = server.Config.Server.AgentToken
		agent.Backend = ""
		agent.Platform = ""
		agent.Capacity = -1
		err := s.store.AgentCreate(agent)
		if err != nil {
			log.Err(err).Msgf("Error creating system agent: %s", err)
			return nil, err
		}
		return agent, nil
	}

	if agentToken == s.agentMasterToken {
		return s.store.AgentFind(agentID)
	}

	return s.store.AgentFindByToken(agentToken)
}