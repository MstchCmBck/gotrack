package provider

type Provider interface {
	Send()
	Result()
}

type ProviderBuilder interface {
	AddToken()
	AddPackageId()
	AddLanguage()
	Build()
}
