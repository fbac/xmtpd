package constants

const (
	JWT_DOMAIN_SEPARATION_LABEL        = "jwt|"
	PAYER_DOMAIN_SEPARATION_LABEL      = "payer|"
	ORIGINATOR_DOMAIN_SEPARATION_LABEL = "originator|"
	NODE_AUTHORIZATION_HEADER_NAME     = "node-authorization"
	MAX_BLOCKCHAIN_ORIGINATOR_ID       = 100
)

type VerifiedNodeRequestCtxKey struct{}
