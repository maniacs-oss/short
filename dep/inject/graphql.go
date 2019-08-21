package inject

import (
	"short/app/adapter/graphql"
	"short/app/usecase/requester"
	"short/app/usecase/url"
	"short/fw"
	"short/modern/mdgraphql"
)

type GraphQlPath string

func GraphGophers(graphqlPath GraphQlPath, logger fw.Logger, tracer fw.Tracer, g fw.GraphQlApi) fw.Server {
	return mdgraphql.NewGraphGophers(string(graphqlPath), logger, tracer, g)
}

func ShortGraphQlApi(
	logger fw.Logger,
	tracer fw.Tracer,
	urlRetriever url.Retriever,
	urlCreator url.Creator,
	requesterVerifier requester.Verifier,
) fw.GraphQlApi {
	return graphql.NewShort(
		logger,
		tracer,
		urlRetriever,
		urlCreator,
		requesterVerifier,
	)
}
