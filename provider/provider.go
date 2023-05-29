package provider

type Provider interface {
	ProviderBuilder
	Send()
	Result()
}

type ProviderBuilder interface {
	AddToken()
	AddPackageId()
	AddLanguage()
	Build()
}
