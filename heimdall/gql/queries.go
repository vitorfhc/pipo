package gql

const getEmployerQuery = `
query GetEmployer($cnpj: String!) {
	employer(where: {cnpj: {_eq: $cnpj}}, limit: 1) {
		cnpj
		password
		uuid
	}
}`
