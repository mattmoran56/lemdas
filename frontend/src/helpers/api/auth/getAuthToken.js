const getAuthToken = async (code) => {
	if (!code) {
		return null;
	}
	const body = {
		code,
	};
	const requestOptions = {
		method: "POST",
		body: JSON.stringify(body),
	};

	const response = await fetch(
		`${process.env.REACT_APP_AUTH_API_URL}/token`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to authorise user");
	}

	return response.json();
};

export default getAuthToken();
